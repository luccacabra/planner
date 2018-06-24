package models

type Priority int
const (
	High   Priority = 0
	Medium Priority = 1
	Low    Priority = 2
)

type Check struct {
	Description string ``
	Done bool ``
}

type CheckList struct {

}

type SubTask struct {
	ID int `db:"id"`
	TaskID int `db:"task_id"`

	Title       string `db:"title"`
	Description string `db:"description"`

	Labels []string `db:"labels"`

	Priority Priority `db:"priority"`
}

type Task struct {
	ID int `db:"id"`

	Title       string `db:"title"`
	Description string `db:"description"`

	Labels []string `db:"labels"`

	Priority Priority `db:"priority"`

	SubTasks []*SubTask
}