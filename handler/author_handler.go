package handler

import (
	"encoding/json"
	dto "library-app/dtos"
	"library-app/model"
	"library-app/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AuthorHandler struct {
	service service.AuthorService
}

func NewAuthorHandler(service service.AuthorService) *AuthorHandler {
	return &AuthorHandler{service}
}

// CreateAuthor godoc
// @Summary Create a new author
// @Description Create a new author with the input payload
// @Tags authors
// @Accept  json
// @Produce  json
// @Param author body dto.AuthorCreateDTO true "Author"
// @Success 201 {object} dto.AuthorCreateDTO
// @Failure 400 {object} string
// @Router /authors [post]
func (h *AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var authorDTO dto.AuthorCreateDTO
	if err := json.NewDecoder(r.Body).Decode(&authorDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateAuthor(&authorDTO); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetAuthorByID godoc
// @Summary Get a author by ID
// @Description Get details of a author by ID
// @Tags authors
// @Produce  json
// @Param id path int true "Author ID"
// @Success 200 {object} dto.AuthorViewDTO
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /authors/{id} [get]
func (h *AuthorHandler) GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := h.service.GetAuthorByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}

// GetAllAuthors godoc
// @Summary Get all authors
// @Description Get a list of all authors
// @Tags authors
// @Produce  json
// @Success 200 {array} dto.AuthorViewDTO
// @Failure 500 {object} string
// @Router /authors [get]
func (h *AuthorHandler) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetAllAuthors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}

// UpdateAuthor godoc
// @Summary Update a author by ID
// @Description Update details of a author by ID
// @Tags authors
// @Accept  json
// @Produce  json
// @Param id path int true "Author ID"
// @Param author body model.Author true "Author"
// @Success 200 {object} model.Author
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /authors/{id} [put]
func (h *AuthorHandler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	var author model.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	author.ID = string(id)
	if err := h.service.UpdateAuthor(&author); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteAuthor godoc
// @Summary Delete a author by ID
// @Description Delete a author by ID
// @Tags authors
// @Param id path int true "Author ID"
// @Success 204
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /authors/{id} [delete]
func (h *AuthorHandler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteAuthor(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
