package game_profiles

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var port = GetEnvOrDefault("PORT", "8080")

func Run() error {
	gamerRepo := NewProfileRepository()
	gameService := NewProfileService(gamerRepo)
	gameHandler := NewProfilehandler(gameService)

	gameRoutes := createGameProfileRoutes(gameHandler)
	corsHandler := createCORSHandler(gameRoutes)
	fmt.Println("all setup and listening", "port", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), corsHandler)
}

func createGameProfileRoutes(handler ProfileHandler) *mux.Router {
	r := mux.NewRouter()
	userRoutes := r.PathPrefix("/api/user/{{userID}}").Subrouter()
	userRoutes.HandleFunc("/list-games", handler.ListGames).Methods(http.MethodGet)
	return r
}

func createCORSHandler(rootHandler http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Content-Disposition"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions},
		MaxAge: 300, // 5 minutes
	}).Handler(rootHandler)
}

func GetEnvOrDefault(name, defaultTo string) string {
	env := os.Getenv(name)
	if len(env) == 0 {
		env = defaultTo
	}
	return env
}
