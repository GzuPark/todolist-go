package db

import "errors"

// ErrorNotFound cannot find on the database
var ErrorNotFound = errors.New("db: not found")
