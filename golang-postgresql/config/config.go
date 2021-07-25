package config

import (
	"book-service/src/models"
	"fmt"
	"github.com/spf13/viper"
	"strconv"
)

const (
	ApplicationPort        = "port"
	HostName               = "host_name"
	PostGreSQLHOST         = "postgresql_host"
	PostGreSQLPort         = "postgresql_port"
	PostGreSQLUsername     = "postgresql_username"
	PostGreSQLPassword     = "postgresql_password"
	PostGreSQLDatabase     = "postgresql_database"
	PostGreSQLSSLMode      = "postgresql_sslmode"
	PostGreSQLMaxIdleConns = "postgresql_max_idle_conns"
	PostGreSQLMaxOpenConns = "postgresql_max_open_conns"
)

// Init configurations
func Init() {
	viper.AutomaticEnv()
	viper.AddConfigPath("./config/app/")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error read in config .env")
	}
}

// Database returns database settings
func Database() (*models.Database, error) {
	database := &models.Database{
		Host:         viper.GetString(PostGreSQLHOST),
		Port:         viper.GetInt(PostGreSQLPort),
		Username:     viper.GetString(PostGreSQLUsername),
		Password:     viper.GetString(PostGreSQLPassword),
		Database:     viper.GetString(PostGreSQLDatabase),
		SSLMode:      viper.GetString(PostGreSQLSSLMode),
		MaxIdleConns: viper.GetInt(PostGreSQLMaxIdleConns),
		MaxOpenConns: viper.GetInt(PostGreSQLMaxOpenConns),
	}

	return database, nil
}

func GetAppPort() string {
	return ":" + viper.GetString(ApplicationPort)
}

func GetBookServicePort() int {
	port := viper.GetString(ApplicationPort)
	p, _ := strconv.Atoi(port)
	return p
}

func GetUserHostName() string {
	return viper.GetString(HostName)
}
