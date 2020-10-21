/*
About variables, constants, and some types in go
*/

package main
import "fmt"

func main() {
    var age int = 40
    var favNum float64 = 1234.5
    randNum := 1 // Go is smart enough to make this an int

    // randNum = "Hello" // This would cause an error, because you cant change the type of a var.

    fmt.Println(age, favNum, randNum) // Automatic spacing here

    const pi float64 = 3.141592

    var ( // Declare and init multiple variables
        varA = 2
        varB = 3
        varC = 4
    )

    fmt.Println("Vars A, B, C:", varA, varB, varC)

    // strings can be enlosed in "" or ``

    var myName string = "Hayden"
    fmt.Println("My name:", myName + `\n`+ "len(myName)", len(myName))

    var isOver40 bool = true

    // Like C printf in a lot of ways
    fmt.Printf("%T \n", pi) // Prints the data type of pi
    fmt.Printf("%t \n", isOver40) // Prints the type of a bool
    fmt.Printf("%x \n", 17)  // Print 17 in hex
    fmt.Printf("%d \n", 17)  // Print 17 as an int




}
