package application

import (
	"github.com/dvhthomas/meez/pkg/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// WebApp is the meez web app
type WebApp struct {
	Recipes models.RecipeModel
}

// BuildRoutes adds all the routes and middleware to those routes. This
// is called during app setup.
func (app *WebApp) BuildRoutes(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.GET("/", app.home)
	e.GET("/recipes/:id", app.getRecipe).Name = "get-recipe"
	e.POST("/recipes", app.createRecipe).Name = "create-recipe"
}
