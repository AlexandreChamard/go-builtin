package builtin

type Integer interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int8 | ~int16 | ~int32 | ~int64 | ~int | ~uint | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	Integer | Float
}

type Ord interface {
	Number | ~string
}

// Function type to do ordering operation
type Ordf[T any] func(a, b T) bool

// Function type to do comparing operation
type Compf[T any] func(a, b T) bool

// Function type to do equality operation
type Eqf[T any] func(a T) bool

// Generic function that performs a == b
func Equal[T comparable](a, b T) bool { return a == b }

// Generic function that performs a < b
func Less[T Ord](a, b T) bool { return a < b }

// Generic function that performs a <= b
func LessEqual[T Ord](a, b T) bool { return a <= b }

// Generic function that performs a > b
func Greater[T Ord](a, b T) bool { return a > b }

// Generic function that performs a >= b
func GreaterEqual[T Ord](a, b T) bool { return a >= b }
