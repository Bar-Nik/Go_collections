package main

import "fmt"

// 5. Сравнение структур: Создайте два экземпляра одной и той же
// структуры. Сравните их с помощью оператора == и выведите результат.

type Dog struct {
	Name string
	Age  int
}

func main() {
	dog1 := Dog{
		Name: "Baly",
		Age:  3,
	}
	dog2 := Dog{
		Name: "Reks",
		Age:  3,
	}
	fmt.Println(dog1 == dog2)

}
