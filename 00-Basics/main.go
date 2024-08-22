// Go es un lenguaje compilado
// Go es de tipado estático
// Go tiene un recolector de basura
//  esto quiere decir que libera memoria de variables que no están siendo utilizadas.

package main // Go se ejecuta en el paquete main

import "fmt" // Aqui hacemos uso de la palabra imports para importar librerias externas o tambien internas de Go

// Datatypes
var (
	boolean bool           = true         // boolean
	integer int            = 10           // integer
	float   float32        = 3.14         // float
	string1 string         = "Hola Mundo" // string
	char    rune           = 'a'          // char
	array   [5]int                        // array
	slice   []int                         // slice
	map1    map[string]int                // map
	pointer *int                          // pointer
)

func main() { // Y su punto de entrada es esta función main ubicada en este paquete main
	fmt.Println("Hola Mundo")

	// Conditionals
	if boolean { // Si boolean es verdadero
		fmt.Println("Boolean es verdadero") // Imprime esto
	} else { // Sino
		fmt.Println("Boolean es falso") // Imprime esto
	}

	for i := 0; i < 5; i++ { // Para i que es igual a 0, mientras i sea menor a 5, incrementa i en 1
		fmt.Println(i) // Y imprime cada valor de i
	}
}
