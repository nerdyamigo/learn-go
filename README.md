![alt text](https://golang.org/doc/gopher/talks.png)
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
	bulbasor.Pokedex()
}
// Here the struct is created with the 3 fields
type Pokemon struct {
	Name string
	Type string
	Evolve bool
}
// here the method is created using a dereferencer operator, we'll be doing some
modifications to the data later
func (p *Pokemon) Pokedex(){ // the method is named Pokedex
	fmt.Printf("<-- POKEDEX -->\nMy Name Is: %s\nMy Typing Is: %s\nI Evolve: %t\n", p.Name,p.Type,p.Evolve)
}
``` 

#### What about constructors
Structures do not have constructors instead, you create a function that returns
an instance of the desired type

```golang
// let's set up a method inside a struct and use it
package mainpackage main
import(
"fmt"
)

func main() {
  bulbasor := &Pokemon{"Bulbasor", "Grass", true}
	bulbasor.Pokedex()

	// let's use our constructor
	charmander := NewPokemon("charmander", "fire", true)
	charmander.Pokedex()
}
// Here the struct is created with the 3 fields
type Pokemon struct {
	Name string
	TypeP string
	Evolve bool
}
// here the method is created using a dereferencer operator, we'll be doing some
func (p *Pokemon) Pokedex(){ 
	fmt.Printf("\n<-- POKEDEX -->\nMy Name Is: %s\nMy Typing Is: %s\nI Evolve: %t\n\n", p.Name,p.TypeP,p.Evolve)
}
// Let's add a constructor --> factory
func NewPokemon(name string,typeP string, evolve bool) *Pokemon {
	return &Pokemon {
		Name: name,
		TypeP: typeP,
		Evolve: evolve,
	}
}
``` 

#### Maps, Arrays & Slices
In Go arrays are fixed. Declaring an array requires we specify size, and once
the size is specified it cannot grow

```golang
/*
here we declared a specific size for our
array(remeber arrays start from 0, so this array will be from 0-150)
*/
var originalPokedex [151]int
// Let's assing something to the array, this is a single assignment
originalPokedex[0] = "Bulbasor"

// Let's also assing multiple things at once
originalPokedex := [3]string{"Bulbasor", "Ivysaur", "Venusaur"}
```
### Iterate through the array
Just like in other languages we can use the length of the array using len and
the range can be used to iterate over it
```golang
for index, value := range originalPokedex {
// Do stuff here
}
```
### Slices
A slice is a lightweight structure that wraps and represents a portion of an
array. Theere are a few ways to create a slice
```golang
/*
 * The first way to create a slice is a variation on how we created the array
 * earlier
/*

// unlike the array declaration we do not declare a length ahead of time
originalPokedex := []int{"Bulbasor", "Ivysaur", "Venusaur"}
// Let's create a slice using the make keyword
originalPokedex := make([]int, 151) // slice with the lenght and capacity of 10
// Here we use the make keyword because it allocates and initialized the slide
// for us
// you can get more specific and specify both length and capacity separetely
// like so:
originalPokedex := make([]int, 0, 151) // slice with the length of 0 but
// capacity of 151


```
Let's talk about the difference between length  and capacity
