package main

import "fmt"

// 4. Указатели на структуры: Создайте структуру и указатель на эту
// структуру. Измените значение поля структуры, используя указатель в
// методе.

type Car struct {
	Model  string
	Color  string
	Number int
}

func (c *Car) ChangeNumber(x int) {
	c.Number = x
}

func main() {
	p := Car{
		Model:  "VW",
		Color:  "Green",
		Number: 100,
	}
	fmt.Println(p)
	p.ChangeNumber(128)
	fmt.Println(p)

}
