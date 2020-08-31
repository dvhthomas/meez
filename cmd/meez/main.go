package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/dvhthomas/meez/internal/application"
	r "github.com/dvhthomas/meez/pkg/models/redis"
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network port like ':4000'")
	flag.Parse()

	pool := &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	app := &application.WebApp{
		Recipes: &r.RecipeModel{
			Pool: pool,
		},
	}

	e := echo.New()
	app.BuildRoutes(e)
	s := &http.Server{
		Addr:         *addr,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	e.Logger.Fatal(e.StartServer(s))
}
