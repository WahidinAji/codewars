package main

import "testing"

func TestMain(m *testing.T) {
	noSpace := NoSpace("Hello World")
	if noSpace != "HelloWorld" {
		m.Errorf("Expected 'HelloWorld', got '%s'", noSpace)
	}

	positiveSum := []struct {
		numbers []int
		want    int
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

	repeatStr := []struct {
		repeat       int
		str, wantStr string
	}{
		{0, "", ""},
		{4, "a", "aaaa"},
		{3, "hello ", "hello hello hello "},
		{2, "abc", "abcabc"},
	}
	for _, v := range repeatStr {
		got := RepeatStr(v.repeat, v.str)
		if got != v.wantStr {
			m.Errorf("Expected %s, got %s", v.wantStr, got)
		}
	}
}
