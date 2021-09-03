package inventory

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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

// type FetchParam struct {
// 	PageNumber uint64
// }

func GetProducts(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pageStr := r.URL.Query().Get("page")
		defaultLimitStr := r.URL.Query().Get("defaultLimit")
		page, err := strconv.Atoi(pageStr)
		if err != nil && pageStr != "" {
			env.NotFound(w)
			return
		}
		if pageStr == "" {
			page = 1
		}

		defaultLimit, err := strconv.Atoi(defaultLimitStr)
		if err != nil && defaultLimitStr != "" {
			env.NotFound(w)
			return
		}
		if defaultLimitStr == "" {
			defaultLimit = 1
		}

		ts := env.Inventory.CountProduct()

		sortName := r.URL.Query().Get("sortName")
		sortSKU := r.URL.Query().Get("sortSKU")
		sortEAN := r.URL.Query().Get("sortEAN")
		sortOnHand := r.URL.Query().Get("sortOnHand")

		if r.Method != http.MethodGet {
			env.NotFound(w)
			w.WriteHeader(405)
			w.Write([]byte("Method Not Allowed"))
			return
		}

		s, err := env.Inventory.Latest(page, defaultLimit, sortName, sortSKU, sortEAN, sortOnHand)
		if err != nil {
			env.ServerError(w, err)
			return
		}

		fp := env.FilterProducts{}

		type TemplateData = config.TemplateData
		env.Render(w, r, "inventory.page.html", &TemplateData{
			Inventorys: s, FilterProducts: fp,
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
