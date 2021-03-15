package main

import "fmt"

func main() {
    nat := make(chan int, 5)
    quad := make(chan int, 10)

    go func(n chan<- int) {
        for x := 0; x < 100; x++ {
            n <- x
        }
        close(n)
    }(nat)

    go func(n <-chan int, q chan<- int) {
        for x := range n {
            quad <- x*x
        }
        close(quad)
    }(nat, quad)

    for q := range quad {
        fmt.Println(q)
    }
}

//func main() {
//    n := make(chan int)
//    s := make(chan int)
//
//    go naturals(n)
//    go squa(n, s)
//    printer(s)
//}

//func naturals(in chan<-int) {
//    for x := 1; x < 100; x++ {
//        in<-x
//    }
//    close(in)
//}
//
//func squa(out <-chan int, in chan<- int) {
//    for x := range out {
//        in<-x*x
//    }
//    close(in)
//}
//
//func printer(out <-chan int) {
//    for s := range out {
//        fmt.Println(s)
//    }
//}
