package store

import (
	"github.com/openmind13/http-api-chat/app/model"
)

// FindUserByUsername ...
func (s *SQLStore) FindUserByUsername(username string) (*model.User, error) {
	user := &model.User{}

	if err := s.db.QueryRow(
		"SELECT id, username, created_at FROM users WHERE username = $1;",
		username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.CreatedAt,
	); err != nil {
		return nil, err
	}

	return user, nil

	// row, err := s.db.Query(
	// 	"SELECT id, username, created_at FROM users WHERE username = $1;",
	// 	username,
	// )
	// if err != nil {
	// 	return nil, err
	// }
	// defer row.Close()

	// for row.Next() {
	// 	if err := row.Scan(
	// 		&user.ID,
	// 		&user.Username,
	// 		&user.CreatedAt,
	// 	); err != nil {
	// 		return nil, err
	// 	}

	// 	return user, nil
	// }
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
func (s *SQLStore) GetAllUserChats(userID int) (*model.User, error) {

	return nil, nil
}

// GetAllChatMessages ...
func (s *SQLStore) GetAllChatMessages(chatID int) ([]string, error) {

	return nil, nil
}
