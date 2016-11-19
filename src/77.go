package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Users struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
}

type Response struct {
	Result []Users
}

func listUsers(c echo.Context) error {
	pat := Users{
		"THAWATCHAI",
		"SINGNGAM",
		"P@ssw0rd",
	}
	jerd := Users{
		"BANJERD",
		"PRATUMHOM",
		"123456",
	}
	result := []Users{pat, jerd}
	return c.JSON(http.StatusOK, result)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/api/users", listUsers)
	e.Logger.Fatal(e.Start(":8888"))
}
