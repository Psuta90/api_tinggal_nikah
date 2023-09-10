package db

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedisConnection() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"), // Ganti dengan alamat dan port Redis Anda
		Password: "",                      // Ganti dengan kata sandi Redis Anda jika diperlukan
		DB:       0,                       // Nomor database Redis yang akan digunakan
	})

	// Cek koneksi Redis
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
