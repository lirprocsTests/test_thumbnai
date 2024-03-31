package main

import (
	"log"
	"thumbnail/internal/database"
	"thumbnail/internal/server"
)

func main() {
	err := database.InitDatabase()
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	// Создаем новый экземпляр gRPC сервера
	s := server.NewServer()

	// Запускаем сервер на порту ":50051"
	s.Run(":50051")

	// В случае ошибки сервера выводим сообщение об ошибке
	log.Fatalf("server stopped unexpectedly")
}
