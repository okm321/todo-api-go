package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func handleOption(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Add("Access-Control-Allow-Origin", "https://e7txy0.csb.app")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Headers", "Origin")
	w.Header().Add("Access-Control-Allow-Headers", "X-Requested-With")
	w.Header().Add("Access-Control-Allow-Headers", "Accept")
	w.Header().Add("Access-Control-Allow-Headers", "Accept-Language")
	w.Header().Set("Content-Type", "application/json")
}

func (app *Application) routes() *httprouter.Router {
	router := httprouter.New()
	router.OPTIONS("/*path", handleOption)
	router.HandlerFunc(http.MethodGet, "/api/todos", app.getAllTodos)
	router.HandlerFunc(http.MethodPost, "/api/todos", app.editTodo)
	router.HandlerFunc(http.MethodGet, "/api/todos/:id", app.getOneTodo)
	router.HandlerFunc(http.MethodDelete, "/api/todos/:id", app.deleteTodo)

	return router
}
