package store

import (
	"errors"
	"sort"
	"strconv"
	"sync"
	"time"
)

var ErrNotFound = errors.New("task not found")

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

type Memory struct {
	mu    sync.RWMutex
	tasks map[string]Task
	next  int
}

func NewMemory() *Memory {
	return &Memory{tasks: make(map[string]Task)}
}

func (m *Memory) List() []Task {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]Task, 0, len(m.tasks))
	for _, t := range m.tasks {
		out = append(out, t)
	}
	sort.Slice(out, func(i, j int) bool {
		ai, _ := strconv.Atoi(out[i].ID)
		aj, _ := strconv.Atoi(out[j].ID)
		return ai < aj
	})
	return out
}

func (m *Memory) Create(title string) Task {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.next++
	id := strconv.Itoa(m.next)
	t := Task{
		ID:        id,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now().UTC(),
	}
	m.tasks[t.ID] = t
	return t
}

func (m *Memory) Get(id string) (Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	t, ok := m.tasks[id]
	if !ok {
		return Task{}, ErrNotFound
	}
	return t, nil
}

func (m *Memory) Toggle(id string) (Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	t, ok := m.tasks[id]
	if !ok {
		return Task{}, ErrNotFound
	}
	t.Done = !t.Done
	m.tasks[id] = t
	return t, nil
}

func (m *Memory) Delete(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.tasks[id]; !ok {
		return ErrNotFound
	}
	delete(m.tasks, id)
	return nil
}
