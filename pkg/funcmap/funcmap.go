package funcmap

import (
	"fmt"
	"html/template"
	"reflect"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/labstack/gommon/random"

	"github.com/tiny-blob/tinyblob/config"
)

var (
	CacheBuster = random.String(10)
)

// GetFuncMap provides a template function map
func GetFuncMap() template.FuncMap {
	// http://masterminds.github.io/sprig/
	funcMap := sprig.FuncMap()

	// Below we can provide a list of custom functions.
	// We should be careful to avoid using a name already used by sprig.
	f := template.FuncMap{
		"hasField": HasField,
		"file":     File,
		"link":     Link,
	}

	for k, v := range f {
		funcMap[k] = v
	}

	return funcMap
}

// HasField checks if an interface contains a given field.
func HasField(v interface{}, name string) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return false
	}
	return rv.FieldByName(name).IsValid()
}

// File appends a cache buster to a given filepath so it can remain cached until
// the app is restarted.
func File(filepath string) string {
	return fmt.Sprintf("/%s/%s?v=%s", config.StaticPrefix, filepath, CacheBuster)
}

// Link outputs HTML for a link element, providing the ability to dynamically
// set the active class.
func Link(url, text, currentPath string, classes ...string) template.HTML {
	classes = append(classes, "text-blue-600", "hover:underline")
	if currentPath == url {
		classes = append(classes, "font-bold")
	}

	html := fmt.Sprintf(
		`<a class="%s" href="%s">%s</a>`,
		strings.Join(classes, " "),
		url,
		text,
	)
	return template.HTML(html)
}
