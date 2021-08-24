package base

import (
	"net/http"

	"github.com/MDPaun/goPaun/cmd/config"
)

func Routes(env *config.Env, mux *http.ServeMux) *http.ServeMux {

	mux.HandleFunc("/", home(env))

	return mux
}
