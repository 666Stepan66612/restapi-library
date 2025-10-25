package mylibrary

import "errors"

var ErrBookNotFound = errors.New("book not found")
var ErrBookAlreadyInLibrary = errors.New("book already in library")