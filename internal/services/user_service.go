package services

import (
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/models"
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/repositories"
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/requests"
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/responses"
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/utils"
)

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	CreateUser(userRequest *requests.UserCreateRequest, userResponse *responses.UserResponse) error
	UpdateUser(userRequest *requests.UserUpdateRequest, userResponse *responses.UserResponse, id uint) error
	DeleteUser(id uint) error
}

type UserService struct {
	repo repositories.UserRepositoryInterface
}

func NewUserService(repo repositories.UserRepositoryInterface) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) CreateUser(userRequest *requests.UserCreateRequest, userResponse *responses.UserResponse) error {
	var createdUser models.User
	err := utils.AutoMap[*requests.UserCreateRequest, models.User](userRequest, &createdUser)
	if err != nil {
		return err
	}

	err = s.repo.Create(&createdUser)
	if err != nil {
		return err
	}

	err = utils.AutoMap[*models.User, *responses.UserResponse](&createdUser, &userResponse)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) UpdateUser(userRequest *requests.UserUpdateRequest, userResponse *responses.UserResponse, id uint) error {
	var updatedUser models.User
	err := utils.AutoMap[*requests.UserUpdateRequest, models.User](userRequest, &updatedUser)
	if err != nil {
		return err
	}
	updatedUser.ID = id

	err = utils.AutoMap[*models.User, *responses.UserResponse](&updatedUser, &userResponse)
	if err != nil {
		err.Error()
	}

	return s.repo.Update(&updatedUser)
}

func (s *UserService) DeleteUser(id uint) error {
	user, err := s.GetUserByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(user)
}
