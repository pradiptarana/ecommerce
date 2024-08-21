package warehouse

import (
	"github.com/pradiptarana/warehouse/repository"
)

type WarehouseUC struct {
	repository.WarehouseRepository
}

func NewWarehouseUC(repo repository.WarehouseRepository) *WarehouseUC {
	return &WarehouseUC{repo}
}

func (uc *WarehouseUC) TransferStock(fromWarehouseID int, toWarehouseID int, productID int, quantity int) (error) {
	err := uc.WarehouseRepository.TransferStock(fromWarehouseID, toWarehouseID, productID, quantity)
	if err != nil {
		return err
	}
	return nil
}
