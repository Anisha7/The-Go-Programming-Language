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

