package service

import (
	"encoding/csv"
	"github.com/go-redis/redis/v8"
	"golangTask/src/varve/impl/entities"
	"io"
	"log"
	"os"
	"path/filepath"
)

func ReadFolder(rdb *redis.Client, folderPath string) {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Println("Folder read error:", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".csv" {
			// if .csv file found, read its content
			readFile(rdb, filepath.Join(folderPath, file.Name()))

			// delete read file
			err = os.Remove(filepath.Join(folderPath, file.Name()))
			log.Println("File deletion error:", err)

		}
	}
}

func readFile(rdb *redis.Client, filePath string) {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		promotion := buildPromotion(record)

		save(promotion, rdb)
	}
	log.Println(file.Name(), " file was successfully read")
}

func buildPromotion(row []string) entities.Promotion {

	var promotion = entities.Promotion{
		ID:             row[0],
		Price:          row[1],
		ExpirationDate: row[2],
	}

	return promotion
}
