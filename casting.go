/*
A demo of type casting in Go.
*/

package main
import (
    "fmt"
    "strconv"
)

func main() {
    randInt := 5
    randFloat := 10.5
    randString := "100"
    randString2 := "250.5"

    fmt.Println(float64(randInt))  // int -> float64
    fmt.Println(int(randFloat)) // float64 -> int
    newInt, _ := strconv.ParseInt(randString, 0, 64) // 64 bit
    fmt.Println(newInt)

    newFloat, _ := strconv.ParseFloat(randString2, 64) // 64 bit
    fmt.Println(newFloat)
}
