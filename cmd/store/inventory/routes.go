package inventory

import (
	"net/http"

	"github.com/MDPaun/goPaun/cmd/config"
)

func Routes(env *config.Env, mux *http.ServeMux) *http.ServeMux {

	// mux.HandleFunc("/staff", readAll(env))
	mux.HandleFunc("/inventory", GetItems(env))
	mux.HandleFunc("/inventory/update", UpdateStock(env))

	return mux
}
