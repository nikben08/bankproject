package routers

import (
	"bankproject/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"bankproject/handlers"
)

func SetupRoutes(DB *gorm.DB, router *mux.Router) {
	h := handlers.New(DB)

	bank := router.PathPrefix("/api").Subrouter()
	bank.HandleFunc("/createBank", h.CreateBank).Methods(http.MethodPost)
	bank.HandleFunc("/banks", h.DeleteBankById).Methods(http.MethodDelete)
	bank.Use(middleware.Authenticated)

	public := router.PathPrefix("/api").Subrouter()
	public.HandleFunc("/banks", h.GetAllBanks).Methods(http.MethodGet)
	public.HandleFunc("/banks/{id}", h.GetBankById).Methods(http.MethodGet)
	public.HandleFunc("/interests", h.RelevantInterests).Methods(http.MethodGet)

	users := router.PathPrefix("/api").Subrouter()
	users.HandleFunc("/createUser", h.CreateUser).Methods(http.MethodPost)
	users.Use(middleware.Authenticated)

	router.HandleFunc("/api/login", h.Login).Methods(http.MethodGet)

	interests := router.PathPrefix("/api").Subrouter()
	interests.HandleFunc("/interests", h.DeleteInterest).Methods(http.MethodDelete)
	interests.HandleFunc("/addInterest", h.AddInterest).Methods(http.MethodPost)
	interests.Use(middleware.Authenticated)
}
