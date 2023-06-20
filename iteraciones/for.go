package iteraciones

import (
	"fmt"
)

func Iterar() {
	for i := 10; i > 1; i-- {
		if i == 6 {
			fmt.Println("seis")
		}
		fmt.Println(i)
	}
}
