package store

import (
	"time"

	"github.com/openmind13/http-api-chat/app/model"
)

// AddUser to database
func (s *SQLStore) AddUser(user *model.User) (int, error) {
	user.CreatedAt = time.Now()

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
func (s *SQLStore) AddUsersIntoChat(chat *model.Chat) error {
	if err := s.CreateChat(chat); err != nil {
		return errChatNotCreated
	}

	users := []*model.User{}
	for _, username := range chat.Users {
		u, err := s.FindUserByUsername(username)
		if err != nil {
			// return 0, errUsersNotFound
			return err
		}

		users = append(users, u)
	}

	for _, user := range users {
		_, err := s.db.Exec(
			"INSERT INTO chat_users (user_id, chat_id) VALUES ($1, $2)",
			user.ID,
			chat.ID,
		)
		if err != nil {
			return errInAddToChatUsers
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

	return nil
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

// AddMessageIntoChat ...
func (s *SQLStore) AddMessageIntoChat(message *model.Message) error {
	message.CreatedAt = time.Now()

	if err := s.db.QueryRow(
		"INSERT INTO messages (chat_id, user_id, text, created_at) VALUES ($1, $2, $3, $4) RETURNING id;",
		message.Chat,
		message.Author,
		message.Text,
		message.CreatedAt,
	).Scan(&message.ID); err != nil {
		return err
	}

	return nil
}
