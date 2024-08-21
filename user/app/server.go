package app

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	auth "github.com/pradiptarana/user/internal/auth"
	dbd "github.com/pradiptarana/user/internal/db"
	env "github.com/pradiptarana/user/internal/env"
	usersRepo "github.com/pradiptarana/user/repository/user"
	usersTr "github.com/pradiptarana/user/transport/api/user"
	usersUC "github.com/pradiptarana/user/usecase/user"
)

func SetupServer() *gin.Engine {
	ctx := context.Background()
	err := env.LoadEnv()
	if err != nil {
		fmt.Println(err)
		log.Fatalf("error when load env file")
	}
	db := dbd.NewDBConnection()
	userRepo := usersRepo.NewUserRepository(db)
	userUC := usersUC.NewUserUC(userRepo)
	userTr := usersTr.NewUsersTransport(userUC)
	router := gin.Default()
	r := router.Group("/api/v1")
	r.POST("/signup", userTr.SignUp)
	r.POST("/login", userTr.Login)

	protected := router.Group("/api/v1")
	protected.Use(auth.JwtAuthMiddleware(ctx))
	// protected.GET("/profile", userTr.GetProfile)

	return router
}
