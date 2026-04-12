package store

import (
	"testing"
)

func TestMemory_Create_List_Order(t *testing.T) {
	m := NewMemory()
	m.Create("a")
	m.Create("b")
	tasks := m.List()
	if len(tasks) != 2 {
		t.Fatalf("len = %d, want 2", len(tasks))
	}
	if tasks[0].ID != "1" || tasks[0].Title != "a" {
		t.Fatalf("first task: %+v", tasks[0])
	}
	if tasks[1].ID != "2" || tasks[1].Title != "b" {
		t.Fatalf("second task: %+v", tasks[1])
	}
}

func TestMemory_Toggle_Delete(t *testing.T) {
	m := NewMemory()
	x := m.Create("x")
	if _, err := m.Toggle(x.ID); err != nil {
		t.Fatal(err)
	}
	got, err := m.Get(x.ID)
	if err != nil {
		t.Fatal(err)
	}
	if !got.Done {
		t.Fatal("expected done")
	}
	if err := m.Delete(x.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := m.Get(x.ID); err != ErrNotFound {
		t.Fatalf("Get: %v", err)
	}
}

func TestMemory_ErrNotFound(t *testing.T) {
	m := NewMemory()
	if _, err := m.Toggle("nope"); err != ErrNotFound {
		t.Fatalf("Toggle: %v", err)
	}
	if err := m.Delete("nope"); err != ErrNotFound {
		t.Fatalf("Delete: %v", err)
	}
}
