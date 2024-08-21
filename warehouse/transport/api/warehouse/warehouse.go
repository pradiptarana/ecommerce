package warehouse

import (
	"net/http"

	model "github.com/pradiptarana/warehouse/model/warehouse"
	"github.com/pradiptarana/warehouse/usecase"

	"github.com/gin-gonic/gin"
)

type WarehouseTransport struct {
	usecase.WarehouseUsecase
}

func NewWarehouseTransport(uc usecase.WarehouseUsecase) *WarehouseTransport {
	return &WarehouseTransport{uc}
}

func (ut *WarehouseTransport) TransferStock(c *gin.Context) {
	var req model.TransferStockRequest
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := ut.WarehouseUsecase.TransferStock(req.SourceWarehouseId, req.TargetWarehouseId, req.ProductId, req.Quantity)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Success transfer stock"})
	return
}
