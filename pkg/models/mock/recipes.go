package mock

import (
	"time"

	"github.com/dvhthomas/meez/pkg/models"
	"github.com/rickb777/date/period"
)

// RecipeModel for mock data
type RecipeModel struct{}

func mockData() []*models.Recipe {

	mockSalad := &models.Recipe{
		Author: &models.Person{
			Email: "alice@waters.com",
			Name:  "Alice Waters",
			URL:   "/alice",
		},
		Description: "A wonderful salad recipe",
		Name:        "Apple & walnut salad",
		CookTime:    period.New(0, 0, 0, 0, 0, 0).String(),
		PrepTime:    period.NewHMS(0, 15, 0).String(),
		PublishedOn: time.Now(),
	}

	mockSalad.AddInstruction(
		"Slicing",
		"Cut all ingredients into similar sized pieces",
		"",
	)
	mockSalad.AddInstruction("Tossing",
		"Put everything in a bowl and toss vigorously", "",
	)

	var mockChicken = &models.Recipe{
		Author: &models.Person{
			Email: "thomas@frenchlaundry.com",
			Name:  "Thomas Keller",
			URL:   "/Thomas",
		},
		Description: "Basic roast chicken inspired by family meal at the French Laundry",
		Name:        "Basic roast chicken",
		CookTime:    period.NewHMS(0, 45, 0).String(),
		PrepTime:    period.NewHMS(0, 5, 0).String(),
		PublishedOn: time.Now().AddDate(0, 0, -7),
	}
	mockChicken.AddInstruction("Prep", "Truss the chicken", "")
	mockChicken.AddInstruction("Season", "Liberally salt and pepper the chicken", "")

	return []*models.Recipe{mockSalad, mockChicken}
}

// Get a single mock recipe
func (m *RecipeModel) Get(id int) (*models.Recipe, error) {
	d := mockData()
	return d[0], nil
}

// Latest gets the most recently PublishedOn Recipes with the ability
// to limit the list to a specifc count.
func (m *RecipeModel) Latest(limit int) ([]*models.Recipe, error) {
	d := mockData()
	return d, nil
}

// Save a new or existing recipe. If the recipe already exists
// it will be overwritten
func (m *RecipeModel) Save(recipe *models.Recipe) (int, error) {
	return 0, nil
}
