package repository

import (
	"todo/core/models"
	"todo/core/ports"
	"todo/datastore"

	"github.com/ansel1/merry/v2"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/doug-martin/goqu/v9/exp"
)

var userTable = goqu.T("users")

type userRepository struct {
	db        *goqu.Database
	userTable exp.IdentifierExpression
}

var UserRepository ports.UserRepositoryInterface = &userRepository{
	db:        &datastore.GoquDB,
	userTable: userTable,
}

func (u *userRepository) GetUserInfo(loginRequest *models.LoginRequest) (*models.UserInfo, error) {
	var userInfo models.UserInfo
	ok, err := u.db.From(u.userTable).Where(
		u.userTable.Col("password").Eq(loginRequest.Password),
		u.userTable.Col("username").Eq(loginRequest.User),
	).ScanStruct(&userInfo)
	if !ok {
		return nil, merry.New("user not found")
	}
	return &userInfo, merry.Wrap(err)
}
