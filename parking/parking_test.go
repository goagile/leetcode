package parking

import (
	"reflect"
	"testing"
)

func Test_110(t *testing.T) {
	want := []bool{true, true, false, false}
	ps := Constructor(1, 1, 0)

	got := []bool{
		ps.AddCar(1),
		ps.AddCar(2),
		ps.AddCar(3),
		ps.AddCar(1),
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot :%v\n", want, got)
	}
}

func Test_000(t *testing.T) {
	want := []bool{false, false, false}
	ps := Constructor(0, 0, 0)

	got := []bool{
		ps.AddCar(1),
		ps.AddCar(2),
		ps.AddCar(3),
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot :%v\n", want, got)
	}
}

func Test_200(t *testing.T) {
	want := []bool{true, true, false}
	ps := Constructor(2, 0, 0)

	got := []bool{
		ps.AddCar(1),
		ps.AddCar(1),
		ps.AddCar(1),
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot :%v\n", want, got)
	}
}
