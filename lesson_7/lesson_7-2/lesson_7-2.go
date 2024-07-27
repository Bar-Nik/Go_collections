package main

import (
	"fmt"
	"math"
	"time"
)

// 2. Использование нескольких пакетов: Импортируйте пакеты "math" и
// "time", используйте "math" для вычисления квадратного корня текущего
// времени в секундах.

func main() {
	t := time.Now().Second()
	fmt.Println(math.Sqrt(float64(t)))
}
