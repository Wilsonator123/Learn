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

	e.GET("/task", func(c echo.Context) error {
		
		response, err := handlers.ListAll()

		if c.Request().Header.Get("HX-Request") == "true" {
			if err != nil {
				return c.Render(http.StatusOK, "TaskList", map[string]interface{}{"Error": err})
			}
			
			return c.Render(http.StatusOK, "TaskList", map[string]interface{}{"Data": response, "Error": err})
		}
		
		return c.Render(http.StatusOK, "index", map[string]interface{}{"Data": response, "Error": err})
	})

	e.POST("/task", func(c echo.Context) error {
		
		positionStr := c.FormValue("position")
		positionInt, err := strconv.ParseInt(positionStr, 10, 16)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid position value")
		}
		positionInt16 := int16(positionInt)
		newTask := model.NewTask{
			Title: c.FormValue("title"),
			Description: c.FormValue("description"),
			Position: &positionInt16,
		}
		
		if err := validate.Struct(newTask); err != nil {
			fmt.Printf("Error parsing form data %v\n", err)
			return c.String(http.StatusBadRequest, "Failed to parse form data")
		}

		response, err := handlers.CreateTask(newTask)

		fmt.Printf("User created with id %v\n", response)

		if err != nil {
			fmt.Printf("Error creating user %v\n", err)

			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.Render(http.StatusCreated, "task", response)
	})

	e.GET("/task/:id", func(c echo.Context) error {

		id := c.Param("id")

		response, err := handlers.GetTask(id)
		
		return c.Render(http.StatusOK, "TaskDetails", map[string]interface{}{"Data": response, "Error": err})
	})

	e.DELETE("/task/:id", func(c echo.Context) error {
		
		id := c.Param("id")

		handlers.DeleteTask(id)

		return c.HTML(http.StatusOK,"<div>Deleted</div>" )

	})
	

}

