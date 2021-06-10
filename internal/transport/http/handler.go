package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler - store pointer to  our comments services
type Handler struct {
	Router *mux.Router
}

// NewHandler - return a pointer to Handler
func NewHandler() *Handler {
	return &Handler{}
}
func (h *Handler) SetUphandler() {
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive!")
	})
}
