package main
import (
	"fmt"
)

func main() {
	// Instatiate the new class
	goku := NewSaiyan("goku", 9000)
	fmt.Printf("Here comes %s\nHis current power level is %d\n", goku.Name, goku.Power)
	// Use our method inside the struct
	goku.Super()
	fmt.Printf("Oh no! %s has powered up! His current level is %d\n", goku.Name, goku.Power)
}
// Here we make the struct and pre define the values 
type Saiyan struct {
	Name string
	Power int
}
// method called super inside the Saiyan struct
func (s *Saiyan) Super() {
	s.Power += 10000
}
/*
 *Structures do not have constructors. Instead you create a function that
 * returns an instance of the desired type(like a factory)
*/

func NewSaiyan (name string, power int) *Saiyan {
	return &Saiyan{
		Name: name,
		Power: power,
	}
}


