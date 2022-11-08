package controllers

import (
	"net/http"

	"github.com/baguseka01/golang-jwt-authentication/database"
	auth "github.com/baguseka01/golang-jwt-authentication/middlewares"
	"github.com/baguseka01/golang-jwt-authentication/models"
	"github.com/gin-gonic/gin"
)

func Register(g *gin.Context) {
	var user models.User
	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		g.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		g.Abort()
		return
	}
	record := database.DB.Create(&user)
	if record.Error != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		g.Abort()
		return
	}
	g.JSON(http.StatusCreated, gin.H{"message": "Register is successfully", "user": user})
}

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(g *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := g.ShouldBindJSON(&request); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		g.Abort()
		return
	}
	// check if email exists and password is correct
	record := database.DB.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		g.Abort()
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		g.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		g.Abort()
		return
	}
	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		g.Abort()
		return
	}
	g.JSON(http.StatusOK, gin.H{"token": tokenString})
}
