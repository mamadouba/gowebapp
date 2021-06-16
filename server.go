package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		host := r.RemoteAddr
		method := r.Method
		uri := r.URL.String()
		status := http.StatusOK
		end := time.Since(start)
		fmt.Printf("host=%s method=%s uri=%s status=%d time=%s\n", host, method, uri, status, end)
	})
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Autentication middleware")
		w.Header().Add("Content-Type", "application/json")
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Could not find authorization header",
			})
		}
		next.ServeHTTP(w, r)
	})
}
func handleRoot(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "Hellow world")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"username": "foo",
		"email":    "foo@dev.com",
		"age":      30,
	})
}
func main() {
	//mux := http.DefaultServeMux
	//mux.Handle("/", Auth(Logger(http.HandlerFunc(handleRoot))))
	//log.Fatal(http.ListenAndServe(":8000", mux))

	str := "id=100 email=foo@bar.com role=1"
	var email string
	fmt.Sscanf(str, "id=%d email=%s role=%d", nil, &email, nil)
	arr := strings.Split(strings.Split(str, " ")[1], "=")[1]
	fmt.Print(email, arr)
}
