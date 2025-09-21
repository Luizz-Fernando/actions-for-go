package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingEndpoint(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "pong")
}

func TestGetUsers(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Luiz")
	assert.Contains(t, w.Body.String(), "Fernando")
}

func TestGetUserByIDSuccess(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Luiz")
}

func TestGetUserByIDNotFound(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/users/99", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Contains(t, w.Body.String(), "user not found")
}

func TestCreateUser(t *testing.T) {
	router := setupRouter()

	body := []byte(`{"id":3,"name":"Matheus"}`)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Contains(t, w.Body.String(), "Matheus")
}
