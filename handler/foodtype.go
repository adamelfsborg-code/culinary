package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adamelfsborg-code/food/culinary/data"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type FoodTypeHandler struct {
	Data data.DataConn
}

func (u *FoodTypeHandler) GetFoodTypeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	foodType, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Failed to parse id: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	foodTypes, err := u.Data.GetFoodTypeById(foodType)
	if err != nil {
		fmt.Println("Failed to get foodType: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(foodTypes)
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (u *FoodTypeHandler) ListFoodTypes(w http.ResponseWriter, r *http.Request) {
	foodTypes, err := u.Data.ListFoodTypes()
	if err != nil {
		fmt.Println("Failed to get foodType: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(foodTypes)
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (u *FoodTypeHandler) CreateFoodType(w http.ResponseWriter, r *http.Request) {
	headerId := r.Header.Get("X-USER-ID")

	userId, err := uuid.Parse(headerId)
	if err != nil {
		fmt.Println("Failed to parse id: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var body struct {
		Name     string `json:"name"`
		Category string `json:"category"`
	}

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	category, err := uuid.Parse(body.Category)
	if err != nil {
		fmt.Println("Failed to parse category: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	foodType, err := data.NewFoodType(userId, body.Name, category)
	if err != nil {
		fmt.Println("Failed to extract foodType details: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = u.Data.CreateFoodType(*foodType)
	if err != nil {
		fmt.Println("Failed to create foodType: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(map[string]string{"message": "FoodType Created"})
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonBytes)
}

func (u *FoodTypeHandler) EditFoodType(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	foodType, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Failed to parse id: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var body struct {
		Name     string `json:"name"`
		Category string `json:"category"`
	}

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	category, err := uuid.Parse(body.Category)
	if err != nil {
		fmt.Println("Failed to parse id: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = u.Data.EditFoodType(foodType, body.Name, category)
	if err != nil {
		fmt.Println("Failed to delete foodType: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(map[string]string{"message": "FoodType Edited"})
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (u *FoodTypeHandler) DeleteFoodType(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	foodType, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Failed to parse id: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = u.Data.DeleteFoodType(foodType)
	if err != nil {
		fmt.Println("Failed to delete foodType: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(map[string]string{"message": "FoodType Deleted"})
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
