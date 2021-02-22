package minstack

import (
	"testing"
)

//
// Push
//
func Test_Push_To_Empty(t *testing.T) {
	want := -2

	s := NewStack()
	s.Push(-2)

	got := s.Head.Val

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_Push_To_Not_Empty_Head(t *testing.T) {
	want := 0

	s := NewStack()
	s.Push(-2)
	s.Push(0)

	got := s.Head.Val

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_Push_To_Not_Empty_Count(t *testing.T) {
	want := 2

	s := NewStack()
	s.Push(-2)
	s.Push(0)

	got := s.Count

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

//
// Pop
//
func Test_Pop(t *testing.T) {
	want := -2

	s := NewStack()
	s.Push(-2)

	got, ok := s.Pop()
	if !ok {
		t.Fatal("Need to be ok!")
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_Pop_Empty(t *testing.T) {
	s := NewStack()

	_, ok := s.Pop()

	if ok {
		t.Fatal("Need to be !ok")
	}
}

//
// Top
//
func Test_Top(t *testing.T) {
	want := -2

	s := NewStack()
	s.Push(-2)

	got, ok := s.Top()
	if !ok {
		t.Fatal("Need to be ok!")
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

//
// MinStack
//
func Test_MinStack_GetMin(t *testing.T) {
	want := -3

	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)

	got := s.GetMin()

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_MinStack_Top_After_Pop(t *testing.T) {
	want := 0

	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	s.Pop()

	got := s.Top()

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_MinStack_GetMin_After_Pop(t *testing.T) {
	want := -2

	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	s.Pop()

	got := s.GetMin()

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_MinStack_GetMin_Empty(t *testing.T) {
	want := 0

	s := Constructor()

	got := s.GetMin()

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_MinStack_Top_Empty(t *testing.T) {
	want := 0

	s := Constructor()

	got := s.Top()

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_MinStack_Case_1_GetMin(t *testing.T) {
	want := 0

	s := Constructor()
	s.Push(2)
	s.Push(0)
	s.Push(3)
	s.Push(0)

	got := s.GetMin()

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_MinStack_Case_1_GetMin_After(t *testing.T) {
	want := 0

	s := Constructor()
	s.Push(2)
	s.Push(0)
	s.Push(3)
	s.Push(0)
	s.Pop()

	got := s.GetMin()

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}
