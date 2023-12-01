# Programming GO

- [Programming GO](#programming-go)
  - [Createing a new module](#createing-a-new-module)
    - [Adding a dependency](#adding-a-dependency)
    - [Upgrading dependencies](#upgrading-dependencies)
  - [Go Project Layout](#go-project-layout)
  - [Go Language Fundamentals](#go-language-fundamentals)
    - [Variable Declaration](#variable-declaration)
    - [Anonymous Function](#anonymous-function)
    - [Strcuts](#strcuts)
    - [Interface](#interface)
  - [Useful Libraries](#useful-libraries)
    - [Cli](#cli)


Take a look at this short summary [here](https://go.dev/blog/using-go-modules), then follow the rest of this doc. You 
may also take a look at the [GO libraries](./go-libraries.md) doc.

`go.mod`: defines the module's module path, import path used for the root directory and its dependency requirements.
`go.sum`: contains the expected cryptographic hashes of the content of specific module versions.

## Createing a new module

Start outside of $GOPATH/src direcotry:

```bash
mkdir -p ~/mymodule/hello && cd mymodule
touch hello/hello.go main.go
```

and `hello/hello.go`:
```go
package hello

func Hello() string {
    return "Hello, world."
}
```

`main.go`:
```go
package main

import (
    "myorg/myproject/hello"
)

func main() {
    fmt.Println(hello.Hello())
}
```

`hello_test.go`:
```go
package hello

import "testing"

func TestHello(t *testing.T) {
    want := "Hello, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
```

To make a go module we would need to make this directory the root of a module:

```bash
# go mod init github.com/org/module
go mod init myorg/myproject
go mod tidy
# then test the module
go test

go build -o main main.go
```

### Adding a dependency

`hello.go`:
```go
package hello

import "rsc.io/quote"

func Hello() string {
    return quote.Hello()
}
```

Now use `go get rsc.io/quote` to fetch the module and its dependencies. 
List the main module and its dependencies by: `go list -m all`

### Upgrading dependencies

Running `go get <module>` will update the package. Sometimes this causes issues. Run `go test` afterwards to see if everything looks fine.

```bash
go list -m -versions rsc.io/sampler
go get rsc.io/sampler@v1.3.1
```

```bash
# clean up the project from unused dependencies
go mod tidy
```

## Go Project Layout
Take a look at this [page](https://github.com/golang-standards/project-layout).
- `internal`: internal is a special directory dedicated to only the parent directory import. 
That means only `github.com/org/app` or `github.com/org/app/internal/sibilings` can import 
`github.com/org/app/internal/config` and no others.

## Go Language Fundamentals

### Variable Declaration
```go
var age int = 25
// Type inference
var age = 25
age := 25
fmt.Println("Age:", age)

var name string = "John"
// Type inference
var name = "John"
name := "John"
fmt.Println("Name:", name)

var (
    age  = 25
    name = "John"
)

var count int  // Zero value for int is 0
var flag bool  // Zero value for bool is false
var text string // Zero value for string is ""
```

### Anonymous Function
```go
add := func(x, y int) int {
    return x + y
}

// Using the anonymous function
result := add(3, 5)
fmt.Println(result) // Output: 8
```


### Strcuts
Look at the the [struct example](./examples/ex-struct/) folder.
```go
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// declared as a value
person1 := Person{
    FirstName: "John",
    Age:       30,
}

// declared as a pointer
person2 := &Person{
    FirstName: "Jane",
    Age:       25,
}

// Dereference the pointer
fmt.Println("P2 Name:", (*person2).FirstName) 
// Alternatively, use Shorthand notation
fmt.Println("P2 Name:", person2.LastName)
```

### Interface
An interface is a type that specifies a set of methods signature. Interfaces in Go are particularly useful 
when you want to define a function (e.g. `printArea`) that multiple types (e.g. `Circle`, `Rectangle`) must 
adhere to. For example, a simple printer 
application that can print information about different shapes without knowing the specific type of each shape.
Look at the [exinterface example](./examples/ex-interface/) folder.
```go
type Circle struct {
	Radius float64
}
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of the circle.
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
// Area calculates the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// We want a single function that prints the area of each shape
// We need to define Shape as interface
func printArea(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
}

// Shape is an interface that defines a method for calculating the area.
type Shape interface {
	Area() float64
}
``` 

## Useful Libraries
### Cli
Take a look at the [cli example](./examples-libraries/cli/) to learn how to manage arguments.