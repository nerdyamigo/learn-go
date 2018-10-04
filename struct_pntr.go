package main
import (
	"fmt"
)

func main() {
	goku := &Saiyan{"Goku", 9000}

	Super(goku)
	fmt.Printf("Here is an alien named %s & he has a power level of %d!\n", goku.Name, goku.Power)
}

type Saiyan struct {
	Name string 
	Power int
}

func Super(s *Saiyan) {
	s = &Saiyan{"Gohan", 1000}
}
