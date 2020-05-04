package main

import "fmt"

func main() {
	var numero int
	multiplicando := 25
	fmt.Println(numero * multiplicando)
	numero = 40
	fmt.Println(numero)
	nombre := "ABC"
	fmt.Println(nombre)
	nombre, numero = "Alberto", 123
	nombre2 := "Julio"
	fmt.Println(nombre, numero)
	fmt.Println(nombre, nombre2)
	nombre, nombre2 = nombre2, nombre
	fmt.Println(nombre, nombre2)
	nombre3, nombre := "Rafael", "Juan"
	fmt.Println(nombre, nombre3)

}
