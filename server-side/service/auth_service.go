package service

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"server-side/model"
	"server-side/repository"
	"server-side/response"
	"server-side/utility"
	"strconv"
	"strings"
	"time"
)

func generateToken(userId int64) (string, error) {
	jwtSecret := []byte(viper.GetString("JWT_SECRET"))
	expireDuration := time.Hour * time.Duration(viper.GetInt("TOKEN_EXPIRE_HOUR"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  strconv.FormatInt(userId, 10),
		"exp": time.Now().Add(expireDuration).Unix(),
	})
	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	_, err = repository.TokenRepositoryInstance.CreateToken(userId, jwtToken, expireDuration)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func validateToken(tokenString string) (*jwt.Token, error) {
	jwtSecret := []byte(viper.GetString("JWT_SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func getTokenFromBearerHeader(c *gin.Context) (string, error) {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		return "", errors.New("no Authorization header provided")
	}
	const bearerSchema = "Bearer "
	if !strings.HasPrefix(authorizationHeader, bearerSchema) {
		return "", errors.New("authorization header format must be 'Bearer {token}'")
	}
	token := strings.TrimPrefix(authorizationHeader, bearerSchema)
	if token == "" {
		return "", errors.New("no token found in Authorization header")
	}
	return token, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := getTokenFromBearerHeader(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		token, err := validateToken(tokenString)
		if token == nil || err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userId", claims["id"])
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		_, err = repository.TokenRepositoryInstance.GetTokenByTokenString(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		c.Next()
	}
}

func GetUserInfoFromId(userId int64) (user *model.User, err error) {
	user, err = repository.UserRepositoryInstance.GetUser(userId)
	if err != nil {
		return nil, response.NewErrorResponse("user not found", http.StatusNotFound)
	}
	return user, nil
}

func Login(username string, password string) (jwtToken string, err error) {
	user, err := repository.UserRepositoryInstance.GetUserByUsername(username)
	if err != nil {
		return "", response.NewErrorResponse("invalid credential", http.StatusUnauthorized)
	}
	if user == nil {
		return "", response.NewErrorResponse("invalid credential", http.StatusUnauthorized)
	}
	if utility.CheckPasswordHash(password, user.Password) {
		return generateToken(user.ID)
	} else {
		return "", response.NewErrorResponse("invalid credential", http.StatusUnauthorized)
	}
}

func Logout(token string) (err error) {
	return repository.TokenRepositoryInstance.DeleteToken(token)
}

func SignUp(create model.UserCreate) (user *model.User, err error) {
	existingUser, err := repository.UserRepositoryInstance.GetUserByUsername(create.Username)
	if err != nil {
		return nil, response.NewErrorResponse("unable to check for existing user", http.StatusInternalServerError)
	}
	if existingUser != nil {
		return nil, response.NewErrorResponse("username has been taken already", http.StatusConflict)
	}
	return repository.UserRepositoryInstance.InsertUser(create)
}
