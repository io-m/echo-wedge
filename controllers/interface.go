package controllers

import (
	"github.com/labstack/echo/v4"
)
// ControllerDevice is an interface that defines handler methods
type ControllerDevice interface {
	GetAll(c echo.Context) error
	GetOne(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}