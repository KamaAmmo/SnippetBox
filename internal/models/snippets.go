package models

import (
	"database/sql"
	"time"
)

// Define a Snippet type to hold the data for an individual snippet. Notice how
// the fields of the struct correspond to the fields in our MySQL snippets
// table?
type Snippet struct {
	ID int
	Title string 
	Content string
	Created time.Time
	Expires  time.Time
}

// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert (title string, content string, expires int) (int , error){
	stmt := `INSERT INTO snippets (title , content, created, expires)
		VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil{
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil{
		return 0, err
	}

	return int(id), nil
}

func (m *SnippetModel) Get (id int) (*Snippet, error){
	return nil, nil
}

//return 10 latest created snippets 
func (m *SnippetModel) Latest () (*[]Snippet, error){
	return nil, nil
}