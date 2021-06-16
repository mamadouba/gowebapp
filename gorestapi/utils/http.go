package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ValidationError struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

type ResponseWriter struct {
	http.ResponseWriter
	status   int
	size     int
	username string
}

func (rw *ResponseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += size
	return size, err
}

func (rw *ResponseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *ResponseWriter) SetUsername(usr string) {
	rw.username = usr
}

func (rw *ResponseWriter) Size() int {
	return rw.size
}

func (rw *ResponseWriter) Username() string {
	if rw.username == "" {
		return "-"
	}
	return rw.username
}

func (rw *ResponseWriter) Status() int {
	return rw.status
}
func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{ResponseWriter: w, status: http.StatusOK}
}

func RespondJSON(w http.ResponseWriter, code int, payload interface{}) {
	switch v := payload.(type) {
	case string:
		payload = Response{code, v}
	case error:
		payload = Response{code, v.Error()}
	}
	response, err := json.Marshal(payload)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondValidationError(w http.ResponseWriter, errs []string) {
	RespondJSON(w, 400, ValidationError{400, "Payload validation error", errs})
}

func DecodeJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	if r.ContentLength == 2 {
		return fmt.Errorf("Request body must not be empty")
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1024*1024)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		switch {
		case errors.Is(err, io.ErrUnexpectedEOF):
			return fmt.Errorf("Request body contains mal-formed json")
		case errors.As(err, &syntaxError):
			return fmt.Errorf("Request body contains mal-formed json at position %d", syntaxError.Offset)
		case errors.As(err, &unmarshalTypeError):
			return fmt.Errorf("Request body contains an invalid value for the %q filed at position %d", unmarshalTypeError.Field, unmarshalTypeError.Offset)
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			field := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("Request body contains unknown field %s", field)
		case errors.Is(err, io.EOF):
			return fmt.Errorf("Request body must not be empty")
		case err.Error() == "http: request body too large":
			return fmt.Errorf("Request body must not be larger than 1MB")
		default:
			log.Println(err.Error())
			return fmt.Errorf("An error occured while decoding request body")
		}
	}
	return nil
}
