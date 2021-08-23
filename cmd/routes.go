package main

import (
	"net/http"

	"github.com/MDPaun/goPaun/cmd/config"
)

func (env *config.Env) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", env.home)
	mux.HandleFunc("/read", env.read)

	fileServer := http.FileServer(http.Dir("./../ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return mux

}
