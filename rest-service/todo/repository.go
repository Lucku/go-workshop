package todo

import (
	"fmt"
	"time"
)

// Repository is an interface for data operations
type Repository interface {
	findAll() []*ToDo
	findByID(id int) *ToDo
	save(newToDo *ToDo) *ToDo
	deleteByID(id int) bool
}

// DefaultMapSize is the default size of the underlying map at initialization
const DefaultMapSize = 15

// String returns a string representation of a map repository
func (r *MapRepository) String() string {
	return fmt.Sprintf("MapRepository: %v", r.todos)
}

type MapRepository struct {
	todos map[int]*ToDo
}

// NewMapRepository initializes a new repository based on an internal map
//
// This is how you create a new map repository
//
//  var m *MapRepository = NewMapRepository()
//
// Input parameters
//
// -
//
// Return values
//
// - Pointer to newly initialized MapRepository
//
func NewMapRepository() *MapRepository {
	return &MapRepository{todos: map[int]*ToDo{
		1: {
			ID:          1,
			Description: "Do some Go training",
			IsDone:      false,
			DueDate:     time.Now().Add(time.Hour * 24),
		},
	}}
}

func (r *MapRepository) findAll() []*ToDo {

	todosList := make([]*ToDo, 0, len(r.todos))

	for _, v := range r.todos {
		todosList = append(todosList, v)
	}

	return todosList
}

func (r *MapRepository) findByID(id int) *ToDo {

	if elem, ok := r.todos[id]; ok {
		return elem
	}

	return nil
}

func (r *MapRepository) save(newToDo *ToDo) *ToDo {

	newID := len(r.todos) + 1
	newToDo.ID = newID
	r.todos[newID] = newToDo

	return newToDo
}

func (r *MapRepository) deleteByID(id int) bool {

	if _, ok := r.todos[id]; ok {
		delete(r.todos, id)
		return true
	}

	return false
}
