package dto

type UpdateUserRequest struct {
	Username     string `json:"username" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Nickname     string `json:"nickname" validate:"required"`
	AvatarID     uint64 `json:"avatarId" validate:"required"`
	PhoneNumber  string `json:"phoneNumber" validate:"required"`
	StudentYear  int    `json:"studentYear" validate:"required"`
	StudentGroup string `json:"studentGroup" validate:"required"`
	MajorCode    int    `json:"majorCode" validate:"required"`
}
