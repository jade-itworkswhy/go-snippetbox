package models

import (
	"database/sql"
	"errors"
	"time"
)

type SnippetModelInterface interface {
	Insert(title string, content string, expires int) (int, error)
	Get(id int) (Snippet, error)
	Latest() ([]Snippet, error)
}

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

/**
SQL -> Go types
CHAR, VARCHAR and TEXT map to string.
BOOLEAN maps to bool.
INT maps to int; BIGINT maps to int64.
DECIMAL and NUMERIC map to float.
TIME, DATE and TIMESTAMP map to time.Time. <- only if with the param parseTime=true, without it, it returns []byte objects.
ref: https://github.com/go-sql-driver/mysql#parameters
*/
// read specific snippet based on its id
func (m *SnippetModel) Get(id int) (Snippet, error) {

	/**
	shorter, cleaner version
	err := m.DB.QueryRow("SELECT ...", id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
	if errors.Is(err, sql.ErrNoRows) {
	return Snippet{}, ErrNoRecord
	} else {
	return Snippet{}, err
	}
	}
	*/
	// Write the SQL statement we want to execute. Again, I've split it over two
	// lines for readability.
	stmt := `
		SELECT id, title, content, created, expires FROM snippets
		WHERE expires > UTC_TIMESTAMP() AND id = ?
	`
	// Use the QueryRow() method on the connection pool to execute our
	// SQL statement, passing in the untrusted id variable as the value for the
	// placeholder parameter. This returns a pointer to a sql.Row object which
	// holds the result from the database.
	row := m.DB.QueryRow(stmt, id)
	// Initialize a new zeroed Snippet struct.
	var s Snippet
	// Use row.Scan() to copy the values from each field in sql.Row to the
	// corresponding field in the Snippet struct. Notice that the arguments
	// to row.Scan are *pointers* to the place you want to copy the data into,
	// and the number of arguments must be exactly the same as the number of
	// columns returned by your statement.
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// If the query returns no rows, then row.Scan() will return a
		// sql.ErrNoRows error. We use the errors.Is() function check for that
		// error specifically, and return our own ErrNoRecord error
		// instead (we'll create this in a moment).
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}
	// If everything went OK, then return the filled Snippet struct.
	return s, nil
}

// get the latest 10 snippets.
func (m *SnippetModel) Latest() ([]Snippet, error) {
	// Write the SQL statement we want to execute.
	stmt := `
		SELECT id, title, content, created, expires 
		FROM snippets
		WHERE expires > UTC_TIMESTAMP() 
		ORDER BY id DESC 
		LIMIT 10
	`
	// Use the Query() method on the connection pool to execute our
	// SQL statement. This returns a sql.Rows resultset containing the result of
	// our query.
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	// We defer rows.Close() to ensure the sql.Rows resultset is
	// always properly closed before the Latest() method returns. This defer
	// statement should come *after* you check for an error from the Query()
	// method. Otherwise, if Query() returns an error, you'll get a panic
	// trying to close a nil resultset.
	defer rows.Close()
	// Initialize an empty slice to hold the Snippet structs.
	var snippets []Snippet
	// Use rows.Next to iterate through the rows in the resultset. This
	// prepares the first (and then each subsequent) row to be acted on by the
	// rows.Scan() method. If iteration over all the rows completes then the
	// resultset automatically closes itself and frees-up the underlying
	// database connection.
	for rows.Next() {
		// Create a pointer to a new zeroed Snippet struct.
		var s Snippet
		// Use rows.Scan() to copy the values from each field in the row to the
		// new Snippet object that we created. Again, the arguments to row.Scan()
		// must be pointers to the place you want to copy the data into, and the
		// number of arguments must be exactly the same as the number of
		// columns returned by your statement.
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		// Append it to the slice of snippets.
		snippets = append(snippets, s)
	}
	// When the rows.Next() loop has finished we call rows.Err() to retrieve any
	// error that was encountered during the iteration. It's important to
	// call this - don't assume that a successful iteration was completed
	// over the whole resultset.
	if err = rows.Err(); err != nil {
		return nil, err
	}
	// If everything went OK then return the Snippets slice.
	return snippets, nil
}
