![alt text](https://golang.org/doc/gopher/talks.png)
# Learning go using OpenSouce Books
I have downloaded and installed golang using MacOS 

This is the current book I am using [The Little Go
Book](https://www.openmymind.net/The-Little-Go-Book/)
Will also use [The Go Bootcamp](http://www.golangbootcamp.com/book?ref=cybrhome)
Also started using the [excercism.io](https://exercism.io/tracks/go) for Golang


### Go Basics
Primitive types in Golang
```golang
// string, numeric type, bool, error
```
How to declare a variable + initializing it in memory

```golang
// here we init. a variable with data, go can infer the type
name := "MyName"
// we can set the var without data and go will set it to 0 || "" | false
var str string
```
### Function Declarations
Functions can return multiple values, let's look at a few
```golang
// no return value
func log(message string) {
}
// one return value, number 
func add(a int, b int) int {
	// this could have been written a bit different since params use the same type (a, b int)
}
// 2 return values number and boolean
func power(name string) (int, bool) {
}
```
Sometimes you only care about one of the return values. In these cases, you assign the other values to `_`
```golang
_, exists := evolves("charizard")
if exists == false {
	// handle error here
}
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
type Pokemon struct {
	name string
	HP int
	Evolves bool
}

// let's use this struct to init a new car
Charizard  := {"Charizard", 400, true}
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
// here we create the Pokedex method on the Pokemon struct using the * 
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
###### Other ways of using "constructors" in go
Like stated above Golng does not have constructors but it does posses a `new` keyword
The result of 
```golang
charizard := new(Pokemon)
// is the same as
chariard := &Pokemon{}
// which one to use? Most people use the latter when initializing data
charizard := new(Pokemon)
charizard.name = "Charizard"
charizard.evolved = true
// vs
charizard := &Pokemon {
	name: "Charizard",
	evolved: true
}
```

#### Fields of a Struct
We can put anything inside a structure including a pointer to another struct or an array, functions or interfaces
```golang
// lets start with our pokemon struct but we will expand it this time
type Pokemon struct {
	Name string
	HP int
	Evolves bool
	Pre-Evolution *Pokemon
}
// and now we initialize this
charizard := &Pokemon{
	Name: "Charizard",
	HP: 500,
	Evolves: false,
	Pre-Evolution: &Pokemon {
		Name: "Charmeleon",
		HP: 300,
		Evolves: true,
		Pre-Evolution: nil // We all know this should be another pokemon lol 
	},
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
*/

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
### Four ways to initialize a s slice
```golang
name := []string{"pikachu", "bulbasor", "charmeleon"}
checks := make([]bool, 10)
var name []string
hp := make([]int, 0, 99)
```
