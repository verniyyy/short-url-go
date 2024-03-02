package src

import (
	"fmt"
	"log/slog"
	"net/http"
)

// Router ...
type Router struct {
	mux *http.ServeMux
}

// MakeRouter ...
func MakeRouter() Either[Router] {
	return Right(Router{
		mux: http.NewServeMux(),
	})
}

// BuildRouter ...
func BuildRouter(r Router) Either[Router] {
	return Right(r).Bind(Get("/shorten-url", URLShortenHandler))
}

// Serve ...
func Serve(port uint16) Transform[Router, Either[Router]] {
	return func(r Router) Either[Router] {
		slog.Info("Listening on", "port", port)
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), r.mux)
		if err != nil {
			return Left[Router](err)
		}
		return Right(r)
	}
}

// Get add route of get method
func Get(path string, h http.HandlerFunc) Transform[Router, Either[Router]] {
	return func(r Router) Either[Router] {
		router := r
		router.mux.HandleFunc(fmt.Sprintf("%s %s", http.MethodGet, path), h)
		return Right(router)
	}
}

// URLShortenHandler ...
var URLShortenHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success!\n"))
}
