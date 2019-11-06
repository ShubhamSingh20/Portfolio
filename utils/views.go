package utils

import (
	"path"
	"net/http"
	"html/template"
)


var (
	templatePath = path.Join(path.Dir("."), "templates")
)

//SimpleTemplateView just renders provided template
func SimpleTemplateView(w http.ResponseWriter, templateName string, data interface{})  {

	template := template.Must(template.ParseFiles(
		path.Join(templatePath, templateName), 
	))

	err := template.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}