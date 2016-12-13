package views

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var ViewsDir = "views"
var ViewsExt = ".gohtml"
var LayoutDir = ViewsDir + "/layouts"

func NewView(layout string, files ...string) *View {
	addViewsDirPrefix(files)
	addViewsExtSuffix(files)
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*" + ViewsExt)
	if err != nil {
		log.Println(err)
	}

	return files
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, nil)
}

// Takes in a slice of strings representing file paths
// and prepends the ViewsDir directory to each string
// in the slice.
//
// Eg the input {"home"} would be updated to match
// {"views/home"} if ViewDir == "views"
func addViewsDirPrefix(files []string) {
	for i, f := range files {
		files[i] = ViewsDir + "/" + f
	}
}

// Takes in a slice of strings representing file paths
// and appends the ViewsExt extension to each string in
// the slice.
//
// Eg the input {"home"} would be updated to match
// {"home.gohtml"} if ViewsExt == ".gohtml"
func addViewsExtSuffix(files []string) {
	for i, f := range files {
		files[i] = f + ViewsExt
	}
}
