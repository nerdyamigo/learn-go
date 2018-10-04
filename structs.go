package main
import (
	"fmt"
)

func main() {
	// We declared a struct that is made up of a string(Name) & an integer(Power)
	type Saiyan struct {
		Name string 
		Power int
	}

/* 
	*This is one way of creating somthing from the struct we declared
	* this a bit of the long way to do this
	goku := Sayian {
		Name: "Goku",
		Power: 9000,
	}
*/

// Here is a shorter way, you can skip the name of the property and go based on
// order
	goku := Saiyan{"Goku", 9000}
	fmt.Printf("Here is an alien named %s & he has a power level of %d!\n", goku.Name, goku.Power)
}
