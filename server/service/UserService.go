package service

import (
	"gofiber-idp/server/dto"
	"gofiber-idp/server/model"
	"gofiber-idp/server/repository"
	"gorm.io/gorm"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{repo: repository.NewUserRepository(db)}
}

func (s *UserService) GetUser(id int) (*model.User, error) {
	return s.repo.Select(id)
}

func (s *UserService) CreateUser(userCreateRequest dto.CreateUserRequest) (*model.User, error) {
	return s.repo.Insert(userCreateRequest)
}

func (s *UserService) UpdateUser(id int, userCreateRequest dto.UpdateUserRequest) (*model.User, error) {
	return s.repo.Update(id, userCreateRequest)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
