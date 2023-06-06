package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

var config Config

type Config struct {
	DBHost     string `env:"DB_HOST"`
	DBPort     string `env:"DB_PORT,default=3306"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBName     string `env:"DB_NAME,required"`
	Migrate    bool   `env:"MIGRATE,default=false"`
	Secret     string `env:"SECRET,required"`
}

func Dns() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
}

func Secret() []byte {
	return []byte(config.Secret)
}

func MysqlConnectionString() *string {
	conn := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?multiStatements=true", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	return &conn
}

func SetupConfig() error {
	ctx := context.Background()

	if err := envconfig.Process(ctx, &config); err != nil {
		return err
	}
	return nil
}

func C() Config {
	return config
}

func TestConnection() {
	config.DBHost = "localhost"
	config.DBName = "todo"
	config.DBPassword = "todo.123"
	config.DBPort = "3306"
	config.DBUser = "latadigital"
}
