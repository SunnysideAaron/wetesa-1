package server

import (
	"context"
	"html/template"
	"log/slog"
	"path/filepath"
)

// template caching examples

// Let's Go! by Alex Edwards Chapter 5.3
// [How To Create A Template Cache For Your Golang Web Application](https://andrew-mccall.com/blog/2022/06/create-a-template-cache-for-a-go-application/)

func NewTemplateCache(ctx context.Context,
	logger *slog.Logger,
	dir string,
) map[string]*template.Template {
	// Initialize a new map to act as the cache.
	cache := map[string]*template.Template{}
	// Use the filepath.Glob function to get a slice of all filepaths with
	// the extension '.page.tmpl'. This essentially gives us a slice of all the
	// 'page' templates for the application.
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		slog.LogAttrs(
			ctx,
			slog.LevelError,
			"template cache error",
			slog.Any("error", err),
		)
		return nil
	}
	// Loop through the pages one-by-one.
	for _, page := range pages {
		// Extract the file name (like 'home.page.tmpl') from the full file pat
		// and assign it to the name variable.
		name := filepath.Base(page)
		// Parse the page template file in to a template set.
		ts, err := template.ParseFiles(page)
		if err != nil {
			slog.LogAttrs(
				ctx,
				slog.LevelError,
				"template cache error",
				slog.Any("error", err),
			)
			return nil
		}
		// Use the ParseGlob method to add any 'layout' templates to the
		// template set (in our case, it's just the 'base' layout at the
		// moment).
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			slog.LogAttrs(
				ctx,
				slog.LevelError,
				"template cache error",
				slog.Any("error", err),
			)
			return nil
		}
		// Use the ParseGlob method to add any 'partial' templates to the
		// template set (in our case, it's just the 'footer' partial at the
		// moment).
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			slog.LogAttrs(
				ctx,
				slog.LevelError,
				"template cache error",
				slog.Any("error", err),
			)
			return nil
		}
		// Add the template set to the cache, using the name of the page
		// (like 'home.page.tmpl') as the key.
		cache[name] = ts
	}
	// Return the map.
	return cache
}
