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


Si GOLANG detecta que una de las importacines no esta siendo utilizada en la logica del archivo, es decir en alguna de las funciones
GOLANG no compilar- el archivo y imprimir- (imported and not used: "XXXX") donde XXXX es el nombre del paquete que no esta siendo utilizado y que hay que remover
para poder compilar el codigo.
*/
import (
    "fmt"
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

/*
Una funcion puede tomar cero o muchos par-metros.. ¿Cuanto es el n-umero de par-metros. permitido para lograr una buena calidad de c-digo?

Esta funci-n cuenta con dos par-metros.

GOLANG define el tipo de los par-metros despues del la declaraci-n del nombre del par-metro
(más informacion acerca la sintaxis en http://blog.golang.org/gos-declaration-syntax).
*/
func add(x int, y int) int {
    return x + y
}

/*
Cuando los par-ametros comparten la misma definicion de tipo se puede omitir el tipo de estos exepto el del -ltimo par-metro. 

en esta funcion acortamos la defici-n  "x int, y int" a "x, y int"
*/
func add2(x, y int) int {
    return x + y
}

func Exersice6() {
     fmt.Println(add(42, 13))
     fmt.Println(add2(42, 13))
}