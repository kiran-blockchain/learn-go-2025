package main

import (
    "fmt"
    "golang.org/x/exp/constraints"
)


// Generic function: Print elements of any slice
func PrintSlice[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}

// Generic struct: Stack
type Stack[T any] struct {
    data []T
}

func (s *Stack[T]) Push(val T) {
    s.data = append(s.data, val)
}

func (s *Stack[T]) Pop() T {
    if len(s.data) == 0 {
        var zero T
        return zero
    }
    last := s.data[len(s.data)-1]
    s.data = s.data[:len(s.data)-1]
    return last
}

// Generic function with constraints: Max of two values
func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func main() {
    fmt.Println("🔹 Generic PrintSlice:")
    PrintSlice([]string{"apple", "banana", "cherry"})
    PrintSlice([]int{1, 2, 3})

    fmt.Println("\n🔹 Generic Stack[int]:")
    var intStack Stack[int]
    intStack.Push(10)
    intStack.Push(20)
    fmt.Println("Pop:", intStack.Pop())
    fmt.Println("Pop:", intStack.Pop())

    fmt.Println("\n🔹 Max with constraints:")
    fmt.Println("Max(3, 7):", Max(3, 7))
    fmt.Println("Max(4.5, 2.3):", Max(4.5, 2.3))
}