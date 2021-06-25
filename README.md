# elm-go-embed-example

A "Hello, world!" example showing how to use `go embed` to bundle an Elm frontend as part of a Go binary.

Package [`embed`](https://golang.org/pkg/embed/) lets you direct the Go compiler to bundle up files as part of the compiled executable.

It's common to build a web application with frontend assets and a Go backend, using the `//go:embed` directive to package the whole program up as a single executable. This example illustrates how to do just that, where the frontend assets are an [Elm](https://elm-lang.org) application.


`Makefile` has the commands to build the frontend and backend. `embed-elm-example` is a standalone executable that serves the Elm application on port 8000.
