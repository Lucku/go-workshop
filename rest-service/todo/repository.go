package todo

import (
	"fmt"
	"time"
)

type Repository interface {
	FindAll() []*Todo
	FindByID(id int) *Todo
	Save(newTodo *Todo) *Todo
	DeleteByID(id int) bool
}

func (r *MapRepository) String() string {
	return fmt.Sprintf("%v", r.todos)
}

type MapRepository struct {
	todos map[int]*Todo
}

func NewMapRepository() *MapRepository {
	return &MapRepository{todos: map[int]*Todo{
		1: {
			ID:          1,
			Description: "Do some Go training",
			IsDone:      false,
			DueDate:     time.Now().Add(time.Hour * 24),
		},
	}}
}

func (r *MapRepository) FindAll() []*Todo {

	todosList := make([]*Todo, 0, len(r.todos))

	for _, v := range r.todos {
		todosList = append(todosList, v)
	}

	return todosList
}

func (r *MapRepository) FindByID(id int) *Todo {

	if elem, ok := r.todos[id]; ok {
		return elem
	}

	return nil
}

func (r *MapRepository) Save(newTodo *Todo) *Todo {

	newID := len(r.todos) + 1
	newTodo.ID = newID
	r.todos[newID] = newTodo

	return newTodo
}

func (r *MapRepository) DeleteByID(id int) bool {

	if _, ok := r.todos[id]; ok {
		delete(r.todos, id)
		return true
	}

	return false
}
