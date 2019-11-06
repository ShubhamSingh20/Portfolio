package router

import (
	"net/http"
	view "github.com/ShubhamSingh20/Portfolio/utils"
)

//Custom404ErrorView Raise not found error
func Custom404ErrorView(w http.ResponseWriter, r *http.Request) {
	view.SimpleTemplateView(w, "router/404_error.html", nil)

}


//Custom403ErrorView Raise bad request error
func Custom403ErrorView(w http.ResponseWriter, r *http.Request) {
	view.SimpleTemplateView(w, "router/403_error.html", nil)

}