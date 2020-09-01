package redis

import (
	"encoding/json"
	"fmt"

	"github.com/dvhthomas/meez/pkg/models"

	"github.com/gomodule/redigo/redis"
)

// RecipeModel wraps datastore calls for a Redis backend
type RecipeModel struct {
	Pool *redis.Pool
}

// All redis strings have a prefix. We'll use 'recipe'
const objPrefix string = "recipe:"

// The key name used to keep track of the next Recipe ID field value
var ids string = fmt.Sprintf("%sid", objPrefix)

// Save a new recipe to the datastore, or store a new one if the ID field
// is empty.
func (rm *RecipeModel) Save(r *models.Recipe) (int, error) {
	conn := rm.Pool.Get()
	defer conn.Close()

	// Default value of an int in Go is zero, unless an ID
	// was explicitly set, a zero value means this is a new
	// Recipe and needs a new ID to use as the Redis key.
	if r.ID == 0 {
		nextID, err := redis.Int(conn.Do("INCR", ids))
		if err != nil {
			return -1, err
		}

		r.ID = nextID
	}

	jsonRecipe, err := json.Marshal(r)
	if err != nil {
		return -1, err
	}

	_, err = conn.Do("SET", fmt.Sprintf("%s%d", objPrefix, r.ID), jsonRecipe)
	if err != nil {
		return -1, err
	}

	return r.ID, nil
}

// Get a single recipe or nothing
func (rm *RecipeModel) Get(id int) (*models.Recipe, error) {
	return nil, nil
}

// Latest recipes meaning most recent published, limited to a max
// number of results
func (rm *RecipeModel) Latest(limit int) ([]*models.Recipe, error) {
	return nil, nil
}
