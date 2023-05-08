package main

import (
	"github.com/go-redis/redis/v8"
	"golangTask/src/varve/impl/apis"
	"golangTask/src/varve/impl/config"
	"golangTask/src/varve/impl/service"
	"log"
	"net/http"
	"time"
)

func main() {

	const folderPath = "...fileStorage\\impl"

	go fileReadingSchedule(config.ConnToRedis(), folderPath)

	http.HandleFunc("/promotions/", apis.GetPromotionByIDHandler)

	log.Fatal(http.ListenAndServe(":1321", nil))

}

func fileReadingSchedule(rdb *redis.Client, folderPath string) {

	ticker := time.NewTicker(time.Hour / 2)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			service.ReadFolder(rdb, folderPath)
		}
	}
}
