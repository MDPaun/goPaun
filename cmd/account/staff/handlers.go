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
			fmt.Fprintf(w, "%v", s)
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
			for _, staff := range s {
				fmt.Fprintf(w, "%v\n", staff)
			}
		}
	}
}

// func readAll(env *config.Env) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method != http.MethodGet {
// 			// env.notFound(w)
// 			// env.ErrorLog.Fatal()
// 			w.WriteHeader(405)
// 			w.Write([]byte("Method Not Allowed"))
// 			return
// 		}

// 		s, err := env.Staff.Latest()
// 		if err != nil {
// 			// app.serverError(w, err)
// 			env.ErrorLog.Fatal(err)
// 			return
// 		}
// 		for _, staff := range s {
// 			fmt.Fprintf(w, "%v\n", staff)
// 		}

// 	}
// }

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
