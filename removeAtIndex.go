package main

import(
	"fmt"
)

func main() {
	scores := []int{1,2,3,4}
	scores = removeAtIndex(scores, 2)
	fmt.Println(scores)
}

func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1
	// swap the last value and the value we want to remove
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}
