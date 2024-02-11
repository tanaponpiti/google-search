package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"os"
	"server-side/model"
	"server-side/repository"
	"server-side/repository/mock"
	"server-side/response"
	"server-side/utility"
	"strconv"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	err := os.Setenv("JWT_SECRET", "ji3ij1nm109aspa")
	err = os.Setenv("TOKEN_EXPIRE_HOUR", "24")
	assert.Nil(t, err)
	defer os.Unsetenv("JWT_SECRET")
	userID := int64(1)
	mockTokenRepo := new(repository_test.TokenRepositoryMock)
	repository.TokenRepositoryInstance = mockTokenRepo

	mockTokenRepo.On("CreateToken", userID, mock.AnythingOfType("string"), mock.AnythingOfType("time.Duration")).Return(&model.Token{}, nil)

	token, err := generateToken(userID)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	tokenData, err := validateToken(token)
	assert.Nil(t, err)
	claims, ok := tokenData.Claims.(jwt.MapClaims)
	assert.Equal(t, true, ok)
	err = claims.Valid()
	assert.Nil(t, err)
	assert.Equal(t, strconv.FormatInt(userID, 10), claims["id"])

	mockTokenRepo.AssertExpectations(t)
}

func TestGetUserInfoFromId(t *testing.T) {
	// Setup
	userID := int64(1)
	mockUserRepo := new(repository_test.UserRepositoryMock)
	repository.UserRepositoryInstance = mockUserRepo

	// Expectations
	mockUserRepo.On("GetUser", userID).Return(&model.User{ID: userID}, nil)

	// Test
	user, err := GetUserInfoFromId(userID)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userID, user.ID)
	mockUserRepo.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	err := os.Setenv("JWT_SECRET", "ji3ij1nm109aspa")
	assert.Nil(t, err)
	defer os.Unsetenv("JWT_SECRET")
	username := "testUser"
	password := "testPassword"
	mockTokenRepo := new(repository_test.TokenRepositoryMock)
	repository.TokenRepositoryInstance = mockTokenRepo
	mockUserRepo := new(repository_test.UserRepositoryMock)
	repository.UserRepositoryInstance = mockUserRepo
	userID := int64(1)
	hashedPassword, err := utility.HashPassword(password)
	mockUser := model.User{ID: userID, Username: "testUser", Password: hashedPassword}
	assert.Nil(t, err)

	mockUserRepo.On("GetUserByUsername", username).Return(&mockUser, nil)
	mockTokenRepo.On("CreateToken", userID, mock.AnythingOfType("string"), mock.AnythingOfType("time.Duration")).Return(&model.Token{}, nil)

	// Test
	token, err := Login(username, password)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	// Test wrong password
	token, err = Login(username, "wrong_password")
	assert.Empty(t, token)
	assert.NotNil(t, err)

	mockUserRepo.AssertExpectations(t)
}

func TestAuthMiddleware(t *testing.T) {
	err := os.Setenv("JWT_SECRET", "ji3ij1nm109aspa")
	assert.Nil(t, err)
	defer os.Unsetenv("JWT_SECRET")

	r := gin.New()
	r.Use(AuthMiddleware())
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"userId": c.MustGet("userId")})
	})

	// Test case: Valid token
	t.Run("Valid Token", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		// Assuming you have a function to generate a valid token for testing
		userID := int64(1)
		mockTokenRepo := new(repository_test.TokenRepositoryMock)
		repository.TokenRepositoryInstance = mockTokenRepo
		mockTokenRepo.On("CreateToken", userID, mock.AnythingOfType("string"), mock.AnythingOfType("time.Duration")).Return(&model.Token{}, nil)
		token, err := generateToken(userID)
		mockTokenRepo.On("GetTokenByTokenString", token).Return(&model.Token{Token: token}, nil)

		assert.Nil(t, err)
		req.Header.Set("Authorization", "Bearer "+token)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	// Test case: Invalid token
	t.Run("Invalid Token", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer invalidToken")

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	// Test case: Missing token
	t.Run("Missing Token", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}

func TestLogout(t *testing.T) {
	// Mock the token repository
	mockTokenRepo := new(repository_test.TokenRepositoryMock)
	repository.TokenRepositoryInstance = mockTokenRepo

	t.Run("Successful Logout", func(t *testing.T) {
		// Setup expectations
		testToken := "testToken"
		mockTokenRepo.On("DeleteToken", testToken).Return(nil)

		// Call the service function
		err1 := Logout(testToken)

		// Assert expectations
		assert.Nil(t, err1)
		mockTokenRepo.AssertExpectations(t)
	})
}

func TestSignUp(t *testing.T) {
	username := "newUser"
	password := "newPassword"

	t.Run("Successful SignUp", func(t *testing.T) {
		mockUserRepo := new(repository_test.UserRepositoryMock)
		repository.UserRepositoryInstance = mockUserRepo
		mockUserRepo.On("GetUserByUsername", username).Return(nil, nil) // No existing user
		mockUserRepo.On("InsertUser", mock.AnythingOfType("model.UserCreate")).Return(&model.User{Username: username}, nil)

		user, err := SignUp(model.UserCreate{Username: username, Password: password})

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, username, user.Username)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("User Already Exists", func(t *testing.T) {
		mockUserRepo := new(repository_test.UserRepositoryMock)
		repository.UserRepositoryInstance = mockUserRepo
		mockUserRepo.On("GetUserByUsername", username).Return(&model.User{}, nil)

		_, err := SignUp(model.UserCreate{Username: username, Password: password})

		assert.NotNil(t, err)
		assert.Equal(t, "status 409: username has been taken already", err.Error())
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Error Checking Existing User", func(t *testing.T) {
		mockUserRepo := new(repository_test.UserRepositoryMock)
		repository.UserRepositoryInstance = mockUserRepo
		mockUserRepo.On("GetUserByUsername", username).Return(nil, response.NewErrorResponse("database error", http.StatusInternalServerError))

		_, err := SignUp(model.UserCreate{Username: username, Password: password})

		assert.NotNil(t, err)
		assert.Equal(t, "status 500: unable to check for existing user", err.Error())
		mockUserRepo.AssertExpectations(t)
	})
}
