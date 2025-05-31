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

	for name := range templates {
		fmt.Println(name)
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
	if err != nil {
		return "", fmt.Errorf("error executing template %s: %w", templateName, err)
	}
	return buf.String(), nil
}
