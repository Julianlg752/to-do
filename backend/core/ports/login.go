package ports

import "todo/core/models"

type LoginPortInterface interface {
	Login(loginRequest *models.LoginRequest) (*models.LoginResponse, error)
}

type UserRepositoryInterface interface {
	GetUserInfo(loginRequest *models.LoginRequest) (*models.UserInfo, error)
}
