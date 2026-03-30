package repository

import (
	"context"

	"backend/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id int64) (*entity.User, error)
	List(ctx context.Context, limit, offset int) ([]entity.User, error)
}

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) EnsureSchema(ctx context.Context) error {
	return r.db.WithContext(ctx).AutoMigrate(&entity.User{})
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *PostgresUserRepository) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) List(ctx context.Context, limit, offset int) ([]entity.User, error) {
	var users []entity.User
	err := r.db.WithContext(ctx).Order("id DESC").Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
