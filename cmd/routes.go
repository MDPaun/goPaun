package main

import (
	"net/http"

	staff "github.com/MDPaun/goPaun/cmd/account/staff"
	base "github.com/MDPaun/goPaun/cmd/base"
	"github.com/MDPaun/goPaun/cmd/config"
)

func routes(env *config.Env) http.Handler {
	mux := http.NewServeMux()

	// mux.HandleFunc("/", home(env))
	base.Routes(env, mux)
	staff.Routes(env, mux)

	fileServer := http.FileServer(http.Dir("./../ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return recoverPanic(env, logRequest(env, secureHeaders(mux)))
}
