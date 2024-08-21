package usecase

import (
	"github.com/pradiptarana/product/model/product"
)

type ProductUsecase interface {
	GetProducts() ([]*product.Product, error)
	GetProduct(productId int) (*product.Product, error)
}
