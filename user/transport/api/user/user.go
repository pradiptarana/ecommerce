package users

import (
	"net/http"

	model "github.com/pradiptarana/user/model/user"
	"github.com/pradiptarana/user/usecase"

	"github.com/gin-gonic/gin"
)

type UsersTransport struct {
	usecase.UsersUsecase
}

func NewUsersTransport(uc usecase.UsersUsecase) *UsersTransport {
	return &UsersTransport{uc}
}

func (ut *UsersTransport) SignUp(c *gin.Context) {
	var req model.User
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := ut.UsersUsecase.SignUp(&req)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "sign up success"})
	return
}

func (ut *UsersTransport) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}
	token, err := ut.UsersUsecase.Login(&req)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "login success", "token": token})
	return
}
