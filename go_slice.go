package main
import (
	"fmt"
)

func main() {
	// here we use make to allocate memory for the underlying array and also
	// init the slice, using make we can do both
	originalPokedex := make([]string, 0, 150) // length, capacity
	// when we want to set a specific item in an index let's re-slice our slice
	originalPokedex = originalPokedex[0:150]
	originalPokedex[141] = "Mew"
	originalPokedex[149] = "Charizard"
	fmt.Println(originalPokedex)
}

