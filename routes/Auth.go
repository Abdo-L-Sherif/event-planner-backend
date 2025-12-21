package routes

import (
	"go-auth-api/database"
	"go-auth-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("supersecretkey")

func Signup(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input"})
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	user := models.User{Username: input.Username, Email: input.Email, Password: string(hashed)}

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"Error": "User already exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signup sucessful"})
}

func Login(c *gin.Context) {
	var input models.User
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input"})
		return
	}

	result := database.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Email"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokentString, _ := token.SignedString(jwtKey)

	c.JSON(http.StatusOK, gin.H{"token": tokentString})
}
