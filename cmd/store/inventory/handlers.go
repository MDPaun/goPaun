package inventory

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MDPaun/goPaun/cmd/config"
)

func GetFromDecoCraft(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, err := env.InventoryDC.GetProducts()
		if err != nil {
			env.ServerError(w, err)
			return
		}

		for _, inventory := range s {
			_, err := env.Inventory.GetBySKU(inventory.SKU)
			if err != nil {
				image := inventory.Image
				name := inventory.Name
				sku := inventory.SKU
				ean := inventory.EAN
				quantity := inventory.Quantity

				err = env.Inventory.AddProduct(image, name, sku, ean, quantity)
				if err != nil {
					log.Fatal(err)
					return
				}

			} else {
				fmt.Println("Product SKU:", inventory.SKU, "already exist")
			}
		}

		http.Redirect(w, r, fmt.Sprintln("/inventory"), http.StatusSeeOther)

	}
}

func GetProducts(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// id, err := strconv.Atoi(r.URL.Query().Get("id"))
		// if id != 0 {
		// 	if err != nil || id < 1 {
		// 		env.NotFound(w)
		// 		return
		// 	}
		// 	log.Println(id)
		// 	s, err := env.Inventory.GetByID(id)
		// 	if err != nil {
		// 		if errors.Is(err, models.ErrNoRecord) {
		// 			env.NotFound(w)
		// 		} else {
		// 			env.ServerError(w, err)
		// 		}
		// 		return
		// 	}
		// 	// Use the new render helper.
		// 	type TemplateData = config.TemplateData
		// 	env.Render(w, r, "show.page.tmpl", &TemplateData{
		// 		Inventory: s,
		// 	})

		// } else {
		env.Inventory.CountProduct()
		page := r.URL.Query().Get("page")
		// if err != nil {
		// 	env.NotFound(w)
		// 	return
		// }
		if r.Method != http.MethodGet {
			env.NotFound(w)
			w.WriteHeader(405)
			w.Write([]byte("Method Not Allowed"))
			return
		}

		if page == "" {
			page = "1"
		}

		s, err := env.Inventory.Latest(page)
		if err != nil {
			env.ServerError(w, err)
			return
		}
		// for _, Inventory := range s {
		// 	fmt.Fprintf(w, "%v\n", Inventory)
		// }
		type TemplateData = config.TemplateData
		env.Render(w, r, "inventory.page.html", &TemplateData{
			Inventorys: s,
		})
		// }
	}
}

func UpdateStock(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			log.Println("Not Allowed")
			return
		}
		err := r.ParseForm()
		if err != nil {
			env.ClientError(w, http.StatusBadRequest)
			return
		}
		quantity := r.PostForm.Get("stock")
		sku := r.PostForm.Get("sku")

		// Pass the data to the InventoryModel.Create() method
		err = env.Inventory.UpdateStock(sku, quantity)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = env.InventoryDC.UpdateStockDecoCraft(sku, quantity)
		if err != nil {
			log.Fatal(err)
			return
		}
		// Redirect the user to the relevant page for the snippet.
		http.Redirect(w, r, fmt.Sprintln("/inventory"), http.StatusSeeOther)
	}
}
