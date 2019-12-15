package router

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestIndex(t *testing.T) {
	r := SetupRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, w.Body.String(), "Hello gin!")
}

var postForm = func(r *gin.Engine, url string, values url.Values) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, url, bytes.NewBufferString(values.Encode()))
	// format: form-urlencoded
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	r.ServeHTTP(w, req)
	return w
}

func TestUserRegister(t *testing.T) {
	r := SetupRouter()

	values := url.Values{}
	values.Add("email", "rchenhyy@outlook.com")
	values.Add("password", "abc123")
	values.Add("password-again", "-") // wrong

	var w *httptest.ResponseRecorder
	w = postForm(r, "/user/register", values)
	assert.Equal(t, w.Code, http.StatusBadRequest)
	t.Logf("Response Text: %s", w.Body.String())

	values.Set("password-again", "abc123") // right
	w = postForm(r, "/user/register", values)
	assert.Equal(t, w.Code, http.StatusOK)
	t.Logf("Response Text: %s", w.Body.String())
}

func TestUserLogin(t *testing.T) {
	r := SetupRouter()

	values := url.Values{}
	values.Add("email", "rchenhyy@outlook.com")
	values.Add("password", "-")

	var w *httptest.ResponseRecorder
	w = postForm(r, "/user/login", values)
	assert.Equal(t, w.Code, http.StatusUnauthorized)

	values.Set("password", "abc123")
	w = postForm(r, "/user/login", values)
	assert.Equal(t, w.Code, http.StatusOK)
}
