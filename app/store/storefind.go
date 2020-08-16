package store

import "github.com/openmind13/http-api-chat/app/model"

// FindUserByUsername ...
func (s *SQLStore) FindUserByUsername(username string) (*model.User, error) {
	user := &model.User{}

	if err := s.db.QueryRow(
		"SELECT id, username, created_at FROM users WHERE username = $1",
		username,
	).Scan(user); err != nil {
		return nil, err
	}

	return user, nil
}
