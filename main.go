package main

import "fmt"

// Add function returns sum of two numbers
func Add(a, b int) int {
    return a + b
}

func main() {
    fmt.Println("Sum:", Add(3, 5))
}
