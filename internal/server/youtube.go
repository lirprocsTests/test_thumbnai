package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"thumbnail/internal/database"
)

// DownloadThumbnail загружает thumbnail видео с YouTube по его идентификатору
func DownloadThumbnail(videoID string) error {
	thumbnailURL := fmt.Sprintf("https://img.youtube.com/vi/%s/default.jpg", videoID)

	response, err := http.Get(thumbnailURL)
	if err != nil {
		return fmt.Errorf("failed to download thumbnail: %v", err)
	}
	defer response.Body.Close()

	name := fmt.Sprintf("%s.jpg", videoID)
	file, err := os.Create("IMG/" + name)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return fmt.Errorf("failed to save thumbnail: %v", err)
	}

	if err := database.Cache(videoID); err != nil {
		return fmt.Errorf("failed to cache thumbnail: %v", err)
	}

	return nil
}
