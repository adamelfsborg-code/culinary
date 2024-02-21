package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adamelfsborg-code/food/culinary/data"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type CategoryHandler struct {
	Data data.DataConn
}

func (u *CategoryHandler) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	category, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Failed to parse id: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	catgories, err := u.Data.GetCategoryById(category)
	if err != nil {
		fmt.Println("Failed to get category: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(catgories)
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (u *CategoryHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	catgories, err := u.Data.ListCategories()
	if err != nil {
		fmt.Println("Failed to get category: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(catgories)
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (u *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	headerId := r.Header.Get("X-USER-ID")

	userId, err := uuid.Parse(headerId)
	if err != nil {
		fmt.Println("Failed to parse id: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var body struct {
		Name string `json:"name"`
	}

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	category, err := data.NewCategoryDto(userId, body.Name)
	if err != nil {
		fmt.Println("Failed to extract category details: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = u.Data.CreateCategory(*category)
	if err != nil {
		fmt.Println("Failed to create category: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(map[string]string{"message": "Category Created"})
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonBytes)
}

func (u *CategoryHandler) EditCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	category, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Failed to parse id: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var body struct {
		Name string `json:"name"`
	}

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = u.Data.EditCategory(category, body.Name)
	if err != nil {
		fmt.Println("Failed to delete category: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(map[string]string{"message": "Category Edited"})
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (u *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	category, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Failed to parse id: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = u.Data.DeleteCategory(category)
	if err != nil {
		fmt.Println("Failed to delete category: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(map[string]string{"message": "Category Deleted"})
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
