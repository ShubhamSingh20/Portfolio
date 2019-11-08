package blog

import (
	"fmt"
	"net/http"

	auth "github.com/ShubhamSingh20/Portfolio/auth"
	view "github.com/ShubhamSingh20/Portfolio/utils/view"
)

func adminLoginHandler(w http.ResponseWriter, r *http.Request) {
	view.SimpleTemplateView(w, nil, "blog", "admin_login.html")
}

func authenticateUserHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	cred := &auth.Credentials{
		Username: username,
		Password: password,
	}

	isUserValid := auth.AuthenticateUser(cred)

	if !isUserValid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

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
