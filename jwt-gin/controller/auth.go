package controller

import (
	"fmt"
	"jwt-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(rtx *gin.Context) {
	var input RegisterInput
	if err := rtx.ShouldBindJSON(&input); err != nil {
		rtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	_, err := u.CreateUser()

	if err != nil {
		rtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rtx.JSON(http.StatusOK, gin.H{"message": "register success"})
}

func Login(rtx *gin.Context) {
	var input LoginInput
	if err := rtx.ShouldBindJSON(&input); err != nil {
		rtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password
	token, err := models.CheckLogin(u.Username, u.Password)
	fmt.Println("token: ", token)
	if err != nil {
		rtx.JSON(http.StatusBadRequest, gin.H{"Username or Password incorrect. ": err.Error()})
		return
	}

	rtx.JSON(http.StatusOK, gin.H{"token": token})
}
