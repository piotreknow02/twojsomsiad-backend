package config

type Config struct {
	Host                 string
	Port                 string
	DBUser               string
	DBPassword           string
	DBHost               string
	DBPort               string
	JwtSecret            string
	JwtAccessExpireTime  uint // in minutes
	JwtRefreshExpireTime uint // in minutes
	DefaultPoints        uint
}
