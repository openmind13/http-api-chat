package store

import (
	"errors"
	"fmt"

	"github.com/openmind13/http-api-chat/app/model"
)

// User Errors
var (
	errUserAlreadyExist = errors.New("User already exists")

	errUsernameIncorrect = fmt.Errorf(
		"Incorrect username. Min length = %d, max = %d",
		model.UsernameMinLength,
		model.UsernameMaxLength,
	)

	errUsersNotFound = errors.New("Not all users found. Create users at first")
)

// Chat Errors
var (
	errChatNameIncorrect = fmt.Errorf(
		"Incorrect chat name. Min length = %d, max = %d",
		model.ChatNameMinLength,
		model.ChatNameMaxLength,
	)

	errChatNotCreated = errors.New("Chat not created")
)

var (
	errInAddToChatUsers = errors.New("Error in add to chat users")
)
