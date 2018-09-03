package watchr

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Watchr contains application logic for Watchr
type Watchr struct {
	errors chan error
	host   string

	clients map[string]*connInfo
	rooms   map[string]*roomInfo
}

// NewWatchr generates a new Watchr application
func NewWatchr(host string) *Watchr {
	w := &Watchr{
		errors: make(chan error),
		host:   host,

		clients: make(map[string]*connInfo),
		rooms:   make(map[string]*roomInfo),
	}
	w.rooms["lobby"] = &roomInfo{
		mediaType: MediaText,
		media:     "welcome to my lobby :)",
	}
	return w
}

// Run the Watchr server
func (w *Watchr) Run() {
	r := chi.NewRouter()
	r.HandleFunc("/gateway", w.onGateway)

	srv := http.Server{
		Addr:    w.host,
		Handler: r,
	}
	e := srv.ListenAndServe()
	if e != nil {
		w.errors <- e
	}
}

// Errors returns the application's errors
func (w *Watchr) Errors() <-chan error {
	return w.errors
}
