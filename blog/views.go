package blog

import (
	"net/http"
	"fmt"

	view "github.com/ShubhamSingh20/Portfolio/utils"
)


func adminLoginHandler(w http.ResponseWriter, r *http.Request) {
	view.SimpleTemplateView(w, nil, "blog", "admin_login.html")
}

func authenticateUser(w http.ResponseWriter, r *http.Request) {
	
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println(username, password)

}


func adminDashboardHandler(w http.ResponseWriter, r *http.Request) {
}



func createPostHandler(w http.ResponseWriter, r *http.Request) {
}



func previewPostHandler(w http.ResponseWriter, r *http.Request) {
}



func editPostHandler(w http.ResponseWriter, r *http.Request) {
}



func deletePostHandler(w http.ResponseWriter, r *http.Request) {
}


