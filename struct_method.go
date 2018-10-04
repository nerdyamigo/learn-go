package main
import (
	"fmt"
)

func main() {
	 goku := &Saiyan{"Goku", 9000}
	 goku.Super()
	 fmt.Printf("Goku has gone Super Saiyan! His power has increased to %d!!\n", goku.Power)
}
type Saiyan struct {
 Name string
 Power int
}

func (s *Saiyan) Super() {
 s.Power += 10000
}


