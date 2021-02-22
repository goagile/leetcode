package stack

import (
	"testing"
)

//
// minOperations
//
func Test_minOperations_Example1(t *testing.T) {
	want := 2

	got := minOperations([]string{"d1/","d2/","../","d21/","./"})

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_minOperations_Example2(t *testing.T) {
	want := 3

	got := minOperations([]string{"d1/","d2/","./","d3/","../","d31/"})

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_minOperations_Example3(t *testing.T) {
	want := 0

	got := minOperations([]string{"d1/","../","../","../"})

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_minOperations_One(t *testing.T) {
	want := 1

	got := minOperations([]string{"d1/"})

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_minOperations_Two(t *testing.T) {
	want := 2

	got := minOperations([]string{"d1/", "d2/"})

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_minOperations_Back(t *testing.T) {
	want := 0

	got := minOperations([]string{"d1/", "../"})

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_minOperations_Curr(t *testing.T) {
	want := 1

	got := minOperations([]string{"d1/", "./"})

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

//
// Count
//
func Test_Count_Empty(t *testing.T) {
	want := 0

	s := NewStack()

	got := s.Count() 

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

//
// Push
//
func Test_Push(t *testing.T) {
	want := 1

	s := NewStack()

	s.Push("./")

	got := s.Count() 

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_Push_Push(t *testing.T) {
	want := 2

	s := NewStack()

	s.Push("./")
	s.Push("./")

	got := s.Count() 

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

//
// Pop
//
func Test_Pop_Empty(t *testing.T) {
	want := 0

	s := NewStack()
	
	s.Pop()

	got := s.Count() 

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_Pop_One(t *testing.T) {
	want := 0

	s := NewStack()
	s.Push("./")
	
	s.Pop()

	got := s.Count() 

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_Pop_Two(t *testing.T) {
	want := 0

	s := NewStack()
	s.Push("./")
	s.Push("./")
	
	s.Pop()
	s.Pop()

	got := s.Count() 

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}
