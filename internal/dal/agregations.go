package dal

// import (
// 	"encoding/json"
// 	"hot/internal/pkg/config"
// 	"hot/models"
// 	"io"
// 	"os"
// 	"path/filepath"
// )

// func GetAggregations() (models.Aggregation, error) {
// 	pathAggr := filepath.Join(config.Dir, "aggregations.json")
// 	jsonFileAggr, err := os.Open(pathAggr)
// 	if err != nil {
// 		return models.Aggregation{}, err
// 	}
// 	defer jsonFileAggr.Close()

// 	byteValueAggr, err := io.ReadAll(jsonFileAggr)
// 	if err != nil {
// 		return models.Aggregation{}, err
// 	}

// 	var aggreg models.Aggregation
// 	err = json.Unmarshal(byteValueAggr, &aggreg)
// 	if err != nil {
// 		return models.Aggregation{}, err
// 	}
// 	return aggreg, nil
// }
