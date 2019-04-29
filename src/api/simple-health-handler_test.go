package api_test

import (
	"github.com/hgsgtk/health-endpoint/src/api"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestSimpleHealthCheck(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/sample", nil)

	api.SimpleHealthCheck(w, r)
	res := w.Result()
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("ioutil.ReadAll() caused unexpected error '%#v'", err)
	}

}
