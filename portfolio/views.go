package portfolio

import (
	"net/http"
 	view "github.com/ShubhamSingh20/Portfolio/utils"

)



func homePageHandler(w http.ResponseWriter, r *http.Request) {
	view.SimpleTemplateView(w, "portfolio/index.html", nil)
}