package app

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	auth "github.com/pradiptarana/product/internal/auth"
	"github.com/pradiptarana/product/internal/cache"
	dbd "github.com/pradiptarana/product/internal/db"
	env "github.com/pradiptarana/product/internal/env"
	productRepo "github.com/pradiptarana/product/repository/product"
	productTr "github.com/pradiptarana/product/transport/api/product"
	productUC "github.com/pradiptarana/product/usecase/product"
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
	productUC := productUC.NewProductUC(productRepo, *myCache)
	productTr := productTr.NewProductTransport(productUC)
	router := gin.Default()

	protected := router.Group("/api/v1")
	protected.Use(auth.JwtAuthMiddleware(ctx))
	protected.GET("/product", productTr.GetProducts)
	protected.GET("/product/:id", productTr.GetProduct)

	return router
}
