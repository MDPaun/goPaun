package staff

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"
// )

// func read(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	if err != nil || id < 1 {
// 		http.NotFound(w, r)
// 		return
// 	}
// 	// Use the fmt.Fprintf() function to interpolate the id value with our response
// 	// and write it to the http.ResponseWriter.
// 	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
// }
