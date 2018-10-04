// I want to test out the value for pointers

package main
import (
	"fmt"
)

func main() {
	name := "Carlos"
	pntr := &name
	backToStr := *pntr
	fmt.Println("Here is the pointer to the var: ",pntr," Here is the text from the pointer: ", backToStr);
}
