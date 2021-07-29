package endpoint

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type FilmEndpoint struct {
}

func filmHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf("Film %s", id)
	fmt.Fprint(w, response)
}

func (*FilmEndpoint) Register(handle func(path string, handler func(http.ResponseWriter, *http.Request)) *mux.Route) {
	handle("/film/{id:[0-9]+}", filmHandler)
}
