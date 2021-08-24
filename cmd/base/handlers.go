package base

import (
	"html/template"
	"log"
	"net/http"

	"github.com/MDPaun/goPaun/cmd/config"
)

func home(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			// env.notFound(w)
			env.ErrorLog.Fatal()
			return
		}

		files := []string{
			"./../ui/html/home.page.tmpl",
			"./../ui/html/base.layout.tmpl",
			"./../ui/html/footer.partial.tmpl",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			// app.serverError(w, err)
			return
		}

	}
}
