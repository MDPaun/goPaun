package main

// func (app *application) home(w http.ResponseWriter, r *http.Request) {
// func home(env *config.Env) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if r.URL.Path != "/" {
// 			// env.notFound(w)
// 			return
// 		}

// 		files := []string{
// 			"./../ui/html/home.page.tmpl",
// 			"./../ui/html/base.layout.tmpl",
// 			"./../ui/html/footer.partial.tmpl",
// 		}

// 		ts, err := template.ParseFiles(files...)
// 		if err != nil {
// 			log.Println(err.Error())
// 			http.Error(w, "Internal Server Error", 500)
// 			return
// 		}
// 		err = ts.Execute(w, nil)
// 		if err != nil {
// 			// app.serverError(w, err)
// 			return
// 		}
// 	}
// }

// func (app *application) read(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	if err != nil || id < 1 {
// 		http.NotFound(w, r)
// 		return
// 	}
// 	// Use the fmt.Fprintf() function to interpolate the id value with our response
// 	// and write it to the http.ResponseWriter.
// 	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
// }
