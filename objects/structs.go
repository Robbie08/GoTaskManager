package objects

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	DateCreated string `json:"datecreated"`
	DateDue     string `json:"datedue"`
	Assignee    string `json:"assignee"`
}
