package utils

import (
	"path"
	"net/http"
	"html/template"
)


var (
	templatePath = path.Join(path.Dir("."), "templates")
)


//GetTemplate get the pointer to the template
func GetTemplate(templateName string) *template.Template {

	return template.Must(template.ParseFiles(
		path.Join(templatePath, templateName), 
	))

}

//SimpleTemplateView just renders provided template
func SimpleTemplateView(w http.ResponseWriter, templateName string, data interface{})  {
	template := GetTemplate(templateName)

	err := template.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}