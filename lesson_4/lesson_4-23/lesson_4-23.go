package main

import "fmt"

// 2. Изменение значений в map: Создайте map с несколькими парами ключ-
// значение. Измените одно из значений и выведите всю map на экран.

func main() {
	cars_speed := map[string]int{
		"car_1": 150,
		"car_2": 160,
	}
	cars_speed["car_2"] = 200
	fmt.Println(cars_speed)
}
