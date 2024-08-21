package repository

import (
	productModel "github.com/pradiptarana/warehouse/model/product"
)

//go:generate mockgen -destination=../mocks/mock_warehouse.go -package=mocks github.com/pradiptarana/warehouse/repository WarehouseRepository
type WarehouseRepository interface {
	TransferStock(fromWarehouseID int, toWarehouseID int, productID int, quantity int) error
	ActivateWarehouse(warehouseID int) error
	DeactivateWarehouse(warehouseID int) error
}

//go:generate mockgen -destination=../mocks/mock_product.go -package=mocks github.com/pradiptarana/warehouse/repository ProductRepository
type ProductRepository interface {
	GetProducts(filter *productModel.GetProductFilter) ([]*productModel.Product, error)
	GetProduct(id int) (*productModel.Product, error)
}
