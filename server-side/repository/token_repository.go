package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"server-side/database"
	"server-side/model"
	"strconv"
	"time"
)

type ITokenRepository interface {
	GetTokenByTokenString(tokenString string) (*model.Token, error)
	CreateToken(userId int64, token string, duration time.Duration) (*model.Token, error)
	DeleteToken(tokenString string) error
}

var TokenRepositoryInstance ITokenRepository

type TokenRepository struct {
	client *redis.Client
}

func NewTokenRepository(client *redis.Client) *TokenRepository {
	return &TokenRepository{client: client}
}

func InitTokenRepository() {
	TokenRepositoryInstance = NewTokenRepository(database.RedisClient)
}

func (r *TokenRepository) CreateToken(userId int64, token string, duration time.Duration) (*model.Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := r.client.SetNX(ctx, token, userId, duration).Err()
	return &model.Token{UserId: userId, Token: token}, err
}

func (r *TokenRepository) GetTokenByTokenString(token string) (*model.Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	userId, err := r.client.Get(ctx, token).Result()
	if err != nil {
		return nil, err
	}
	userIdNum, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, err
	}
	return &model.Token{UserId: userIdNum, Token: token}, nil
}

func (r *TokenRepository) DeleteToken(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return r.client.Del(ctx, token).Err()
}
