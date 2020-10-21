/*
This file demos control flow in go.
*/
package main
import "fmt"

func main() {

    // While loops ("spelled \'for\' ")
    i := 1
    for i <= 10 {
        fmt.Println(i)
        i++
    }

    for j := 0; j < 5; j++ {
        fmt.Println(j)
    }


    // if-elseif-else statements
    yourAge := 17
    if yourAge >= 21 {
        fmt.Println("You can drink and drive")
    } else if yourAge >= 16 {
        fmt.Println("You can drive")
    } else {
        fmt.Println("You can't drink or drive")
    }

    // switch statements
    switch yourAge {
        case 16 : fmt.Println("Go drive")
        case 18 : fmt.Println("Go vote")
        default : fmt.Println("Go have fun")
    }
}
