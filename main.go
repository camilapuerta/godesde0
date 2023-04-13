package main

//importaciones
import (
	"fmt"

	"github.com/ptilotta/godesde0/variables"
)

// funcion, declaracion de variables, resto de codigo etc..
func main() {
	Estado, texto := variables.ConviertoaTexto(1588)

	fmt.Println(Estado)
	fmt.Println(texto)
}
