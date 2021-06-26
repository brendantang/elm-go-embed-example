# elm-go-embed-example

A "Hello, world!" example showing how to use `go embed` to bundle an Elm frontend as part of a Go binary.

Package [`embed`](https://golang.org/pkg/embed/) lets you direct the Go compiler to bundle up files as part of the compiled executable.

It's common to build a web application with frontend assets and a Go backend, using the `//go:embed` directive to package the whole program up as a single executable. This example illustrates how to do just that, where the frontend assets are an [Elm](https://elm-lang.org) application.



## Build and run


Run `make` and find the standalone executable `dist/elm-go-embed-example`. 
See the Makefile for the frontend and backend build commands.



## Other examles

The `main` branch of this repo shows a bare minimum setup to bundle an Elm frontend in a Go binary.
Other branches have more involved examples:

- Branch `elm-spa` bundles a frontend built using [`elm-spa`](https://www.elm-spa.dev/).
  - Backend handles routes prefixed with `/api`, while others are handled by the elm-spa client-side router. Try visiting `/`, `/blargh`, `/echo/blargh`, or `/api/echo/blargh`.

