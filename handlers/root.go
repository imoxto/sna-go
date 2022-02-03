package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Root(c echo.Context) error {

	data := map[string]string{
		"message": "Hello World!",
	}

	c.JSON(http.StatusOK, data)

	return nil
}
