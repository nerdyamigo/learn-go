package main
import (
	"fmt"
	"runtime"
)

func main() {
	slicing()
	slicing2()

}

func slicing() {
	// here we use make to allocate memory for the underlying array and also
	// init the slice, using make we can do both
	originalPokedex := make([]string, 0, 150) // length, capacity
	// when we want to set a specific item in an index let's re-slice our slice
	originalPokedex = originalPokedex[0:150]
	originalPokedex[141] = "Mew"
	originalPokedex[149] = "Charizard"

	fmt.Println(originalPokedex)

}

func slicing2() {
	originalPokedex := make([]int, 0, 5) // length, capacity
	c := cap(originalPokedex)
	fmt.Println(c)

	for i := 0; i <25; i++ {
		originalPokedex = append(originalPokedex, i)

		// if our capacity has changed
		// Go had to grow to accomodate for the new data
		if cap(originalPokedex) != c {
			c = cap(originalPokedex)
			fmt.Println(c)
		}
	}

}


