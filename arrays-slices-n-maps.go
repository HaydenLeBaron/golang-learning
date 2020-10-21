/*
Example code for arrays, slices, and maps
*/

package main
import "fmt"


func main() {

    ///////////////////////////////////////
    // Arrays
    ///////////////////////////////////////
   fmt.Println("////////////////////////////\n// ARRAYS\n///////////////////////////\n");



    var favNums2[3] float64
    favNums2[0] = 0.12
    favNums2[1] =  123
    favNums2[2] =  3.14

    fmt.Println(favNums2[2])

    favNums3 := [3]int {-1, 4, 3}

    // iterate through array printing
    // <value> <idx>
    for i, value := range favNums3 {
        fmt.Println(value, i)
    }


    ///////////////////////////////////////
    // Slices
    ///////////////////////////////////////
    fmt.Println("////////////////////////////\n// SLICES\n///////////////////////////\n");

    // Slices work like lists and slices in Python I think.
    numSlice := []int {5, 4, 3, 2, 1}

    numSlice2 := numSlice[3:5]  // [start:end] corresponds to interval [starti, endi)
                                // AKA [starti, endi-1] 
    // numSlice2 := numSlice[3:] // Go from idx 3 to end

    numSlice3 := make([]int, 5, 10) // Not exactly sure what this does
    copy(numSlice3, numSlice)
    fmt.Println(numSlice3[0])
    numSlice3 = append(numSlice3, 0, -1) // Append more values to numSlice 3

    fmt.Println("numSlice =", numSlice)
    fmt.Println("numSlice = ", numSlice2)
    fmt.Println("numSlice = ", numSlice3)

    ///////////////////////////////////////
    // MAPS
    ///////////////////////////////////////
    fmt.Println("////////////////////////////\n// MAPS\n///////////////////////////\n");

    presAge := make(map[string] int)    // Create a map (aka dictionary)
    presAge["Obama"] = 42               // Add mapping
    fmt.Println(presAge)
    presAge["Obama"] = 100              // Edit mapping
    fmt.Println(presAge)
    delete(presAge, "Obama")            // Delete mapping
    fmt.Println(presAge)


}
