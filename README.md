# golang_webserver

This is a simple web server written in Go.

## Features

- Serves static files from the `static/` directory
- Simple routing for GET requests
- Graceful shutdown on CTRL+C

## Usage


go run main.go


This will start the server on port 8080. 

To change the port:


go run main.go -port=8000


## Routing

The server has some simple routing built in:

- `/` - Serves `static/index.html` 
- `/hello` - `return hello in return`
- `/form` - Serves `static/forms.html`

## Deployment

The server is just a single Go binary, so you can build it for your platform and deploy it anywhere.

For example, to build for Linux:

GOOS=linux GOARCH=amd64 go build


