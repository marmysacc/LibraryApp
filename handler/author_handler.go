package handler

import (
	"encoding/json"
	dto "library-app/dtos"
	"library-app/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type AuthorHandler struct {
	service   service.AuthorService
	validator *validator.Validate
}

func NewAuthorHandler(service service.AuthorService) *AuthorHandler {
	return &AuthorHandler{
		service:   service,
		validator: validator.New(),
	}
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

	// if err := h.validator.Struct(authorDTO); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	if err := h.validator.Struct(authorDTO); err != nil {
		// Przygotowanie odpowiedzi JSON z błędami walidacji
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Error())
		}
		responseData := map[string]interface{}{
			"errors": validationErrors,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responseData)
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
// @Param id path string true "Author ID"
// @Success 200 {object} dto.AuthorViewDTO
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /authors/{id} [get]
func (h *AuthorHandler) GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	author, err := h.service.GetAuthorByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ") // Ustawienie wcięć dla lepszego formatowania
	if err := encoder.Encode(author); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
	authors, err := h.service.GetAllAuthors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ") // Ustawienie wcięć dla lepszego formatowania
	if err := encoder.Encode(authors); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateAuthor godoc
// @Summary Update a author by ID
// @Description Update details of a author by ID
// @Tags authors
// @Accept  json
// @Produce  json
// @Param id path string true "Author ID"
// @Param author body dto.AuthorCreateDTO true "Author"
// @Success 200 {object} dto.AuthorCreateDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /authors/{id} [put]
func (h *AuthorHandler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var author dto.AuthorCreateDTO
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateAuthor(&author, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteAuthor godoc
// @Summary Delete a author by ID
// @Description Delete a author by ID
// @Tags authors
// @Param id path string true "Author ID"
// @Success 204
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /authors/{id} [delete]
func (h *AuthorHandler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if err := h.service.DeleteAuthor(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
