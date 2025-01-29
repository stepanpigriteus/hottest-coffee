package main

import (
	"encoding/json"
	"log/slog"
	"os"
	"path/filepath"

	handler "hot/internal/handlers"
	"hot/internal/pkg/config"
	"hot/internal/pkg/flags"
	"hot/models"
)



func main() {
	config.Addr, config.Dir = flags.Flags()
	config.Logger = config.NewLogger()

	config.Logger.Info("Application is starting...", slog.String("directory", config.Dir))

	err := os.Mkdir(config.Dir, 0o777)
	if err != nil {
		if os.IsExist(err) {
			config.Logger.Info("Directory already exists", slog.String("directory", config.Dir))
		} else {
			config.Logger.Error("Error creating directory", "directory", config.Dir, "error", err)
			os.Exit(1)
		}
	} else {
		config.Logger.Info("Directory created successfully", slog.String("directory", config.Dir))
	}

	files := []string{"aggregations.json", "inventory.json", "menu_items.json", "orders.json"}
	for _, r := range files {
		path := filepath.Join(config.Dir, r)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			file, err := os.Create(path)
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
			if err != nil {
				config.Logger.Error("Error creating file", slog.String("file", r), slog.Any("error", err))
				os.Exit(1)
			}
			file.Close()
			config.Logger.Info("File created", slog.String("file", r))
		} else {
			config.Logger.Info("File already exists", slog.String("file", r))
		}
	}

	config.Logger.Info("All files are ready. Starting the server...")
	config.Logger.Info("Server is running", slog.String("address", config.Dir), slog.Int("port", config.Addr))

	// Start the server
	handler.StartServer()
}
