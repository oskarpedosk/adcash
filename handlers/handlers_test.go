package handlers

import (
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/go-chi/chi/v5"
)

type postData struct {
	key   string
	value string
}

var tests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"post-apply", "/apply", "POST", []postData{
		{key: "personal_id", value: "2"},
		{key: "name", value: "Testmaster 3000"},
		{key: "amount", value: "999"},
		{key: "term", value: "12"},
	}, 200},
}

func TestHandlers(t *testing.T) {
	mux := chi.NewRouter()
	mux.Get("/", Repo.Instructions)
	mux.Post("/loans", Repo.Loans)
	mux.Post("/apply", Repo.PostApply)
	
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	for _, test := range tests {
		if test.method == "GET" {
			response, err := ts.Client().Get(ts.URL + test.url)
			if err != nil {
				t.Fatal(err)
			}
			if response.StatusCode != test.expectedStatusCode {
				t.Errorf("For %s, expected %d, but got %d", test.name, test.expectedStatusCode, response.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range test.params {
				values.Add(x.key, x.value)
			}
			response, err := ts.Client().PostForm(ts.URL+test.url, values)
			if err != nil {
				t.Fatal(err)
			}
			if response.StatusCode != test.expectedStatusCode {
				t.Errorf("For %s, expected %d, but got %d", test.name, test.expectedStatusCode, response.StatusCode)
			}
		}
	}
}
