package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RouteTest(c echo.Context) error {
	color := c.Param("color")
	weight := c.Param("weight")
	qa := c.QueryParam("qa")
	return c.String(http.StatusOK, color+" "+weight+" "+qa)
}
