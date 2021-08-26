package staff

import (
	"net/http"

	"github.com/MDPaun/goPaun/cmd/config"
)

func Routes(env *config.Env, mux *http.ServeMux) *http.ServeMux {

	// mux.HandleFunc("/staff", readAll(env))
	mux.HandleFunc("/staff", GetStaff(env))
	mux.HandleFunc("/staff/create", CreateStaff(env))

	return mux
}
