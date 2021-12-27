package main

import "fmt"

func main() {

	//Instanciando estructuras
	Go := &course.Course{
		Name:    "Go desde cero",
		Price:   123.34,
		IsFree:  false,
		UserIDs: []uint{12, 56, 89},
		Classes: map[uint]string{
			1: "Introduccion",
			2: "Estructuras",
			3: "Variables",
		},
	}

	Go.PrintClasses()

}

//Funcion normal
func PrintClassesNormal(c Course) {
	text := "Las clases del curso %s son: "
	for _, class := range c.Classes {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}
