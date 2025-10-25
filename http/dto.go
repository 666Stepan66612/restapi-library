package http

import (
	"encoding/json"
	"errors"
	"time"
)

type ReadBookDTO struct {
	Read bool
}

type BookDTO struct {
	Title string
	Author string
	Pages int
	Text  string
}

func (b BookDTO) ValidateForCreate() error {
	if b.Title == "" {
		return errors.New("title is empty")
	}

	if b.Text == "" || b.Pages == 0{
		return errors.New("book is empty")
	}

	return nil
}

type ErrorDTO struct {
	Message string
	Time time.Time
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil{
		panic(err)
	}

	return string(b)
}