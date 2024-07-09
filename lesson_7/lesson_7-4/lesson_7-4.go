package main

import (
	"fmt"

	"math"
)

// 4. Использование специфической функции из пакета: Импортируйте
// только функцию Sqrt из пакета "math" и используйте ее в вашей
// программе.

func main() {
	var s float64
	fmt.Scan(&s)
	fmt.Println(math.Sqrt(s))
}

// не нашел как импортировать только функцию sqrt
