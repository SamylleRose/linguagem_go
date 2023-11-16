package main

import "fmt"

func main() {
	var variavel1 string = "variavel 1"
	variavel2 := "variavel 2" //inferencia de tipo

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var ( //declara mais de uma variavel
		variavel3 string = "variavel 3"
		variavel4 string = "variavel 4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel5", "variavel6" // outra forma de declarar mais de uma variavel

	fmt.Println(variavel5, variavel6)

	const constante1 string = "constante" // declaração de uma constante
	fmt.Println(constante1)

	variavel1, variavel2 = variavel2, variavel1 //troca as variaveis
	fmt.Println(variavel1, variavel2)
}
