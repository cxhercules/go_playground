package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    var i uint32 = 4
    var d float32 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewReader(os.Stdin)

    // Declare second integer, double, and String variables.
    var x uint32
    var c float32
    var t string
    
    // Read and save an integer, double, and String to your variables.
    fmt.Scan(&x)
    fmt.Scan(&c)
    t, _ = scanner.ReadString('\n')    
    
    // Print the sum of both integer variables on a new line.
    fmt.Printf("%d\n", i + x )
 
    // Print the sum of the double variables on a new line.
    fmt.Printf("%.1f\n", d + c )
    
    // Concatenate and print the String variables on a new line
    // The 's' variable above should be printed first.
    fmt.Printf("%s\n", s + t )
}
