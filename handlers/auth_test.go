package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"products/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"golang.org/x/crypto/bcrypt"
)

func TestRegisterHandler(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("successful registration", func(mt *mtest.T) {
		router := gin.Default()
		router.POST("/register", RegisterHandler(mt.Client))

		user := models.User{
			Username: "testuser",
			Password: "testpass",
		}
		jsonValue, _ := json.Marshal(user)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		var response models.ResponseUser
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response.Status)
		assert.Equal(t, "User registered successfully", response.Message)
	})

	mt.Run("registration with invalid data", func(mt *mtest.T) {
		router := gin.Default()
		router.POST("/register", RegisterHandler(mt.Client))

		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(`{"username":""}`)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var response models.ResponseUser
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "error", response.Status)
	})
}

func TestLoginHandler(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	jwtSecret := []byte("testsecret")

	mt.Run("successful login", func(mt *mtest.T) {
		router := gin.Default()
		router.POST("/login", LoginHandler(mt.Client, jwtSecret))

		password, _ := bcrypt.GenerateFromPassword([]byte("testpass"), bcrypt.DefaultCost)
		user := models.User{
			Username: "testuser",
			Password: string(password),
		}
		mt.AddMockResponses(mtest.CreateSuccessResponse(bson.E{}))
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "products.users", mtest.FirstBatch, bson.D{
			{Key: "username", Value: user.Username},
			{Key: "password", Value: user.Password},
		}))

		loginData := map[string]string{
			"username": "testuser",
			"password": "testpass",
		}
		jsonValue, _ := json.Marshal(loginData)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var response models.ResponseUser
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response.Status)
		assert.NotEmpty(t, response.Token)

		token, err := jwt.Parse(response.Token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})
		assert.NoError(t, err)
		assert.True(t, token.Valid)
	})

	mt.Run("login with invalid credentials", func(mt *mtest.T) {
		router := gin.Default()
		router.POST("/login", LoginHandler(mt.Client, jwtSecret))

		loginData := map[string]string{
			"username": "wronguser",
			"password": "wrongpass",
		}
		jsonValue, _ := json.Marshal(loginData)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		var response models.ResponseUser
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "error", response.Status)
		assert.Equal(t, "Invalid username or password", response.Message)
	})
}
