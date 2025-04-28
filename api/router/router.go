package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Wilsonator123/Learn/handlers"
	"github.com/Wilsonator123/Learn/model"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()

func New(e* echo.Echo ){
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", "World")
	})

	e.GET("/list", func(c echo.Context) error {
		
		response, err := handlers.ListAll()
		
		return c.Render(http.StatusOK, "index", map[string]interface{}{"Data": response, "Error": err})
	})

	e.POST("/list", func(c echo.Context) error {

		var newItem model.NewItem
		
		if err := json.NewDecoder(c.Request().Body).Decode(&newItem); err != nil {
			fmt.Printf("Error parsing body %v\n", err)
			return err
		}

		if err := validate.Struct(newItem); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			fmt.Printf("Error parsing body %v\n", err)
			return c.Render(http.StatusBadRequest, "error", fmt.Sprintf("Validation failed: %v", validationErrors))
		}


		response, err := handlers.CreateItem(newItem)

		fmt.Printf("User created with id %v\n", response)

		if err != nil {
			fmt.Printf("Error creating user %v\n", err)

			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.Redirect(http.StatusSeeOther, "/list/" + response)
	})

	e.GET("/list/:id", func(c echo.Context) error {

		id := c.Param("id")

		response, err := handlers.GetItem(id)
		
		return c.Render(http.StatusOK, "ItemDetails", map[string]interface{}{"Data": response, "Error": err})
	})
	

}

