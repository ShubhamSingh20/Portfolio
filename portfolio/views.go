package portfolio

import (
	"net/http"
	"html/template"
	"path"
	//"fmt"
)

var (
	templatePath = path.Join("portfolio", "templates")
)

//HomePageHandler handles the request for homepage of portfolio
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	portfolioTemplate := template.Must(template.ParseFiles(
		path.Join(templatePath, "index.html"),
	))

	err := portfolioTemplate.Execute(w, nil)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}