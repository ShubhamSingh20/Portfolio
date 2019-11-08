package view

import (
	"net/http"
)

//Custom404ErrorView Raise not found error
func Custom404ErrorView(w http.ResponseWriter, r *http.Request) {
	SimpleTemplateView(w, nil, "router", "404_error.html")

}


//Custom403ErrorView Raise bad request error
func Custom403ErrorView(w http.ResponseWriter, r *http.Request) {
	SimpleTemplateView(w, nil, "router", "403_error.html")

}