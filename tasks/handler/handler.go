package handler

import (
	"api_tinggal_nikah/db"
	"api_tinggal_nikah/repository"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hibiken/asynq"
)

func HandlerDeleteImageTask(ctx context.Context, t *asynq.Task) error {
	cwd, _ := os.Getwd()
	fmt.Println("task delete image jalan")

	conn := db.GetDB()
	GalleryPhotosRepo := repository.NewGalleryPhotosRepository(conn)

	data, err := GalleryPhotosRepo.GetAllGalleryPhotos()
	if err != nil {
		return err
	}

	for _, items := range data {

		pathtempimage := filepath.Join(cwd, "..", "user", "temp_image")
		filename := strings.Split(items.Path, "/")
		pathfile := filepath.Join(pathtempimage, filename[1])

		if fileExists(pathfile) {
			err := deleteFile(pathfile)
			if err != nil {
				return fmt.Errorf("error deleting file: %v", err)
			} else {
				return nil
			}
		} else {
			fmt.Println("file not exist")
		}

	}

	return nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func deleteFile(filename string) error {
	err := os.Remove(filename)
	return err
}
