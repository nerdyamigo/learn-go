![alt text](https://golang.org/doc/gopher/talks.png)
# Learning go using OpenSouce Books
I have downloaded and installed golang using MacOS 

This is the current book I am using [The Little Go
Book](https://www.openmymind.net/The-Little-Go-Book/)
Will also use [The Go Bootcamp](http://www.golangbootcamp.com/book?ref=cybrhome)
Also started using the [excercism.io](https://exercism.io/tracks/go) for Golang


### Go Basics
### Packages
Every Go program is made up of `packages`.
Programs start running in the package `main`.
By convention, the package name is the same as the last element of the import path. For instance, the `"math/rand"` package comprises files that begin
with the statement `package rand`.

### Imports
This code groups the imports into a parenthisized "factored" import statement. 
You can also write multiple import statements like: 
`import "fmt"`
`import "math"`

### Exported Names
In Go a name is exported ig it begins with a `capital letter` using `lowercase letteer` will make it be ony usable inside the package but not outside
of it

### Primitive types in Golang
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
### Types

A `type` determines a set of values together with operations and methods specific to those values. A type may be denoted by a `type` name, if it has
one or specified  by a `type literal`, which composes a type from existing types
The language predeclares certain type names. Composite types - array,struct,pointer,function,interface,slice,map,and channel types-- may be
constructed using `type literals`

Each `type T` has an underlying type: If T is one of the predeclared boolean, numeric, or string types, or a type literal, the corresponding
underlying type is T itself. 
#### Boolean Types
A `boolean` type represents the set of `Boolean` truth values denoted by the predeclared constants `true` and `false`. The predeclared `boolean` type is `bool`
#### Numeric Types
A `numeric` type represents sets of `integer` or `floating-point` values. The predeclared architecture-independent numer types are:

| Numeric Types | Description |
|---|---|
| unit8 | the set of all unsigned 8bit integers (0-255) |
| unit16 | the set of all the unsigned 16bit integers (0-65535) |
| unit32 | the set of all the unsigned 32bit integers (0-4294967295) |
| unit64 | the set of all the unsignes 64bit integers (0 -18446744073709551615) |
| int8 | the set of all the signed 8bit integers (-128-127) |
| int16 | the set of all the signed 16bit integers (-32768 to 32767) | 
| int32 | the set of all the signed 32bit integers (-2147483648 to 2147483647) |
| int64 | the set of all the signed 64 bit integers (-9223372036854775808 to 9223372036854775807) |
| float32 | the set of all IEEE-754 32-bit floating-point numbers | 
| float64 | the set of all IEEE-754 64-bit floating-point numbers | 
| complex64 | the set of all complex numbers with float32 real and imaginary parts |
| complex128 | the set of all complex numbers with float64 real and imaginary parts | 
| byte | alias for uint8 | 
|rune | alias for int32 | 

The value of an `n-integer` is `n` bits wide and represented using `two's complement arithmetic`
There is also a set of predeclared numeric types with implementation-specific sizes

#### String Types
A `string` type represents the set of string values. A string value is a sequence of bytes. Strings are immutable, `once created it is impossible to
change the contents of a string`.
The length of a string `s` can be discovered using the built-in function `len`. The length is
a compile-time constant if the string is a constant. A string's bytes can be accessed by integer `indices` 0 through `len(s)-1`. It is illegal to take
the address of such element; if `s[i]` is the i'th byte of a string, `&s[i]` is invalid.

#### Array Types
An `array` is a numbered sequence of elements of a single type, called the element type. The number of elements is called the length and is never
negative

The length is part of the `array` type; it must evaluate to a non-negative constant representable by a value of type `int`. The length of an array can
be discovered using the built-in function `len`. The elements can be addressed by ineteger indices 0 through `len(a)-1`. Array types are always
one-dimensional but may be composed to form multi-dimensional types

#### Slice Types
A `slice` is a descriptor for a contigous segment of an `underlying array` and provides access to a numbered sequence of elements from that array. A
`slice` type denotes the set of all slices of arrays of its element type. The value of an unitialized `slice` is `nil`
```golang
SliceType = "[" "]" ElementType
```
Like arrays slices are indexable and have a length. The length of `slice s` can be discovered by the built-in `len` function; unlike arrays it may
change during execution. The elemsts can be addressed by integer indices 0 through `len(s)-1`. The slice index of a given element may be less than the
index of the same element in the underlying array 

A slice once initialized is always associated with an underlying array that holds its elements. A slice therefore shares storage with its array and
with other slices of the same array; by contrast, distinct arrays always represent distinct storage.

The array underlying a slice may extend past tghe end of the slice. The `capacity` is a mersure of that extent: it is the sum of the length of the
slice and the length of the the array beyond the slice; a slice of length up to that capacity can be created by `slicing` a new one from the original
slice. The `capacity` of a slice can be discovered using the built-in funtion `cap(a)`.

A new, initialized slice value for a given element type T is made using the built-in function `make`, which takes a slice type and parameters
speciifying the length and optionally the cap. A slide created with the `make` keyword always allocates a new, hidden array to which the returned
slice value refers
```golang
make([]int, 50, 100)
new([100]int)[0:50]
```
Like arrays, slices are always one-dimensional but may be composed to construct higher-dimensional objects. 

#### Struct Types 
A `struct` is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified
explicitly(IdentifierList) or implicitly(EmbeddedField). Within a struct non-blank field names must be unique
```golang
// empty struct
struct {}

// a struct with 6 fields
struct {
	x,y int
	u float32
	_ float32 // padding
	A *[]int
	F func()
}
```
A field declared with a type but not explicit fieldname is called an `embedded field`. An embedded field must be specified as a type T or as a pointer
to a non-interface type name *T and T itself may not be a pointer type. The unqualified type name acts as a field name
```golang
// struct with 4 embedded fields of types: T1, *T2, P.T3 and *P.T4
struct {
	T1		// fieldname is T1
	*T2		// fieldname is *T2
	P.T3	// fieldname is T3
	*P.T4	// fieldname is T4
	x,y int	// fieldnames are x and y
}
```
The following declaration is illegal becase field names must be unique in a struct type
```golang
struct {
	T			// conflicts with embedded field *T and *P.T
	*T		// conflicts with embedded field *T and *P.T
	*P.T	// conflicts with embedded field T and *T
}
```
A field or method `f` of an embedded field in a struct `x` is called `promoted` if `x.f` is a legal selector that denotes that field of method`f`

##### Promoted Fields
Promoted fields act like ordinary fields of a struct except that they cannot be used as fieldnames in composite literals of the struct

Given a struct type S and a defined type T, promoted methods are included in the method set of the struct as follows
	* If `S` contains an embedded field T, the method sets of `S` and `*S` both include promoted methods with receiver T. The method of set `*S` also
	  includes promoted methods with receiver *T.
	* If `S` contains an embedded method field `*T`, the method sets of `S` and `*S` both include promoted methods with receiver `T` or `*T`

A field declaration may be followed by an optional string literla tag, which becomes an attribute for all the fields in the corresponding field
declaration. An empty tag string is equivalent to an absent tag. The tags are made visible trough a `reflection interface` and take part in `type
identity` for structs but are otherwise ignored

#### Pointer Types
A pointer denotes the set of all pointers to variables of a given type, called the `base type` of the pointer. The value of uninitialized pointter is
nil
```golang
PointerType = "*"BaseType .
BaseType = Type .

*Point
*[4]int
```

### Function Types
A function type denotes the set of all functions with the same parameter and result types. The value of an unitialize varibale of function type is nil

### Interface Types
An interface type specifies a method set called its interfaces. A variable of interface type can store a value of any type with a method set that is
any superset of the interface. Such a type is said to `implement the interface`. The value of an unitialized variable of interface type is nil

As with all method sets, in an interface type, each method must have a unique non-blank name.
More than one type may implement an interface. For instance if two types S1 and S2 have the method set
```golang
func (p T) Read(b Buffer) bool {return ...}
func (p T) Write(b Buffer) bool {return ...}
func (p T) Close() {...}
// Where T stands for either S1 || S2 then the file interface is implemented by both S1 and S2, regardless of what other methods S2 and S2 may have or
share
```
A type implements any interface comprising any subset of its methods and may therefore impement distinct interfaces. For isntance all types
implement the empty interface
```golang
interface{}
```
An interface T may use a interface type name E in place of a method specification. This is called embedding interface E in T; it adds all methods of E
to the interface T
```golang
type ReadWrite interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
}

type File interface {
	ReadWrite	// same as adding the methods from ReadWrite
	Locker		// same as adding the methods from Locker
	Close()
}

type LockedFile interface {
	Locker
	File	// illegal: Lock, Unlock Unlock not unique
	Lock()// illegal: Lock not unique
}
```
An interface type T may not embed itself or any interface type that embeds T recursively

### Map Types
A `map` is an unordered group of elements of one type, called the element type, indexed by a set of unique keys of another type, called the key type.
The value of an unitialized map is `nil`

The comparison ops `==` & `!=` must be fully defined for operands of the key type; thus the key type must not be a function, map, or slice. If the
key type is an interface type, these comparison operators must be defined for the dynamic key values; failure will cause a run-time panic
```golang
map[string]int
map[*T]struct{x,y float64}
map[string]interface{}
```
The number of map elements is called its length. For map `m` it can be discovered using the built-in function `len` and may change during execution.
Elements may be added during execution using assignments and retrieved with index expressions; they may be removed using the `delete` built-in
function

A new empty map value is made using the built-in `make`, which takes the map type and an optional capacity hint as args
```golang
make(map[string]int)
make(map[string]int, 100)
```
The initial cap does not bound its size; maps grow to accomodate the number of items stored in them, with the exception of nil maps. A nil map is
equivalent to an empty map except that no elemets may be added

### Channel Type
A `channel` type provides a mechanism for concurrently executing functions to communicate by sending and receiving values of a specified element type.
The value of an unitilized channel is nil
```golang
ChannelType = ( "chan" | "chan" "<-" | "<-" "chan" ) ElementType
```
The optional `<-` operator specifies the channel `direction`, `send` or `receive`. If no direction, the channel is `bidirectional`. A channel may be
considered contrained only to send or only to receive by conversion or aassignment
```golang
chan T							// can be use to send and receive values of type T
chan<- float64			// can only be used to send float64s
<-chan int					// can only be used to receive ints

// the `<-` operator associatres with the leftmost chan possible
chan<- chan int			// same as chan<- (chan int)
chan <- <-chan int	// same as chan<- (<-chan int)
<-chan <-chan int 	// same as <-chan (<-chan int)
```
A new initialized channel value can be made using the built-in function `make`, which takes the channel type and an optional cap as arguments
```golang
make(chan int, 100)
```
The capacity, in number of elements, sets the size of the buffer in the channel. If the cap is zero or absent the channel is unbuffered and
communication succeeds only when both sender and receiver are ready. Otherwise, the channel is buffered and communication succeeds without blocking if
the buffer is not full or empty. A nil channel is never ready for communication.

A channel may be closed with the built-in function `close`. The multi-valued assignment form of the receive operator reports wether a received value
was sent before the channel was closed. 

A single channel may be used in send statements, receive operations, and calls to the built-in functions `cap` and `len` by any number of goroutines
without further synchronization. Channels act as first-in-first-out queues. For example if one goroutine sends values on a channel and a second
goroutine receives them, the values are received in the order sent.

### Type Conversions
The expression `T(v)` convertes the value `v` to type `T`
```golang
// Some numeric conversion
var i int = 42
var f float64 = float64(i);
var u uint = uint(f)
``` 
Unlike in C, in Go assignment between different items of different type requires an explicit converision. 

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
##### Composition
Go supports composition whichs is the act of including one structure into another
```golang
// lets create a mixin using composition
type Animal struct {
	Name string
}
func (p *Animal) Screech() {
	fmt.Printf("SRRRRRRRKTTTT, %s\n", p.Animal)
}

type Pokemon struct {
	*Animal
	HP int
}
// let's use our struct:
charizard := &Pokemon {
	Animal: &Animal("Charizard"),
	HP: 400,
}
charizard.Screech()
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
// here we know what is going into the slice
name := []string{"pikachu", "bulbasor", "charmeleon"}
// here we create a slice with the length & capacity ot 10
checks := make([]bool, 10)
// this is a slice that has not been initialize nor does it have a specific size || cap
var name []string
// here we declare a slice with a specific length and capacity
hp := make([]int, 0, 99)

// Example when writting into specific indexes of a slice
// here we loop all the pokemons and extract only the data we need
func extractNames (pokemons []*Pokemons) []string {
	names := make ([]string, len(pokemons))
	for index, pokemon := range pokemons {
		names[index] = pokemon.Name
	}
	return names
}
// we can use a nil slice when the number of elemensts is unkown
// let's refactor our last function to use it with append
func extractNames (pokemons, []*Pokemons) []string {
	names := make([]string, 0, len(pokemons))
	for _, pokemon := range pokemons {
		names = append(names, pokemon.Names)
	}
	return names
}
```
##### COPY
Normally,  method that copies values from one array to another has 5 params, `source`, `sourceStart`, `count`, `destination`, `destinationStart`. With
slice we only need two
```golang
import ( 
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	scores = make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)
	worsr := make([]int, 5)
	copy(worst, scores[:5]
	fmt.Println(worst)
}
```
### MAPS
Maps in Go are what other languages call hashtables or dictionaries, you define a key and a value, and can get, set and delete values from it
Maps like slices are created using the `make` function
```golang
fun main() {
	lookup := make(map[string][int])
	lookup["charizard"] = 400
	hp, exists := lookup["pikachu"]
	
	// prints 0, false
	// 0 is the default value for an integer
	fmt.Println(power, exists)
	// to get the number of keys we can use len
	total := len(lookup) // returns 1
	// delete has no return, can be called on a non-existent key
	delete(lookup, "goku")
}
```
Maps grow dynamically. However we can supply a second argument to make to set an initial size
```golang
	lookup := make(map[string][int], 100)
```
If you have somne idea of how many keys your map will have, defining an initial size can help with performance.
When you need a map as a field of structure you define as:
```golang
type Pokemon struct {
	Name string
	PartyPokemon map[string]*Pokemon
}
// lets initialize the above struct
charizard := &Pokmeon {
	Name: "Charizard",
	PartyPokemon: make(map[string]*Pokemon),
}

charizard.PartyPokemon["bulbasour"] = ... // @TODO create or load pokemon

// here's another way to create a map, like `make` this approach is specific to maps and arrays. We can declare as a composite literal:

lookup := map[string]int {
	"charizard": 400,
	"bulasour": 200,
}
// we can then iterate over a map using a `for loop` combined with the `range` keyword
for key, value := range lookup {
	// CODE GOES HERE
}
// this iteration is not in order it will return the value key pair in random order 
```
### Packages
In Go package names follow the direcotry structure of your Go workspace. If we were bullding a shopping system, we'd prpbably start with a package
named `shopping` and put our source files in `$GOPATH/src/shopping/`

We don't want to put everything inside this folder though, for example maybe we want to isolate some database logic inside its own folder. To achieve
this we create a subfolder at `$GOPATH/src/shopping/db`. The package name of the files within this subfolder is simply `db`, but to access it from
another package, including the `shopping` package we need to `import` `shopping/db`

When we name our package using the `package` keyword we provide a single value, when we import we specofy the complete path

### Visibility
Go uses a simple rule to define what types and functions are visible outside of a package. If the name of the type or function starts with an
uppercase letter, it's visible. If it starts with a lowercase letter, it isn't. 
This also applies to struc fields, it a struct field name starts with a lowercase, only code inside the same package will be able to access them.
### Interfaces
Interfaces define a contgract but no implementation
```golang
type Logger interface {
	Log(message string)
}
```
#### What purpose could interfaces have?
Interfaces help decouple your code from specific implementations
```golang
type SqlLogger struct {...}
type ConsoleLogger struct {...}
type FileLogger struct {...}
```

`@TODO LOOK UP HOW TO IMPLEMENT AND HOW TO USE INTERFACES IN GO`

How would you use an `interface`?
```golang
// it could be a structs field
type Server struct {
	logger Logger
}
// or a function param
func process(logger Logger) {
	logger.Log("hello!")
}
```
### Error Handling
Go's preferred way to deal with errors is through return values, not exceptions. 
```golang
// let's consider strconv.Atoi
package main
import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("not a valid number")
	} else {
		fmt.Println(n)
	}
}
// create your own error type; the only requirement is that it fulfills the built-in `error` interface
type error interface {
	Error() string
}
// more commonly we can create our own errors by importing the errors package and using it in the New function

import (
	"errors"
)

func process(count int) error {
	if count < 1 {
		return errors.New("Invalid Count")
	}
	// some code here
	return nil
}
```
### Defer
Even though Go has a garbage collector, some resources required that we explicitly release them. For example we need to `Close()` files after we're
done with them. This sort of code is always dangerous. For one thing, as we're writting a function, it's easy to forget to Close something we declared
10 lines up, Go's solutions for this is the `defer` keyword
```golang
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("a_file_to_read")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// read file here
}

```

### Empty Interface & Conversions
Go has an mepty interface with no methods `interface{}`. Since every type implements all 0 of the empty interface's methods, and since interfaces are
implicitly implemented, every type fulfills the contract of the emtpy interface
```golang
// if we wanted to we could write an add function with the following signature
func add(a interface{}, b interface{}, c interface{}) {
	// code here
}
// To convert an interface variable to an explicit type user `.`
return a.(int) + b.(int)
```

### Strings & Byte Arrays
Strings and byte arrays are closely related. We can convert one to another
```golang
stra := "the spice must flow"
byts := []bytes(stra)
strb := string(byts)
```
In fact, this way of converting is common across various types as well. Some functions explicitly expect and int32 or int64 or their unsigned
counterparts. You might find yourself having to do something like this:
```golang
int64(count)
```
### Function Type
Functions are first-class types: 
```golang
type Add func(a int, b int) int
// which can be used anywhere - as a field type, param or a return val

package main
import (
	"fmt"
)

func main() {
	fmt.Println(process(func(a int, b int) int {
		return a + b
	}))
}

func process(adder Add) int {
	return adder(1,2)
}
// Using functions like this can help decouple code from specific implementations much like we achieve with interfaces
```

# Concurrency
### Goroutines

A goroutine is similar to a thread, but it is scheduled by `Go` not the OS. Code that runs in a goroutine can run in concurrency with other code

```golang
package main

import(
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	go process()
	time.Sleep(time.Millisecond * 10) // this is bad do not do this
	fmt.Println("done"
}

func process() {
	fmt.Println("processing")
}
/* There are a few interesting things here but the most important is how we start a goroutine. We simply use the go keyword followed by the function
we want to execute. If we just want to run a bit of code, such as the above, we can use an anonymous function. Do not that anon funcs are not only
used with goroutines however.
go func() {
	fmt.Println("processing")	
}()
Goroutines are easy to create and have little overhead. Multiple goroutines end up running on the same underlying OS thread. This is often called M:N
threading model because we have M applications threads(goroutines) running on N OS threads. 
/*
```
### Synchronization
Creating goroutines is trivial, and they are so cheap that we can start many; however concurrent code needs to be coordinated. To help with this
problem, Go provides `channels`.
##### Concurrent Programming
Writing concurrent code requires that you pay specific attention to where and how you read and write values. In some ways, it's like programming
without a garbage collector - it requires that you think about your data from a new angle, always watchful for possible danger, condifer:
```golang
package main

import (
	"fmt"
	"time"
)

var counter = 0

func main() {
	for i := 0; i < 20; i++ {
		go incr()
	}
	time.Sleep(time.Millisecond * 10) 
}

func incr() {
	counter++
	fmt.Println(counter)
}
```
What do you think will happen above?
The behavior is actually undefined, because we potentially have multiple(two in this case) goroutines writting to the same variable, `counter` at the
same time. Ot just as bad one goroutine would be reading `counter` while another writes to it

Is it really that bad?
Yes, `counter++` might seem like a simple line of code, but it actually gets broken down into multiple assembly statements - the exact number depends
on the platform you are running it. If you run this example you'll see very often the numbers are printed in weird order and/or numbers are
duplicated/missings. 
### Channels
The challenge with concurrent programming stems from sharing data. If your goroutines share no data you needn't worry about synchronizing them. 
Channels help make concurrent programming saner by taking shared data out of the picture. A channel is a communication pipe between goroutines which
is used to passed data. In othwr words a goroutine that has data can pass it to anothwr goroutine via a channel. The result is that at any point in
time only one goroutine has access to the data

A channel like everything else has a typ. This is the type of data that we'll be passing though our channel. 
```golang
// create a channel which can be used to pass an integer around
c := make(chan int)
// The type of this channel is chan int. Therefore, to pass this chanel to a function, our signature looks like 
func worker(c chan int) {
	// code goes here 
}
// Channels support two ops: receiving and sending. We send to a channel by doing:
CHANNEL <- DATA
// and we receive from one by doing
VAR := <-CHANNEL
// The arrow points in the direction the data flows. When sending, the data flows into the channel
// When receiving the data flows out of the channel
```
One more thing to know is that receiving and sending from a channel is blocking. That is when we receive from a channel, execution of the goroutine
won't continue until the data is received
```golang
type Worker struct {
	id int
}

func (w Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
	}
}

c := make(chan int)
for i := 0; i < 4; i++ {
	worker := Worker{id: i}
	go worker.process(c)
}
// give them some work
for {
	c <- rand.Int()
	time.Sleep(time.Millisecond * 50)
}
```
We don't know which worker is going to get what data. What we do know, what Go gurantees is that the data we send to a channel will only be received
by s single receiver.

#### Buffered Channels
What happends if we have more data coming in that we can handle?

In cases where you need high gurantees that the data is being processed, you probably will want to start blocking the client. In other cases,  you
might be willing to loosen those gurantees. 
The first way to do this is to buffer the data. If no worker is available, we want to temporarily store the data in some sort of queue. Channels have
this buffering capability built in. Wehen we created our channel with make, we can give our channel a length: 
```golang
c := make(chan int, 100)
```
However buffer channels do not add more capacity; they merely provide a queue for pending work and a good way to deal with a sudden spike. 
#### Select
Evem with buffering, there comes a popint where we need to start dropping messages. For this we can use `select`
Syntastically `select` looks like a switch. With it we can provide code for when the channel is not available to send to
