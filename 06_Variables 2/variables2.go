package main

import "fmt"

var razaYo = "Divinidad"

// Fuera de función no se puede declarar con :=

func main() {
	// Nombrar variables
	numero := 25
	_nombre := "Alberto"
	numero2 := 54
	nombreUsuario := "admin"
	Numero := 46
	fmt.Println(numero, numero2, _nombre,
		nombreUsuario, Numero)
	var (
		dios               = "Yo"
		enemigo1, enemigo2 = "Nadie", "Nadie2"
	)
	var enemigo3 = 69
	fmt.Println(dios, enemigo1, enemigo2, enemigo3)

	//Scope

	fmt.Println("La raza de Yo es: " + razaYo)
	imprimir()
}

func imprimir() {
	fmt.Println("La raza de Yo es: " + razaYo)
}
