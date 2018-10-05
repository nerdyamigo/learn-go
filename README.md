[logo]: https://golang.org/doc/gopher/talks.png
# Learning go using OpenSouce Books
I have downloaded and installed golang using MacOS 

This is the current book I am using [The Little Go
Book](https://www.openmymind.net/The-Little-Go-Book/)

Also started using the [excercism.io](https://exercism.io/tracks/go) for Golang


### Go Basics
How to declare a variable + initializing it in memory

```golang
// here we init. a variable with data, go can infer the type
name := "MyName"
// we can set the var without data and go will set it to 0 || "" | false
var str string
```

#### Just like in C Go starts the program whe it gets to the main function
```golang
func main() {
// work inside here
}
```

#### Structs
A struct is a piece of memory that gets reserved to be used with data of
diff. types, in go we define a struct like this:
```golang
type Car struct {
	model string
	year int
	owner string
}

// let's use this struct to init a new car
fordFocus := {"Focus", 2013, "John Smith"}
```
#### Pointers
A pointer is just an address in memory; it's the location of where to find the
actiual value, this is important becuse if you want to change a variables
value(not type) you have to know where it is in memory, otherwise if you try to
simple reference the variable you'll get a copy if it and not the actual
variable

Loosely you can say it's the difference between being at a house and having
directions to the house

So how do I define a pointer?
```golang
name := "John"
pointerToName := &name // here we saved a pointer into the pointerToName
variable
```

#### Methods
Go isn’t an object-oriented (OO) language like C++, Java, Ruby and C#. It doesn’t have objects nor inheritance and
thus, doesn’t have the many concepts associated with OO such as polymorphism and overloading.

What Go does have is structs which can be associated with methods. Go also
supports a simple method of composition
```golang
// let's set up a method inside a struct and use it
package main
import(
"fmt"
)

func main() {
  bulbasor := &Pokemon{"Bulbasor", "Grass", true}
	entry := Pokedex(bulbasor)
	fmt.Println(entry)
	
}

type Pokemon struct {
	Name string
	Type string
	Evolve boolean
}

func Pokedex(p *Pokemon) {
	sentence := fmt.Printf("<-- POKEDEX ENTRY -->\n Pokemon Name: %s\nPokemon
Type: %d\nPokemon Evolved: %t\n")	
}
``` 
