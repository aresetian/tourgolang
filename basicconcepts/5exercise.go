/*
  Para los usuarios de Java la paqueteria no indica  la ruta fisica que se acostumbra a crear en los archivos .java, en GOLANG.
  la paqueteria es un nombre l-gico que no necesariamente debe estar en creada en el disco duro.
  
  Tip: El nombre del paquete no puede empezar con numeros.
*/
package basicconcepts

/*
 Este ejercicio esta utilizando dos paquetes el paquete fmt y math/rand
 
 Por convenci-n el nombre del paquete es el nombre del -ltimo elemento del camino importado donde se encuentra el archivo. por ejemplo,
 el paquete math/rand comprende archivos que comienzan con la declaraci-n package rand.
 
 
 Se puede utilizar una agrupacion de importaciones para cada paquete como se muestra a continuaci-n o 
 se puede utilizar un agrupacion en este ejercicio:
 
 import  "fmt"
 import  "math"

*/
import (
    "fmt"
    "math"
)



/* En GOLANG no existen clases como en otros lenguajes(Java), en GOLANG solo existen funciones y estructuras, para este caso solo 
 se define una funci-n, las estructuras se ver-an en otros ejemplos.
 recordar que la funci-n con nombre main es especial y es el punto de entrada a una aplicaci-n, 
 el equivalente en Java a el m-todo"public static void main(String[] arg)"
 
 
 
Todos los programas de GOLANG estan construidos de paquetes, el paquete principal es main y es el utilizado para iniciar un programa GOLANG,
este paquete principal debe contar con la funci-n main que es el punto de entrada para la ejecuci-n del proyecto.

Si el nombre de la funci-n  inicia con letra may-scula indica que la funcion es publica y puede ser vista desde otros paquetes,
si el nombre de la funci-n inicia con minuscula indica que la fucnion es privada al archivo donde se declaro.
*/




func Exersice5() {
    // debe dar error se se descomenta la linea : ./5exercise.go:50: cannot refer to unexported name math.pi, el anterior error se presenta porque las funciones exportadas por los paquetes
    // deben tener la primera letra en mayuscula de lo constrario la funcion no es importada.
    //fmt.Println(math.pi)
    
    fmt.Println(math.Pi) // Imprime 3.141592653589793 en la consola
}