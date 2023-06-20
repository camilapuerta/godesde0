package mapas

import (
	"fmt"
)

func MostrarMapas() {
	paises := make(map[string]string)

	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	//fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 38,
	}

	fmt.Print(campeonato)
	/*
		for equipo, puntaje := range campeonato {
			fmt.Printf("Equipo %s, tiene un puntaje %d \n", equipo, puntaje)
		} */

	delete(campeonato, "Real Madrid")
	fmt.Print(campeonato)

	puntaje, existe := campeonato["Juventus"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)
}
