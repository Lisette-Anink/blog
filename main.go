package main

import (
	"blog/models"
	"blog/templates"
	"context"
	"log"
	"os"
	"path"

	"github.com/a-h/templ"
)

func main() {
	posts := models.GetJsonPosts()
	// Output path.
	rootPath := "public"
	if err := os.Mkdir(rootPath, 0755); err != nil && !os.IsExist(err) {
		log.Fatalf("failed to create output directory: %v", err)
	}
	// Create an index page.
	name := path.Join(rootPath, "index.html")
	f, err := os.Create(name)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}
	// Write it out.
	err = templates.IndexPage(posts).Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write index page: %v", err)
	}

	// Create the Static pages
	addStatic(rootPath, templates.MapPath, templates.MapPage())
	addStatic(rootPath, templates.AboutPath, templates.AboutPage())

	// Create a page for each post.
	for _, post := range posts {
		// Create the output directory.
		dir := path.Join(rootPath, post.Slug())
		f := createDirAndIndex(dir)
		// Create the output file.

		data := templates.BlogData{Post: post}
		// Use templ to render the template containing the raw HTML.
		err = templates.ContentPage(data).Render(context.Background(), f)
		if err != nil {
			log.Fatalf("failed to write output file: %v", err)
		}
	}
}

func addStatic(rootPath string, pagePath string, template templ.Component) {
	dir := path.Join(rootPath, pagePath)
	f := createDirAndIndex(dir)

	err := template.Render(context.Background(), f)

	if err != nil {
		log.Fatalf("failed to write index page: %v", err)
	}
}

func createDirAndIndex(dir string) *os.File {
	if err := os.MkdirAll(dir, 0755); err != nil && !os.IsExist(err) {
		log.Fatalf("failed to create output directory: %v", err)
	}
	name := path.Join(dir, "index.html")
	f, err := os.Create(name)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}
	return f
}
