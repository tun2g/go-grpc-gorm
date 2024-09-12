package config

import (
	"app/src/lib/logger"

	"github.com/spf13/viper"
)

type Config struct {
	// Application
	AppHost string `mapstructure:"APP_HOST"`
	AppPort int    `mapstructure:"APP_PORT"`
	ReflectionEnabled bool   `mapstructure:"REFLECTION_ENABLED"`

	// Database
	DbName         string `mapstructure:"DATABASE_NAME"`
	DbPort         int    `mapstructure:"DATABASE_PORT"`
	DbHost         string `mapstructure:"DATABASE_HOST"`
	DbDriver       string `mapstructure:"DATABASE_DRIVER"`
	DbUser         string `mapstructure:"DATABASE_USER"`
	DbPassword     string `mapstructure:"DATABASE_PASSWORD"`
	DbMaxOpenConns int    `mapstructure:"DATABASE_MAX_OPEN_CONNS"`
	DbMaxIdleConns int    `mapstructure:"DATABASE_MAX_IDLE_CONNS"`
	DbConnMaxLife  int    `mapstructure:"DATABASE_CONN_MAX_LIFE"`

	// Authentication
	JwtAccessTokenExpirationTime  int    `mapstructure:"JWT_ACCESS_TOKEN_EXPIRATION_TIME"`
	JwtAccessTokenSecret          string `mapstructure:"JWT_ACCESS_TOKEN_SECRET"`
	JwtRefreshTokenExpirationTime int    `mapstructure:"JWT_REFRESH_TOKEN_EXPIRATION_TIME"`
	JwtRefreshTokenSecret         string `mapstructure:"JWT_REFRESH_TOKEN_SECRET"`

	// Redis broker
	RedisBrokerUrl string `mapstructure:"REDIS_BROKER_URL"`
}

var log = logger.NewLogger("AppConfiguration")

func loadConfiguration(path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("Load config failed %s", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Errorf("Load config failed %s", err)
	}

	return config
}

var AppConfiguration = loadConfiguration(".")
