package leaderboard

import "sort"

// Entry represents a user and their score.
type Entry struct {
	UserID string
	Score  int
}

// Board keeps scores by user.
type Board struct {
	scores map[string]int
}

// NewBoard constructs an empty leaderboard.
func NewBoard() *Board {
	return &Board{scores: make(map[string]int)}
}

// Submit records a user's score, keeping the highest value.
func (b *Board) Submit(userID string, score int) {
	if current, ok := b.scores[userID]; !ok || score > current {
		b.scores[userID] = score
	}
}

// Top returns the top n entries ordered by score descending.
func (b *Board) Top(n int) []Entry {
	entries := make([]Entry, 0, len(b.scores))
	for id, score := range b.scores {
		entries = append(entries, Entry{UserID: id, Score: score})
	}
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Score == entries[j].Score {
			return entries[i].UserID < entries[j].UserID
		}
		return entries[i].Score > entries[j].Score
	})
	if n > len(entries) || n <= 0 {
		return entries
	}
	return entries[:n]
}
