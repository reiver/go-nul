# go-nul

Package **nul** implements a **nullable** **optional-type**, for the Go programming language.

In other programming languages, an **optional-type** might be know as: a **option-type**, or a **maybe-type**.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/sourcecode.social/reiver/go-nul

[![GoDoc](https://godoc.org/sourcecode.social/reiver/go-nul?status.svg)](https://godoc.org/sourcecode.social/reiver/go-nul)

## Examples

Here is an example **nullable** **optional-type** that can hold a string:
```go
import "sourcecode.social/reiver/go-nul"

//

var op nul.Optional[string] // the default value is nothing.

// ...

if nul.Nothing[string]() == op {
	fmt.Println("nothing")
}

// ...

op = nul.Null[string]()

// ...

if nul.Null[string]() == op {
	fmt.Println("null")
}

// ...

op = nul.Something("Hello world! 👾")

// ...

switch op {
case op.Nothing[string]():
	//@TODO
case op.Null[string]():
	//@TODO
case op.Something("apple"):
	//@TODO
case op.Something("banana"):
	//@TODO
case op.Something("cherry"):
	//@TODO
default:
	//@TODO
}

// ...

value, found := op.Get()
if found {
	fmt.Println("VALUE:", value)
} else {
	fmt.Println("either nothing or null")
}
```
## Import

To import package **nul** use `import` code like the follownig:
```
import "sourcecode.social/reiver/go-nul"
```

# Installation

To install package **nul** do the following:
```
GOPROXY=direct https://sourcecode.social/reiver/go-nul
```

## Author

Package **nul** was written by [Charles Iliya Krempeaux](http://changelog.ca)
