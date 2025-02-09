package srv

import (
	"net/http"

	"github.com/go-chi/render"
)

func readyHandler(w http.ResponseWriter, r *http.Request) {
	render.PlainText(w, r, "OK")
}

func liveHandler(w http.ResponseWriter, r *http.Request) {
	render.PlainText(w, r, "OK")
}
