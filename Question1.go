package main

import (
	"fmt"
	"sort"
)

// O(N^2) time and O(1) space
func TwoNumberSum(array []int, target int) []int {
	for i := 0; i < len(array); i++ {
		firstNum := array[i]
		for j := i + 1; j < len(array); j++ {
			secondNum := array[j]
			if firstNum+secondNum == target {
				return []int{firstNum, secondNum}
			}
		}
	}
	// Write your code here.
	return []int{}
}

//O(n) time and O(N) Space    (Optimal Solution)
func TwoNumberSum2(array []int, target int) []int {
	m := make(map[int]bool)
	for i := 0; i < len(array); i++ {
		potentialMatch := target - array[i]
		if _, found := m[potentialMatch]; found {
			return []int{potentialMatch, array[i]}
		} else {
			m[array[i]] = true
		}
	}
	return []int{}
}

// O(NlogN) time and O(1) space complexity
func TwoNumberSum3(array []int, target int) []int {
	sort.Ints(array)
	left, right := 0, len(array)-1
	for left < right {
		currentSum := array[left] + array[right]
		if currentSum == target {
			return []int{array[left], array[right]}
		} else if currentSum < target {
			left += 1
		} else {
			right -= 1
		}
	}
	return []int{}
}

func main() {
	var targ int
	fmt.Scanf("%d", &targ)
	check := []int{10, -10, 2, 34, 55, 56, 44, 23, -60}
	fmt.Println(TwoNumberSum3(check, targ))

}
