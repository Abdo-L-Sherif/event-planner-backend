package routes

import (
	"go-auth-api/database"
	"go-auth-api/models"
	"net/http"
<<<<<<< HEAD
	"os"
=======
>>>>>>> 15fa3d9ca933de6a2f9567693bff307b392c5d8c
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

<<<<<<< HEAD
var jwtKey = []byte(getJWTSecret())

func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// Fallback for development, but log warning
		secret = "supersecretkey"
	}
	return secret
}
=======
var jwtKey = []byte("supersecretkey")
>>>>>>> 15fa3d9ca933de6a2f9567693bff307b392c5d8c

func Signup(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
<<<<<<< HEAD
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Basic input validation
	if input.Email == "" || input.Password == "" || input.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email, username, and password are required"})
		return
	}

	if len(input.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 6 characters long"})
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
=======
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input"})
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
>>>>>>> 15fa3d9ca933de6a2f9567693bff307b392c5d8c
	user := models.User{Username: input.Username, Email: input.Email, Password: string(hashed)}

	result := database.DB.Create(&user)
	if result.Error != nil {
<<<<<<< HEAD
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signup successful"})
=======
		c.JSON(http.StatusConflict, gin.H{"Error": "User already exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signup sucessful"})
>>>>>>> 15fa3d9ca933de6a2f9567693bff307b392c5d8c
}

func Login(c *gin.Context) {
	var input models.User
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
<<<<<<< HEAD
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Basic input validation
	if input.Email == "" || input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
=======
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input"})
>>>>>>> 15fa3d9ca933de6a2f9567693bff307b392c5d8c
		return
	}

	result := database.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
<<<<<<< HEAD
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
=======
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Email"})
>>>>>>> 15fa3d9ca933de6a2f9567693bff307b392c5d8c
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
<<<<<<< HEAD
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
=======
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Password"})
>>>>>>> 15fa3d9ca933de6a2f9567693bff307b392c5d8c
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

<<<<<<< HEAD
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
=======
	tokentString, _ := token.SignedString(jwtKey)

	c.JSON(http.StatusOK, gin.H{"token": tokentString})
>>>>>>> 15fa3d9ca933de6a2f9567693bff307b392c5d8c
}
