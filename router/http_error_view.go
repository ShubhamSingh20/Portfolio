package router

import (
	"path"
	"net/http"
	"html/template"
)


var (
	templatePath = path.Join("router", "templates")
)


//Custom404ErrorView Raise not found error
func Custom404ErrorView(w http.ResponseWriter, r *http.Request) {
	
	custom404 := template.Must(template.ParseFiles(
		path.Join(templatePath, "404_error.html"),
	))

	err := custom404.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}


//Custom403ErrorView Raise bad request error
func Custom403ErrorView(w http.ResponseWriter, r *http.Request) {
	
	custom403 := template.Must(template.ParseFiles(
		path.Join(templatePath, "403_error.html"),
	))

	err := custom403.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}