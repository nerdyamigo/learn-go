package main
import (
	"fmt"
)


func main() {

}
func extractPowers(saiyans []*Saiyans) []int {
	powers := make([]int, 0, len(saiyans))

	for index, saiyan := range saiyans {
		powers[index] = saiyan.Power
	}
	return powers
}
