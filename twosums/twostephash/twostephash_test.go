package twosums

import (
	"testing"
	"reflect"
)

func Test_twoSum_empty(t *testing.T) {
	want := []int{}

	got := twoSum([]int{}, 3)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_twoSum_1_2(t *testing.T) {
	want := []int{0, 1}

	got := twoSum([]int{1, 2}, 3)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_twoSum_1_2_3(t *testing.T) {
	want := []int{0, 1}

	got := twoSum([]int{1, 2, 3}, 3)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_twoSum_3_2_1(t *testing.T) {
	want := []int{1, 2}

	got := twoSum([]int{3, 2, 1}, 3)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_twoSum_repeated(t *testing.T) {
	want := []int{2, 4}

	got := twoSum([]int{3, 3, 2, 2, 1, 1}, 3)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_twoSum_from_task(t *testing.T) {
	want := []int{0, 1}

	got := twoSum([]int{2, 7, 11, 15}, 9)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

// func Test_twoSum_zero_sums(t *testing.T) {
// 	want := []int{0, 3}

// 	got := twoSum([]int{0,4,3,0}, 0)

// 	if !reflect.DeepEqual(want, got) {
// 		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
// 	}
// }

func Test_twoSum_non_zero_sums(t *testing.T) {
	want := []int{1, 3}

	got := twoSum([]int{4,0,3,1,8}, 1)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_twoSum_less_than_zero(t *testing.T) {
	want := []int{0, 2}

	got := twoSum([]int{-3,4,3,90}, 0)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func BenchmarkTwostephash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum([]int{-3,4,3,90}, 0)
	}
}

func twoSum(nums []int, target int) []int {
    
	m := make(map[int]int, len(nums))    

	for i, n := range nums {
		if _, ok := m[n]; !ok {
			m[n] = i
		}
	}

	for i, n := range nums {
		complement := target - n
		if j, ok := m[complement]; ok && i != j {
			return []int{i, j}
		}
	}

    return []int{}
}
