package models

// RecipeModel provides a repository-style interface to the recipe datastore
type RecipeModel interface {
	// Get a single recipe based on a unique ID.
	Get(id int) (*Recipe, error)
	Latest(limit int) ([]*Recipe, error)
	Save(recipe *Recipe) (int, error)
}
