package variables

import "fmt"

func MuestroEnteros() {
	//3 variables que si inicializan
	//a la variable intde64 le voy a guardar lo que me devuelva en la funcion int64

	var intcomun int
	intde32 := int32(10)
	intde64 := int64(100)

	//3 display por consola
	fmt.Println("Intcomun = ", intcomun)
	fmt.Println("Intde32 = ", intde32)
	fmt.Println("Intde64 = ", intde64)

}
