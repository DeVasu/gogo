package main

import (
	"net/http"

	"github.com/unrolled/render"
)

func createMatchHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// w.Header().Add("Location", "value of the created resource")
		formatter.JSON(w,
			http.StatusCreated,
		struct { Test string}{"This is a test"})
	}
}