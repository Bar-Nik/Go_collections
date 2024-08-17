package main

import (
	"encoding/json"
	"fmt"
)

// 3. Использование тегов: Создайте структуру Employee с полями ID , Name , и
// Product . Используйте теги для изменения имен полей в JSON.
// Сериализуйте экземпляр структуры в JSON и убедитесь, что имена полей
// соответствуют тегам.

type Employee struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Product string `json:"product"`
}

func main() {
	emp := Employee{ID: 1, Name: "Tom", Product: "Milk"}
	empJSON, err := json.Marshal(emp)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}
	fmt.Println("JSON:", string(empJSON))
}
