package repository

import (
	"gofiber-idp/server/dto"
	"gofiber-idp/server/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Select(id int) (*model.User, error) {
	var user = &model.User{}
	if err := r.db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Insert(userCreateRequest dto.CreateUserRequest) (*model.User, error) {
	user := getUserFromBody(userCreateRequest)

	if err := r.db.Create(user).Error; err != nil {
		return &model.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Update(id int, userUpdateRequest dto.UpdateUserRequest) (*model.User, error) {
	var user model.User
	if result := r.db.First(&user, id); result.Error != nil {
		return &model.User{}, nil
	}

	updateUserFields(&user, userUpdateRequest)

	if err := r.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Delete(id int) error {
	var user model.User

	if result := r.db.First(&user, id); result.Error != nil {
		return result.Error
	}

	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func getUserFromBody(body dto.CreateUserRequest) *model.User {
	user := &model.User{}

	user.Username = body.Username
	user.Name = body.Name
	user.Nickname = body.Nickname
	user.AvatarID = body.AvatarID
	user.PhoneNumber = body.PhoneNumber
	user.StudentYear = body.StudentYear
	user.StudentGroup = body.StudentGroup
	user.MajorCode = body.MajorCode

	return user
}

func updateUserFields(user *model.User, body dto.UpdateUserRequest) {
	user.Username = body.Username
	user.Name = body.Name
	user.Nickname = body.Nickname
	user.AvatarID = body.AvatarID
	user.PhoneNumber = body.PhoneNumber
	user.StudentYear = body.StudentYear
	user.StudentGroup = body.StudentGroup
	user.MajorCode = body.MajorCode
}
