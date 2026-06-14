package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// JSON is repsonse helper functions
func JSON(w http.ResponseWriter, code int, payload any, headers map[string]string) error {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(payload)
}

// Bind binds the request body to given model
func Bind[M any](r *http.Request) (M, error) {
	var m M
	return m, json.NewDecoder(r.Body).Decode(&m)
}

// Custom Handler func that Defines error on top of http.HandlerFunc
type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		if errors.As(err, &AppErr{}) {
			e := err.(AppErr)
			fmt.Println(err)
			JSON(w, e.status, e, map[string]string{"Content-Type": "application/json"})
		}
		return
	}
}

// app err which implments custom
// JsonMarshaler and Error interface
type AppErr struct {
	err    error
	msg    string
	status int
}


func HTTPError(e error, m string, code int) AppErr {
	return AppErr{e, m, code}
}

func (e AppErr) Unwrap() error {
	return e.err
}

func (e AppErr) Error() string {
	return e.err.Error()
}

func (e AppErr) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Err        string `json:"err,omitempty"`
		Msg        string `json:"msg,omitempty"`
		StatusCode int    `json:"status_code,omitempty"`
	}{
		e.Error(), e.msg, e.status,
	})
}

// api struct
type API struct {
	router *http.ServeMux
}

func (api *API) GetE(w http.ResponseWriter, r *http.Request) error {
	return HTTPError(errors.New("this is error"), "just error", 400)
}

func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.router.ServeHTTP(w, r)
}

func (api *API) Routes() {
	api.router.Handle("/", HandlerFunc(api.GetE))
}

func main() {
	mux := http.NewServeMux()
	api := &API{mux}
	api.Routes()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: api,
	}
	srv.ListenAndServe()
}
