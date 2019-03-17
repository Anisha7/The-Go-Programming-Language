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

Check surface files for example. Challenges completed:
1. Modify the program to skip invalid programs (non-finite float64 values)
2. Experiment with visualizations of other functions from the math package. Can you produce egg box, moguls, or a saddle? (TODO)
3. Color each polygon based on its height, so that the peaks are colored red (#ff0000) and the valleys blue (#0000ff) (TODO)
4. Following the lissajous example approach, construct a web server that computes serfaces and writes SVG data to the client. The server must set the Content-Type header like this:
    ``` w.Header().Set("Content-Type", "image/svg+xml")


