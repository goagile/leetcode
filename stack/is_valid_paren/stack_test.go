package stack

import (
	"testing"
)

//
// push
//
func Test_stack_Push_sym(t *testing.T) {
	want := '{'

	s := NewStack()
	s.Push('{')

	got := s.head.val

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_stack_Push_count(t *testing.T) {
	want := 1

	s := NewStack()
	s.Push('{')

	got := s.count

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

// 
// empty
//
func Test_stack_Empty_false(t *testing.T) {
	want := false

	s := NewStack()
	s.Push('{')

	got := s.Empty()

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_stack_Empty_true(t *testing.T) {
	want := true

	s := NewStack()

	got := s.Empty()

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

//
// pop
//
func Test_Pop_rune(t *testing.T) {
	want := '{'

	s := NewStack()
	s.Push('{')

	got := s.Pop()

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_Pop_count(t *testing.T) {
	want := 0

	s := NewStack()
	s.Push('{')
	s.Pop()

	got := s.count

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

//
// isValid
//
func Test_isValid_empty(t *testing.T) {
	want := true

	got := isValid("")

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_isValid_LRBRAK_true(t *testing.T) {
	want := true

	got := isValid("{}")

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_isValid_LRBRAK_false(t *testing.T) {
	want := false

	got := isValid("{{}")

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_isValid_LRSQUARE_true(t *testing.T) {
	want := true

	got := isValid("{[]}")

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_isValid_LRSQUARE_false(t *testing.T) {
	want := false

	got := isValid("{[}]")

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_isValid_LRPAREN_true(t *testing.T) {
	want := true

	got := isValid("({[]})")

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_isValid_LRPAREN_false(t *testing.T) {
	want := false

	got := isValid("({[)]}")

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_isValid_ex1(t *testing.T) {
	want := true

	got := isValid("()")

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_isValid_ex2(t *testing.T) {
	want := true

	got := isValid("()[]{}")

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_isValid_ex3(t *testing.T) {
	want := false

	got := isValid("(]")

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_isValid_ex4(t *testing.T) {
	want := false

	got := isValid("([)]")

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_isValid_ex5(t *testing.T) {
	want := true

	got := isValid("{[]}")

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}