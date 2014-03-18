package concurrency

import "fmt"

func sum(a []int, c chan int) {
    sum := <- c
    for _, v := range a {
        sum += v
    }
    c <- sum // send sum to c
}



func Exercise66() {
    
    a := []int{7, 2, 8, -9, 4, 0}
    c := make(chan int)
    
    go sum(a[:len(a)/2], c)
    go sum(a[len(a)/2:], c)
    
   
    
    x, y := <-c , <-c // receive from c
    fmt.Println(x, y, x+y)
}