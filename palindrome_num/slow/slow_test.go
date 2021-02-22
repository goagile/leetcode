package slow

import (
	"testing"
)

func Test_IsPalindrome_0(t *testing.T) {
	want := true

	got := IsPalindrome(0)

	if got != want {
		t.Fatalf("\ngot:%v\nwant:%v\n", got, want)
	}
}

func Test_IsPalindrome_11(t *testing.T) {
	want := true

	got := IsPalindrome(11)

	if got != want {
		t.Fatalf("\ngot:%v\nwant:%v\n", got, want)
	}
}

func Test_IsPalindrome_121(t *testing.T) {
	want := true

	got := IsPalindrome(121)

	if got != want {
		t.Fatalf("\ngot:%v\nwant:%v\n", got, want)
	}
}

func Test_IsPalindrome_123(t *testing.T) {
	want := false

	got := IsPalindrome(123)

	if got != want {
		t.Fatalf("\ngot:%v\nwant:%v\n", got, want)
	}
}

func Test_IsPalindrome_minus1(t *testing.T) {
	want := false

	got := IsPalindrome(-1)

	if got != want {
		t.Fatalf("\ngot:%v\nwant:%v\n", got, want)
	}
}
