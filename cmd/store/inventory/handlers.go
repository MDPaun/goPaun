package inventory

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/MDPaun/goPaun/cmd/config"
	models "github.com/MDPaun/goPaun/pkg/store/inventory"
)

func GetItems(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if id != 0 {
			if err != nil || id < 1 {
				env.NotFound(w)
				return
			}
			log.Println(id)
			s, err := env.Inventory.FindByID(id)
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
				Inventory: s,
			})

		} else {
			if r.Method != http.MethodGet {
				env.NotFound(w)
				w.WriteHeader(405)
				w.Write([]byte("Method Not Allowed"))
				return
			}

			s, err := env.Inventory.Latest()
			if err != nil {
				env.ServerError(w, err)
				return
			}
			// for _, Inventory := range s {
			// 	fmt.Fprintf(w, "%v\n", Inventory)
			// }
			type TemplateData = config.TemplateData
			env.Render(w, r, "home1.page.tmpl", &TemplateData{
				Inventorys: s,
			})
		}
	}
}

func UpdateStock(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			log.Println("Not Allowed")
			return
		}
		// Create some variables holding dummy data. We'll remove these later on
		// during the build.
		stock := 16
		id := 100557

		// Pass the data to the InventoryModel.Create() method
		err := env.Inventory.UpdateStock(id, stock)
		if err != nil {
			log.Fatal(err)
			return
		}
		// Redirect the user to the relevant page for the snippet.
		http.Redirect(w, r, fmt.Sprintln("/inventory"), http.StatusSeeOther)
	}
}
