// I want to test out the value for pointers

package main
import (
	"fmt"
)

func main() {
	name := "Carlos"
	pntr := &name
	fmt.Println(pntr);
}
