package custommiddleware

import "github.com/labstack/echo/v4"

type Mode struct {
	Dev  bool
	Prod bool
}

func NewMode(mode string) *Mode {
	if mode == "development" {
		return &Mode{Dev: true, Prod: false}
	}
	return &Mode{Dev: false, Prod: true}
}

func (m *Mode) SetMode(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("dev", m.Dev)
		c.Set("prod", m.Prod)
		return next(c)
	}
}
