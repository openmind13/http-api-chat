package store

import (
	"fmt"
	"time"

	"github.com/openmind13/http-api-chat/app/model"
)

// AddUser to database
func (s *SQLStore) AddUser(user *model.User) (int, error) {
	user.BeforeAddUser()

	if err := user.Validate(); err != nil {
		return 0, errUsernameIncorrect
	}

	if err := s.db.QueryRow(
		"INSERT INTO users (username, created_at) VALUES ($1, $2) RETURNING id;",
		user.Username,
		user.CreatedAt,
	).Scan(&user.ID); err != nil {
		return 0, errUserAlreadyExist
	}

	return user.ID, nil
}

// AddUsersIntoChat - create chat and add users into it
func (s *SQLStore) AddUsersIntoChat(chat *model.Chat) (int, error) {
	if err := s.CreateChat(chat); err != nil {
		return 0, errChatNotCreated
	}

	users := []*model.User{}
	for _, username := range chat.Users {
		u, err := s.FindUserByUsername(username)
		if err != nil {
			// return 0, errUsersNotFound
			return 0, err
		}

		users = append(users, u)
	}

	fmt.Println()
	for _, user := range users {
		fmt.Println(user)
	}

	for _, user := range users {
		_, err := s.db.Exec(
			"INSERT INTO chat_users (user_id, chat_id) VALUES ($1, $2) RETURNING id;",
			user.ID,
			chat.ID,
		)
		if err != nil {
			return 0, errInAddToChatUsers
		}

		// if err := s.db.QueryRow(
		// 	"INSERT INTO chat_users (user_id, chat_id) VALUES ($1, $2) RETURNING id;",
		// 	user.ID,
		// 	chat.ID,
		// ).Scan(nil); err != nil {
		// 	// return 0, errInAddToChatUsers
		// 	return 0, err
		// }
	}

	return chat.ID, nil
}

// CreateChat ...
func (s *SQLStore) CreateChat(chat *model.Chat) error {
	chat.CreatedAt = time.Now()

	if err := chat.Validate(); err != nil {
		return errChatNameIncorrect
	}

	if err := s.db.QueryRow(
		"INSERT INTO chats (name, created_at) VALUES ($1, $2) RETURNING id;",
		chat.Name,
		chat.CreatedAt,
	).Scan(&chat.ID); err != nil {
		return err
	}

	return nil
}
