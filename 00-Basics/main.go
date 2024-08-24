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

func incrementa(x *int) { // Esta funcion recibe un puntero de tipo int
	*x++ // Obtiene el valor de x con el operador * y lo incrementa en 1
}

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

	switch integer { // Switch es como un if pero con multiples casos
	case 1: // Si integer es igual a 1
		fmt.Println("Es 1") // Imprime esto
	case 2: // Si integer es igual a
		fmt.Println("Es 2") // Imprime esto
	case 3: // Si integer es igual a 3
		fmt.Println("Es 3") // Imprime esto
	default: // Si no se cumple ningun caso
		fmt.Println("No es 1, 2 o 3")
	}

	pointer = &integer // Aqui se le asigna el valor de 10

	// Este es un ejemplo basico sobre como se puede hacer uso de un puntero
	x := 10
	incrementa(&x) // Aqui se manda a llamar la funcion incrementa y se le pasa la direccion de memoria de x
	fmt.Println(x) // Luego se imprime el valor de x y listo, el valor de x se incremento en 1
	//*pointer = 500
	fmt.Println(*pointer)
}
