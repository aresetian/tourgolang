package concurrency

import "testing"
import "fmt"
import "runtime"

func BenchmarkHello(b *testing.B) {
    runtime.GOMAXPROCS(4)
    //fmt.Println(runtime.NumCPU())
    
    fmt.Println("hello")
   
    
    for i := 0; i < b.N; i++ {
        
    }
    
    go say("peru")
    go say("colombia")
    go say("ecuador")
    go say("argentina")
    go say("brazil")
    go say("venuela")
    
    for i := 0; i < 500; i++ {
     //   time.Sleep(100 * time.Millisecond)
        fmt.Printf("Este es el mensaje %d : %s. \n",i,"ultimo")
    }
}



func say(s string) {
    for i := 0; i < 1000; i++ {
     //   time.Sleep(100 * time.Millisecond)
        fmt.Printf("Este es el mensaje %d : %s. \n",i,s)
    }
}

