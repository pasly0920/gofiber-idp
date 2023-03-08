package dto

type CreateUserRequest struct {
	Username     string `json:"username" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Nickname     string `json:"nickname" validate:"required"`
	AvatarID     uint64 `json:"avatarId" validate:"required,number"`
	PhoneNumber  string `json:"phoneNumber" validate:"required,phone"`
	StudentYear  int    `json:"studentYear" validate:"required,number"`
	StudentGroup string `json:"studentGroup" validate:"required"`
	MajorCode    int    `json:"majorCode" validate:"required,number"`
}
