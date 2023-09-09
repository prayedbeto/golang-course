package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicar() {
	scanner := bufio.NewScanner(os.Stdin)
	var numero int
	var err error

	fmt.Print("Ingrese un número: ")

	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("Error con el numero " + err.Error())
		}
	}

	for i := 1; i < 11; i++ {
		fmt.Println("\nTabla de multiplicar del ", numero, "es igual a", numero*i)
	}
}
