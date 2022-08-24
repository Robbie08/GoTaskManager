package objects

import (
	"database/sql"
	"fmt"
)

func (t *Task) Init(title, dateCreated, dateDue, assignee string) {
	t.Title = title
	t.DateCreated = dateCreated
	t.DateDue = dateDue
	t.Assignee = assignee
}

func (t *Task) AddTask(db *sql.DB) {
	fmt.Printf("Adding Task...")
	sqlStatement := `
	INSERT INTO tasks (title, datecreated, datedue, assignee)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 1

	err := db.QueryRow(sqlStatement, t.Title, t.DateCreated, t.DateDue, t.Assignee).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}

func (t *Task) insertQuery(table string) string {
	return fmt.Sprintf("INSERT INTO %s (title, datecreated, datadue, assignee) VALUES('%s', '%s', '%s', '%s')", table, t.Title, t.DateCreated, t.DateDue, t.Assignee)
}
