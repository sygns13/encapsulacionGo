package course

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

//Metodo de la estructura
func (c *Course) PrintClasses() {
	text := "Las clases del curso %s son: "
	for _, class := range c.Classes {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}

func (c *Course) ChangePrice(newPrice float64) {
	c.Price = newPrice
}
