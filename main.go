package main

import (
	"embed"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io/fs"
	"log"
	"net/http"
	"os"
)

// The special `//go:embed ...` directive tells the Go compiler to bundle up the
// files at the provided path in the compiled binary. In this case, we bundle
// the build directory of our Elm frontend and make it available as an `fs.FS`
// value.
//go:embed elm-frontend/public/*
var assets embed.FS

func main() {

	// `frontend` is an `fs.FS` value with its "root" inside the Elm
	// frontend's build directory.
	frontend, err := fs.Sub(assets, "elm-frontend/public")
	if err != nil {
		log.Fatal(err)
	}

	// Any paths not prefixed with "/api" serve the frontend
	frontendFileSystem := singlePageFileSystem{http.FS(frontend)}
	router := chi.NewRouter()
	router.Mount("/", http.FileServer(frontendFileSystem))

	// Paths prefixed with "/api" are handled by the backend
	apiRouter := chi.NewRouter()
	apiRouter.HandleFunc("/echo/{toEcho}", handleEcho)
	router.Mount("/api/", apiRouter)

	log.Fatal(
		http.ListenAndServe(
			":8000",
			router,
		),
	)
}

// singlePageFileSystem satisfies the `http.FileSystem` interface, but unlike
// the type you get as a result of passing a `fs.FS` to `http.FS`, its Open
// method defaults to an index page when a static file is not found with the
// given name.
//
// You can pass a `singlePageFileSystem` to `http.FileServer` to serve some
// static assets, and also allow the frontend app mounted on the index page to
// perform client-side routing.
type singlePageFileSystem struct {
	http.FileSystem
}

// Open(name string) serves the static file at `name` if it exists, and
// otherwise serves `index.html`.
func (fsys singlePageFileSystem) Open(name string) (http.File, error) {

	f, err := fsys.FileSystem.Open(name)

	// Handle an unexpected error.
	if err != nil && !os.IsNotExist(err) {

		return nil, err
	}

	// If the file isn't found, assume it's a route for the SPA to handle.
	if err != nil && os.IsNotExist(err) {
		return fsys.Open("index.html")
	}

	// Serve the static file if it exists.
	return http.File(f), nil

}

func handleEcho(w http.ResponseWriter, r *http.Request) {
	toEcho := chi.URLParam(r, "toEcho")
	echo := fmt.Sprintf("'%v' echoed from the backend", toEcho)
	w.Write([]byte(echo))
}
