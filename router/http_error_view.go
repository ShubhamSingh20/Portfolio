package router

import (
	"path"
	"net/http"

	view "github.com/ShubhamSingh20/Portfolio/utils"
)


var (
	templatePath = path.Join("router", "templates")
)


//Custom404ErrorView Raise not found error
func Custom404ErrorView(w http.ResponseWriter, r *http.Request) {
	view.SimpleTemplateView(w, path.Join(templatePath, "404_error.html"), nil)

}


//Custom403ErrorView Raise bad request error
func Custom403ErrorView(w http.ResponseWriter, r *http.Request) {
	view.SimpleTemplateView(w, path.Join(templatePath, "403_error.html"), nil)

}