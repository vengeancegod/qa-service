package config

import (
	"github.com/joho/godotenv"
)

// type DBConfig struct {
// 	Host     string
// 	Port     string
// 	User     string
// 	Password string
// 	DBName   string
// 	SSLMode  string
// }

// type Config struct {
// 	DB          DBConfig
// 	ServerPort  string
// 	Environment string
// }

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
