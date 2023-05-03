package main

//importaciones
import (
	"fmt"
	//"github.com/ptilotta/godesde0/variables"
	"github.com/ptilotta/godesde0/ejercicios"
)

// funcion, declaracion de variables, resto de codigo etc..
func main() {
	/*Estado, texto := variables.ConviertoaTexto(1588)

	fmt.Println(Estado)
	fmt.Println(texto) */

	numero, texto := ejercicios.ConverNumerico("50")
	fmt.Println(numero)
	fmt.Println(texto)
}
