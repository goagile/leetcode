package linklist

import (
	"testing"
	// "fmt"
)

//
// ListNode Eq
//
func Test_ListNode_Eq_false_nil(t *testing.T) {
	want := false

	a := &ListNode{Val:1}
	var b *ListNode

	got := a.Eq(b)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_ListNode_Eq_false(t *testing.T) {
	want := false

	a := &ListNode{Val:1}
	b := &ListNode{Val:2}

	got := a.Eq(b)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_ListNode_Eq_true(t *testing.T) {
	want := true

	a := &ListNode{Val:1}
	b := &ListNode{Val:1}

	got := a.Eq(b)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// Eq
//
func Test_Eq_empty_true(t *testing.T) {
	want := true

	a := New()

	b := New()

	got := a.Eq(b)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_Eq_true(t *testing.T) {
	want := true

	a := New()
	a.Append(1)
	a.Append(2)
	a.Append(3)

	b := New()
	b.Append(1)
	b.Append(2)
	b.Append(3)

	got := a.Eq(b)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_Eq_false(t *testing.T) {
	want := false

	a := New()
	a.Append(1)

	b := New()
	b.Append(1)
	b.Append(2)
	b.Append(3)

	got := a.Eq(b)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_NewFromNode_Eq(t *testing.T) {
	want := true

	a := New()
	a.Append(1)
	a.Append(2)
	a.Append(3)

	b := NewFromNode(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	})

	got := a.Eq(b)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_deleteDuplicates_nil(t *testing.T) {
	want := true

	a := New()

	b := a.RemoveDuplicates()

	got := a.Eq(b)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_deleteDuplicates_1(t *testing.T) {
	want := true

	a := New()
	a.Append(1)

	b := a.RemoveDuplicates()

	got := a.Eq(b)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_deleteDuplicates_1_2(t *testing.T) {
	want := true

	a := New()
	a.Append(1)
	a.Append(2)

	b := a.RemoveDuplicates()

	got := a.Eq(b)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_deleteDuplicates_1_1_2(t *testing.T) {
	want := true

	d := New()
	d.Append(1)
	d.Append(2)

	a := New()
	a.Append(1)
	a.Append(1)
	a.Append(2)

	b := a.RemoveDuplicates()

	got := b.Eq(d)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_deleteDuplicates_1_1_2_3_3(t *testing.T) {
	want := true

	d := New()
	d.Append(1)
	d.Append(2)
	d.Append(3)

	a := New()
	a.Append(1)
	a.Append(1)
	a.Append(2)
	a.Append(3)
	a.Append(3)

	b := a.RemoveDuplicates()

	got := b.Eq(d)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_deleteDuplicates_1_1_1(t *testing.T) {
	want := true

	d := New()
	d.Append(1)

	a := New()
	a.Append(1)
	a.Append(1)
	a.Append(1)

	b := a.RemoveDuplicates()

	got := b.Eq(d)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_deleteDuplicates_1_2_2_2(t *testing.T) {
	want := true

	d := New()
	d.Append(1)
	d.Append(2)

	a := New()
	a.Append(1)
	a.Append(2)
	a.Append(2)
	a.Append(2)

	b := a.RemoveDuplicates()

	got := b.Eq(d)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// reverse
//
func Test_reverse_nil(t *testing.T) {
	want := true

	d := New()

	a := New()

	b := a.Reversed()

	got := b.Eq(d)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_reverse_1(t *testing.T) {
	want := true

	d := New()
	d.Append(1)

	a := New()
	a.Append(1)

	b := a.Reversed()

	got := b.Eq(d)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_reverse_1_2(t *testing.T) {
	want := true

	d := New()
	d.Append(2)
	d.Append(1)

	a := New()
	a.Append(1)
	a.Append(2)

	b := a.Reversed()

	got := b.Eq(d)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_reverse_1_1_1_2_2_2(t *testing.T) {
	want := true

	d := New()
	d.Append(2)
	d.Append(2)
	d.Append(2)
	d.Append(1)
	d.Append(1)
	d.Append(1)

	a := New()
	a.Append(1)
	a.Append(1)
	a.Append(1)
	a.Append(2)
	a.Append(2)
	a.Append(2)

	b := a.Reversed()

	got := b.Eq(d)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_reverse_from_reverse(t *testing.T) {
	want := NewFromElems(1, 2, 3, 4)

	a := NewFromElems(1, 2, 3, 4)

	got := NewFromNode(reverseList(reverseList(a.Head)))

	if !got.Eq(want) {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// copy
//
func Test_copy(t *testing.T) {
	want := true

	d := New()
	d.Append(1)
	d.Append(2)
	d.Append(3)

	a := New()
	a.Append(1)
	a.Append(2)
	a.Append(3)

	b := a.Copy()

	got := b.Eq(d)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// palindrome
//
func Test_IsPalindrome_1(t *testing.T) {
    want := true

	a := New()
	a.Append(1)

	got := a.IsPalindrome()

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_IsPalindrome_1_1(t *testing.T) {
    want := true

	a := New()
	a.Append(1)
	a.Append(1)

	got := a.IsPalindrome()

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_IsPalindrome_1_2(t *testing.T) {
    want := false

	a := New()
	a.Append(1)
	a.Append(2)

	got := a.IsPalindrome()

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_IsPalindrome_1_2_1(t *testing.T) {
    want := true

	a := New()
	a.Append(1)	
	a.Append(2)
	a.Append(1)

	got := a.IsPalindrome()

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_IsPalindrome_1_2_2_1(t *testing.T) {
    want := true

	a := New()
	a.Append(1)	
	a.Append(2)
	a.Append(2)
	a.Append(1)

	got := a.IsPalindrome()

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_IsPalindrome_1_2_3_2_1(t *testing.T) {
    want := true

	a := New()
	a.Append(1)	
	a.Append(2)
	a.Append(3)
	a.Append(2)
	a.Append(1)

	got := a.IsPalindrome()

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// Bin to Dec
//
func Test_Decimal_00_0(t *testing.T) {
    want := 0

	b := New()
	b.Append(0)
	b.Append(0)

	got := b.Decimal()

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_Decimal_001_1(t *testing.T) {
    want := 1

	b := New()
	b.Append(0)
	b.Append(0)
	b.Append(1)

	got := b.Decimal()

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_Decimal_10_2(t *testing.T) {
    want := 2

	b := New()
	b.Append(1)	
	b.Append(0)

	got := b.Decimal()

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_Decimal_00101_5(t *testing.T) {
    want := 5

	b := New()
	b.Append(0)
	b.Append(0)
	b.Append(1)	
	b.Append(0)
	b.Append(1)

	got := b.Decimal()

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// Get
//
func Test_Get(t *testing.T) {
    want := 1

	a := New()
	a.Append(4)
	a.Append(5)
	a.Append(1)
	a.Append(9)

	five := a.Get(5)

	got := five.Next.Val

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// delete node
//
func Test_delete_node(t *testing.T) {
    want := true

    d := New()
    d.Append(4)
    d.Append(1)
	d.Append(9)

	a := New()
	a.Append(4)
	a.Append(5)
	a.Append(1)
	a.Append(9)

	deleteNode(a.Get(5))
	a.Count--

	got := a.Eq(d)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_delete_node_2(t *testing.T) {
    want := true

    d := New()
    d.Append(4)
    d.Append(5)
	d.Append(9)

	a := New()
	a.Append(4)
	a.Append(5)
	a.Append(1)
	a.Append(9)

	deleteNode(a.Get(1))
	a.Count--

	got := a.Eq(d)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_delete_node_3(t *testing.T) {
    want := true

    d := New()
    d.Append(1)
    d.Append(2)
	d.Append(4)

	a := New()
	a.Append(1)
	a.Append(2)
	a.Append(3)
	a.Append(4)

	deleteNode(a.Get(3))
	a.Count--

	got := a.Eq(d)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_delete_node_4(t *testing.T) {
    want := true

    d := New()
    d.Append(1)

	a := New()
	a.Append(0)
	a.Append(1)

	deleteNode(a.Get(0))
	a.Count--

	got := a.Eq(d)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_delete_node_5(t *testing.T) {
    want := true

    d := New()
    d.Append(5)
    d.Append(-99)

	a := New()
	a.Append(-3)
	a.Append(5)
	a.Append(-99)

	deleteNode(a.Get(-3))
	a.Count--

	got := a.Eq(d)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// removeElements
//
func Test_removeElements(t *testing.T) {
    want := true

    d := New()
    d.Append(1)
    d.Append(2)
    d.Append(3)
    d.Append(4)
    d.Append(5)

	a := New()
	a.Append(1)
	a.Append(2)
	a.Append(6)
	a.Append(3)
	a.Append(4)
	a.Append(5)
	a.Append(6)

	b := NewFromNode(removeElements(a.Head, 6))
	b.Count = 5

	got := b.Eq(d)

	// fmt.Println("a:", a)
	// fmt.Println("b:", b)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_removeElements_2(t *testing.T) {
    want := true

	d := New()

	a := New()
	a.Append(2)
	a.Append(2)
	a.Append(2)

	b := NewFromNode(removeElements(a.Head, 2))
	// b.Count

	got := b.Eq(d)

	// fmt.Println("a:", a)
	// fmt.Println("d:", d)

	if got != want {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// middleNode
//
func Test_middleNode_nil(t *testing.T) {
	
	got := middleNode(nil)

	if nil != got {
		t.Fatalf("\ngot:%v", got)
	}
}

func Test_middleNode_head(t *testing.T) {
	want := NewFromElems(3)

	a := NewFromElems(3)

	got := NewFromNode(middleNode(a.Head))
	got.Count = 1

	if !want.Eq(got) {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_middleNode_2_elems(t *testing.T) {
	want := NewFromElems(2)

	a := NewFromElems(1,2)

	got := NewFromNode(middleNode(a.Head))
	got.Count = 1

	if !want.Eq(got) {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_middleNode_odd(t *testing.T) {
	want := NewFromElems(3,4,5)

	a := NewFromElems(1,2,3,4,5)

	got := NewFromNode(middleNode(a.Head))
	got.Count = 3

	if !want.Eq(got) {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_middleNode_even(t *testing.T) {
	want := NewFromElems(4,5,6)

	a := NewFromElems(1,2,3,4,5,6)

	got := NewFromNode(middleNode(a.Head))
	got.Count = 3

	if !want.Eq(got) {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// hasCycle
//
func Test_hasCycle_nil(t *testing.T) {
	want := false

	got := hasCycle(nil)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_hasCycle_false_1(t *testing.T) {
	want := false

	a := NewFromElems(1)

	got := hasCycle(a.Head)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}


func Test_hasCycle_false_2(t *testing.T) {
	want := false

	a := NewFromElems(3,2,0,-4)

	got := hasCycle(a.Head)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_hasCycle_true_1(t *testing.T) {
	want := true

	a := NewFromElems(3,2,0,-4)

	to := a.Get(2)
	tail := a.Get(-4)
	tail.Next = to

	got := hasCycle(a.Head)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_hasCycle_true_2(t *testing.T) {
	want := true

	a := NewFromElems(1, 2)

	to := a.Get(1)
	tail := a.Get(2)
	tail.Next = to

	got := hasCycle(a.Head)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// cmpcount
//
func Test_cmpcount_nil(t *testing.T) {
	want := 0

	a := New()
	b := New()

	got := cmpcount(a.Head, b.Head)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_cmpcount_a_eq_b(t *testing.T) {
	want := 0

	a := NewFromElems(1)
	b := NewFromElems(1)

	got := cmpcount(a.Head, b.Head)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_cmpcount_b_lt_a(t *testing.T) {
	want := 1

	a := NewFromElems(1, 2, 3, 4)
	b := NewFromElems(1, 2, 3)

	got := cmpcount(a.Head, b.Head)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_cmpcount_a_lt_b(t *testing.T) {
	want := -1

	a := NewFromElems(1, 2, 3)
	b := NewFromElems(1, 2, 3, 4)

	got := cmpcount(a.Head, b.Head)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// scroll
//
func Test_scroll_nil(t *testing.T) {
	got := scroll(nil, -1000000)

	if got != nil {
		t.Fatalf("\n is not nil got:%v", got)
	}
}

func Test_scroll_gt_count(t *testing.T) {
	a := NewFromElems(1, 2, 3, 4, 5)

	got := scroll(a.Head, 1000)

	if got != nil {
		t.Fatalf("\n is not nil got:%v", got)
	}
}

func Test_scroll_2(t *testing.T) {
	want := NewFromElems(3, 4, 5)

	a := NewFromElems(1, 2, 3, 4, 5)

	got := NewFromNode(scroll(a.Head, 2))

	if !want.Eq(got) {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_scroll_minus_2(t *testing.T) {
	want := NewFromElems(3, 4, 5)

	a := NewFromElems(1, 2, 3, 4, 5)

	got := NewFromNode(scroll(a.Head, -2))

	if !want.Eq(got) {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// getIntersectionNode
//
func Test_getIntersectionNode_Ex1_8(t *testing.T) {
	a := NewFromElems(   4, 1, 8, 4, 5)
	b := NewFromElems(5, 6, 1  )
	b.Get(1).Next = a.Get(8)

	want := a.Get(8)

	got := getIntersectionNode(a.Head, b.Head)
	if got == nil {
		t.Fatal("nil intersection")
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_getIntersectionNode_not_inx(t *testing.T) {
	a := NewFromElems(4, 2, 3)
	b := NewFromElems(5, 6, 1)

	got := getIntersectionNode(a.Head, b.Head)
	
	if got != nil {
		t.Fatal("not nil intersection")
	}
}

func Test_getIntersectionNode_common_one(t *testing.T) {
	want := &ListNode{Val:1}

	a := want
	b := want

	got := getIntersectionNode(a, b)
	if got == nil {
		t.Fatal("nil intersection")
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_getIntersectionNode_nil(t *testing.T) {
	got := getIntersectionNode(nil, nil)
	
	if got != nil {
		t.Fatal("not nil intersection")
	}
}

func Test_getIntersectionNode_1_2_3(t *testing.T) {
	l := NewFromElems(1, 2, 3)
	a := l.Get(1)
	b := l.Get(1)

	want := l.Get(1)

	got := getIntersectionNode(a, b)
	if got == nil {
		t.Fatal("nil intersection")
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_getIntersectionNode_end(t *testing.T) {
	a := &ListNode{Val:1}
	b := &ListNode{Val:1}
	want := &ListNode{Val:2}
	a.Next, b.Next = want, want

	got := getIntersectionNode(a, b)
	if got == nil {
		t.Fatal("nil intersection")
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// sortedadd
//
func Test_sortedadd_nil(t *testing.T) {
	want := &ListNode{Val:1}

	var s *ListNode

	got := sortedadd(s, 1)

	if !want.Eq(got) {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_sortedadd_1(t *testing.T) {
	want := NewFromElems(-1, 1, 2)

	s := NewFromElems(1, 2)

	got := NewFromNode(sortedadd(s.Head, -1))

	if !want.Eq(got) {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_sortedadd_minus_1(t *testing.T) {
	want := NewFromElems(1, 2, 4)

	a := &ListNode{Val:4}
	a = sortedadd(a, 1)
	a = sortedadd(a, 2)
	// a = sortedadd(a, 3)

	got := NewFromNode(a)

	if !want.Eq(got) {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_sortedadd_3(t *testing.T) {
	want := NewFromElems(1, 2, 2, 3)

	s := NewFromElems(1, 2, 3)

	got := NewFromNode(sortedadd(s.Head, 2))

	if !want.Eq(got) {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_sortedadd_4(t *testing.T) {
	want := NewFromElems(1, 2, 3, 4)

	s := NewFromElems(1, 2, 3)

	got := NewFromNode(sortedadd(s.Head, 4))

	if !want.Eq(got) {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_sortedadd_multy(t *testing.T) {
	want := NewFromElems(-1, 0, 1, 2, 3)

	a := &ListNode{Val:2}
	a = sortedadd(a, 1)
	a = sortedadd(a, 3)
	a = sortedadd(a, 0)
	a = sortedadd(a, -1)

	got := NewFromNode(a)

	if !want.Eq(got) {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// merge
//
func Test_merge_1(t *testing.T) {
	want := NewFromElems(1, 1, 2, 3, 4, 4)

	a := NewFromElems(1, 2, 4)
	b := NewFromElems(1, 3, 4)

	got := NewFromNode(merge(a.Head, b.Head))

	if !want.Eq(got) {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_merge_b_nil(t *testing.T) {
	want := NewFromElems(1, 3, 4)

	a := NewFromElems(1, 3, 4)

	got := NewFromNode(merge(a.Head, nil))

	if !want.Eq(got) {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_merge_a_nil(t *testing.T) {
	want := NewFromElems(1, 3, 4)

	b := NewFromElems(1, 3, 4)

	got := NewFromNode(merge(nil, b.Head))

	if !want.Eq(got) {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_merge_5(t *testing.T) {
	want := NewFromElems(1, 1, 2, 3, 4)

	a := NewFromElems(1, 2)
	b := NewFromElems(1, 3, 4)

	got := NewFromNode(merge(a.Head, b.Head))

	if !want.Eq(got) {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_merge_6(t *testing.T) {
	want := NewFromElems(1, 2, 3, 4)

	a := NewFromElems(4)
	b := NewFromElems(1, 2, 3)

	got := NewFromNode(merge(a.Head, b.Head))

	if !want.Eq(got) {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}
