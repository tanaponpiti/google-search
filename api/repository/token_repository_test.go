package repository_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"server-side/repository"
)

func setupRedis() *redis.Client {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return client
}

func TestTokenRepositoryCreateToken(t *testing.T) {
	redisClient := setupRedis()
	repo := repository.NewTokenRepository(redisClient)

	tokenString := "token123"
	userId := int64(1)
	duration := time.Hour

	token, err := repo.CreateToken(userId, tokenString, duration)
	require.NoError(t, err)
	require.NotNil(t, token)
	assert.Equal(t, userId, token.UserId)
	assert.Equal(t, tokenString, token.Token)

	ctx := context.Background()
	storedUserId, err := redisClient.Get(ctx, tokenString).Result()
	require.NoError(t, err)
	assert.Equal(t, strconv.FormatInt(userId, 10), storedUserId)
}

func TestTokenRepositoryGetTokenByTokenString(t *testing.T) {
	redisClient := setupRedis()
	repo := repository.NewTokenRepository(redisClient)

	tokenString := "token123"
	userId := int64(1)

	ctx := context.Background()
	err := redisClient.Set(ctx, tokenString, userId, time.Hour).Err()
	require.NoError(t, err)

	token, err := repo.GetTokenByTokenString(tokenString)
	require.NoError(t, err)
	require.NotNil(t, token)
	assert.Equal(t, userId, token.UserId)
	assert.Equal(t, tokenString, token.Token)
}

func TestTokenRepositoryDeleteToken(t *testing.T) {
	redisClient := setupRedis()
	repo := repository.NewTokenRepository(redisClient)

	tokenString := "token123"
	userId := int64(1)

	ctx := context.Background()
	err := redisClient.Set(ctx, tokenString, userId, time.Hour).Err()
	require.NoError(t, err)

	err = repo.DeleteToken(tokenString)
	require.NoError(t, err)

	_, err = redisClient.Get(ctx, tokenString).Result()
	assert.Equal(t, redis.Nil, err)
}
