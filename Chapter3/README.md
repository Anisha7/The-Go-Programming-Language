# Basic Data Types
- Go's Types: Basic types, aggregate types, reference types, and interface types

## 1. Integers
- int8, int16, int32, int64
- uint8, uint16, uint32, uint64
- int, uint: 32 or 64 bits (unknown, depends)
- rune: int32
- byte: uint8
- uintptr: big enough to hold pointer value
- Signed integers: -2^(n-1) to 2^(n-1) - 1
- Unsigned integers: non-negative values, 0 to 2^(n) - 1
- Overflow: when the result of an arithmetic operation (signed/unsigned) has more bits than can be represented in the result type.
- page 53 for bitwise operations
- %d, %o, %x for radix and format control with printing numbers (page 56)

## 2. Floating-Point Numbers
- Float32, Float64 types; Use math.MaxFloat32 or math.MaxFloat64 for math num
- math.NaN() checks if value is not a number

Check surface files for example. 
```
cd surface
go run main.go
```

Challenges completed:
1. Modify the program to skip invalid programs (non-finite float64 values)
2. Experiment with visualizations of other functions from the math package. Can you produce egg box, moguls, or a saddle? (TODO)
3. Color each polygon based on its height, so that the peaks are colored red (#ff0000) and the valleys blue
4. Following the lissajous example approach, construct a web server that computes serfaces and writes SVG data to the client. The server must set the Content-Type header like this:
    ``` w.Header().Set("Content-Type", "image/svg+xml") ```

```
cd surfacedraw
go run main.go
```

## 3. Complex Numbers
- complex64, complex128
- Use built-in function complex (returns a complex number from real and imaginary components) ex. complex(1,2) for 1+2i
- Use built-in real and imag functions to get those components
- library math/cmplx

Challenges (TODO) :
5. Implement a full-color Mandelbrot using image/NewRGBA and type color.RGBA or color.YCbCr.
6. Supersampling reduces the effect of pixelation by computing the color value at several points within each pixel and taking the average. The simplest method is to divide each pixel into four "subpixes." Implement this.
7. Another simple fractal uses Newton's method to find complex solutions to a function such as z^(4) - 1 = 0. Shade each starting point by the number of iterations required to get close to one of the four roots. Color each point by the root it approaches.
8. Rendering fractals at high zoom levels demands great arithmetic precision. Implement this fractal using 4 different representations of numbers: complex64, complex128, big.Float, and big.Rat (math/big package). Compare performance and memory usage. At what zoom levels do rendering artifacts become visible?
9. Write a web server that renders fractals and writes image data to the client. Allow the client to specify the x, y, and zoom values as parameters to the HTTP request.

## 4. Booleans
- true, false

## 5. Strings
- Try these uses of string to remove prefix/suffixes and format integer strings
`
cd basename1
go run main.go '/loc/hello.py' // prints hello

cd basename2
go run main.go '/loc/hello.py' // prints hello

cd comma
go run main.go '1234523' // prints 1,234,523
`

- Building up strings can be a lot of allocation and copying. In these cases, using bytes.Buffer type can be more efficient.
- Strings can be converted to bytes slices and back:
`
s := "abc"
b := []byte(s)
s2 := string(b)
`

Check out Exercises 3.10 to 3.12 by running the go files in the respective folders.

- strconv package can be used for string <-> int conversions (Atoi, ParseInt, ParseUint)

## 6. Constants
- iota: A constant generator, used to create a sequence of related values without spelling each one out. Ex. weekdays.
- See how its used in netflag
- Check exercise 3.13 for more examples on how to use iota
- Untyped Constants : super precisely represented by compiler and can by used in expressions even though some may be too big to be stored in a specific variable like int.

