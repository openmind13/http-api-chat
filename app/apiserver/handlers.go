package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/openmind13/http-api-chat/app/model"
)

// http://localhost:9000/users/add
func (s *server) handleAddUser(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	type request struct {
		Username string `json:"username"`
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, r, http.StatusBadRequest, err)
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

func (s *server) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.store.GetAllUsers()
	if err != nil {
		s.error(w, r, http.StatusNotFound, err)
		return
	}

	var response bytes.Buffer

	if err := json.NewEncoder(&response).Encode(users); err != nil {
		fmt.Println(err)
	}

	for i, u := range users {
		fmt.Printf("%d %v\n", i, u.Username)
	}

	s.respondJSON(w, r, http.StatusOK, response)
}

// http://localhost:9000/chat/add
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

	id, err := s.store.AddUsersIntoChat(chat)
	if err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
		return
	}
	chat.ID = id
	s.respondJSON(w, r, http.StatusCreated, chat.ID)
}

// url
// http://localhost:9000/messages/add
func (s *server) handleAddMessage(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Chat   int    `json:"chat"`
		Author int    `json:"author"`
		Text   string `json:"text"`
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		return
	}

	// message := &model.Message{
	// 	Chat:   req.Chat,
	// 	Author: req.Author,
	// 	Text:   req.Text,
	// }

	// id, err := s.store.AddMessage(message)
	// if err != nil {
	// 	s.error
	// }
	s.respondJSON(w, r, http.StatusInternalServerError, nil)
}

func (s *server) handleGetUserChats(w http.ResponseWriter, r *http.Request) {

	s.respondJSON(w, r, http.StatusInternalServerError, nil)
}

func (s *server) handleGetChatMessages(w http.ResponseWriter, r *http.Request) {

	s.respondJSON(w, r, http.StatusInternalServerError, nil)
}
