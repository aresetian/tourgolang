/*
  Para los usuarios de Java la paqueteria no indica  la ruta fisica que se acostumbra a crear en los archivos .java, en GOLANG.
  la paqueteria es un nombre l-gico que no necesariamente debe estar en creada en el disco duro.
  
  Tip: El nombre del paquete no puede empezar con numeros.
*/
package basicconcepts

// El paquete fmt que puede ser encontrado en http://golang.org/pkg/fmt/ para su an-lisis  es un paquete que permite formatear 
// la escritura y entrada de cadenas en una consola.
// Todos los paquetes de GOLANG pueden ser encontrados en http://golang.org/pkg/, la anterior ruta es el equivalente a la API de JAVA
import (
    "fmt"
)

/* En GOLANG no existen clases como en otros lenguajes(Java), en GOLANG solo existen funciones y estructuras, para este caso solo se
 define una funci-n, las estructuras se ver-an en otros ejemplos.
 recordar que la funci-n con nombre main es especial y es el punto de entrada a una aplicaci-n, el equivalente en Java a el m-todo 
"public static void main(String[] arg)"

Si el nombre de la funci-n  inicia con letra may-scula indica que la funcion es publica y puede ser vista desde otros paquetes,
si el nombre de la funci-n inicia con minuscula indica que la fucnion es privada al archivo donde se declaro.
*/ 
func Exercice() string {
    // debido a que ya importamos el paquete para usarlo se debe volver colocar el nombre del paquete seguido de un punto y la 
    // primera letra de la funci-n a invocar debe estar en mayuscula
    fmt.Println("Hello my firts line on GOLANG")
    
    return  "retorno de ejercicio uno"
}