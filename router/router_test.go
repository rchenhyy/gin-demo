package router

import (
	"bytes"
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

func TestUserRegister(t *testing.T) {
	r := SetupRouter()

	values := url.Values{}
	values.Add("email", "rchenhyy@outlook.com")
	values.Add("password", "abc123")
	values.Add("password-again", "-") // wrong

	serve := func(values url.Values) *httptest.ResponseRecorder {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewBufferString(values.Encode()))
		// format: form-urlencoded
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
		r.ServeHTTP(w, req)
		return w
	}

	var w *httptest.ResponseRecorder
	w = serve(values)
	assert.Equal(t, w.Code, http.StatusBadRequest)
	t.Logf("Response Text: %s", w.Body.String())

	values.Set("password-again", "abc123") // right
	w = serve(values)
	assert.Equal(t, w.Code, http.StatusOK)
	t.Logf("Response Text: %s", w.Body.String())
}
