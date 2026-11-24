package main

import "testing"

func TestMinStack(t *testing.T) {
	var s MinStack
	if _, ok := s.Min(); ok {
		t.Error("Min on empty must be not-ok")
	}
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	if m, _ := s.Min(); m != -3 {
		t.Errorf("Min = %d, want -3", m)
	}
	if v, _ := s.Pop(); v != -3 {
		t.Errorf("Pop = %d, want -3", v)
	}
	if v, _ := s.Top(); v != 0 {
		t.Errorf("Top = %d, want 0", v)
	}
	if m, _ := s.Min(); m != -2 {
		t.Errorf("Min = %d, want -2 after pop", m)
	}
	// Duplicate minimums survive pops correctly.
	var d MinStack
	d.Push(1)
	d.Push(1)
	d.Pop()
	if m, _ := d.Min(); m != 1 {
		t.Errorf("Min = %d, want 1 with duplicate minimums", m)
	}
}
