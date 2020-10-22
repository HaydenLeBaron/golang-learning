/*
This file contains a demo of file IO in Go
*/
package main
import (
    "fmt"
    "os"
    "log"
    "io/ioutil"
)

func main() {

    // Write to file
    file, err := os.Create("samp.txt")
    if err != nil {
        log.Fatal(err)
    }
    file.WriteString("Here is some text")
    file.Close()

    // Read from file
    stream, err := ioutil.ReadFile("samp.txt")
    if err != nil {
        log.Fatal(err)
    }
    readString := string(stream)
    fmt.Println(readString)
}
