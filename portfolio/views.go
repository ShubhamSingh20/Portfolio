package portfolio

import (
	view "github.com/ShubhamSingh20/Portfolio/utils/view"
	"net/http"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	view.SimpleTemplateView(w, nil, "portfolio", "index.html")
}
