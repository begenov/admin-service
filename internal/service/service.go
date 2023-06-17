package service

import (
	"context"

	"github.com/begenov/admin-service/internal/config"
	"github.com/begenov/admin-service/internal/domain"
	"github.com/begenov/admin-service/internal/repository"
	"github.com/begenov/admin-service/pkg/auth"
	"github.com/begenov/admin-service/pkg/hash"
	"github.com/begenov/admin-service/pkg/kafka"
)

type Admin interface {
	SignUp(ctx context.Context, admin domain.Admin) error
	SignIn(ctx context.Context, email string, password string) (domain.Token, error)
	GetByRefreshToken(ctx context.Context, refreshToken string) (domain.Token, error)
}

type Kafka interface {
	SendMessages(topic string, message string) error
	ConsumeMessages(topic string, handler func(message string)) error
}

type Service struct {
	Admin *AdminService
	Kafka Kafka
}

func NewService(repo *repository.Repository, hash hash.PasswordHasher, tokenManager auth.TokenManager, cfg *config.Config, producer *kafka.Producer, consumer *kafka.Consumer) *Service {
	return &Service{
		Admin: NewAdminService(repo.Admin, hash, tokenManager, cfg.JWT.AccessTokenTTL, cfg.JWT.RefreshTokenTTL),
		Kafka: NewKafkaSerivce(producer, consumer),
	}
}
