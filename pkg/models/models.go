package models

import (
	"errors"
	"math"
	"time"
)

// ErrNoRecipe is used when a query for a specific recipe
// fails to return one recipe
var ErrNoRecipe = errors.New("no recipe found")

// Recipe is something you make to eat.
// The structure is strongly influenced by https://developers.google.com/search/docs/data-types/recipe#nutrition.calories
// and the underlyling https://schema.org/Recipe thing.
// Time periods are stored as text-based ISO 8601 durations (https://en.wikipedia.org/wiki/ISO_8601#Durations).
type Recipe struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Author       *Person   `json:"author"`
	PublishedOn  time.Time `json:"datePublished"`
	Instructions []*Step   `json:"recipeInstructions"`
	Description  string    `json:"description"`
	PrepTime     string    `json:"prepTime"`
	CookTime     string    `json:"cookTime"`
	URL          string    `json:"url"`
}

// AddInstruction appends a new Step to the Recipe.Instructions
// at the end of the list.
func (r *Recipe) AddInstruction(name, text, url string) {
	stepCount := len(r.Instructions)
	stepNum := stepCount + 1
	s := &Step{
		Name:     name,
		Text:     text,
		ImageURL: url,
		Order:    stepNum,
	}
	r.Instructions = append(r.Instructions, s)
}

// Person is inspired by https://schema.org/Person and can be an
// Author of a recipe, a User, etc.
type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	URL   string `json:"url"`
}

// Step is a single how-to step in a guide to completing
// a Recipe.
type Step struct {
	Name     string `json:"name"`
	Order    int    `json:"stepNumber"`
	Text     string `json:"text"`
	ImageURL string `json:"url"`
}

var millisecondsPerMinute int64 = 1000000000

func secondsToDuration(seconds int) time.Duration {
	s := math.Abs(float64(seconds))
	nanoseconds := int64(s) * millisecondsPerMinute
	return time.Duration(nanoseconds)
}

func minutesToDuration(minutes int) time.Duration {
	m := math.Abs(float64(minutes))
	s := int(m * 60)
	return secondsToDuration(s)
}
