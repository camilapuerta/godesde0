package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ptilotta/godesde0/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.TablaDeMultiplicar()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.TablaDeMultiplicar()
	if Append(fileName, texto) == false {
		fmt.Println("Error al concatenar conetido")
	}

}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append" + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString " + err.Error())
		return false
	}

	arch.Close()
	return true
}

func leoArchivo() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al leer archivo" + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}

	archivo.Close()
}
