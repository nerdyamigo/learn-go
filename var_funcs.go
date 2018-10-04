package main

import (
	"fmt"
)

func main() {
	power := getPower()
	fmt.Printf("Look! It's over %d!!!\n", power);
}

func getPower() int {
	return 9001
}
