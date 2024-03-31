package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sync"

	"google.golang.org/grpc"

	"thumbnail/internal/proto/gen"
)

func main() {
	// Определение флага для асинхронной загрузки файлов
	async := flag.Bool("async", false, "enable asynchronous download")
	flag.Parse()

	// Получение аргументов командной строки
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Usage: thumbnail-cli [--async] <video_link>")
		return
	}

	// Подключение к серверу gRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := gen.NewThumbnailServiceClient(conn)

	// Асинхронная загрузка файлов, если установлен флаг --async
	if *async {
		var wg sync.WaitGroup
		for _, link := range args {
			wg.Add(1)
			go func(link string) {
				defer wg.Done()
				if err := downloadThumbnail(client, link); err != nil {
					log.Printf("failed to download thumbnail for %s: %v", link, err)
				}
			}(link)
		}
		wg.Wait()
	} else { // Синхронная загрузка файлов
		for _, link := range args {
			if err := downloadThumbnail(client, link); err != nil {
				log.Printf("failed to download thumbnail for %s: %v", link, err)
			}
		}
	}
}

// Функция для загрузки миниатюры видео с помощью gRPC-сервера
func downloadThumbnail(client gen.ThumbnailServiceClient, link string) error {
	// Отправка запроса на сервер
	resp, err := client.GetThumbnail(context.Background(), &gen.ThumbnailRequest{VideoLink: link})
	if err != nil {
		return fmt.Errorf("failed to get thumbnail from gRPC server: %v", err)
	}

	// Вывод результата загрузки
	log.Printf("Thumbnail for video %s downloaded: %s", link, resp.Thumbnail)
	return nil
}
