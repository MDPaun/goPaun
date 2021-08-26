package staff

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/MDPaun/goPaun/cmd/config"
	models "github.com/MDPaun/goPaun/pkg/account/staff"
)

func GetStaff(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if id != 0 {
			if err != nil || id < 1 {
				env.NotFound(w)
				return
			}
			log.Println(id)
			s, err := env.Staff.FindByID(id)
			if err != nil {
				if errors.Is(err, models.ErrNoRecord) {
					env.NotFound(w)
				} else {
					env.ServerError(w, err)
				}
				return
			}
			// Use the new render helper.
			type TemplateData = config.TemplateData
			env.Render(w, r, "show.page.tmpl", &TemplateData{
				Staff: s,
			})
			// data := &templateData{Staff: s}

			// files := []string{
			// 	"./../ui/html/show.page.tmpl",
			// 	"./../ui/html/base.layout.tmpl",
			// 	"./../ui/html/footer.partial.tmpl",
			// }

			// ts, err := template.ParseFiles(files...)
			// if err != nil {
			// 	env.ServerError(w, err)
			// 	return
			// }

			// err = ts.Execute(w, data)
			// if err != nil {
			// 	env.ServerError(w, err)
			// }

		} else {
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
		}
	}
}

func CreateStaff(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			log.Println("Not Allowed")
			return
		}
		// Create some variables holding dummy data. We'll remove these later on
		// during the build.
		email := "O snail"
		fullname := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
		image := "7"
		status := true
		date_added := time.Now()

		// Pass the data to the StaffModel.Create() method
		err := env.Staff.Create(email, fullname, image, status, date_added)
		if err != nil {
			log.Fatal(err)
			return
		}
		// Redirect the user to the relevant page for the snippet.
		http.Redirect(w, r, fmt.Sprintln("/staff"), http.StatusSeeOther)
	}
}
