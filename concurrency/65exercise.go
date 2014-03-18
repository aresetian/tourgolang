package concurrency

import (
    "fmt"
    "time"
    "runtime"
)

func say(s string) {
    for i := 0; i < 1000; i++ {
     //   time.Sleep(100 * time.Millisecond)
        fmt.Printf("Este es el mensaje %d : %s. \n",i,s)
    }
    /*for i := 0; i < 5; i++ {
       // time.Sleep(1 * time.Millisecond)
        fmt.Printf("Este es el mensaje %d : %s. \n",i,s)
    }
    for i := 0; i < 5; i++ {
    //    time.Sleep(100 * time.Millisecond)
        fmt.Printf("Este es el mensaje %d : %s. \n",i,s)
    }
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Printf("Este es el mensaje %d : %s. \n",i,s)
    }*/
}

func Exercise65() {
    tiempo1 := time.Now()
   // runtime.GOMAXPROCS(1)
    fmt.Println(runtime.GOMAXPROCS(2))
    go say("peru")
    go say("colombia")
    go say("ecuador")
    go say("argentina")
    go say("brazil")
    go say("venuela")
    
    //say("hello")
    
    // fmt.Println("Este es el mensaje")
   for i := 0; i < 500; i++ {
     //   time.Sleep(100 * time.Millisecond)
     fmt.Printf("Este es el mensaje %d : %s. \n",i,"ultimo")
    } 
    
   //time.Sleep(100 * time.Millisecond)
    fmt.Println(time.Now().Sub(tiempo1))
    /*fmt.Println(runtime.NumCPU())
    fmt.Println(runtime.NumGoroutine())
    fmt.Println(runtime.GOMAXPROCS(4))*/
    
}