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

/**
DB.Query() is used for SELECT queries which return multiple rows.
DB.QueryRow() is used for SELECT queries which return a single row.
DB.Exec() is used for statements which don’t return rows (like INSERT and DELETE).
*/

// insert function
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	// ? acted as a placeholder for the data we want to insert.

	/**
		The placeholder parameter syntax differs depending on your database. MySQL, SQL Server
	and SQLite use the ? notation, but PostgreSQL uses the $N notation. For example, if you
	were using PostgreSQL instead you would write:
	*/
	// todo: any better way for this?
	stmt := `
		INSERT INTO snippets (title, content, created, expires)
		VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))
	`

	// return sql.Result type
	result, err := m.DB.Exec(stmt, title, content, expires)

	if err != nil {
		return 0, err
	}

	// get the last one's id

	/**
		tip: Not all drivers and databases support the LastInsertId() and
	RowsAffected() methods. For example, LastInsertId() is not supported by
	PostgreSQL. So if you’re planning on using these methods it’s important to check the
	documentation for your particular driver first.
	*/
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// read specific snippet based on its id
func (m *SnippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

// get the latest 10 snippets.
func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
