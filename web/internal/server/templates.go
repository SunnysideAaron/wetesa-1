package server

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
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

	templates = make(map[string]*template.Template)

	pages, err := filepath.Glob(filepath.Join("./templatepages", "*.tmpl"))
	if err != nil {
		return err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		name = name[:len(name)-5] // Remove ".tmpl" extension. Doesn't really matter but makes the call that much shorter later.

		ts, err := template.ParseFiles(page)
		if err != nil {
			return err
		}

		// yes this means that every page has all the shared templates.
		// the alternatives get wonky. See ADR-1017 Template Cache
		// good enough for now.
		ts, err = ts.ParseGlob(filepath.Join("./templateshared", "*.tmpl"))
		if err != nil {
			return err
		}

		templates[name] = ts
	}

	// Debugging code to see what templates are loaded.
	// DON't DELETE maybe someday put in debug and log
	// fmt.Println("Templates:")
	// for name, ts := range templates {
	// 	fmt.Printf("Template %s includes:\n", name)
	// 	for _, t := range ts.Templates() {
	// 		fmt.Printf("  - %s\n", t.Name())
	// 	}
	// }

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
