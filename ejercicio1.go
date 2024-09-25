package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func imprimir() {
	fmt.Println("Texto desde la funcion imprimir")
}
func hola_string(s string) string {
	return "Hola " + s
}

func sumar(a int, b int) int {
	return a * b
}

func comparar_bool() {
	var a bool
	b := false
	a = false
	if a == b {
		fmt.Println("A y B son iguales")
	} else {
		fmt.Println("A y B son diferentes")
	}

}

func arreglo() {
	var aprendices [5]string
	aprendices[0] = "Leonardo"
	aprendices[1] = "Juan Sebastian"
	aprendices[2] = "Franck"
	aprendices[3] = "Juan jose"
	aprendices[4] = "Jaider"

	fmt.Println(aprendices[1])
}

func tipos_datos() {
	var Texto string = "Sebastian"
	var numero int = 1000
	var decimal float64 = 0.00001

	fmt.Println(reflect.TypeOf(Texto))
	fmt.Println(reflect.TypeOf(numero))
	fmt.Println(reflect.TypeOf(decimal))
}

func converstringtoboolean() {
	var palabra string = "true"

	boolean, err := strconv.ParseBool(palabra)
	fmt.Println(boolean, reflect.TypeOf(boolean))
	fmt.Println("Este es el error: ", err)
}

func convertirtexto() {
	var palabra_bool bool = true
	strbool := strconv.FormatBool(palabra_bool)
	fmt.Println(strbool, reflect.TypeOf(strbool))
}

func variables_1() {
	var nombre, apellido string = "Yashido", "Salsa"
	fmt.Println(nombre, apellido)
}

func variables_2() {
	var (
		nombre     string = "Yashido"
		apellido   string = "Salsa"
		edad       int    = 18
		estudiando bool   = true
	)
	fmt.Println("Nombre: ", nombre)
	fmt.Println("Apellido: ", apellido)
	fmt.Println("Edad: ", edad)
	fmt.Println("Estudiando: ", estudiando)

}

func valor_cero() {
	var nombre string
	var edad int
	var peso float64
	var studiante bool
	fmt.Println("Nombre: ", nombre)
	fmt.Println("Edad: ", edad)
	fmt.Println("Peso: ", peso)
	fmt.Println("Estudiante: ", studiante)
}

func declaracion_corta_variable() {
	nombre := "Yashido"
	edad := 88
	peso := 80
	estudiando := false
	fmt.Println("Nombre: ", nombre)
	fmt.Println("Edad: ", edad)
	fmt.Println("Peso: ", peso)
	fmt.Println("Estudiante: ", estudiando)

}

var profesion = "Deportista"

func variable_global() {
	sueldo := 1000000
	fmt.Println("Profesion: ", profesion)
	fmt.Println("Sueldo: ", sueldo)
}

var var1 = "Este es el nivel 1"

func scope() {
	var var2 = "aeste es el nivel 2"

	{
		var var3 = "Este es el nivel 3"
		fmt.Println(var3)
	}
	fmt.Println(var1)
	fmt.Println(var2)

}

func punteros_1() {
	color := "rojo"
	fmt.Println(color, &color)

}

func punteros_2() {

	computador1 := "rojo"
	fmt.Println("El computador1 es ", computador1)

	computador2 := computador1
	fmt.Println("El computador2 es ", computador2)

	computador3 := &computador1
	fmt.Println("El computador3 es ", *computador3)

	computador1 = "amarillo"

	fmt.Println("El computador1 es ", computador1, &computador1)
	fmt.Println("El computador2 es ", computador2, &computador2)
	fmt.Println("El computador3 es ", *computador3, computador3)

}

func argumentos(altura float32) float32 {
	altura = altura * 3.28
	return altura
}

func argumentos2(altura *float32) float32 {
	*altura = *altura - 0.10
	return *altura
}

var altura float32 = 1.70


const Pi = 3.1415




func main() {
	// fmt.Println("Texto desde la funcion main")
	// imprimir()
	// fmt.Println(hola_string("Yashido"))
	// fmt.Println(sumar(3, 5))
	// comparar_bool()
	// arreglo()
	// tipos_datos()
	// converstringtoboolean()
	// convertirtexto()
	// variables_1()
	// variables_2()
	// valor_cero()
	// declaracion_corta_variable()
	// variable_global()}
	// scope()
	// punteros_1()
	// punteros_2()
	// fmt.Println("La altura es: ", altura, "mts")
	// fmt.Println("La altura es:", argumentos(altura), " pies")
	// fmt.Println("La nueva altura es: ", altura, "mts")
	// fmt.Println("La altura es: ", altura, "mts")
	// fmt.Println("La altura es:", argumentos2(&altura), " pies")
	// fmt.Println("La nueva altura es: ", altura, "mts")
	// fmt.Println("El valor de Pi es: ", Pi)
}
