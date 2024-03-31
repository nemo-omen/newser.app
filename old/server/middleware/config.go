package middleware

import "github.com/labstack/echo/v4"

type Config struct {
	// DSN  string
	Dev  bool
	Prod bool
	// Logger
}

// func NewConfig(dev bool, dsn string) Config {
func NewConfig(dev bool) Config {
	return Config{Dev: dev, Prod: !dev}
	// return Config{Dev: dev, Prod: !dev, DSN: dsn}
}

func (conf *Config) SetConfig(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("dev", conf.Dev)
		c.Set("prod", conf.Prod)
		// c.Set("dsn", conf.DSN)
		return next(c)
	}
}
