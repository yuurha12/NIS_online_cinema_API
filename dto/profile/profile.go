package profiledto

import "server/models"

type CreateProfile struct {
	Phone  string             `json:"phone" form:"phone" validate:"required"`
	Image  string             `json:"image" form:"id" validate:"required"`
	UserID int                `json:"user_id"`
	User   models.UserProfile `json:"user"`
}

type UpdateProfile struct {
	Phone string `json:"phone" form:"phone"`
	Image string `json:"image" form:"image"`
}

type ProfileResponse struct {
	Phone  string             `json:"phone" form:"phone"`
	Image  string             `json:"image" form:"image"`
	UserID int                `json:"user_id"`
	User   models.UserProfile `json:"user"`
}
