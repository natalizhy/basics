package apiserver

import (
	"github.com/natalizhy/basics/http-rest-api/internal/app/store/teststore"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestServer_HandleUsersCreate(t *testing.T){
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	s := newServer(teststore.New())
	s.ServeHTTP(rec, req)
	assert.Equel(t, rec.Code, http.StatusOK)
}
