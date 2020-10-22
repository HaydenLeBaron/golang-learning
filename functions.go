/*
This file contains info and examples related to functions in Go.
*/

package main
import "fmt"

func main() {
    listOfNums := []float64 {1.1, 2.2, 3.3, 4.4, 5.5}
    var result float64 = sumFloatSlice(listOfNums)
    fmt.Println("Sum of  {1.1, 2.2, 3.3, 4.4, 5.5} :", result)

    num1, num2 := next2Values(1, 10) // Call function that returns 2 vals
    fmt.Println("Next 2 values of 1 and 10:", num1, num2)

    // Call variatic function
    fmt.Println("Subtract, 1 through 5:", subtractIntSlice(1, 2, 3, 4, 5))

    // Closure
    num3 := 3
    doubleNum := func() int {
        num3 *= 2 // This function can see num3 because it is nested in the same scope as main. The value of num3 is now changed.
        return num3
    }

    // num3 == 3
    fmt.Println(doubleNum()) // num3 == 6
    fmt.Println(doubleNum()) // num3 == 12

    // Call recursive function
    fmt.Println("4! == ", factorial(4))

    defer printTwo() // Because printTwo is deferred, it will be called only after main finishes.
    printOne()

    fmt.Println(safeDiv(3, 0)) // Will cause error, but caught by recover in safeDiv function
    fmt.Println(safeDiv(3, 2))

    // Call panic demo function
    demoPanic()
}

// Simple function
// func funcname(argname argtype) returntype { ... }
func sumFloatSlice(floatSlice []float64) float64 {
    sum := 0.0
    for _, val := range floatSlice {
        sum += val
    }
    return sum
}


// Function that returns two values
func next2Values(num1 int, num2 int) (int, int) {
    return num1 + 1, num2 + 1
}

// Variatic function (takes undefined number of values)
func subtractIntSlice(args ...int) int {

    sub := 0
    for _, val := range args {
        sub -= val
    }
    return sub
}

// Recursive function
func factorial(n int) int {
    if n == 1 {
        return 1
    }
    return n * factorial(n-1)
}

func printOne() {fmt.Println(1)}
func printTwo() {fmt.Println(2)}

func safeDiv(num1, num2 int) int {
    defer func() {
        fmt.Println(recover()) // If num2 == 0, then div by 0 will cause a fatal error. Recover will basically "catch" this error--and allow execution to continue.
    }()

    solution := num1 / num2
    return solution
}

func demoPanic() {
    defer func() {
        fmt.Println(recover())
    }()

    panic("Panic") // Throws an exception that will be caught by recover
}

