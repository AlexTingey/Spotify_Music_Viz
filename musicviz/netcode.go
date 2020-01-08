package main

import "net/http"

type server struct{}

func (s *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	switch request.Method {
	case "GET":
		// w.WriteHeader(http.StatusOK)
		// w.Write([]byte)
	}
	writer.Write([]byte(`{"message": "hello world"}`))
}
