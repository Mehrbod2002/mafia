package queue

import "errors"

// Task defines a unit of work executed by the queue.
type Task func()

// Dispatcher provides the minimal interface used by services.
type Dispatcher interface {
	Enqueue(Task) error
	Close()
}

// BackgroundQueue executes tasks on a dedicated goroutine.
type BackgroundQueue struct {
	tasks   chan Task
	running bool
}

// NewBackgroundQueue spins up a queue with the given buffer size.
func NewBackgroundQueue(buffer int) *BackgroundQueue {
	q := &BackgroundQueue{
		tasks:   make(chan Task, buffer),
		running: true,
	}
	go q.worker()
	return q
}

// Enqueue schedules a task for execution.
func (q *BackgroundQueue) Enqueue(task Task) error {
	if !q.running {
		return errors.New("queue stopped")
	}
	select {
	case q.tasks <- task:
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
