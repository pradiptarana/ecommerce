package app

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	auth "github.com/pradiptarana/order/internal/auth"
	"github.com/pradiptarana/order/internal/cache"
	dbd "github.com/pradiptarana/order/internal/db"
	env "github.com/pradiptarana/order/internal/env"
	orderRepo "github.com/pradiptarana/order/repository/order"
	productRepo "github.com/pradiptarana/order/repository/product"
	orderTr "github.com/pradiptarana/order/transport/api/order"
	productTr "github.com/pradiptarana/order/transport/api/product"
	orderUC "github.com/pradiptarana/order/usecase/order"
	productUC "github.com/pradiptarana/order/usecase/product"
)

func SetupServer() *gin.Engine {
	ctx := context.Background()
	err := env.LoadEnv()
	if err != nil {
		fmt.Println(err)
		log.Fatalf("error when load env file")
	}
	db := dbd.NewDBConnection()
	myCache := cache.New[int, []byte]()
	productRepo := productRepo.NewProductRepository(db)
	orderRepo := orderRepo.NewOrderRepository(db)
	productUC := productUC.NewProductUC(productRepo, *myCache)
	orderUC := orderUC.NewOrderUC(orderRepo, productUC)
	productTr := productTr.NewProductTransport(productUC)
	orderTr := orderTr.NewOrderTransport(orderUC)
	router := gin.Default()

	protected := router.Group("/api/v1")
	protected.Use(auth.JwtAuthMiddleware(ctx))
	protected.GET("/product", productTr.GetProducts)
	protected.GET("/product/:id", productTr.GetProduct)
	protected.POST("/order/cart", orderTr.AddToCart)
	protected.GET("/order/cart", orderTr.GetCurrentCart)
	protected.PUT("/order/cart/:id", orderTr.UpdateCart)
	protected.POST("/order/checkout", orderTr.Checkout)
	protected.GET("/order/history", orderTr.GetOrderHistory)
	protected.GET("/order/:id", orderTr.GetOrderById)

	return router
}
