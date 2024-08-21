package warehouse

type TransferStockRequest struct {
	SourceWarehouseId int `json:"source_id"`
	TargetWarehouseId int `json:"target_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
