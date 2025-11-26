package queue

import "errors"

// Task is a type alias for func() – we keep it only for readability/documentation
type Task func()

// Dispatcher is your old internal interface (you can keep it or delete it later)
type Dispatcher interface {
	Enqueue(Task) error
	Close()
}

// BackgroundQueue executes tasks on a dedicated goroutine.
type BackgroundQueue struct {
	tasks   chan Task
	running bool
}

// NewBackgroundQueue creates a new queue with the given buffer size.
func NewBackgroundQueue(buffer int) *BackgroundQueue {
	q := &BackgroundQueue{
		tasks:   make(chan Task, buffer),
		running: true,
	}
	go q.worker()
	return q
}

// Enqueue – the ONLY method you need now.
// It satisfies BOTH your old Dispatcher AND ports.Queue (which expects func()).
func (q *BackgroundQueue) Enqueue(task func()) error {
	if !q.running {
		return errors.New("queue stopped")
	}
	select {
	case q.tasks <- Task(task): // explicit conversion only for clarity
		return nil
	default:
		return errors.New("queue full")
	}
}

// Close stops the worker and drains pending tasks.
func (q *BackgroundQueue) Close() {
	if !q.running {
		return
	}
	q.running = false
	close(q.tasks)
}

func (q *BackgroundQueue) worker() {
	for task := range q.tasks {
		if task != nil {
			task()
		}
	}
}
