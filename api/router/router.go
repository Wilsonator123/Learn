package router

import (
	"net/http"

	"github.com/Wilsonator123/Learn/handlers"
	"github.com/labstack/echo/v4"
)

func New(e* echo.Echo ){
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", "World")
	})

	e.GET("/list", func(c echo.Context) error {
		
		response := handlers.ListAll()
		
		return c.Render(http.StatusOK, "index", response)
	})

	e.GET("/list:id", func(c echo.Context) error {

		response := handlers.GetItem()

		return c.Render(http.StatusOK, "ItemDetails", response)
	})
	

}

