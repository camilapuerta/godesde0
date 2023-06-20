package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Inicio")

	resultado := operacionesMidd(sumar)(2, 3)
	fmt.Println(resultado)

	resultado = operacionesMidd(restar)(10, 3)
	fmt.Println(resultado)

	resultado = operacionesMidd(multiplicar)(2, 3)
	fmt.Println(resultado)
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a, b)
	}
}
