package main
import(
"fmt"
)

func main() {
  bulbasor := &Pokemon{"Bulbasor", "Grass", true}
	bulbasor.Pokedex()
}

type Pokemon struct {
	Name string
	Type string
	Evolve bool
}

func (p *Pokemon) Pokedex(){
	fmt.Printf("<-- POKEDEX -->\nMy Name Is: %s\nMy Typing Is: %s\nI Evolve: %t\n", p.Name,p.Type,p.Evolve)
}
