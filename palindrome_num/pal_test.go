package main

import (
	"testing"

	"github.com/goagile/mcd/leet_code/easy/palindrome_num/fast"
	"github.com/goagile/mcd/leet_code/easy/palindrome_num/slow"
)

func Benchmark_slow_IsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slow.IsPalindrome(12345654321)
	}
}

func Benchmark_fast_IsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fast.IsPalindrome(12345654321)
	}
}
