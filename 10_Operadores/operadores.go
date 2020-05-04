package main

import "fmt"

func main() {
	// Operadores Aritméticos
	// + Suma
	// - Resta
	// * Multiplicación
	// / División
	// % Mod

	a := 5
	b := 3

	fmt.Println("5 + 3 = ", a+b)
	fmt.Println("5 - 3 = ", a-b)
	fmt.Println("5 * 3 = ", a*b)
	fmt.Println("5 / 3 = ", a/b)
	fmt.Println("5 % 3 = ", a%b)

	// Operadores de Comparación
	// == Igual que
	// != Diferente que
	// < Menor que
	// <= Menor o Igual que
	// > Mayor que
	// >= Mayor o Igual que

	c := 3
	fmt.Println("********************************")

	fmt.Println("5 == 3 = ", a == b)
	fmt.Println("3 == 3 = ", b == c)
	fmt.Println("5 != 3 = ", a != b)
	fmt.Println("3 != 3 = ", c != b)
	fmt.Println("5 < 3 = ", a < b)
	fmt.Println("5 <= 3 = ", a <= b)
	fmt.Println("3 <= 3 = ", c <= b)
	fmt.Println("5 > 3 = ", a > b)
	fmt.Println("5 >= 3 = ", a >= b)
	fmt.Println("3 >= 3 = ", c >= b)

	// Operadores de Asignación
	// += || a += b || a = a + b
	// -= || a -= b || a = a - b
	// *= || a *= b || a = a * b
	// /= || a /= b || a = a / b
	// %= || a %= b || a = a % b

	// Operadores Lógicos
	// && AND
	// || OR
	// ! Negación

	fmt.Println("********************************")
	fmt.Println("AND - &&")
	fmt.Println("true && true = ", true && true)
	fmt.Println("false && true = ", false && true)
	fmt.Println("true && false = ", true && false)
	fmt.Println("false && false = ", false && false)
	fmt.Println("OR - ||")
	fmt.Println("true || true = ", true || true)
	fmt.Println("false || true = ", false || true)
	fmt.Println("true || false = ", true || false)
	fmt.Println("false || false = ", false || false)
	fmt.Println("!true = ", !true)
	fmt.Println("!false = ", !false)

	// 	Jerarquía de Operadores
	// ()
	// * / %
	// + -
	// == != < <= > >=
	// &&
	// ||
	fmt.Println("********************************")
	fmt.Println((5 - 3) * 5)

	// Operadores de Incremento y Decremento
	// ++ || a++ || a = a + 1
	// -- || a-- || a = a - 1
	a++
	fmt.Println("a++ ", a)

}
