package utils

import (
	"net/http"
	"html/template"
)


//SimpleTemplateView just renders provided template
func SimpleTemplateView(w http.ResponseWriter, templateName string, data interface{})  {

	template := template.Must(template.ParseFiles(
		templateName,
	))

	err := template.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}