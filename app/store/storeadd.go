package store

import (
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

// AddChat - create chat and add users into it
func (s *SQLStore) AddChat(chat *model.Chat) (int, error) {
	chat.CreatedAt = time.Now()

	if err := chat.Validate(); err != nil {
		return 0, errChatNameIncorrect
	}

	if err := s.createChat(chat); err != nil {
		return 0, errChatNotCreated
	}

	users := []*model.User{}
	for _, username := range chat.Users {
		u, err := s.FindUserByUsername(username)
		if err != nil {
			return 0, errUsersNotFound
		}

		users = append(users, u)
	}

	for _, user := range users {
		_, err := s.db.Exec(
			"INSERT INTO chat_users (user_id, chat_id) VALUES ($1, $2) RETURNING id;",
			user.ID,
			chat.ID,
		)
		if err != nil {
			// delete chat from database
			//

			return 0, errInAddToChatUsers
		}
	}

	return chat.ID, nil
}

func (s *SQLStore) createChat(chat *model.Chat) error {
	if err := s.db.QueryRow(
		"INSERT INTO chats (name, created_at) VALUES ($1, $2) RETURNING id;",
		chat.Name,
		chat.CreatedAt,
	).Scan(&chat.ID); err != nil {
		return err
	}

	return nil
}
