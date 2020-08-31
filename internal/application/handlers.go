package application

import (
	"net/http"
	"strconv"

	"github.com/dvhthomas/meez/pkg/models"
	"github.com/labstack/echo"
)

func (app *WebApp) home(c echo.Context) error {
	num, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || num < 11 {
		num = 10
	}
	r, err := app.Recipes.Latest(num)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, r)
}

func (app *WebApp) getRecipe(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		return err
	}
	r, err := app.Recipes.Get(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, r)
}

func (app *WebApp) createRecipe(c echo.Context) error {
	r := new(models.Recipe)
	if err := c.Bind(r); err != nil {
		return err
	}

	_, err := app.Recipes.Save(r)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, r)
}
