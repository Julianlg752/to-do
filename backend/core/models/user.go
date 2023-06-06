package models

type UserInfo struct {
	ID       int64  `json:"userId" db:"id"`
	UserName string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
