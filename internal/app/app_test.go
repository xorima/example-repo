package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodoHandlers(t *testing.T) {
	t.Run("it should add a new todo", func(t *testing.T) {
		app := NewApp()

		todo := "Test Todo"
		todoJSON, _ := json.Marshal(todo)
		req, err := http.NewRequest("POST", "/api/v1/todo", bytes.NewBuffer(todoJSON))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		app.router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code, "Expected status code to be 201 Created")
		assert.Contains(t, app.todos, todo, "Expected todo to be added to the list")
	})

	t.Run("it should get all todos", func(t *testing.T) {
		app := NewApp()

		app.todos = append(app.todos, "Test Todo")

		req, err := http.NewRequest("GET", "/api/v1/todo", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		app.router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code, "Expected status code to be 200 OK")

		var todos []string
		err = json.NewDecoder(rr.Body).Decode(&todos)
		assert.NoError(t, err)

		assert.Contains(t, todos, "Test Todo", "Expected todo to be in the response")
	})
	t.Run("it should return an error for invalid JSON", func(t *testing.T) {
		app := NewApp()

		invalidJSON := []byte(`{invalid json}`)
		req, err := http.NewRequest("POST", "/api/v1/todo", bytes.NewBuffer(invalidJSON))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		app.router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code, "Expected status code to be 400 Bad Request")
		assert.Contains(t, rr.Body.String(), "invalid character", "Expected error message for invalid JSON")
	})
}
