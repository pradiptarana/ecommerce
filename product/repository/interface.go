package repository

import (
	productModel "github.com/pradiptarana/product/model/product"
)

//go:generate mockgen -destination=../mocks/mock_product.go -package=mocks github.com/pradiptarana/product/repository ProductRepository
type ProductRepository interface {
	GetProducts(filter *productModel.GetProductFilter) ([]*productModel.Product, error)
	GetLatestProducts() ([]*productModel.Product, error)
}
