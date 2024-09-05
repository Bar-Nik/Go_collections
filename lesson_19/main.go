package main

import (
	"log"
	"log/slog"
	"os"
)

// 1. Попробуйте вариант установки минимального уровня в “error” и
// посмотрите на результат логирования
//
// 2. Запись в файл ( сложное): изучите, что такое параметр os.Stdout и то,
// как можно его заменит на файл? Сделайте запись логов в файл

func main() {
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	miniLevel := slog.LevelInfo
	logger := slog.New(slog.NewTextHandler(f, &slog.HandlerOptions{
		Level: miniLevel,
	}))
	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error")
}
