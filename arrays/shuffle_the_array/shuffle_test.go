package main

import (
	"reflect"
	"testing"
)

func shuffle(a []int, n int) []int {
	z := make([]int, 0)
	x := a[:n]
	y := a[n:]
	for i := 0; i < n; i++ {
		z = append(z, x[i])
		z = append(z, y[i])
	}
	return z
}

func Test_shuffle_n_1(t *testing.T) {
	want := []int{1, 2}

	got := shuffle([]int{1, 2}, 1)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\ngot:%v\nwant:%v\n", got, want)
	}
}

func Test_shuffle_n_2(t *testing.T) {
	want := []int{1, 3, 2, 4}

	got := shuffle([]int{1, 2, 3, 4}, 2)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\ngot:%v\nwant:%v\n", got, want)
	}
}

func Test_shuffle_n_3(t *testing.T) {
	want := []int{1, 4, 2, 5, 3, 6}

	got := shuffle([]int{1, 2, 3, 4, 5, 6}, 3)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\ngot:%v\nwant:%v\n", got, want)
	}
}
