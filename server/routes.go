package server

import (
	"net/http"

	"github.com/adamelfsborg-code/food/culinary/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (a *Server) loadRoutes() {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/api/v1/categories", a.loadCategoryRoutes)
	router.Route("/api/v1/brands", a.loadBrandRoutes)
	router.Route("/api/v1/foodtypes", a.loadFoodTypeRoutes)
	a.router = router
}

func (a *Server) loadCategoryRoutes(router chi.Router) {
	categoryHandler := &handler.CategoryHandler{
		Data: a.data,
	}

	router.Group(func(r chi.Router) {
		r.Use(CustomAuthMiddleware())

		r.Get("/list", categoryHandler.ListCategories)
		r.Post("/", categoryHandler.CreateCategory)

		r.Get("/{id}", categoryHandler.GetCategoryById)
		r.Put("/{id}", categoryHandler.EditCategory)
		r.Delete("/{id}", categoryHandler.DeleteCategory)
	})
}

func (a *Server) loadBrandRoutes(router chi.Router) {
	brandHandler := &handler.BrandHandler{
		Data: a.data,
	}

	router.Group(func(r chi.Router) {
		r.Use(CustomAuthMiddleware())

		r.Get("/list", brandHandler.ListBrands)
		r.Post("/", brandHandler.CreateBrand)

		r.Get("/{id}", brandHandler.GetBrandById)
		r.Put("/{id}", brandHandler.EditBrand)
		r.Delete("/{id}", brandHandler.DeleteBrand)
	})
}

func (a *Server) loadFoodTypeRoutes(router chi.Router) {
	brandHandler := &handler.FoodTypeHandler{
		Data: a.data,
	}

	router.Group(func(r chi.Router) {
		r.Use(CustomAuthMiddleware())

		r.Get("/list", brandHandler.ListFoodTypes)
		r.Post("/", brandHandler.CreateFoodType)

		r.Get("/{id}", brandHandler.GetFoodTypeById)
		r.Put("/{id}", brandHandler.EditFoodType)
		r.Delete("/{id}", brandHandler.DeleteFoodType)
	})
}
