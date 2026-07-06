package main

import (
	"fmt"
)

func slice(scores []int) {
	for i := 0; i < len(scores); i++ {
		if scores[i] == 66 {
			scores = append(scores[:i+1], append([]int{88}, scores[i+1:]...)...)
			break
		}
	}
	fmt.Println(scores)

	for i := range scores {
		fmt.Println(scores[i])
	}

	for i := 0; i < len(scores); i++ {
		if scores[i] == 66 {
			nums1 := scores[i]
			nums2 := scores[i+1]
			scores[i] = scores[i-1]
			scores[i+1] = scores[i-2]
			scores[i-1] = nums1
			scores[i-2] = nums2
		}
	}
	fmt.Println(scores)
	l := 0
	r := len(scores)
	for l < r {
		fmt.Println(scores[l])
		l++
	}
}

func main() {
	scores := []int{50, 75, 66, 20, 32, 90}
	slice(scores)
}
