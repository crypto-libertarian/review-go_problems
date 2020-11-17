package main

import (
    "fmt"
)

func main() {
    var a [10]int

    for i := range a {
        a[i] = i
    }

    var b [10]struct{n int; nn int}

    for i, n := range a {
        b[i].n  = n
        b[i].nn = n * n
    }

    for _, v := range b {
        fmt.Printf("%v: %v\n", v.n, v.nn)
    }
}
