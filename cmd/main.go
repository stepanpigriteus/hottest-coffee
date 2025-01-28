package main

import (
	"encoding/json"
	"fmt"
	"hot/config"
	"hot/flags"
	"hot/internal/server"
	"hot/models"
	"log/slog"
	"os"
	"path/filepath"
)

func main() {
	config.Addr, config.Dir = flags.Flags()
	config.Logger = config.NewLogger()
	fmt.Println(config.Dir)
	// папка мне запили!
	err := os.Mkdir(config.Dir, 0o777)
	if err != nil {
		if os.IsExist(err) {
			config.Logger.Info("Directory exists", slog.String("directory", config.Dir))
		} else {
			config.Logger.Error("Error creating directory", "directory", config.Dir, "error", err)
			os.Exit(1)
		}
	}
	files := []string{"aggregations.json", "inventory.json", "menu_items.json", "orders.json"}

	for _, r := range files {
		path := filepath.Join(config.Dir, r)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			file, err := os.Create(path)
			if err != nil {
				config.Logger.Error("Error creating file", slog.String("file", r), slog.Any("error", err))
				os.Exit(1)
			}
			defer file.Close()
			if r == "aggregations.json" {
				empty := models.Aggregation{TotalSales: 0, ItemSales: []models.AggregationItem{}}
				byteValue, err := json.MarshalIndent(empty, "", "\t")
				if err != nil {
					print(err.Error())
				}
				err = os.WriteFile(path, byteValue, 0o755)
				if err != nil {
					print(err.Error())
				}
			}
			config.Logger.Info("Created file", slog.String("file", path))
		} else if err != nil {
			config.Logger.Error("Error checking file", slog.String("file", r), slog.Any("error", err))
			os.Exit(1)
		} else {
			config.Logger.Info("File already exists", slog.String("file", path))
		}

	}

	server.Start(config.Addr, config.Dir)
}
