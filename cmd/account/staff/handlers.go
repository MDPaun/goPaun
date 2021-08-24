package staff

import (
	"fmt"
	"net/http"

	"github.com/MDPaun/goPaun/cmd/config"
)

func read(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/read" {
			// env.notFound(w)
			env.ErrorLog.Fatal()
			return
		}

		s, err := env.Staff.Latest()
		if err != nil {
			// app.serverError(w, err)
			env.ErrorLog.Fatal(err)
			return
		}
		for _, staff := range s {
			fmt.Fprintf(w, "%v\n", staff)
		}

	}
}
