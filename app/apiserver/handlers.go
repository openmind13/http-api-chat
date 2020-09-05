package apiserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/openmind13/http-api-chat/app/model"
)

// POST
// http://localhost:9000/users/add
func (s *server) handleAddUser(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	type request struct {
		Username string `json:"username"`
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, r, http.StatusBadRequest, errParseRequest)
		return
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

// GET
// http:/localhost:9000/users/get
func (s *server) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	users, err := s.store.GetAllUsers()
	if err != nil {
		s.error(w, r, http.StatusNotFound, err)
		return
	}

	s.respondJSON(w, r, http.StatusOK, users)
}

// POST
// http://localhost:9000/chats/add
func (s *server) handleAddChat(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name  string `json:"name"`
		Users []int  `json:"users"`
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, r, http.StatusBadRequest, errParseRequest)
		return
	}

	chat := &model.Chat{
		Name: req.Name,
	}

	if err := s.store.AddUsersIntoChat(chat, req.Users); err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	s.respondJSON(w, r, http.StatusCreated, chat.ID)
}

// POST
// http://localhost:9000/chats/get
func (s *server) handleGetChats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type request struct {
		User int `json:"user"`
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, r, http.StatusBadRequest, errParseRequest)
		return
	}

	chats, err := s.store.GetAllUserChats(req.User)
	if err != nil {
		s.error(w, r, http.StatusNotFound, err)
		return
	}

	s.respondJSON(w, r, http.StatusOK, chats)
}

// POST
// http://localhost:9000/messages/add
func (s *server) handleAddMessage(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Chat   int    `json:"chat"`
		Author int    `json:"author"`
		Text   string `json:"text"`
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, r, http.StatusBadRequest, errParseRequest)
		return
	}

	message := &model.Message{
		Chat:   req.Chat,
		Author: req.Author,
		Text:   req.Text,
	}

	if err := s.store.AddMessageIntoChat(message); err != nil {
		// handle error
		s.error(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	s.respondJSON(w, r, http.StatusOK, message.ID)
}

// POST
// http://localhost:9000/messages/get
func (s *server) handleGetChatMessages(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Chat int `json:"chat"`
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, r, http.StatusBadRequest, errParseRequest)
		return
	}

	s.respondJSON(w, r, http.StatusInternalServerError, nil)
}
