package main

import (
	"errors"
	"fmt"
)

func dividir(a float32, b float32) (float32, error) {

	var err error = nil
	if b == 0 {

		err = errors.New("Error en el valor de divison")
		return 0, err
	}

	resultado := a / b

	return resultado, err

}

func main() {

	div, er := dividir(6, 0)

	if er != nil {

		fmt.Println(er.Error())

	} else {
		fmt.Println(div)
	}

	//fmt.Println(div, er)

}
