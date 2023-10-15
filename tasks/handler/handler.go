package handler

import (
	"api_tinggal_nikah/config"
	"api_tinggal_nikah/db"
	"api_tinggal_nikah/repository"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hibiken/asynq"
	"github.com/minio/minio-go/v7"
	"github.com/samber/lo"
)

func HandlerDeleteImageTask(ctx context.Context, t *asynq.Task) error {
	var imagebucket []string
	var imagedatabase []string
	bucket := config.GetClientMinio()

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
				fmt.Println("berhasil deleted file")
			}
		} else {
			fmt.Println("file not exist")
		}

		imagedatabase = append(imagedatabase, filename[1])
	}

	objectCh := bucket.ListObjects(ctx, os.Getenv("WASABI_BUCKET_NAME"), minio.ListObjectsOptions{
		Prefix:    "image/",
		Recursive: true,
	})

	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
		}
		filename := strings.Split(object.Key, "/")

		imagebucket = append(imagebucket, filename[1])
	}

	ImageNotInDatabase := lo.Without(imagebucket, imagedatabase...)

	for _, items := range ImageNotInDatabase {
		imagebucketpath := filepath.Join("image", items)
		fmt.Println("image bucket path : ", imagebucketpath)

		err := bucket.RemoveObject(ctx, os.Getenv("WASABI_BUCKET_NAME"), imagebucketpath, minio.RemoveObjectOptions{
			ForceDelete: true,
		})

		if err != nil {
			return fmt.Errorf("error deleting file bucket : %v", err)
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
