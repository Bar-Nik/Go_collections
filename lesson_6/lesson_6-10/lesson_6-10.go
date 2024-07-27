package main

import "fmt"

// 10. Функция как аргумент другой функции (сложное): Исследуйте,
// можем ли мы в Go передать функцию как аргумент другой функции. Если
// нет, то объясните, как можно достичь похожего результата.

func main() {
	fmt.Println(func1(func2))

}

func func1(fn func(int) int) int {
	r := fn(2)
	return r
}

func func2(x int) int {
	return x * 2
}
