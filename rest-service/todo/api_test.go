package todo

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func TestAPI_GetTodoByID(t *testing.T) {

	testRepo := &MapRepository{todos: map[int]*Todo{
		1: {
			ID:          1,
			Description: "First test Todo",
			IsDone:      false,
			DueDate:     time.Time{},
		},
		2: {
			ID:          2,
			Description: "Second test Todo",
			IsDone:      true,
			DueDate:     time.Date(2020, time.April, 22, 15, 0, 0, 0, time.UTC),
		}},
	}

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
							"Description": "First test Todo", 
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
							"Description": "Second test Todo",
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

			testAPI.GetTodoByID(c)

			actualHTTPCode := w.Code
			actualBody := w.Body.String()

			assert.Equal(t, tt.wantHTTPCode, actualHTTPCode)
			assert.JSONEq(t, tt.wantBody, actualBody)
		})

	}
}
