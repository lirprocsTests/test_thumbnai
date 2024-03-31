package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDatabase() error {
	var err error
	DB, err = sql.Open("sqlite3", "./internal/database/cache.db")
	if err != nil {
		//fmt.Println("Error opening database:", err)
		return err
	}
	return err
}

func Cache(videoID string) error {
	if DB == nil {
		return fmt.Errorf("database connection is nil")
	}

	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS thumbnails (id INTEGER PRIMARY KEY AUTOINCREMENT, video_id TEXT, thumbnail BLOB)")
	if err != nil {
		fmt.Println("Error creating table:", err)
		return err
	}

	// SQL запрос для вставки данных в таблицу
	stmt, err := DB.Prepare("INSERT INTO thumbnails (video_id, thumbnail) VALUES (?, ?)")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()

	image, err := os.ReadFile("IMG/" + videoID + ".jpg")
	if err != nil {
		fmt.Println("Error reading image file:", err)
		return err
	}

	// Выполнение запроса для сохранения изображения
	_, err = stmt.Exec(videoID, image)
	if err != nil {
		fmt.Println("Error inserting image into database:", err)
		return err
	}

	fmt.Println("Image cached successfully.")
	return nil
}

func GetFromCache(videoID string) error {
	if DB == nil {
		return fmt.Errorf("database connection is nil")
	}

	var imageData []byte
	err := DB.QueryRow("SELECT thumbnail FROM thumbnails WHERE video_id = ?", videoID).Scan(&imageData)
	if err != nil {
		fmt.Println("Error retrieving image from database:", err)
		return err
	}

	// Сохранение изображения в файл
	err = os.WriteFile("IMG/"+videoID+"_FromCache"+".jpg", imageData, 0644)
	if err != nil {
		fmt.Println("Error saving retrieved image to file:", err)
		return err
	}

	fmt.Println("Image retrieved from cache successfully.")

	return nil
}
