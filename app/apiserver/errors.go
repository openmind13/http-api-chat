package apiserver

import "errors"

var (
	errConnDB = errors.New("Error in connection to database. Maybe database server is not running")
)
