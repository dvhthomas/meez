# Meez

[Mise en place](https://en.wikipedia.org/wiki/Mise_en_place) is about being prepared for cooking. You might here it called 'meez'.

And since I'm learning Go into the bargain, there's a surprising amount of 'prep' going into my 'meez'

## Getting Started

You'll need a Go compiler. The backing data store is Redis. On a Mac you can do this with:

```sh
brew update
brew install redis
```

And try it out using `redis-server`.

meez uses Go modules so you can stick the code anywhere you like. Just keep things sane by 
also setting an environment variable for that, e.g., `export PROJECT=/the/place/you/chose`.

## Running Meez

I'm kind of a tmux lover, so I tend to do something like `cd $PROJECT;tmux new -s meez` before I get started. That way I can
easily switch between a bunch of terminals running things like Redis and the web server. But do whatever works
best for you. But tmux is good...

1. `cd $PROJECT`
1. Start Redis: `./scripts/run-db.sh`.
1. The web app is `cmd/meez/main.go` so the following will work: `go run ./cmd/meez`. That will start on https://localhost:4000 by
   default, but you can pass in a port with the `-addr ":8081"` switch (obviously choose whatever port works).

Once that's running, try creating a recipe. I like [HTTPie](https://httpie.org/) but use `curl` or Postman is that's your thing.

```sh
http POST http://localhost:4000/recipes name="Sourdough biscuits" \
  description="Wonderful bread" datePublished="2020-08-30T12:00:00Z" \
  author:='{"name": "Thomas Keller", "email": "thomas@frenchlaundry.com"}
```

Check out the recipe ID in the response. You can use that to see the recipe in redis. So assuming
you started redis on it's default port you can connect to that server and see recipe #1 (or whatever)
using the following.

```sh
$ redis-cli
redis> GET recipe:id 1
```
