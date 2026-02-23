package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUser     string
	SMTPPassword string
	AdminEmail   string
	FromEmail    string
	SiteName     string
}

type ConfigBD struct {
	BD_USER     string
	BD_PASSWORD string
	BD_HOST     string
	BD_PORT     string
	DB_NAME     string
}

func LoadConfig() (*Config, error) {
	godotenv.Load()

	config := &Config{
		SMTPHost:     GetEnv("SMTP_HOST", "smtp.gmail.com"),
		SMTPPort:     getEnvAsInt("SMTP_PORT", 587),
		SMTPUser:     GetEnv("SMTP_USER", ""),
		SMTPPassword: GetEnv("SMTP_PASSWORD", ""),
		AdminEmail:   GetEnv("ADMIN_EMAIL", "admin@example.com"),
		FromEmail:    GetEnv("FROM_EMAIL", "no-reply@example.com"),
		SiteName:     GetEnv("SITE_NAME", "Contact Service"),
	}

	if config.SMTPUser == "" || config.SMTPPassword == "" {
		return nil, fmt.Errorf("SMTP_USER и SMTP_PASSWORD обязательны")
	}

	return config, nil
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		var intValue int
		fmt.Sscanf(value, "%d", &intValue)
		return intValue
	}
	return defaultValue
}

func LoadConfigBD() (*ConfigBD, error) {
	godotenv.Load()

	config := &ConfigBD{
		BD_USER:     GetEnv("BD_USER", "root"),
		BD_PASSWORD: GetEnv("BD_PASSWORD", "1234"),
		BD_HOST:     GetEnv("BD_HOST", "localhost"),
		BD_PORT:     GetEnv("BD_PORT", "3306"),
		DB_NAME:     GetEnv("DB_NAME", ""),
	}
	if config.BD_USER == "" || config.BD_PASSWORD == "" {
		return nil, fmt.Errorf("SMTP_USER и SMTP_PASSWORD обязательны")
	}
	return config, nil

}
