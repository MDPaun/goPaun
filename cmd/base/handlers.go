package base

import (
	"net/http"

	"github.com/MDPaun/goPaun/cmd/config"
)

func home(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			env.NotFound(w)
			w.WriteHeader(405)
			w.Write([]byte("Method Not Allowed"))
			return
		}
		if r.URL.Path != "/" {
			env.NotFound(w)
			return
		}

		s, err := env.Inventory.Latest(1, 10, "", "", "", "")
		if err != nil {
			env.ServerError(w, err)
			return
		}
		// for _, staff := range s {
		// 	fmt.Fprintf(w, "%v\n", staff)
		// }
		type TemplateData = config.TemplateData
		env.Render(w, r, "admin.page.html", &TemplateData{
			Inventorys: s,
		})

	}
}
