package main

import (
    "fmt"
)

func main()  {
    animal := [...]string{"dog", "cat", "parrot"}
    fmt.Println(animal[0])
    fmt.Println(animal[1])
    fmt.Println(animal[2])

    x := 10
    y := 2

    x++
    y--

    fmt.Println(x)
    fmt.Println(y)
}