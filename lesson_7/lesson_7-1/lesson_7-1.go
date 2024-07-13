package main

import (
	"fmt"
	"math"
)

// 1. Использование стандартной библиотеки: Импортируйте пакет "math" и
// используйте его для вычисления квадратного корня числа.

func main() {
	var num float64
	fmt.Scan(&num)
	fmt.Println(math.Sqrt(num))
}
