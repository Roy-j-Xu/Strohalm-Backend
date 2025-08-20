package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	InitDatabase()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// CRUD routes
	// r.Get("/items", listItems)
	// r.Post("/items", createItem)
	// r.Put("/items/{id}", updateItem)

	log.Println("Server started")
	http.ListenAndServe(":8080", r)
}

func registerRouter(r *chi.Mux, route Route) {
	switch route.Method {
	case http.MethodGet:
		r.Get(route.Path, route.Handler)
	case http.MethodPost:
		r.Post(route.Path, route.Handler)
	case http.MethodPut:
		r.Put(route.Path, route.Handler)
	case http.MethodDelete:
		r.Delete(route.Path, route.Handler)
	default:
		panic(fmt.Sprintf("Unsupported method: %s", route.Method))
	}
}

// func listItems(w http.ResponseWriter, r *http.Request) {
// 	json.NewEncoder(w).Encode(items)
// }

// func createItem(w http.ResponseWriter, r *http.Request) {
// 	var data map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
// 		http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 		return
// 	}
// 	mu.Lock()
// 	item := Item{ID: nextID, Name: data["name"]}
// 	nextID++
// 	items = append(items, item)
// 	mu.Unlock()
// 	json.NewEncoder(w).Encode(item)
// }

// func updateItem(w http.ResponseWriter, r *http.Request) {
// 	idStr := chi.URLParam(r, "id")
// 	id, _ := strconv.Atoi(idStr)

// 	mu.Lock()
// 	defer mu.Unlock()
// 	for i, it := range items {
// 		if it.ID == id {
// 			var data map[string]string
// 			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
// 				http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 				return
// 			}
// 			items[i].Name = data["name"]
// 			json.NewEncoder(w).Encode(items[i])
// 			return
// 		}
// 	}
// 	http.Error(w, "Item not found", http.StatusNotFound)
// }
