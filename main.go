package main

import (
	"fmt"
)

func insertSlice(scores []int) []int {
	for i := 0; i < len(scores); i++ {
		if scores[i] == 66 {
			scores = append(scores[:i+1], append([]int{88}, scores[i+1:]...)...)
			break
		}
	}

	fmt.Print("\n\n----PRINT SCORE AFTER INSERT SLICE----\n")

	for i := range scores {
		fmt.Println(scores[i])
	}

	fmt.Println("-------- SELESAI --------\n")

	return scores
}

func acakSlice(scores []int, index int) []int {
	fmt.Println("scores : ", scores)
	left := []int{}
	right := []int{}
	results := []int{}

	i := 0
	j := len(scores)
	for i < j {
		if scores[i] == 66 {
			left = scores[0:i]
			right = scores[i+2 : j]
			i = j
		}
		i++
	}

	insert := []int{88, 66}
	if index < len(left) {
		left = append(left[:index], append(insert, left[index:]...)...)
	}

	if index < len(left)+2 {
		left = append(left, insert...)
	}

	if index > len(left)+2 {
		index = index - (len(left) + 1)
		right = append(right[:index], append(insert, right[index:]...)...)
	}

	// if index > len(scores) {
	// 	results = append(left, right...)
	// 	results = append(results, insert...)
	// 	return results
	// }

	results = append(left, right...)

	return results
}

func main() {
	scores := []int{50, 75, 66, 20, 32, 90}
	insert := insertSlice(scores)
	acak := acakSlice(insert, 5)
	fmt.Println("After Acak  : ", acak)
}
