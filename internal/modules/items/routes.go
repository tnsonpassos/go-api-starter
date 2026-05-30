package items

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /items", ListItemsHandler)
	mux.HandleFunc("GET /items/{id}", GetItemHandler)
	mux.HandleFunc("POST /items", CreateItemHandler)
	mux.HandleFunc("PUT /items/{id}", UpdateItemHandler)
	mux.HandleFunc("DELETE /items/{id}", DeleteItemHandler)
}
