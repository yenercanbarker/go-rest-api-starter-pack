package repositories

import (
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/models"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	BaseRepositoryInterface[models.User]
	FindByEmail(email string) (*models.User, error)
}

type UserRepository struct {
	*GenericRepository[models.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		GenericRepository: NewGenericRepository[models.User](db),
	}
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
