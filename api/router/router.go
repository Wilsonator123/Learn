package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Wilsonator123/Learn/handlers"
	"github.com/Wilsonator123/Learn/model"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()

func New(e* echo.Echo ){
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", map[string]interface{}{});
	})

	e.GET("/list", func(c echo.Context) error {
		
		response, err := handlers.ListAll()

		if c.Request().Header.Get("HX-Request") == "true" {
			if err != nil {
			return c.Render(http.StatusOK, "ColumnList", map[string]interface{}{"Error": err})
				
			}
			return c.Render(http.StatusOK, "ColumnList", map[string]interface{}{"Data": response, "Error": err})
		}
		
		return c.Render(http.StatusOK, "index", map[string]interface{}{"Data": response, "Error": err})
	})

	e.POST("/list", func(c echo.Context) error {
		
		priorityStr := c.FormValue("priority")
		priorityInt, err := strconv.ParseInt(priorityStr, 10, 16)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid priority value")
		}
		priorityInt16 := int16(priorityInt)
		newItem := model.NewItem{
			Title: c.FormValue("title"),
			Description: c.FormValue("description"),
			Priority: &priorityInt16,
		}
		
		if err := validate.Struct(newItem); err != nil {
			fmt.Printf("Error parsing form data %v\n", err)
			return c.String(http.StatusBadRequest, "Failed to parse form data")
		}

		response, err := handlers.CreateItem(newItem)

		fmt.Printf("User created with id %v\n", response)

		if err != nil {
			fmt.Printf("Error creating user %v\n", err)

			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.Render(http.StatusCreated, "item", response)
	})

	e.GET("/list/:id", func(c echo.Context) error {

		id := c.Param("id")

		response, err := handlers.GetItem(id)
		
		return c.Render(http.StatusOK, "ItemDetails", map[string]interface{}{"Data": response, "Error": err})
	})

	e.DELETE("/list/:id", func(c echo.Context) error {
		
		id := c.Param("id")

		handlers.DeleteItem(id)

		return c.HTML(http.StatusOK,"<div>Deleted</div>" )

	})
	

}

