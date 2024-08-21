package usecase

import (
	"github.com/pradiptarana/warehouse/model/product"
)

type WarehouseUsecase interface {
	TransferStock(fromWarehouseID int, toWarehouseID int, productID int, quantity int) error
}

type ProductUsecase interface {
	GetProducts(filter *product.GetProductFilter) ([]*product.Product, error)
	GetProduct(productId int) (*product.Product, error)
}
