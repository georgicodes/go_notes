# GO

## Package Declaration
Two types of Go programs: executables and libraries.

Executable apps can be run from the command line. Libraries are collecitons of code pacakced toegether for reuse in other programs.

## Types
Go is statically typed which means variables always have a specific type that cannot change.

## Variables
Use lower camel case for variable names.

```go
var personName string = "georgi codes"

var y string
y = "pizza"

// no type is actually required
var x = "Hello World"

// short hand syntax for variable assignment
// best practice
x := "georgi codes a lot"
```
There is no type because the Go compiler infers the type based on the literal value you assign the variable.

It's best practice to use this shorter version.

Another shortand for when defining multiple variables

```go
var (
	a=5
	b=10
	c=15
)
const (
	name="mary"
	age=20
)
```

### Constants
Constants are variables whose values cannont be changed later.

```go
const x string = "hello world"
x = "another string" // compile-time error
```

## Scope
Go is lexically scoped using blocks.

## Control Structures

```go
// for loop
for i := 1; i <= 10; i++ {
  fmt.Println(i)
}

// if statments
if i % 2 == 0 {
	fmt.Println("2")
} else if i % 4 == 0 {
	fmt.Println("4")
} else {
	fmt.Println("none")
}

// switch statements
switch i {
case 0: fmt.Println("Zero")
case 1: fmt.Println("One")
default: fmt.Println("Unknown Number")
}
```

## Arrays
A numbered sequence of elements of a single type with a fixed length.

```go
var x [5]int
x[0] // will be 

y := [5]float64{ 98, 93, 77, 82, 83 } // shorthand syntax
```

Looping over an array can be made easy with a modified loop syntax where `index` is the current position and `value` is the same as `x[index]`.

```go
var total float64 = 0
for index, value := range x {
     total += value
}
fmt.Println(total / float64(len(x)))
```
The Go compiler won't allow you to create variables that you never use. Since we don't use `index` - ann error will be thrown and we can change this code to be:

```go
var total float64 = 0
for _, value := range x {
     total += value
}
fmt.Println(total / float64(len(x)))
```
A single underscore tells the compiler that we don't need this.

## Slices
A slice is a segment of an array. Like arrays, slices are indexable and have a length however unlike arrays, this length is allowed to change.

```go
// notice missing length between brackets
// this creates a slice with length 0
var x []float64 

// create a slice with a length of 5
y := make([]float64, 5)
```

Unerneath slices are always associated with some array and although they can never be longer than the array, they can be smaller. The `make` function can take in a 3rd paramater to represent the capacity of the underlying array.

```go
mySlice := make([]float64, 5, 10)
```

Another way to create a slice is to use the `low:high` syntax.

```go
// the following are equivalent
arr := []float64{1,2,3,4,5}
x := arr[0:5]
```

## Maps
A map is an unordered collection of key-value pairs. Also known as an associative array, a hash table or a dictionary.

```go
var x map[string]int // declaration

// but maps need to be initialised before they can be used
y := make(map[string]int)
y["key"] = 10
fmt.Println(y["key"])

// delete values from a map
delete(y, "key")
```

Accessing an element of a map can return two values instead of just one. The first value is the result of the lookup and the second indicates whether or not the lookup was successful.

```go
elements := make(map[string]string)
elements["H"] = "Hydrogen"
name, ok := elements["Un"]
fmt.Println(ok) // prints false
```

There is a handy Go idiom to handle this:
```go
if name, ok := elements["Un"]; ok {
	fmt.Println(name, ok)
}
```
