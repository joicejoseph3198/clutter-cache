package repository

import "errors"

// ErrNotFound is returned when a requested record is not found.
var ErrorNotFound = errors.New("record not found")