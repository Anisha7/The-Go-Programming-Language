# Program Structure

## 1. Names
- Begins with a letter, and may have more letters/digits/underscores
- Case matters
- Keywords cannot be used as names (ex. if, switch, func, etc.)
- If name begins with an upper case letter, it is exported (visible and accessible outside of its own package)
- "camel casing"

## 2. Declarations
- var, const, type, and func

```
./boiling
go run main.go
```

```
./ftoc
go run main.go
```

## 3. Variables
- var name type = expression
- Short variable declarations: used for local variables, ex. t := 0.0
- ':=' is a declaration and '=' is an assignment
- Pointers: &varname to access address, *varname to access pointer

Testing echo4: 
```
➜ cd echo4
➜ go build main.go
➜ ./main a bc def
a bc def
➜ ./main -s / a bc def
a/bc/def
➜ ./main -n a bc def  
a bc def%                                                                                                ➜ ./main -help
Usage of ./main:
  -n    omit trailing newline
  -s string
        separator (default " ")
```

- The built-in 'new' Function: new(T) creates an unnamed variable of type T, initialized to zero value of T, and returns its address of type *T
- Lifetime of Variables: depends of whether it is heap allocated or on the stack

## 4. Assignments
- Tuple Assignment: x,y = y,x or a[i], a[j] = a[j], a[i], or v, ok = map[key], or blank identifiers to discared result _, err = io.Copy(dst, src)
- Assignability: slices - madals := []string{"gold","silver","bronze"} is an array with those values at indices 0, 1, 2. (implicit assignment)

## 5. Type Declarations
- type name underlying-type
- Review tempconv0

## 6. Packages and Files
- Review tempconv to see how packages are cleared and how methods are used from a package file in other package files.
- Exercise 2.1 added in tempconv package: Add types, constants, and functions to tempconv for processing temperatures in the Kelvin scale. 
- Imports: check cf files
```
cd cf
go build main.go
./main 212
```
- Exercise 2.2: Write a general-purpose unit-conversion program analogous to cf that reads numbers from its command-line arguments or from standard input if there are no arguments, and converts each number into units like temperature in Celsius and Fahrenheit, etc. (Answer: Check kfc)

```
cd kfc
go build main.go
./main 212
```

- Package Initialization: func init() {...} can be used. They are executed automatically when program starts, so you should not call or reference it. 
- Exercises 2.3-2.5 in popcount package

```
➜ go run testPopcount.go
Testing Version 1
197
0.00s elapsed
Testing Version 2
197
0.00s elapsed
Testing Version 3
197
0.00s elapsed
```