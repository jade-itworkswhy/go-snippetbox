package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// define the model for wrapping the connection pool

type SnippetModel struct {
	DB *sql.DB
}

// insert function
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

// read specific snippet based on its id
func (m *SnippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

// get the latest 10 snippets.
func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
