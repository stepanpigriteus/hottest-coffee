package models

type Aggregation struct {
	TotalSales float64           `json:"total_sales"`
	ItemSales  []AggregationItem `json:"item_sales"`
}

type AggregationItem struct {
	ID        string `json:"product_id"`
	QuantSold int    `json:"quantity_sold"`
}
