package base

import (
	"net/http"

	"github.com/MDPaun/goPaun/cmd/config"
)

func home(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			env.NotFound(w)
		}
		if r.Method != http.MethodGet {
			env.NotFound(w)
			w.WriteHeader(405)
			w.Write([]byte("Method Not Allowed"))
			return
		}

		s, err := env.Staff.Latest()
		if err != nil {
			env.ServerError(w, err)
			return
		}
		// for _, staff := range s {
		// 	fmt.Fprintf(w, "%v\n", staff)
		// }
		type TemplateData = config.TemplateData
		env.Render(w, r, "home.page.tmpl", &TemplateData{
			Staffs: s,
		})

		// type TemplateData = config.TemplateData //! de verificat daca functioneaza
		// data := &TemplateData{Staffs: s}
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
		// err = ts.Execute(w, data)
		// if err != nil {
		// 	env.ServerError(w, err)
		// 	return
		// }

	}
}
