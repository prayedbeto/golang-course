package main

import (
	"fmt"
	"iusa-auto-deploy/ejercicios"
	// "iusa-auto-deploy/variables"
	// "runtime"
)

func main() {
	// fmt.Println("Hola mundo!")

	// variables.MuestroEnteros()

	// variables.RestoVariables()

	// estado, texto := variables.ConvertirATexto(1588)
	// fmt.Println(estado)
	// fmt.Println(texto)

	// if os := runtime.GOOS; os == "linux" || os == "darwin" {
	// 	fmt.Println("Esto no es Windows, es", os)
	// } else {
	// 	fmt.Print("Esto es Windows ", os)
	// }

	// Switch
	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es Linux")
	// case "darwin":
	// 	fmt.Println("Esto es Darwin")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }

	// Ejercicio 01
	numero, texto := ejercicios.ConvNumerico("50")

	fmt.Println(numero)
	fmt.Println(texto)
}
