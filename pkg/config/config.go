package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
)

type Config struct {
	UseCache       bool
	TemplatesCache map[string]*template.Template
	InProduction   bool
	Session        *scs.SessionManager
}
