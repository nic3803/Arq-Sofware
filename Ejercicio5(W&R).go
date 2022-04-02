package main

import (
	"fmt"
	"os"
)

func main() {

	//FILE WRITE

	f, err := os.Create("C:/Users/nic22/Desktop/Arq-Sofware/n.txt") //*os.File
	fmt.Fprintln(f, "Hola como estasssssssssss")
	err = f.Close()

	if err != nil {

		fmt.Println("EROR")

	}

	//FILE READ

	f2, err1 := os.ReadFile("C:/Users/nic22/Desktop/Arq-Sofware/n.txt")

	fmt.Print(string(f2))

	if err1 != nil {

		fmt.Println("ERROR EN LA LECTURA")
	}

}
