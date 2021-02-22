package revint

import (
	"testing"
	"strconv"
	"fmt"
	"math"
)

func reverse(a int) int {
	b := a
	z := ""
	if b < 0 {
		z = "-"
		b = -b
	}
	for x := b; x > 0; x /= 10 {
		z += fmt.Sprintf("%v", x%10)
	}
	y, _ := strconv.Atoi(z)

	if y < math.MinInt32 || y > math.MaxInt32 {
		return 0
	}

	return y
}

func Test_123(t *testing.T) {
	want := 321

	got := reverse(123)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_minus_321(t *testing.T) {
	want := -321

	got := reverse(-123)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_120(t *testing.T) {
	want := 21

	got := reverse(120)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_int32(t *testing.T) {
	want := 0

				 //9646324351
	got := reverse(1534236469)
	// fmt.Println(math.MaxInt32)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_max_int32(t *testing.T) {
	want := 0

				 //9646324351
	got := reverse(1534236469)
	// fmt.Println(math.MaxInt32)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_min_int32(t *testing.T) {
	want := 0

				 //9646324351
	// got := reverse(1534236469)

	got := reverse(-2147483648)
	fmt.Println(math.MinInt32)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

