package apiserver

import "errors"

var (
	errConnDB = errors.New("Error in connection to database. Maybe database server is not running")
)

// errors in parsing json
var (
	errParseRequest = errors.New("Error in parsing request. Incorrect data")
)
