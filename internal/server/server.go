package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
	"sync"

	"thumbnail/internal/database"
	pb "thumbnail/internal/proto/gen" // Импорт сгенерированных protobuf-файлов
)

// Server представляет gRPC сервер
type Server struct {
	mutex sync.Mutex
	pb.UnimplementedThumbnailServiceServer
}

func (s Server) Run(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterThumbnailServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// NewServer создает новый экземпляр gRPC сервера
func NewServer() *Server {
	return &Server{}
}

func getVideoID(url string) string {
	// Разбиваем URL по символу "v="
	parts := strings.Split(url, "v=")

	// Получаем последнюю часть разбитой строки
	videoID := parts[len(parts)-1]

	return videoID
}

// GetThumbnail реализует метод GetThumbnail gRPC сервиса
func (s Server) GetThumbnail(ctx context.Context, request *pb.ThumbnailRequest) (*pb.ThumbnailResponse, error) {
	videoID := getVideoID(request.VideoLink)

	// Проверяем наличие thumbnail в кеше
	err := database.GetFromCache(videoID)
	if err == nil {
		return &pb.ThumbnailResponse{Thumbnail: videoID}, nil
	}

	// Если в кеше нет thumbnail, загружаем его
	if err := DownloadThumbnail(videoID); err != nil {
		return nil, fmt.Errorf("failed to download thumbnail: %v", err)
	}

	return &pb.ThumbnailResponse{Thumbnail: videoID}, nil
}
