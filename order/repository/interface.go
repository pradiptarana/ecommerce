package repository

import (
	orderModel "github.com/pradiptarana/order/model/order"
	productModel "github.com/pradiptarana/order/model/product"
)

//go:generate mockgen -destination=../mocks/mock_product.go -package=mocks github.com/pradiptarana/order/repository ProductRepository
type ProductRepository interface {
	GetProducts(filter *productModel.GetProductFilter) ([]*productModel.Product, error)
	GetProduct(id int) (*productModel.Product, error)
}

//go:generate mockgen -destination=../mocks/mock_order.go -package=mocks github.com/pradiptarana/order/repository OrderRepository
type OrderRepository interface {
	AddToCart(cart *orderModel.Cart) error
	GetCart(cartId int) (*orderModel.Cart, error)
	UpdateCart(cart *orderModel.Cart) error
	CreateOrder(order *orderModel.Order) error
	GetOrderHistory(filter *orderModel.GetOrderHistoryFilter) ([]*orderModel.Order, error)
	GetOrderById(id int) (*orderModel.Order, error)
	GetCurrentCart(userId int) (*orderModel.Cart, error)
	Checkout(productID int, quantity int, userID int) error
}
