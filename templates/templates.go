package templates

import (
	"embed"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

type (
	Layout string
	Page   string
)

const (
	LayoutMain Layout = "main"
)

const (
	PageDashboard Page = "dashboard"
	PageError     Page = "error"
	PageHome      Page = "home"
)

//go:embed *
var templates embed.FS

// Get returns a file system that contains all templates via embed.FS
func Get() embed.FS {
	return templates
}

// GetOS returns a file system containing all templates which will load
// the files directly from the operating system.
// This should only be used for development in order to facilitate live reloading.
func GetOS() fs.FS {
	// Get the complete templates path.
	// Note: This function maybe called from outside main i.e. tests.
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	p := filepath.Join(filepath.Dir(d), "templates")
	return os.DirFS(p)
}
