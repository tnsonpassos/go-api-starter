package items

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-api-starter/pkg/response"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) ListItemsHandler(w http.ResponseWriter, r *http.Request) {
	items, err := h.Service.ListItems()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Erro ao listar itens")
		return
	}

	response.Success(w, http.StatusOK, "Lista de itens carregada com sucesso", items)
}

func (h *Handler) GetItemHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "ID inválido")
		return
	}

	item, err := h.Service.GetItem(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, "Item não encontrado")
		return
	}

	response.Success(w, http.StatusOK, "Item encontrado com sucesso", item)
}

func (h *Handler) CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	var newItem Item

	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	createdItem, err := h.Service.CreateItem(newItem)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Erro ao criar item")
		return
	}

	response.Success(w, http.StatusCreated, "Item criado com sucesso", createdItem)
}

func (h *Handler) UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
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

	item, err := h.Service.UpdateItem(id, updatedItem)
	if err != nil {
		response.Error(w, http.StatusNotFound, "Item não encontrado")
		return
	}

	response.Success(w, http.StatusOK, "Item atualizado com sucesso", item)
}

func (h *Handler) DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "ID inválido")
		return
	}

	err = h.Service.DeleteItem(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, "Item não encontrado")
		return
	}

	response.Success(w, http.StatusOK, "Item removido com sucesso", nil)
}

func getIDFromRequest(r *http.Request) (int, error) {
	idParam := r.PathValue("id")
	return strconv.Atoi(idParam)
}
