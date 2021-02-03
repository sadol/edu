package shaper

// interface's naming convention os rather weird
type Shaper interface {
	Perimeter() float64
	Area() float64
	String() string
}
