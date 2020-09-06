package store

import (
	"github.com/openmind13/http-api-chat/app/model"
)

// FindUserByID ...
func (s *SQLStore) FindUserByID(userID int) (*model.User, error) {
	user := model.User{}

	if err := s.db.QueryRow(
		"SELECT id, username, created_at FROM users WHERE id = $1",
		userID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetAllUsers ...
func (s *SQLStore) GetAllUsers() ([]model.User, error) {
	rows, err := s.db.Query(
		"SELECT id, username, created_at FROM users ORDER BY created_at DESC",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		u := model.User{}

		if err := rows.Scan(
			&u.ID,
			&u.Username,
			&u.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

// GetAllUserChats ...
func (s *SQLStore) GetAllUserChats(userID int) ([]model.Chat, error) {
	// rows, err := s.db.Query(
	// 	`SELECT chats.id, chats.name, chats.created_at FROM chats
	// 	INNER JOIN
	// 	chat_users ON chat_users.chat_id = chats.id
	// 	WHERE chat_users.user_id = $1;`,
	// 	userID,
	// )

	rows, err := s.db.Query(
		`SELECT chats.id, chats.name, chats.created_at FROM chats
		JOIN chat_users ON chat_users.chat_id = chats.id
		LEFT JOIN messages ON messages.chat_id = chats.id AND 
		messages.created_at = (SELECT MAX(messages.created_at) FROM messages) 
		WHERE chat_users.user_id = $1 ORDER BY messages.created_at;`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []model.Chat
	for rows.Next() {
		chat := model.Chat{}

		if err := rows.Scan(
			&chat.ID,
			&chat.Name,
			&chat.CreatedAt,
		); err != nil {
			return nil, err
		}

		chats = append(chats, chat)
	}

	return chats, nil
}

// GetAllChatMessages ...
func (s *SQLStore) GetAllChatMessages(chatID int) ([]model.Message, error) {
	rows, err := s.db.Query(
		"SELECT * FROM messages WHERE chat_id = $1 ORDER BY created_at;",
		chatID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []model.Message
	for rows.Next() {
		msg := model.Message{}

		if err := rows.Scan(
			&msg.ID,
			&msg.Chat,
			&msg.Author,
			&msg.Text,
			&msg.CreatedAt,
		); err != nil {
			return nil, err
		}

		messages = append(messages, msg)
	}

	return messages, nil
}
