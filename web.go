package main

import (
    "fmt"
    "net/http"
    "os"
    "github.com/aresetian/tourgolang/concurrency"
    "github.com/aresetian/tourgolang/template"
    
)



func main() {

    http.HandleFunc("/", hello)
    fmt.Println("listening...")
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
}

func hello(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "Benchmark Innova4j en la nueva app")
    template.Print()
    fmt.Fprintln(res, concurrency.Exercise65())
    fmt.Fprintln(res, "fin test")
}