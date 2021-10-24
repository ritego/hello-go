# Hello Go

The full source code for this tutorial is available at [Hello Go](https://github.com/ritego/hello-go). 

## Introduction
In this article, we would learn how to write a simple `Hello World` application in Go. Since this is the first RiteGo article, its traditional that we start this way.

More specifically, we would:
- Write a program that console log `Hello World`
- Write a simple HTTP server that respond with `Hello World`

## Start With Dependencies
[Download and install Go](https://golang.org/dl/). To verify your setup is good, run `go version` in your favorite CLI.
```bash
$ go version

// Output
// go version go1.16.4 darwin/amd64
```

## Bootstrap Your Project
To a large extend, Go is both a language and a framework. This is because in addition to the core language features, Go comes out of the box with tools and standard libraries that makes programming beautiful.

For instance, Go has inbuilt CLI tool for bootstrapping a module (a new project). This tool is called `go mod` and is similar to `npm` and `composer` for node and php developers respectively.

Your project can have one or more modules (with only one module at the root), a module can have one or more packages and a package can have one or more source files.

Run `go mod init`. this command corresponds to `npm init` and `composer init` for node and php respectively. The only difference between Go and these other languages is that the dependency management tool is not a third-party tool, it is part of the core of Go.

```bash
$ go mod init github.com/[username]/[hello-go]

// Output
// go: creating new go.mod: module github.com/[username]/[hello-go]
// go: to add module requirements and sums:
//         go mod tidy
```

This command would also create two files at the root of your directory:

`go.mod`
`go.mod` records the modules dependencies just like `package.json` and `composer.json`.
```go
module github.com/ritego/hello-go

go 1.16

require (
	github.com/bketelsen/crypt v0.0.4 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/spf13/viper v1.9.0 // indirect
)

```

`go.sum`
`go.sum` tracks the specific version of installed dependencies just like `package.lock` and `composer.lock`, though it uses a different approach.
```go
...
cloud.google.com/go v0.93.3/go.mod h1:8utlLll2EF5XMAV15woO4lSbWQlk8rer9aLOfLh7+YI=
cloud.google.com/go/bigquery v1.0.1/go.mod h1:i/xbL2UlR5RvWAURpBYZTtm/cXjCha9lbfbpx4poX+o=
...
```


## Main
The root file for go projects is `main.go` and must belong to a package named `main` and have a function called `main`. The `main` function is used by Go to initialize your module.

```go
// main.go

package main

func main() {
    // start coding
}
```

## Console Log Hello World 
Import the standard library for formatting and output a string to standard output (os.Stdout).

```go
// main.go

package main

import "fmt"

func main() {
	// Say Hello World in console
	fmt.Println("Hello World")

    ...
}
```

## Simple HTTP Server
That was simple enough, next we would build a simple HTTP server. 

First, we define our port number, our service would bind to this port when running.
```go
// main.go

...

addr := ":7777"
```

Next, we would define routing patterns for incoming requests using the default router provided by Go. There are more sophisticated routers that can be used, like [Gorilla Mux](https://github.com/gorilla/mux), but we stick to the simple one for now.
```go
// main.go

...

router := http.DefaultServeMux
router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
    rw.Write([]byte("Hello World"))
})
```

Next, we build our server
```go
// main.go

...

srv := &http.Server{
    Handler:      router,
    Addr:         addr,
    WriteTimeout: 5 * time.Second,
    ReadTimeout:  5 * time.Second,
}
```

And lastly, start accepting requests
```go
// main.go

...

log.Printf("Server running on: %s", addr)
if err := srv.ListenAndServe(); err != nil {
    log.Fatal(err)
}
```

Here is everything put together.
```bash
- hello-go
    - go.mod
    - go.sum
    - main.go
```

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Say Hello World in console
	fmt.Println("Hello World")

	// make the port number a reusable variable
	addr := ":7777"

	// define routes
	router := http.DefaultServeMux
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello World"))
	})

	// configure server
	srv := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	// start accepting requests
	log.Printf("Server running on: %s", addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
```

## Run Your Server
```bash
$ go run main.go

// Output
// Hello World
// 2021/10/23 07:39:49 Server running on: :7777
```

You can visit [localhost:7777](localhost:7777) on your browser to see the `Hello World` response from your server.

## Conclusion
The full source code for this tutorial is available at [Hello Go](github.com/ritego/hello-go). You can fork and extend the program with additional features:
- Console log `Hello World` in multiple languages
- Add a different route for each language. e.g `/`, `/yoruba`, `/hausa`, `igbo`

