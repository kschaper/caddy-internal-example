# Caddy http.internal example

[http.internal](https://caddyserver.com/docs/internal) can be used to protect resources. Authentication is handled by a backend.
For this example it's written in Go.

## Set up

1. `git clone` this repo

### caddy

Note: This is not meant to be used in a production environment. For production activate at least [automatic HTTPS](https://caddyserver.com/docs/automatic-https).

1. download Caddy: https://caddyserver.com/download/darwin/amd64?license=personal&telemetry=off
1. unpack and move binary into PATH
1. `$ cd ./web/`
1. `$ caddy`

### go authenticaton app

1. make sure [go](https://golang.org/) is installed
1. `$ cd ./cmd/authenticaton/`
1. `$ go run main.go`

### test

unauthenticated:
1. http://localhost:8080/internal → 404 Not Found
1. http://localhost:8080/private/main.html → 404 Not Found

authenticated:
1. http://localhost:8080/internal → 404 Not Found
1. http://localhost:8080/private/main.html → 200 OK

### notes

- do not use [index files](https://caddyserver.com/docs/index) in internal dir
  - Given there is the file `internal/index.html` and the URL is `/private/index.html` then the response will be a redirect to `/internal/`.
  - see issue [File server redirects not handled by internal middleware](https://github.com/mholt/caddy/issues/1811)
