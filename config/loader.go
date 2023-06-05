package config

import (
	"fmt"
	"os"
	"strconv"
)

var Conf *Config

func Setup() error {
	c := &Config{}
	err := c.GetFromEnv()
	Conf = c
	return err
}

func (c *Config) GetFromEnv() error {
	var err error
	c.IsDev = os.Getenv("DEV") == "true"
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
	c.JwtSecret = os.Getenv("JWT_SECRET")
	if c.JwtSecret == "" {
		return envErr("JWT_SECRET")
	}
	jwtatime, err := strconv.Atoi(os.Getenv("JWT_ACCESS_EXPIRE_TIME"))
	c.JwtAccessExpireTime = uint(jwtatime)
	if c.JwtAccessExpireTime == 0 || err != nil {
		return envErr("JWT_ACCESS_EXPIRE_TIME")
	}
	jwtrtime, err := strconv.Atoi(os.Getenv("JWT_REFRESH_EXPIRE_TIME"))
	c.JwtRefreshExpireTime = uint(jwtrtime)
	if c.JwtRefreshExpireTime == 0 || err != nil {
		return envErr("JWT_REFRESH_EXPIRE_TIME")
	}
	defaultpoints, err := strconv.Atoi(os.Getenv("DEFAULT_POINTS"))
	c.DefaultPoints = uint(defaultpoints)
	if c.JwtAccessExpireTime == 0 || err != nil {
		return envErr("DEFAULT_POINTS")
	}
	return nil
}

func envErr(varableName string) error {
	return fmt.Errorf("env variable %s not set or incorrect value", varableName)
}
