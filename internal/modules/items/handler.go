package items

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-starter/pkg/response"
)

func ListItemsHandler(w http.ResponseWriter, r *http.Request) {
	items := ListItems()

	response.Success(
		w,
		http.StatusOK,
		"Lista de itens carregada com sucesso",
		items,
	)
}

func GetItemHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "ID inválido")
		return
	}

	item, err := GetItem(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, "Item não encontrado")
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"Item encontrado com sucesso",
		item,
	)
}

func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	var newItem Item

	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	createdItem := CreateItem(newItem)

	response.Success(
		w,
		http.StatusCreated,
		"Item criado com sucesso",
		createdItem,
	)
}

func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var updatedItem Item

	err = json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	item, err := UpdateItem(id, updatedItem)
	if err != nil {
		response.Error(w, http.StatusNotFound, "Item não encontrado")
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"Item atualizado com sucesso",
		item,
	)
}

func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "ID inválido")
		return
	}

	err = DeleteItem(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, "Item não encontrado")
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"Item removido com sucesso",
		nil,
	)
}

func getIDFromRequest(r *http.Request) (int, error) {
	idParam := r.PathValue("id")
	return strconv.Atoi(idParam)
}
