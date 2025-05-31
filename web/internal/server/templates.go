package server

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
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

	err := filepath.Walk("./templates", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".tmpl" {
			return nil
		}

		tmpl, err := template.ParseFiles(path)
		if err != nil {
			return fmt.Errorf("error parsing template %s: %w", path, err)
		}

		templateName := filepath.Base(path)
		templates[templateName] = tmpl
		return nil
	})
	return err
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
