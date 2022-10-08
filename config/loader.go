package config

import (
	"fmt"
	"os"
)

var Conf *Config

func Setup() error {
	c := &Config{}
	err := c.GetFromEnv()
	Conf = c
	return err
}

func (c *Config) GetFromEnv() error {
	c.Host = os.Getenv("HOST")
	if c.Host == "" {
		return envErr("HOST")
	}
	c.Port = os.Getenv("PORT")
	if c.Port == "" {
		return envErr("PORT")
	}
	c.DBUser = os.Getenv("DB_USER")
	if c.DBUser == "" {
		return envErr("DB_USER")
	}
	c.DBPassword = os.Getenv("DB_PASSWORD")
	if c.DBPassword == "" {
		return envErr("DB_PASSWORD")
	}
	c.DBHost = os.Getenv("DB_HOST")
	if c.DBHost == "" {
		return envErr("DB_HOST")
	}
	return nil
}

func envErr(varableName string) error {
	return fmt.Errorf("env variable %s not set or incorrect value", varableName)
}
