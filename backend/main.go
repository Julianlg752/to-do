package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo/config"
	"todo/controller"
	"todo/core/models"
	"todo/datastore"
	"todo/router"

	"github.com/doug-martin/goqu/v9"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	if err := config.SetupConfig(); err != nil {
		panic(err)
	}

	if err := datastore.Connection(); err != nil {
		panic(err)
	}

	if config.C().Migrate {
		datastore.RunMigration()
	}

	TestDatabase()
	CreateDemoUsers()
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.Use(cors.New(
		cors.Config{
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			AllowCredentials: false,
			MaxAge:           12 * time.Hour,
			AllowAllOrigins:  true,
		},
	))
	controller := controller.Controller{
		Task:  controller.TaskController,
		Login: controller.LoginController,
	}

	r = router.NewRouter(r, controller)

	go func() {
		if err := r.Run(); err != nil {
			logrus.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

}

func TestDatabase() {
	raw := datastore.DB.QueryRow("SELECT 1")
	var val int
	if err := raw.Scan(&val); err != nil {
		panic(err)
	}
	if val != 1 {
		panic("query error")
	}
}

func CreateDemoUsers() {
	password, err := bcrypt.GenerateFromPassword([]byte(config.C().UserPassword), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	user := &models.UserInfo{
		UserName: "admin",
		Password: string(password),
	}
	if _, err := datastore.GoquDB.Insert(goqu.T("users")).Rows(user).Executor().Exec(); err != nil {
		panic(err)
	}
}
