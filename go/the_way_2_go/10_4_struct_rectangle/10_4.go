package main

import (
    "fmt"
    "./rectangle"
)

func main() {
	r1 := rectangle.Rectangle{Hight: 10, Width: 4.34}
    fmt.Println()
    fmt.Printf("Rectangle `r1': %v.\n", r1)
    if arr, err := r1.Area(); err == nil {
        fmt.Printf("Area : %.2f units.\n", arr)
    }
    if prr, err := r1.Perimeter(); err == nil {
        fmt.Printf("Perimeter : %.2f units.\n", prr)
    }
    fmt.Println()
}
