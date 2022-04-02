package main

import (
	"fmt"
	"os"
)

func CreateFile() *os.File {

	f, _ := os.Create("C:/Users/nic22/Desktop/Arq-Sofware/nEjercio6.txt") //*os.File

	return f
}

func WriteFile(f *os.File) {

	fmt.Fprint(f, "NICO PRANDI CAPO")

}

func CloseFile(f *os.File) {
	fmt.Println("Cerrando")
	err := f.Close()

	if err != nil {

		println("ERROR")
	}

}

func main() {

	f := CreateFile()
	defer CloseFile(f)
	WriteFile(f)

	fmt.Println("Archivos Creado")

}
