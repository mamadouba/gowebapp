package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host         string
	Port         string
	LogLevel     string
	LogFile      string
	SecretKey    string
	DbType       string
	DbName       string
	DbHost       string
	DbPort       string
	DbUser       string
	DbPass       string
	MailUser     string
	MailPass     string
	MailHost     string
	SendGridKey  string
	SendGridHost string
}

func NewConfig() *Config {
	godotenv.Load(".env")
	return &Config{
		Host:         os.Getenv("HOST"),
		Port:         os.Getenv("PORT"),
		DbType:       os.Getenv("DB_TYPE"),
		DbName:       os.Getenv("DB_NAME"),
		DbHost:       os.Getenv("DB_HOST"),
		DbPort:       os.Getenv("DB_PORT"),
		DbUser:       os.Getenv("DB_USER"),
		DbPass:       os.Getenv("DB_PASS"),
		LogFile:      os.Getenv("LOG_FILE"),
		LogLevel:     os.Getenv("LOG_LEVEL"),
		SecretKey:    os.Getenv("SECRET_KEY"),
		MailUser:     os.Getenv("MAIL_USER"),
		MailPass:     os.Getenv("MAIL_PASS"),
		MailHost:     os.Getenv("MAIL_HOST"),
		SendGridKey:  os.Getenv("SENDGRID_KEY"),
		SendGridHost: os.Getenv("SENDGRID_HOST"),
	}
}

func (c *Config) GetConnectStr() string {
	switch c.DbType {
	case "sqlite":
		return fmt.Sprintf("./%s", c.DbName)
	case "postgres":
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			c.DbHost,
			c.DbPort,
			c.DbUser,
			c.DbPass,
			c.DbName,
		)
	default:
		panic("Unknow database type")
	}
}

var (
	Configuration *Config
)

func init() {
	Configuration = NewConfig()
}
