/*
 El ejemplo de las plantillas fue tomado de :http://blog.veracode.com/2013/12/golangs-context-aware-html-templates/ 
 
 la idea del ejemplo es mostrar como GOLANG encode cada uno de los valores de atack XSS en valores seguros, para qu esta transformacion a valores seguros
 sea teniada encuenta lo unico que hay que utilizar es la funcion Execute de html/template, no es necesario ninguan otra configuracion.
*/
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
    // os.Getwd : retorna el path del directorio actual. si el directorio actual es alcanzado por multiples pahtm debido a enlaces simbolicos,
    // Getwd retornara ya sea el enlace simbolico o la ruta
	dir, _ := os.Getwd() 
	// Al path retorna en la linea anterior se le adicina el nombre del la carpeta donde estan las plantillas para este caso el nombre de la carpeta
	// a adicoina es templates(valor de la variable templatesPath)
	templatesPath = filepath.Join(dir, templatesPath)
}

func main() {

    if err := handlers.ServerFunc(templatesPath); err != nil {
		fmt.Println("Error iniciando el servidor!")
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