package shaper

type AreaInterface interface {
	Area() float64
}

type PeriInterface interface {
	Perimeter() float64
}

type Shaper interface {
	AreaInterface
	PeriInterface
	String() string
}
