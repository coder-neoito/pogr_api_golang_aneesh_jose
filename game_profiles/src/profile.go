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
	r.HandleFunc("/api/health", handler.HealthCheck).Methods(http.MethodGet)
	r.HandleFunc("/api/list-all-games", handler.ListAllGames).Methods(http.MethodGet)

	createUserRoutes(handler, r)
	return r
}

func createUserRoutes(handler ProfileHandler, r *mux.Router) {
	userRoutes := r.PathPrefix("/api/user/{userID}").Subrouter()
	userRoutes.HandleFunc("/list-games", handler.ListGames).Methods(http.MethodGet)
	userRoutes.HandleFunc("/game/{gameCode}/get-characteristics", handler.GetCharacteristics).Methods(http.MethodGet)
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
