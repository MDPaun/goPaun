package main

import (
	"net/http"

	staff "github.com/MDPaun/goPaun/cmd/account/staff"
	base "github.com/MDPaun/goPaun/cmd/base"
	config "github.com/MDPaun/goPaun/cmd/config"
	inventory "github.com/MDPaun/goPaun/cmd/store/inventory"
)

func routes(env *config.Env) http.Handler {
	mux := http.NewServeMux()

	// mux.HandleFunc("/", home(env))
	base.Routes(env, mux)
	staff.Routes(env, mux)
	inventory.Routes(env, mux)

	fileServer := http.FileServer(http.Dir("./../static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// return mux
	return recoverPanic(env, logRequest(env, secureHeaders(mux)))
}
