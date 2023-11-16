package main

import (
	"fmt"
	"pacotes/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Avião")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("samylle1brito@gmail.com")
	fmt.Println(erro)
}
