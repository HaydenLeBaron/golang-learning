/*
This file demos pointers, structs, and interfaces in Go.
*/

package main
import (
    "fmt"
    "math"
)

func main() {

    // Pointer demo
    x := 0 // creates an int
    fmt.Println("x =", x)
    changeXValNow(&x, 5)
    fmt.Println("x =", x)

    yPtr := new(int) // Creates a pointer to an int.  Same as: var yPtr *int = new(int)
    *yPtr = 5
    fmt.Println("yPtr =", *yPtr)
    changeXValNow(yPtr, 100)
    fmt.Println("yPtr =", *yPtr)

    rect1 := Rectangle{leftX: 0, topY: 50, height: 10, width: 12}
    // Or do this: rect1 := Rectangle{0, 50, 10, 12}
    fmt.Println("Rectangle width: ", rect1.width)
    fmt.Println("Rectangle area: ", rect1.area())


    // Demo interfaces
    circle1 := Circle{10}
    fmt.Println("Rectangle Area = ", calcArea(rect1))
    fmt.Println("Circle Area = ", calcArea(circle1))

}

// For demoing pointers
func changeXValNow(x *int, newVal int) {
    *x = newVal
}

type Shape interface {
    area() float64
}

type Rectangle struct {
    leftX float64  // These are private fields because their names are lowercase
    topY float64
    height float64
    width float64
    Name string   // this is a public field because its name is capitalized
}

type Circle struct {
    radius float64
}

func (rect Rectangle) area() float64 {
    return rect.width * rect.height
}

func (c Circle) area() float64 {
    return math.Pi * math.Pow(c.radius, 2)
}

func calcArea(shape Shape) float64 {
    return shape.area()
}

// TODO: how would I fulfill an interface where all args expected pointers to Shapes / Rectangles / Circles etc?
