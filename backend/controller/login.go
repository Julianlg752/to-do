package controller

import (
	"net/http"
	"todo/core/interactor"
	"todo/core/models"
	"todo/core/ports"

	"github.com/ansel1/merry/v2"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LoginControllerInterface interface {
	Login(c *gin.Context)
}

type loginController struct {
	loginInteractor ports.LoginPortInterface
}

var LoginController LoginControllerInterface = &loginController{
	loginInteractor: interactor.LoginInteractor,
}

func (l *loginController) Login(c *gin.Context) {
	var request *models.LoginRequest
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, merry.Wrap(err))
		return
	}
	response, err := l.loginInteractor.Login(request)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.LoginResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
