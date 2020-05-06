package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) findAll() []*ToDo {

	args := m.Called()

	if args.Get(0) != nil {
		return args.Get(0).([]*ToDo)
	}

	return nil
}

func (m *MockRepository) findByID(id int) *ToDo {

	args := m.Called(id)

	if args.Get(0) != nil {
		return args.Get(0).(*ToDo)
	}

	return nil
}

func (m *MockRepository) save(newToDo *ToDo) *ToDo {

	args := m.Called(newToDo)

	if args.Get(0) != nil {
		return args.Get(0).(*ToDo)
	}

	return nil
}

func (m *MockRepository) deleteByID(id int) bool {

	args := m.Called(id)

	return args.Bool(0)
}

func TestAPI_GetToDoByID(t *testing.T) {

	testRepo := &MockRepository{}

	testRepo.On("findByID", 1).Return(&ToDo{
		ID:          1,
		Description: "First test ToDo",
		IsDone:      false,
		DueDate:     time.Time{},
	})

	testRepo.On("findByID", 2).Return(&ToDo{
		ID:          2,
		Description: "Second test ToDo",
		IsDone:      true,
		DueDate:     time.Date(2020, time.April, 22, 15, 0, 0, 0, time.UTC),
	})

	testRepo.On("findByID", 3).Return(nil)

/*
	Same as:

		testRepo := &MapRepository{todos: map[int]*ToDo{
			1: {
				ID:          1,
				Description: "First test ToDo",
				IsDone:      false,
				DueDate:     time.Time{},
			},
			2: {
				ID:          2,
				Description: "Second test ToDo",
				IsDone:      true,
				DueDate:     time.Date(2020, time.April, 22, 15, 0, 0, 0, time.UTC),
			}},
	}*/

	testAPI := NewAPI(testRepo)

	tests := []struct {
		giveRequestID string
		wantHTTPCode  int
		wantBody      string
	}{
		{
			giveRequestID: "1",
			wantHTTPCode:  http.StatusOK,
			wantBody: `
				{
					"todo": 
						{
							"ID": 1, 
							"Description": "First test ToDo", 
							"IsDone": false, 
							"DueDate": "0001-01-01T00:00:00Z"
					}
				}`,
		},
		{
			giveRequestID: "2",
			wantHTTPCode:  http.StatusOK,
			wantBody: `
				{
					"todo":
						{
							"ID": 2,
							"Description": "Second test ToDo",
							"IsDone": true,
							"DueDate": "2020-04-22T15:00:00Z"
						}
				}`,
		},
		{
			giveRequestID: "3",
			wantHTTPCode:  http.StatusOK,
			wantBody:      `{"todo": null}`,
		},
		{
			giveRequestID: "",
			wantHTTPCode:  http.StatusBadRequest,
			wantBody:      `{"error": "No ID parameter provided"}`,
		},
		{
			giveRequestID: "test",
			wantHTTPCode:  http.StatusBadRequest,
			wantBody:      `{"error": "ID not provided as a number"}`,
		},
	}

	for _, tt := range tests {

		t.Run(tt.giveRequestID, func(t *testing.T) {

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = gin.Params{
				{
					Key:   "id",
					Value: tt.giveRequestID,
				},
			}

			testAPI.GetToDoByID(c)

			actualHTTPCode := w.Code
			actualBody := w.Body.String()

			assert.Equal(t, tt.wantHTTPCode, actualHTTPCode)
			assert.JSONEq(t, tt.wantBody, actualBody)
		})

	}
}
