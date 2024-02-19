package custommiddleware

import (
	"context"

	"github.com/labstack/echo/v4"
)

func NewMiddlewareContextValue(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(contextValue{c})
	}
}

type contextValue struct {
	echo.Context
}

// type contextkey string

func (c contextValue) Get(key string) interface{} {
	val := c.Context.Get(key)
	if val != nil {
		return val
	}
	return c.Request().Context().Value(key)
}

func (c contextValue) Set(key string, val interface{}) {
	c.SetRequest(
		c.Request().WithContext(
			context.WithValue(
				c.Request().Context(),
				key,
				val,
			),
		),
	)
}
