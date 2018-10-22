package main

import (
	"fmt"
)

func main() {
	/*
		long way of declaring a variable
	 	var power int
		power = 9000
	*/

	/*
		short way
		this method of declaring variables also infers the type of data
	*/
	power := 9000
	fmt.Printf("It's over %d\n", power)
}
