package funciones

import "fmt"

func tabla(valor int) func() int {
	numero1 := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero1 * secuencia
	}
}

func LlamarClosure() {
	tabladel := 5
	Mitabla := tabla(tabladel)
	for i := 1; i <= 10; i++ {
		fmt.Println(Mitabla())
	}
}
