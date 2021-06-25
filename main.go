package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

// The special `//go:embed ...` directive tells the Go compiler to bundle up the
// files at the provided path in the compiled binary. In this case, we bundle
// the build directory of our Elm frontend and make it available as an `fs.FS`
// value.
//go:embed elm-frontend/build/*
var f embed.FS

func main() {
	// `frontend` is an `fs.FS` value with its "root" inside the Elm
	// frontend's build directory.
	frontend, err := fs.Sub(f, "elm-frontend/build")

	if err != nil {
		log.Fatal(err)
	}

	// Serve the compiled Elm frontend as static files.
	log.Fatal(
		http.ListenAndServe(
			":8000",
			http.FileServer(http.FS(frontend)),
		),
	)

}
