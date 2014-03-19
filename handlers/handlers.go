package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

// Estructura donde se almacenar-n los valores de ingresados por el usuario, de esta estructura una vez este completa la plantilla 
// form_results.html tomar- los valores para mostrarlos en pantalla.
type UserContextInput struct {
	Tag              string
	Script           string
	Style            string
	AttributeLink    string
	AttributeDouble  string
	AttributeSingle  string
	AttributeOnEvent string
	AttributeSrc     string
}

// variable al alcance del paquete para mantener el camino de las plantillas
var templates string

/*
 * Se crean los manejadores de las URl que va a tener en el proyecto, para este caso son dos / y /display
 */
func ServerFunc(dir string) (err error) {
	// Se asigna las ruta donde estan las platillas
	templates = dir
    // os.Stat(templates) : verifica que las rutas de las plantillas indicadas en la variable templates, se enceuntra en el disco
    // duro y que son rutas validas, en caso de no encontrar alguna ruta retornar- un *PathError y se imprimir- templates path is incorrect!
	if _, err := os.Stat(templates); err != nil {
		return fmt.Errorf("templates path is incorrect!")
	}
	// Mapea  / a la funci-n rootHandler, y se adiciona un filtro antes del llamado a rootHandler con la funci-n disableXSSProtection
	http.HandleFunc("/", disableXSSProtection(rootHandler))
	// Mapea  /display a la funci-n postHandler  y se adiciona un filtro antes del llamado a rootHandler con la funci-n disableXSSProtection
	http.HandleFunc("/display", disableXSSProtection(postHandler))
	// retorna nil en caso de no encontrar ningun error
	return nil
}

// Nota (no hacer esto(Desactivar la proteccion XSS en IE/Chrome) ) . la  proteccion de XSS para estos navegadores se raeliza para poder ejecutar el ejemplo.
func disableXSSProtection(function func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
    // retorna un clousore wrapper del response con la adicion de la cabeceza
	return func(w http.ResponseWriter, r *http.Request) {
	    // la cabecera X-XSS-Protection no solociona todos los casos de XSS pero si algunos, para este caso se desactiva para poder hacer el ejemplo adecuadamente
	    // y ver como GOLANG al identificar estos posibles valores de ataque los transforma en valores seguros.
		w.Header().Set("X-XSS-Protection", "0")
		function(w, r)
	}
}
 
 
// Maneja las peticiones solicitadas en  / 
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Parsea la plantilla root.html 
	if t, err := template.ParseFiles(filepath.Join(templates, "root.html")); err != nil {
		// Se ha presentado un error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// retorna la respuesta al cliente por medio de t.Execute
		t.Execute(w, nil)
	}
}

// Maneja las peticiones solicitadas en  /display 
func postHandler(w http.ResponseWriter, r *http.Request) {
	// Parsea la plantilla root.html  con los valores ingresdos por el usuario
	if t, err := template.ParseFiles(filepath.Join(templates, "form_results.html")); err != nil {
	    // Se ha presentado un error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// se asegura que el parseo de los datos del  formulario son correcto, este m-todo es valido para 
		// peticiones POST o PUT
		// Documentacion del m-todo: http://golang.org/pkg/net/http/#Request.ParseForm
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form values!", http.StatusInternalServerError)
			return
		}
		// Se obtiene los valores del formulario
		userInput, err := getFormValues(&r.Form)
		if err != nil {
		    // Se ha presentado un error
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Go implicitamente injecta los valores del usuario en el contexto, es decir en el template form_results.html.
		t.Execute(w, userInput)
	}
}

// Parcea los valores del formulario en una estructura personalizada
func getFormValues(userInput *url.Values) (contextInput *UserContextInput, err error) {
	// crea una nueva estrucutura UserContextInput para mantener los valores ingresados por el usuario
	contextInput = new(UserContextInput)
 
	//  itera sobre cada uno de los valores asignados al formulario para asignarlos a la estrucutura UserContextInput(contextInput)
	for key, value := range *userInput {
		switch key {
		case "tag":
			contextInput.Tag = value[0]
		case "script":
			contextInput.Script = value[0]
		case "style":
			contextInput.Style = value[0]
		case "attr_double":
			contextInput.AttributeDouble = value[0]
		case "attr_single":
			contextInput.AttributeSingle = value[0]
		case "attr_onevent":
			contextInput.AttributeOnEvent = value[0]
		case "attr_src":
			contextInput.AttributeSrc = value[0]
		default:
			return nil, fmt.Errorf("El formulario contiene valores desconocidos!") 
		}
	}
	return contextInput, nil //Es seguro retorna un puntero en GO.
}