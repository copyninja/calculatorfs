package infix

import (
	"testing"
)

func TestBasic(t *testing.T) {
	s := NewStack()
	s.Push("2")

	o1 := s.Peek().(string)
	if o1 != "2" {
		t.Errorf("s.Peek() == %q, want 2", o1)
	}

	if s.Count() != 1 {
		t.Errorf("s.Count() == %q, want 1", s.Count())
	}

	o2 := s.Pop().(string)
	if o2 != "2" {
		t.Errorf("s.Pop() == %q, want 2", o2)
	}

	if s.Count() != 0 {
		t.Errorf("s.Count() == %q, want 0", s.Count())
	}

	o3 := s.Peek().(string)
	if len(o3) != 0 {
		t.Errorf("s.Peek() == %q, expecting no element", o3)
	}
}

func TestPushPop(t *testing.T) {
	s := NewStack()

	inp := []string{"1", "2", "3", "4", "5"}
	out := []string{"5", "4", "3", "2", "1"}

	for _, x := range inp {
		s.Push(x)
	}

	if s.Count() != 5 {
		t.Errorf("s.Count() == %q, want 5", s.Count())
	}

	if o1 := s.Peek().(string); o1 != out[0] {
		t.Errorf("s.Peek() == %q, want %q", o1, out[0])
	}

	for _, x := range out {
		if o2 := s.Pop().(string); o2 != x {
			t.Errorf("s.Pop() == %q, want %q", o2, x)
		}
	}
}
