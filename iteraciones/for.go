package iteraciones

import "fmt"

func Iterar() {
	i := 0

	for i < 10 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
		j++
	}
}
