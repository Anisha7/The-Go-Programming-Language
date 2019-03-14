# 1. Hello, World

## package fmt
This package contains functions like Println.

## package main
Function main is where the execution of the program begins. Main will call functions from other packages.
### Note: must import packages used by the program

## declarations
func, var, const, type
A function declaration consists of keyword func, name of function, parameter list (empty for main), a result list (empty), and body of the function (enclosed in braces)

## Semantics
- Semicolons only required when two or more statements/declarations appear on the same line.
- New line can occur after the operator but not before (ex. x + y)

## Run Instructions
Make sure you're in this repo in your terminal and run the below commands.
```
cd Chapter1-Tutorial
go run helloworld.go
```
You should see Hello, world print in your terminal.

# 2. Command-Line Arguments
## package os
Can access command line arguments in program using os.Args, which returns a slice of strings (similar to lists or arrays in other languages).

## Run Instructions
```
go run echo1.go words you want to print
```
You can play around with this by replacing "words you want to print" with whatever you like.

## Loops
The for loop is the only loop statement in Go
```
for initialization; condition; post {
    // zero or more statements
}
```
A traditional while loop would be as follows:
```
for condition {
    // ...
}
```

## The blank identifier: _
The underscore can be used when syntax requires a variable name but the program logic does not, like in the case of our echo program, echo2.go consists of the improvement.
```
go run echo2.go words you want to print
```

## Efficiency concern with string += string
This creates a new string each time and would be costly for large data. 
Use the strings.Join method as follows:
```
func main() {
    fmt.Println(strings.Join(os.Args[1:], " "))
    // or for debugging (and no specific formatting)
    // fmt.Println(os.Args[1:])
}
```

## Test exercises
```go run exercises1-2.go hi there fellas```

# 3. Finding Duplicate Lines
### Semantics
- Parantheses are never used around the condition in an if statement.
- Maps don't have key errors. If non-existent key is accessed, map returns the zero value of its type.
- Important Verbs (string formatting): %d (decimals), %s (string), %t (bool), %% (literal percent sign)
- String literals: '\t' for tab and '\n' for new line
- Functions and other package-level entities may be declared in any order

### Lets try our first duplicate line function
Look over dup1.go

This will take in input from command line and tell us how many duplicate lines exist
```echo "hi there\n moo moo\n moo moo" | go run dup1.go```
This should print 
```2   moo moo```

### Now lets try one that can take in file as input
Look over dup2.go

```go run dup2.go test.txt```
This should print
```
2       Hi there I am a person
2       Why
```

We made sure to allow this function to also take command line input.
Try this:
```echo "hi there\n moo moo\n moo moo" | go run dup2.go```
This should print 
```2   moo moo```

### The above methods are known as "streaming mode, where input is read and broken into lines as needed. Let's read the entire input into memory in one big gulp and spit it into lines all at once in this next approach.
Look over dup3.go

```go run dup3.go test.txt test2.txt```

### Now lets print the file names next to the duplicate lines as well
Look over exercise1-3.go

```go run exercises1-3.go test.txt test2.txt```

## 4. Animated GIFS
### Notes
- Composite literals: compact notation for instantiating any of Go's composite types from a sequence of values. Example from file lissajous.go:
```
[]color.Color{...} // slice
gif.GIF{...} // struct: a group of values v=called fields
// anim is a struct of type gif.GIF
```

### So, we created GIFs using code!! How to run it?

The basic version:
```
go build lissajous.go
./lissajous >out.gif
```

The colorful version:
```
go build exercises1-4.go
./exercises1-4 >multi.gif
```

## 5. Fetching a URL

Test fetch.go by using the a random dog breed api (or use your own)

```
go build fetch.go
go run fetch.go https://dog.ceo/api/breeds/image/random
```

Optimized for space and efficienct:
```
go build exercises1-5.go
go run exercises1-5.go https://dog.ceo/api/breeds/image/random
```

## 6. Fetching URLs Concurrently

```
go build fetchall.go 
go run fetchall.go golang.org gopl.io godoc.org
```

## 7. A Web Server
Run the command:
```
go run server1.go &
```
This prints: 
```
[1] 3988
```

Run the command:
```
go build fetch.go
./fetch http://localhost:8000
```
This prints:
```
URL.Path = "/"
```

Run the command:
```
./fetch http://localhost:8000/help
```
This prints:
```
URL.Path = "/help"
```

### How to kill the running server?
```
Coming soon
```

## 8. Loose Ends
- Control Flow: if, for, Switch statements
- Named types: A type declaration to give a name to an existing type.
- Pointers: & for address of a variable and * for retrieving the pointer
- Methods and interfaces
- Packages: https://golang.org/pkg , https://godoc.org
    ``` go doc http.ListenAndServe```
- Comments: /* ... */ for multiple line comments.