package main

// import (
// 	"fmt"
// 	"net/http"
// 	"runtime/debug"

// 	"github.com/MDPaun/goPaun/cmd/config"
// )

// // The serverError helper writes an error message and stack trace to the errorLog,
// // then sends a generic 500 Internal Server Error response to the user.
// func (env *config.Env) serverError(w http.ResponseWriter, err error) {
// 	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
// 	env.errorLog.Output(2, trace)
// 	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// }

// // The clientError helper sends a specific status code and corresponding description
// // to the user. We'll use this later in the book to send responses like 400 "Bad
// // Request" when there's a problem with the request that the user sent.
// func (app *application) clientError(w http.ResponseWriter, status int) {
// 	http.Error(w, http.StatusText(status), status)
// }

// // For consistency, we'll also implement a notFound helper. This is simply a
// // convenience wrapper around clientError which sends a 404 Not Found response to
// // the user.
// func (app *application) notFound(w http.ResponseWriter) {
// 	app.clientError(w, http.StatusNotFound)
// }
