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
```

#### Just like in C Go starts the program whe it gets to the main function
```golang
func main() {
// work inside here
}
```

#### Structs
A struct is a piece of mempry that gets reserved to be used with data of
diff. types, in go we define a struct like this:
```golang
type Car struct {
	model string
	year int
	owner string
}
```
