package main

import "testing"

func TestMain(m *testing.T) {
	noSpace := NoSpace("Hello World")
	if noSpace != "HelloWorld" {
		m.Errorf("Expected 'HelloWorld', got '%s'", noSpace)
	}

	positiveSum := []struct{
		numbers []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{1, -2, 3, 4, 5}, 13},
		{[]int{}, 0},
		{[]int{-1, -2, -3, -4, -5}, 0},
		{[]int{-1, 2, 3, 4, -5}, 9},
	}
	for _, v := range positiveSum {
		got := PositiveSum(v.numbers)
		if got != v.want {
			m.Errorf("Expected %d, got %d", v.want, got)
		}
	}
}
