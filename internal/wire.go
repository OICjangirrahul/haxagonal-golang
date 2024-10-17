//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/OICjangirrahul/haxagonal-golang/config"
	"github.com/OICjangirrahul/haxagonal-golang/internal/adapters/controller"
	"github.com/OICjangirrahul/haxagonal-golang/internal/adapters/repository"
	"github.com/OICjangirrahul/haxagonal-golang/internal/application/port/out"
	"github.com/OICjangirrahul/haxagonal-golang/internal/application/service"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	Controllers *controller.Controllers
	DB          *gorm.DB
	RedisClient *redis.Client
}

func InitializeApp() (*App, error) {
	wire.Build(
		NewDatabaseConnection,
		NewValidator,
		NewRedisClient,
		repository.NewAccountRepository,
		service.NewAccountRegistrationService,
		controller.NewAccountController,
		controller.NewAuthController,
		repository.NewAuthRepository,
		service.NewAuthRegistrationService,
		controller.NewControllers,
		wire.Bind(new(out.AuthOutPort), new(*repository.AuthRepository)),
		wire.Struct(new(App), "*"),
	)
	return &App{}, nil
}

func NewDatabaseConnection() (*gorm.DB, error) {
	dsn := config.Cfg.Database.DSN
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func NewValidator() *validator.Validate {
	return validator.New()
}

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Cfg.Redis.Addr,
		Password: config.Cfg.Redis.Password,
		DB:       config.Cfg.Redis.DB,
	})
	return rdb
}
