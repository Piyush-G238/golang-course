package models

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"noteTitle"`
	Content   string    `json:"noteContent"`
	CreatedAt time.Time `json:"created_at"`
}

func NewNotes(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title and content are mandatory")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) GetTitle() string {
	return note.Title
}

func (note Note) GetContent() string {
	return note.Content
}

func (note Note) Save() error {

	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName)
	jsonData, _error := json.Marshal(note)
	if _error != nil {
		return _error
	}

	return os.WriteFile(fileName+".json", jsonData, os.ModePerm)
}
