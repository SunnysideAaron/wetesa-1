
Here is an example of how to implement template caching in Go:
1. Create a Template Cache
    - A map[string]*template.Template is used to store parsed templates, where the key is the template name, and the value is the parsed template.
    - A function parseTemplates is defined to parse all templates in a directory and store them in the cache.

``` go
import (
	"fmt"
	"html/template"
	"io/fs"
	"path/filepath"
	"sync"
)

var (
	templates   map[string]*template.Template
	templatesMu sync.RWMutex
)

func parseTemplates(templateDir string) error {
	templatesMu.Lock()
	defer templatesMu.Unlock()
	templates = make(map[string]*template.Template)

	err := filepath.Walk(templateDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".html" {
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
```

2. Load Templates
    - Call parseTemplates once at the application startup to load all the templates into the cache.

``` go
func main() {
    templateDir := "./templates" // Replace with your template directory
    err := parseTemplates(templateDir)
    if err != nil {
        fmt.Println("Error loading templates:", err)
        return
    }
	// ... rest of the application
}
```

3. Render Templates
    - A function renderTemplate is defined to render a template by fetching it from the cache.

``` go
func renderTemplate(templateName string, data interface{}) (string, error) {
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
```

4. Usage
    - Use renderTemplate in your handlers to render the template.

``` go
func handleIndex(w http.ResponseWriter, r *http.Request) {
    data := map[string]string{"Title": "Home Page"}
    rendered, err := renderTemplate("index.html", data)
    if err != nil {
        http.Error(w, "Error rendering template", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(rendered))
}
```