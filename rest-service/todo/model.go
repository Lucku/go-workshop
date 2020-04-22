package todo

import "time"

type Todo struct {
	ID          int
	Description string
	IsDone      bool
	DueDate     time.Time
}
