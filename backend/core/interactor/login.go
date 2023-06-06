package interactor

import (
	"time"
	"todo/config"
	"todo/core/models"
	"todo/core/ports"
	"todo/core/repository"

	"github.com/ansel1/merry/v2"
	"github.com/dgrijalva/jwt-go"
)

type Login struct {
	userRepository ports.UserRepositoryInterface
}

var LoginInteractor ports.LoginPortInterface = &Login{
	userRepository: repository.UserRepository,
}

func (c *Login) Login(loginRequest *models.LoginRequest) (*models.LoginResponse, error) {
	if err := loginRequest.Validate(); err != nil {
		return nil, merry.Wrap(err)
	}
	userInfo, err := c.userRepository.GetUserInfo(loginRequest)
	if err != nil {
		return nil, merry.Wrap(err)
	}
	token, err := c.generateToken(userInfo)
	if err != nil {
		return nil, merry.Wrap(err)
	}

	return &models.LoginResponse{Token: *token}, nil
}

func (c *Login) generateToken(userInfo *models.UserInfo) (*string, error) {
	claims := jwt.MapClaims{
		"userId":   userInfo.ID,
		"username": userInfo.UserName,
		"exp":      time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(config.Secret())
	if err != nil {
		return nil, merry.Wrap(err)
	}
	return &signedToken, nil
}
