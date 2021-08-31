package inventory

import (
	"net/http"

	"github.com/MDPaun/goPaun/cmd/config"
)

func Routes(env *config.Env, mux *http.ServeMux) *http.ServeMux {

	mux.HandleFunc("/inventory", GetProducts(env))
	mux.HandleFunc("/inventory/getfromdecocraft/", GetFromDecoCraft(env))
	mux.HandleFunc("/inventory/update", UpdateStock(env))

	return mux
}
