package main

import (
	"fmt"
	"strings"
)

func main() {
	// mn := MakeNegatve(0)
	// fmt.Println(mn)
	// fmt.Println("Hello")
	// fmt.Println(-11 + -12 + -13 + -14 + -15)
	// a := CountPositivesSumNegatives([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15})
	// fmt.Println(a)
	fmt.Println(IsPalindrome("Abba"))
	// a := "A"
	// b := "a"
	// fmt.Println(a == b)
	// fmt.Println(strings.ToLower(a) == strings.ToLower(b))
}

func IsPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		var firstIndex = i
		var lastIndex = len(str) - i - 1

		//compare per index
		fmt.Println(string(str[firstIndex]), " : ", string(str[lastIndex]))

		if strings.ToLower(string(str[firstIndex])) != strings.ToLower(string(str[lastIndex])) {
			fmt.Println("Palindrome false")
			return false
		}
	}
	fmt.Println("Palindrome true")
	return true
}

func Maps(x []int) []int {
	var res []int
	for _, v := range x {
		sum := v * 2
		res = append(res, sum)
	}
	return res
}

func OddCount(n int) int {
	return n / 2
}

func countSheep(num int) string {
	var str string = ""
	for i := 0; i < num; i++ {
		str += fmt.Sprint(i+1) + " sheep..."
	}
	return str
}

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int

	if len(numbers) != 0 {
		pos := 0
		neg := 0
		for _, n := range numbers {
			if n > 0 {
				pos++
			} else {
				neg += n
			}
		}
		res = []int{pos, neg}
	}
	return res // your code here
}

func MakeNegatve(x int) int {
	var res int
	if x > 0 {
		res = -x
	}
	if x < 0 {
		res = x
	}
	return res
}

func NoSpace(word string) string {
	return strings.ReplaceAll(word, " ", "")
}

func PositiveSum(numbers []int) (sum int) {
	for i := 0; i < len(numbers); i++{
		if numbers[i] > 0{
			sum += numbers[i]
		}
	}
	return
}