// internal/adapters/repository/auth_repository.go

package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/OICjangirrahul/haxagonal-golang/internal/application/domain"
)

type AuthRepository struct {
	RedisClient *redis.Client
}

// NewAuthRepository creates a new AuthRepository instance
func NewAuthRepository(redisClient *redis.Client) *AuthRepository {
	return &AuthRepository{RedisClient: redisClient}
}

// TokenSave saves the token in Redis
func (r *AuthRepository) TokenSave(token *domain.AuthToken) error {
	ctx := context.Background()
	return r.RedisClient.Set(ctx, token.Token, token.Token, 0).Err()
}
