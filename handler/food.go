package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adamelfsborg-code/food/culinary/data"
	"github.com/adamelfsborg-code/food/culinary/lib"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type FoodHandler struct {
	Data data.DataConn
}

func (u *FoodHandler) GetFoodById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	foodType, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Failed to parse id: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	foodTypes, err := u.Data.GetFoodById(foodType)
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

func (u *FoodHandler) ListFoods(w http.ResponseWriter, r *http.Request) {
	pageIndex := r.URL.Query().Get("pageIndex")
	pageSize := r.URL.Query().Get("pageSize")

	pagination, err := lib.NewPagination(pageIndex, pageSize)
	if err != nil {
		fmt.Println("Failed to parse pagination: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	foods, err := u.Data.ListFoods(pagination.PageIndex, pagination.PageSize)
	if err != nil {
		fmt.Println("Failed to get food: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	count, err := u.Data.CountFoods()
	if err != nil {
		fmt.Println("Failed to count food: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := lib.NewPaginatedResponse(foods, count, *pagination)

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (u *FoodHandler) CreateFood(w http.ResponseWriter, r *http.Request) {
	headerId := r.Header.Get("X-USER-ID")

	user, err := uuid.Parse(headerId)
	if err != nil {
		fmt.Println("Failed to parse id: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var body struct {
		Name        string  `json:"name"`
		FoodType    string  `json:"foodtype"`
		Brand       string  `json:"brand"`
		KCAL        float32 `json:"kcal"`
		Protein     float32 `json:"protein"`
		Carbs       float32 `json:"carbs"`
		Fat         float32 `json:"fat"`
		Saturated   float32 `json:"saturated"`
		Unsaturated float32 `json:"unsaturated"`
		Fiber       float32 `json:"fiber"`
		Sugars      float32 `json:"sugars"`
	}

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	foodtype, err := uuid.Parse(body.FoodType)
	if err != nil {
		fmt.Println("Failed to parse foodtype: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	brand, err := uuid.Parse(body.Brand)
	if err != nil {
		fmt.Println("Failed to parse brand: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	food, err := data.NewFood(body.Name, body.KCAL, body.Protein, body.Carbs, body.Fat, body.Saturated, body.Unsaturated, body.Fiber, body.Sugars, user, foodtype, brand)
	if err != nil {
		fmt.Println("Failed to extract food details: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = u.Data.CreateFood(*food)
	if err != nil {
		fmt.Println("Failed to create food: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(map[string]string{"message": "Food Created"})
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonBytes)
}

func (u *FoodHandler) EditFood(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	food, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Failed to parse id: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var body struct {
		Name        string  `json:"name"`
		FoodType    string  `json:"foodtype"`
		Brand       string  `json:"brand"`
		KCAL        float32 `json:"kcal"`
		Protein     float32 `json:"protein"`
		Carbs       float32 `json:"carbs"`
		Fat         float32 `json:"fat"`
		Saturated   float32 `json:"saturated"`
		Unsaturated float32 `json:"unsaturated"`
		Fiber       float32 `json:"fiber"`
		Sugars      float32 `json:"sugars"`
	}

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	foodtype, err := uuid.Parse(body.FoodType)
	if err != nil {
		fmt.Println("Failed to parse id: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	brand, err := uuid.Parse(body.Brand)
	if err != nil {
		fmt.Println("Failed to parse id: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = u.Data.EditFood(body.Name, body.KCAL, body.Protein, body.Carbs, body.Fat, body.Saturated, body.Unsaturated, body.Fiber, body.Sugars, brand, foodtype, food)
	if err != nil {
		fmt.Println("Failed to edit food: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(map[string]string{"message": "Food Edited"})
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (u *FoodHandler) DeleteFood(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	foodType, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Failed to parse id: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = u.Data.DeleteFood(foodType)
	if err != nil {
		fmt.Println("Failed to delete foodType: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(map[string]string{"message": "Food Deleted"})
	if err != nil {
		fmt.Println("Failed to decode json: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
