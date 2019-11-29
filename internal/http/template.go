package http

import (
	"html/template"
	"io"
	"path/filepath"

	"github.com/hashicorp/go-hclog"
	"github.com/labstack/echo"
)

// Renderer fulfills the echo Renderer interface and maintains a cache
// of templates.  It also provides a debug mode where changes to
// templates on disk are reflected on the next page reload rather than
// having to restart the server.
type renderer struct {
	hclog.Logger

	tmpls map[string]*template.Template
}

func newRenderer(baseDir string, logger hclog.Logger) (*renderer, error) {
	r := &renderer{}

	r.Logger = logger.Named("renderer")

	if err := r.loadTmpls(baseDir); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *renderer) loadTmpls(loadFrom string) error {
	base, err := template.ParseGlob(filepath.Join(loadFrom, "base", "*.tpl"))
	if err != nil {
		return err
	}

	r.tmpls = make(map[string]*template.Template)

	// We can safely throw away this error because the pattern is
	// hard coded.
	pages, _ := filepath.Glob(filepath.Join(loadFrom, "pages", "*", "*.tpl"))

	for _, path := range pages {
		page := filepath.Base(filepath.Dir(path))
		r.Debug("Loading page template", "template", page)
		pTmpl, err := base.Clone()
		if err != nil {
			return err
		}
		pTmpl, err = pTmpl.ParseGlob(path)
		if err != nil {
			return err
		}
		r.tmpls[page] = pTmpl
	}

	return nil
}

func (r *renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.tmpls[name].ExecuteTemplate(w, "base", data)
}
