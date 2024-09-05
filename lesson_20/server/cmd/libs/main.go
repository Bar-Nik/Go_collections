package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"server/internal/api"

	"github.com/gorilla/mux"
)

// 1. Добавить новый middleware : который будет логировать информацию о том,
// какой метод ручки был вызван, если вызван POST и PUT , выводить в логи
// “создаем” или обновляем” и аналогично с GET и DELETE
//
// 2. Логируйте body ( сложное): в middleware сделайте логирование body
// для всех типов запросов, содержащие body. Важное уточнение - нужно
// логировать тело запроса, а не само поле структуры :)

// 1. Методы: превратите все функции из пакета API в методы, создав
// структуру “server”, на котором завязаны эти методы.

func main() {
	r := mux.NewRouter()

	miniLevel := slog.LevelInfo
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: miniLevel,
	}))

	r.Use(api.Logging(logger))
	r.HandleFunc("/book", api.GetBook).Methods(http.MethodGet)
	r.HandleFunc("/book", api.AddBook).Methods(http.MethodPost)
	r.HandleFunc("/book", api.DeleteBook).Methods(http.MethodDelete)
	r.HandleFunc("/book", api.UpdateBook).Methods(http.MethodPut)
	r.HandleFunc("/books", api.AllBooks).Methods(http.MethodGet)
	err := http.ListenAndServe("127.0.0.1:8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
