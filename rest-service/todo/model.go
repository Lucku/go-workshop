package todo

import "time"

// ToDo is a data model representing a to do on a bucket list
type ToDo struct {
	ID          int
	Description string
	IsDone      bool
	DueDate     time.Time
}
