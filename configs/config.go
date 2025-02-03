package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Redis    RedisConfig
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type JWTConfig struct {
	Secret             string
	AccessTokenExpiry  time.Duration
	RefreshTokenExpiry time.Duration
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	viper.AutomaticEnv()

	_ = viper.BindEnv("database.password", "DB_PASSWORD")
	_ = viper.BindEnv("database.user", "DB_USER")
	_ = viper.BindEnv("redis.password", "REDIS_PASSWORD")
	_ = viper.BindEnv("JWT_SECRET", "JWT_SECRET")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("⚠️ Не удалось загрузить config.yaml: %v. Используем переменные окружения.", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Ошибка при парсинге конфигурации: %v", err)
		return nil, err
	}

	config.Database.Password = viper.GetString("database.password")
	config.Database.User = viper.GetString("database.user")
	config.Redis.Password = viper.GetString("redis.password")
	config.JWT.Secret = viper.GetString("JWT_SECRET")

	return &config, nil
}
