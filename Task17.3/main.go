package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
		body:           &bytes.Buffer{},
	}
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logFile.Close()

	l := log.New(logFile, "", log.LstdFlags)
	logHandler := logMiddleware(l)

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: logHandler(authHandler(mux)),
	}
	if err := httpServer.ListenAndServe(); err != nil {
		l.Fatalln(fmt.Errorf("Не удалось запустить сервер: %w", err))
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	msg := "Hello, Go!"
	fmt.Fprint(res, msg)
}

func authHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xId := r.Header.Get("x-my-app-id")
		if xId != "my_secret" {
			http.Error(w, "пользователь не авторизован ", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func logMiddleware(l *log.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			crw := newResponseWriter(w)
			h.ServeHTTP(crw, r)
			l.Println("url:", r.URL)
			l.Println("Response Status:", crw.statusCode)
			l.Println("Response Body:", crw.body.String())
		})
	}
}
