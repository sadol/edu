package main

import "fmt"

func main() {
    // --------------NEW---------------------
    var i int
    fmt.Printf("var i int ==>  value: %d, type: %T, representation: %#v\n" ,i ,i, i)
    var j *int
    fmt.Printf("var j *int ==>  value: %d, type: %T, representation: %#v\n" ,j ,j, j)
    var k [5]int
    fmt.Printf("var k [5]int ==>  value: %d, type: %T, representation: %#v\n" ,k ,k, k)
    var l []int
    fmt.Printf("var l []int ==>  value: %d, type: %T, representation: %#v\n" ,l ,l, l)
    var m *int = new(int)
    fmt.Printf("var m *int = new(int) ==>  value: %d, type: %T, representation: %#v\n" ,*m ,m, m)
    (*m) = 10
    fmt.Printf("-----(*m) = 10 ==>  value: %d, type: %T, representation: %#v\n" ,*m ,m, m)
    var n *[5]int = new([5]int)
    fmt.Printf("var n *[5]int = new([5]int) ==>  value: %d, type: %T, representation: %#v\n" ,n ,n, n)
    n[0] = 10
    fmt.Printf("-----n[0] = 10 ==>  value: %d, type: %T, representation: %#v\n" ,n ,n, n)
    var o []int = make([]int, 5)
    fmt.Printf("var o []int = make([]int, 5) ==>  value: %d, type: %T, representation: %#v\n" ,o ,o, o)
    o[0] = 1
    o[1] = 2
    o[2] = 3
    o[3] = 4
    o[4] = 5
    fmt.Printf("-----o[0] = 10 ==>  value: %d, type: %T, representation: %#v\n" ,o ,o, o)
    var p *[]int = new([]int)
    fmt.Printf("var p *[]int = new([]int) ==>  value: %d, type: %T, representation: %#v\n" ,p ,p, p)
    p = &o
    fmt.Printf("-----p = &o ==>  value: %d, type: %T, representation: %#v\n" ,p ,p, p)
    oo := o[3:]
    fmt.Println(oo)
    oo[1] = 100
    fmt.Println(o)
}
