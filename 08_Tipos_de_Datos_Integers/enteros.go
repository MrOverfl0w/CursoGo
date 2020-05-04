package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		//Enteros con signo
		//entero8  int8  //Todos los enteros de 8 bits con signo (-128 to 127)
		//entero16 int16 //Todos los enteros de 16 bits con signo (-32768 to 32767
		entero32 int32 //Todos los enteros de 32 bits con signo (-2147483648 to 2147483647)
		entero64 int64 //Todos los enteros de 64 bits con signo (-9223372036854775808 to 9223372036854775807)

		//Enteros sin signo
		//uentero8  uint8  //Todos los enteros de 8 bits sin signo (0 to 255)
		//uentero16 uint16 //Todos los enteros de 16 bits sin signo (0 to 65535)
		//uentero32 uint32 //Todos los enteros de 32 bits sin signo (0 to 4294967295)
		//uentero64 uint64 //Todos los enteros de 64 bits sin signo (0 to 18446744073709551615)

		//Alias
		//enteroByte byte //alias para uint8
		enteroRune rune //alias para int32

		//Tipos dependientes de la plataforma
		//enteroUint    uint    //32 o 64 bits
		enteroInt int //32 o 64bits
		//enteroUintptr uintptr //Entero sin signo lo suficientemente grande para contener el valor de un puntero
	)

	//******************************************
	//Conversión entre tipos

	//Suma int32 y int64

	entero32 = 25 //int32
	entero64 = 85 //int64

	fmt.Println(entero32 + int32(entero64))

	//Suma int32 y rune
	enteroRune = 46 //rune
	fmt.Println(entero32 + enteroRune)

	//Suma int32 y int
	enteroInt = 56
	fmt.Println(entero32 + int32(enteroInt))

	fmt.Println(unsafe.Sizeof(enteroInt)*8, unsafe.Sizeof(entero32)*8)

}
