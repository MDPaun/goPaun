package main

import (
	"fmt"
	"net/http"

	"github.com/MDPaun/goPaun/cmd/config"
)

func routes(env *config.Env) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home(env))
	// mux.HandleFunc("/read", env.read)

	fileServer := http.FileServer(http.Dir("./../ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return mux

}

func home(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
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

		// files := []string{
		// 	"./../ui/html/home.page.tmpl",
		// 	"./../ui/html/base.layout.tmpl",
		// 	"./../ui/html/footer.partial.tmpl",
		// }

		// ts, err := template.ParseFiles(files...)
		// if err != nil {
		// 	log.Println(err.Error())
		// 	http.Error(w, "Internal Server Error", 500)
		// 	return
		// }
		// err = ts.Execute(w, nil)
		// if err != nil {
		// 	// app.serverError(w, err)
		// 	return
		// }

	}
}
