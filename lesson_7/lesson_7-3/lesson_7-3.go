package main

import (
	"fmt"
	ran "math/rand"
)

// 3. Использование псевдонимов для пакетов: Импортируйте пакет
// "math/rand" с псевдонимом и используйте его для генерации случайного
// числа.

func main() {
	fmt.Println(ran.Intn(100))
}
