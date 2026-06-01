package items

import (
	"database/sql"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, db *sql.DB) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	mux.HandleFunc("GET /items", handler.ListItemsHandler)
	mux.HandleFunc("GET /items/{id}", handler.GetItemHandler)
	mux.HandleFunc("POST /items", handler.CreateItemHandler)
	mux.HandleFunc("PUT /items/{id}", handler.UpdateItemHandler)
	mux.HandleFunc("DELETE /items/{id}", handler.DeleteItemHandler)
}
