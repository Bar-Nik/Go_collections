package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// 1. Создание простого HTTP сервера: Создайте HTTP сервер, который будет
// слушать на порту 8080 и возвращать приветственное сообщение в
// формате JSON при обращении по пути /hello .
// 2. Обработчики маршрутов: Создайте несколько обработчиков маршрутов.
// Например, /user , который будет возвращать информацию о пользователе,
// и /product , который будет возвращать информацию о продукте.
// 3. Обработка ошибок: Добавьте обработку ошибок в ваш сервер. Например,
// если обработчик не может декодировать JSON, он должен возвращать
// соответствующий код ошибки и сообщение. Подумайте, какой код лучше
// всего подходит по эту ситуацию.

type Msg struct {
	Message string `json:"msg"`
}
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type Product struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

func msgHandler(w http.ResponseWriter, r *http.Request) {

	msgForm := Msg{Message: "Hello"}
	msgJSON, err := json.Marshal(msgForm)

	if err != nil {
		log.Println("Ошибка: ", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(msgJSON)
	if err != nil {
		log.Fatal(err)
		return
	}

}

func userHandler(w http.ResponseWriter, r *http.Request) {

	userForm := User{Name: "Jack", Age: 30}
	userJSON, err := json.Marshal(userForm)

	if err != nil {
		log.Println("Ошибка: ", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(userJSON)
	if err != nil {
		log.Fatal(err)
		return
	}
}
func productHandler(w http.ResponseWriter, r *http.Request) {

	productForm := Product{Name: "apple", Amount: 12}
	productJSON, err := json.Marshal(productForm)

	if err != nil {
		log.Println("Ошибка: ", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(productJSON)
	if err != nil {
		log.Fatal(err)
		return
	}

}

func main() {
	http.HandleFunc("/hello", msgHandler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/product", productHandler)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
