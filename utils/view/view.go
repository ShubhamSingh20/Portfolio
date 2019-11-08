package view

import (
	"path/filepath"
	"net/http"
	"html/template"
)


var (
	templatePath = filepath.Join(filepath.Dir("."), "templates")
)


//GetTemplate get the pointer to the template
func GetTemplate(templateName ...string) *template.Template {
	fullTemplatePath := append([]string{templatePath}, templateName...)

	return template.Must(template.ParseFiles(
		filepath.Join(fullTemplatePath...), 
	))

}

//SimpleTemplateView just renders provided template
func SimpleTemplateView(w http.ResponseWriter, data interface{}, templateName ...string)  {
	template := GetTemplate(templateName...)

	err := template.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}