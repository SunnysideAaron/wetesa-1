package server

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"web/internal/config"
)

var (
	templates   map[string]*template.Template
	templatesMu sync.RWMutex
)

func ParseTemplates() error {
	templatesMu.Lock()
	defer templatesMu.Unlock()

	// should not include "./" or "/" before or after. We need this in slightly
	// different ways later on.
	dir := "templates"

	templates = make(map[string]*template.Template)

	pages := []string{}
	err := filepath.Walk("./"+dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".page.tmpl") {
			pages = append(pages, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	for _, page := range pages {
		fmt.Println(page)

		name := strings.TrimPrefix(page, dir+"/")

		ts, err := template.ParseFiles(page)
		if err != nil {
			return err
		}

		// referencedTemplates := findReferencedTemplates(page)

		referencedTemplates := []string{"templates/_site/base.layout.tmpl", "templates/_site/footer.partial.tmpl", "templates/clients/footer.partial.tmpl"}

		if len(referencedTemplates) > 0 {
			fmt.Println("Referenced templates:", referencedTemplates)

			// It appears that ParseFiles
			// 1. creates a template with the base filename (without path).
			// 2. creates templates named from {{define "name"}} code as well.
			// this can result in duplicates. but it makes sense since one template
			// can have multiple {{define}} blocks.
			ts, err = ts.ParseFiles(referencedTemplates...)
			if err != nil {
				return fmt.Errorf("error parsing referenced templates for %s: %w", page, err)
			}
		}

		// ts, err = ts.ParseGlob(filepath.Join("./templates", "*.layout.tmpl"))
		// if err != nil {
		// 	return err
		// }

		// ts, err = ts.ParseGlob(filepath.Join("./templates", "*.partial.tmpl"))
		// if err != nil {
		// 	return err
		// }

		// Add the template set to the cache, using the name of the page
		// (like 'home.page.tmpl') as the key.
		templates[name] = ts
	}

	// for _, t := range templates {
	// 	fmt.Println(t.Name())
	// }

	fmt.Println("Templates:")
	for name, ts := range templates {
		fmt.Printf("Template %s includes:\n", name)
		for _, t := range ts.Templates() {
			fmt.Printf("  - %s\n", t.Name())
		}
	}

	// err := filepath.Walk("./templates", func(path string, info fs.FileInfo, err error) error {
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if info.IsDir() {
	// 		return nil
	// 	}

	// 	if filepath.Ext(path) != ".tmpl" {
	// 		return nil
	// 	}

	// 	tmpl, err := template.ParseFiles(path)
	// 	if err != nil {
	// 		return fmt.Errorf("error parsing template %s: %w", path, err)
	// 	}

	// 	templateName := filepath.Base(path)
	// 	templates[templateName] = tmpl
	// 	return nil
	// })
	// return err
	return nil
}

func renderTemplate(cfg *config.WebConfig, templateName string, data any) (string, error) {
	if !cfg.WebCacheTemplates {
		err := ParseTemplates()
		if err != nil {
			fmt.Println("Error parsing templates:", err)
			return "", err
		}
	}

	templatesMu.RLock()
	tmpl, ok := templates[templateName]
	templatesMu.RUnlock()

	if !ok {
		return "", fmt.Errorf("template %s not found", templateName)
	}

	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	// err := tmpl.ExecuteTemplate(&buf, "body", data)
	if err != nil {
		return "", fmt.Errorf("error executing template %s: %w", templateName, err)
	}
	return buf.String(), nil
}

// findReferencedTemplates recursively finds all templates referenced in a template file
func findReferencedTemplates(templatePath string) []string {
	// Map to track unique template paths
	foundTemplates := make(map[string]bool)

	// Process the initial template
	processTemplate(templatePath, foundTemplates)

	// Convert map to slice
	result := []string{}
	for path := range foundTemplates {
		if path != templatePath { // Don't include the original template
			result = append(result, path)
		}
	}

	return result
}

// processTemplate scans a template file for references to other templates
func processTemplate(templatePath string, foundTemplates map[string]bool) {
	// Read the template file
	content, err := os.ReadFile(templatePath)
	if err != nil {
		fmt.Printf("Warning: Could not read template %s: %v\n", templatePath, err)
		return
	}

	// Look for template references like {{template "name" .}}
	templateDir := filepath.Dir(templatePath)

	// Simple regex-like search for {{template "something"
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if idx := strings.Index(line, "{{template"); idx >= 0 {
			// Extract the template name
			startQuote := strings.Index(line[idx:], "\"")
			if startQuote < 0 {
				continue
			}

			endQuote := strings.Index(line[idx+startQuote+1:], "\"")
			if endQuote < 0 {
				continue
			}

			templateName := line[idx+startQuote+1 : idx+startQuote+1+endQuote]

			// Try to find the actual template file
			possibleExtensions := []string{".layout.tmpl", ".partial.tmpl", ".page.tmpl", ".tmpl"}

			for _, ext := range possibleExtensions {
				// Check if it's a relative path or just a name
				var possiblePath string
				if strings.Contains(templateName, "/") {
					// It's a path, try to resolve it
					possiblePath = filepath.Join("templates", templateName+ext)
				} else {
					// It's just a name, look in the same directory and in common directories
					possiblePath = filepath.Join(templateDir, templateName+ext)

					// Also check in common template directories
					if _, err := os.Stat(possiblePath); os.IsNotExist(err) {
						possiblePath = filepath.Join("templates", "_site", templateName+ext)
					}
				}

				// If file exists and we haven't processed it yet
				if _, err := os.Stat(possiblePath); err == nil {
					if !foundTemplates[possiblePath] {
						foundTemplates[possiblePath] = true
						// Recursively process this template
						processTemplate(possiblePath, foundTemplates)
					}
					break // Found the template, no need to try other extensions
				}
			}
		}
	}
}
