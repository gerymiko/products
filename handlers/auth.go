package handlers

import (
	"context"
	"net/http"
	"time"

	"products/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// RegisterHandler handles user registration
// @Summary      Register
// @Description  Register a new user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User Data"
// @Success      201   	body			models.ResponseUser
// @Failure      400   	body			models.ResponseUser
// @Router       /register [post]
func RegisterHandler(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseUser{
				Status:  "error",
				Message: err.Error(),
			})
			return
		}

		hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hash)

		collection := db.Database("products").Collection("users")
		_, err := collection.InsertOne(context.Background(), user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ResponseUser{
				Status:  "error",
				Message: "Could not register user",
			})
			return
		}

		c.JSON(http.StatusCreated, models.ResponseUser{
			Status:  "success",
			Message: "User registered successfully",
		})
	}
}

// LoginHandler handles user login requests
// @Summary      Login
// @Description  Login a user and return a JWT token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        loginData  body      models.User  true  "Login Data"
// @Success      200        {object}  models.ResponseUser
// @Failure      400        {object}  models.ResponseUser
// @Failure      401        {object}  models.ResponseUser
// @Router       /login [post]
func LoginHandler(db *mongo.Client, jwtSecret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginData struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseUser{
				Status:  "error",
				Message: err.Error(),
			})
			return
		}

		collection := db.Database("products").Collection("users")
		var user models.User
		err := collection.FindOne(context.Background(), bson.M{"username": loginData.Username}).Decode(&user)
		if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)) != nil {
			c.JSON(http.StatusUnauthorized, models.ResponseUser{
				Status:  "error",
				Message: "Invalid username or password",
			})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
			"exp":      time.Now().Add(time.Hour * 2).Unix(),
		})
		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ResponseUser{
				Status:  "error",
				Message: "Could not generate token",
			})
			return
		}

		c.JSON(http.StatusOK, models.ResponseUser{
			Status:  "success",
			Message: "Login successful",
			Token:   tokenString,
		})
	}
}
