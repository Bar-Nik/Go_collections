package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// 1. Методы с разными возвращаемыми значениями: Добавьте к структуре
// "Rectangle" метод "IsSquare", который возвращает bool, указывающий,
// является ли прямоугольник квадратом.

func (r Rectangle) IsSquare() bool {
	if r.Width == r.Height {
		return true
	}
	return false
}

// 2. Методы, изменяющие значения: Добавьте к структуре "Rectangle" метод
// "DoubleSize", который удваивает размеры прямоугольника.

func (r Rectangle) DoubleSize() (float64, float64) {
	return r.Height * 2, r.Width * 2
}

// 3. Использование нескольких методов: Используйте методы "Area" и
// "IsSquare" на нескольких экземплярах структуры "Rectangle".

func main() {
	r := Rectangle{
		Width:  20,
		Height: 20,
	}
	fmt.Println("Площадь прямоугольника:", r.Area())
	fmt.Println("Прямоугольник - квадрат?", r.IsSquare())

	a := Rectangle{
		Width:  10,
		Height: 30,
	}
	fmt.Println("Площадь прямоугольника:", a.Area())
	fmt.Println("Прямоугольник - квадрат?", a.IsSquare())
}
