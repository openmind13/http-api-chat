package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/openmind13/http-api-chat/app/store"
)

// Server struct
type server struct {
	router *mux.Router
	store  *store.SQLStore
}

// New - init server struct
func newServer(store *store.SQLStore) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
	}

	// set middleware
	s.router.Use(commonMiddleware)
	s.configureRouter()

	return s
}

// Start http handling
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/users/add", s.handleAddUser).Methods("POST")
	s.router.HandleFunc("/users/get", s.handleGetUsers).Methods("GET")

	s.router.HandleFunc("/chats/add", s.handleAddChat).Methods("POST")
	s.router.HandleFunc("/chats/get", s.handleGetChats).Methods("POST")

	s.router.HandleFunc("/messages/add", s.handleAddMessage).Methods("POST")
	s.router.HandleFunc("/messages/get", s.handleGetChatMessages).Methods("POST")
}
