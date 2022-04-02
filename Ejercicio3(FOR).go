package main

import (
	"fmt"
)

func main() {
	//lo uso para saber en que posicion esta cada caracter

	for pos, char := range "hola\x80pancho" { //el range interpreta como un arreglo de caracteres el string

		fmt.Printf("Character %#U starts at byte position %d\n", char, pos)

	}

	/*	sum := 0
		array[3] := (3,5,6)
		for _, value := range array{ ----->lo utilizo para ignorar los parametros que devuelve una funcion

			sum += value
		}

	*/

	//El defer ejecuta al final, invierta el orden
	//

}
