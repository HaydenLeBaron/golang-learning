/*
This file demonstrates the use of Go Routines and Channels.

Go Routines allow for parallelism.
Channels allow the passing of data between Go Routines.
TODO: Create example for channels
*/

package main
import (
    "fmt"
    "time"
)

func main() {
    for i := 0; i < 10; i++ {
        go count(i) // This allows versions of count to be run in parallel.
    }

    var userInput string
    fmt.Scanln(&userInput) // This exists to keep the app open long enough for Go Routines to finish
}

func count(id int) {
    for i:= 0; i < 10; i++ {
        fmt.Println(id, ":", i)
        time.Sleep(time.Millisecond * 1000)
    }
}
