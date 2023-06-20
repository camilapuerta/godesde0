package main

import (
	"github.com/ptilotta/godesde0/middleware"
)

func main() {
	/*Estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(Estado)
	fmt.Println(texto)

	numero, texto := ejercicios.ConverNumerico("50")
	fmt.Println(numero)
	fmt.Println(texto)

	funciones.Exponencia(2)

	mapas.MostrarMapas()

	Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)

	defer_panic.EjemploPanic()

	//rutina asincrona
	go gorutines.MiNombreLento("Camila Puerta")

	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)

	//channels
	canal1 := make(chan bool)
	go gorutines.MiNombreLento("Camila Puerta", canal1)
	fmt.Println("Estoy aqui")

	<-canal1

	webserver.MiWebServer()*/

	middleware.MiMiddleware()

}
