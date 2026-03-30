package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"backend/config"
	"backend/handler"
	"backend/middleware"
	"backend/repository"
	"backend/route"
	"backend/service"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Run() error {
	cfg := config.Load()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := gorm.Open(postgres.Open(cfg.PostgresURL), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to postgres via gorm: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}
	defer sqlDB.Close()

	if err := sqlDB.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to ping postgres: %w", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       0,
	})
	defer redisClient.Close()

	if err := redisClient.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("failed to connect redis: %w", err)
	}

	userRepo := repository.NewPostgresUserRepository(db)
	if err := userRepo.EnsureSchema(ctx); err != nil {
		return fmt.Errorf("failed to ensure users schema: %w", err)
	}

	userService := service.NewUserService(userRepo)
	healthService := service.NewHealthService(
		func(ctx context.Context) error { return sqlDB.PingContext(ctx) },
		func(ctx context.Context) error { return redisClient.Ping(ctx).Err() },
	)

	systemHandler := handler.NewSystemHandler()
	healthHandler := handler.NewHealthHandler(healthService)
	userHandler := handler.NewUserHandler(userService)

	router := route.NewRouter(systemHandler, healthHandler, userHandler)
	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      middleware.CORS(cfg.CORSAllowedOrigin)(router),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Printf("backend listening on :%s", cfg.Port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}
