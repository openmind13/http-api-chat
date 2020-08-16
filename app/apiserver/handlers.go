package apiserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/openmind13/http-api-chat/app/model"
)

func (s *server) handleAddUser(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Username string `json:"username"`
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, r, http.StatusBadRequest, err)
	}

	// create user model
	user := &model.User{
		Username: req.Username,
	}
	user.CreatedAt = time.Now()

	// store user in database
	id, err := s.store.AddUser(user)
	if err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
		return
	}
	user.ID = id
	s.respondJSON(w, r, http.StatusCreated, user.ID)
}

func (s *server) handleAddChat(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name  string   `json:"name"`
		Users []string `json:"users"`
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, r, http.StatusBadRequest, err)
	}

	chat := &model.Chat{
		Name:  req.Name,
		Users: req.Users,
	}

	id, err := s.store.AddChat(chat)
	if err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
		return
	}
	chat.ID = id
	s.respondJSON(w, r, http.StatusCreated, chat.ID)
}

func (s *server) handleAddMessage(w http.ResponseWriter, r *http.Request) {
	// type request struct {
	// 	chat   int    `json:"chat"`
	// 	author int    `json:"author"`
	// 	text   string `json:"text"`
	// }
}

func (s *server) handleGetUserChats(w http.ResponseWriter, r *http.Request) {

}

func (s *server) handleGetChatMessages(w http.ResponseWriter, r *http.Request) {

}
