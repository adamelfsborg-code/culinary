package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adamelfsborg-code/food/culinary/data"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type BrandHandler struct {
	Data data.DataConn
}

func (u *BrandHandler) GetBrandById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	brand, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Failed to parse id: ", err)
		http.Error(w, "Failed to parse id", http.StatusBadRequest)
		return
	}

	brands, err := u.Data.GetBrandById(brand)
	if err != nil {
		fmt.Println("Failed to get brand: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(brands)
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (u *BrandHandler) ListBrands(w http.ResponseWriter, r *http.Request) {
	brands, err := u.Data.ListBrands()
	if err != nil {
		fmt.Println("Failed to get brand: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(brands)
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (u *BrandHandler) CreateBrand(w http.ResponseWriter, r *http.Request) {
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

	brand, err := data.NewBrandDto(userId, body.Name)
	if err != nil {
		fmt.Println("Failed to extract brand details: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = u.Data.CreateBrand(*brand)
	if err != nil {
		fmt.Println("Failed to create brand: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(map[string]string{"message": "Brand Created"})
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonBytes)
}

func (u *BrandHandler) EditBrand(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	brand, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Failed to parse id: ", err)
		http.Error(w, "Failed to parse id", http.StatusBadRequest)
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

	err = u.Data.EditBrand(brand, body.Name)
	if err != nil {
		fmt.Println("Failed to delete brand: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(map[string]string{"message": "Brand Edited"})
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (u *BrandHandler) DeleteBrand(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	brand, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Failed to parse id: ", err)
		http.Error(w, "Failed to parse id", http.StatusBadRequest)
		return
	}

	err = u.Data.DeleteBrand(brand)
	if err != nil {
		fmt.Println("Failed to delete brand: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(map[string]string{"message": "Brand Deleted"})
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
