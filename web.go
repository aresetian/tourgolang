package main

import (
    "fmt"
    "net/http"
    "os"
    "github.com/aresetian/tourgolang/concurrency"
    "github.com/aresetian/tourgolang/template"
    "github.com/aresetian/tourgolang/handlers"
	"path/filepath"
    
)

var templatesPath = "templates"
 
// our init function to get our templates path depending on where we are.
func init() {
	dir, _ := os.Getwd() // gives us the source path if we haven't installed.
	templatesPath = filepath.Join(dir, templatesPath)
}

func main() {

    if err := handlers.ServerFunc(templatesPath); err != nil {
		fmt.Println("Error starting server!")
		os.Exit(1)
	}
	
    /*http.HandleFunc("/", hello)*/
    fmt.Println("listening...")
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
}

func hello(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "Benchmark Innova4j en la nueva app")
    template.Print(res)
    fmt.Fprintln(res, concurrency.Exercise65())
    fmt.Fprintln(res, "fin test")
}