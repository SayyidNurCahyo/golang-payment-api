package config

import (
	"fmt"
	"merchant-payment-api/util"
	"os"
)

// struct buat nyimpan konfigurasi database
type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

type APIConfig struct{
	ApiPort string
}

type FileConfig struct{
	FilePath string
}

// embedded struct dari DbConfig, memisahkan logic dari database config
type Config struct {
	DbConfig
	APIConfig
	FileConfig
}

// buat method ReadConfig() punya struct Config = baca informasi konfigurasi dari environment variable
func (c *Config) ReadConfig() error {
	err := util.LoadEnv()
	if err!= nil{
		return err
	}
	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	c.APIConfig = APIConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	c.FileConfig = FileConfig{
		FilePath: os.Getenv("FILE_PATH"),
	}

	if c.DbConfig.Host == "" || c.DbConfig.Driver == "" || c.DbConfig.Name == "" || c.DbConfig.Password == "" || c.DbConfig.Port == "" || c.DbConfig.User == "" || c.APIConfig.ApiPort == "" || c.FileConfig.FilePath == "" {
		return fmt.Errorf("missing required environment variable")
	}

	return nil
}

// constructor buat instance baru dari struct Config
func NewConfig() (*Config, error){
	cfg := &Config{}
	err := cfg.ReadConfig()
	if err != nil {
		return nil,err
	}
	return cfg,nil
}