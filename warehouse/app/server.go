package app

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	auth "github.com/pradiptarana/warehouse/internal/auth"
	// "github.com/pradiptarana/warehouse/internal/cache"
	dbd "github.com/pradiptarana/warehouse/internal/db"
	env "github.com/pradiptarana/warehouse/internal/env"
	warehouseRepo "github.com/pradiptarana/warehouse/repository/warehouse"
	// productRepo "github.com/pradiptarana/warehouse/repository/product"
	warehouseTr "github.com/pradiptarana/warehouse/transport/api/warehouse"
	// productTr "github.com/pradiptarana/warehouse/transport/api/product"
	warehouseUC "github.com/pradiptarana/warehouse/usecase/warehouse"
	// productUC "github.com/pradiptarana/warehouse/usecase/product"
)

func SetupServer() *gin.Engine {
	ctx := context.Background()
	err := env.LoadEnv()
	if err != nil {
		fmt.Println(err)
		log.Fatalf("error when load env file")
	}
	db := dbd.NewDBConnection()
	// myCache := cache.New[int, []byte]()
	// productRepo := productRepo.NewProductRepository(db)
	warehouseRepo := warehouseRepo.NewWarehouseRepository(db)
	// productUC := productUC.NewProductUC(productRepo, *myCache)
	warehouseUC := warehouseUC.NewWarehouseUC(warehouseRepo)
	warehouseTr := warehouseTr.NewWarehouseTransport(warehouseUC)

	router := gin.Default()
	protected := router.Group("/api/v1")
	protected.Use(auth.JwtAuthMiddleware(ctx))
	protected.GET("/warehouse/product/transfer", warehouseTr.TransferStock)

	return router
}
