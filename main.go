package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"github.com/google/go-github/github"
)

func main() {

	client := github.NewClient(nil)

	// handler switches on file type (md)

	panic(http.ListenAndServe(":8080", http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			upath := r.URL.Path
			if !strings.HasPrefix(upath, "./") {
				upath = "./" + upath
				r.URL.Path = upath
			}
			p := path.Clean(upath)
			if filepath.Ext(p) == ".md" {
				f, err := ioutil.ReadFile(p)
				if err != nil {
					http.Error(rw, err.Error(), http.StatusBadRequest)
				}
				//opt := &github.MarkdownOptions{Mode: "gfm", Context: "google/go-github"}

				output, _, err := client.Markdown(context.Background(), string(f), nil)
				if err != nil {
					http.Error(rw, err.Error(), http.StatusBadRequest)
				}
				rw.Write([]byte(output))
			} else {
				http.ServeFile(rw, r, p)
			}

		})))

	panic(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
