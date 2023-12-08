package pkg

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"service/internal/model"
)

func ReadOrdersFromFiles(jsonFiles []string) []model.Order {
	var orders []model.Order

	for _, jsonFile := range jsonFiles {
		jsonData, err := os.ReadFile(jsonFile)
		if err != nil {
			log.Fatalf("Error reading JSON file: %v", err)
		}

		var order model.Order
		err = json.Unmarshal(jsonData, &order)
		if err != nil {
			log.Fatalf("Error unmarshalling JSON: %v", err)
		}

		orders = append(orders, order)
	}

	return orders
}

func ReadOrdersFromDirectory(dir string) []model.Order {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	var orders []model.Order
	var jsonFiles []string

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			jsonFiles = append(jsonFiles, filepath.Join(dir, file.Name()))
		}
	}

	orders = ReadOrdersFromFiles(jsonFiles)

	return orders
}
