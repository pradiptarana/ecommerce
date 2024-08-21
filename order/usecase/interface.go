package usecase

import (
	"github.com/pradiptarana/order/model/order"
	"github.com/pradiptarana/order/model/product"
)

type ProductUsecase interface {
	GetProducts(filter *product.GetProductFilter) ([]*product.Product, error)
	GetProduct(productId int) (*product.Product, error)
}

type OrderUsecase interface {
	GetCurrentCart(userId int) (*order.Cart, error)
	AddToCart(req *order.Cart) error
	UpdateCart(req *order.Cart) error
	Checkout(userId int) error
	GetOrderHistory(filter *order.GetOrderHistoryFilter) ([]*order.Order, error)
	GetOrderById(id int) (*order.Order, error)
}
