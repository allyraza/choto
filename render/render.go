package render

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Default : default render
var Default *Render

const (
	defaultTemplateSuffix = "*.html"
	defaultTemplatePath   = "./templates"
)

func init() {
	Default = New()
}

// Render :
type Render struct {
	TemplatePath   string
	TemplateSuffix string
	Engine         *template.Template
}

// New : creates new render
func New() *Render {
	r := &Render{
		TemplateSuffix: defaultTemplateSuffix,
		TemplatePath:   defaultTemplatePath,
	}

	r.parseTemplates()

	return r
}

func (r *Render) parseTemplates() {

	dir, err := filepath.Abs(r.TemplatePath)
	if err != nil {
		log.Fatal(err)
	}

	path := filepath.Join(dir, r.TemplateSuffix)
	log.Println(path)
	files, err := filepath.Glob(path)
	if err != nil {
		log.Fatal(err)
	}

	tpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatal(err)
	}

	r.Engine = tpl
}

// JSON :
func (r *Render) JSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

// SetTemplatePath :
func (r *Render) SetTemplatePath(path string) {
	r.TemplatePath = path

	r.initTemplateEngine()
}

// GetTemplatePath :
func (r *Render) GetTemplatePath() string {
	return r.TemplatePath
}

// GetTemplatePath : package level alias
func GetTemplatePath() string {
	return Default.GetTemplatePath()
}

// SetTemplatePath : package level alias
func SetTemplatePath(path string) {
	Default.SetTemplatePath(path)
}

// HTML : renders a given template
func (r *Render) HTML(w http.ResponseWriter, name string, data interface{}) {
	r.Engine.ExecuteTemplate(w, name, data)
}

// HTML : package level alias
func HTML(w http.ResponseWriter, name string, data interface{}) {
	Default.HTML(w, name, data)
}

// JSON :
func JSON(w http.ResponseWriter, data interface{}) {
	Default.JSON(w, data)
}
