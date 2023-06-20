package arreglos_slices

import "fmt"

var tablas []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 98, 56, 8, 2, 1, 4, 6}

func MuestroSlices() {
	fmt.Println(tablas)

	porcion := arreglo[3:]   //Slice creado desde un vector, desde la posicion 3
	porsion2 := arreglo[:5]  //Slice creado desde un vector, desde la posicion 0 hasta la posicion 5
	porsion3 := arreglo[6:7] //Slice creado desde la posicion 6 hasta la 7

	fmt.Println(porcion)
	fmt.Println(porsion2)
	fmt.Println(porsion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
		fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))
	}
}
