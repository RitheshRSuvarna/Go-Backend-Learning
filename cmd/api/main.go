package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	platformpostgres "postgres"
	tripservice "trip/application/services"
	triprepository "trip/infrastructure/driven/postgres/repository"
	triprest "trip/infrastructure/driving/rest"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	ctx := context.Background()

	dbURL := os.Getenv("DATABASE_URL")

	db, err := platformpostgres.NewDB(ctx, platformpostgres.PostgresConfig{DATABASE_URL: dbURL})
	if err != nil {
		log.Fatalf("database: %v", err)
	}
	defer db.Close()

	port := os.Getenv("PORT")

	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", initTripHandler(db)))
	mux.HandleFunc("/health", healthHandler(db))

	log.Printf("Trips API: http://localhost:%s/api/trips", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))

}

func loadEnv() {
	for _, path := range []string{".env", "../.env", "../../.env"} {
		if err := godotenv.Load(path); err == nil {
			log.Printf("Loaded env from %s", path)
			return
		}
	}
}

func initTripHandler(db *pgxpool.Pool) http.Handler {
	repo := triprepository.NewTripRepository(db)
	createSvc := tripservice.NewCreateTripService(repo)
	listSvc := tripservice.NewTripListService(repo)
	return triprest.NewHandler(createSvc, listSvc)
}

func healthHandler(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		status := "healthy"
		checks := map[string]string{}

		if err := db.Ping(ctx); err != nil {
			status = "unhealthy"
			checks["database"] = err.Error()
		} else {
			checks["database"] = "healthy"
		}

		w.Header().Set("Content-Type", "application/json")
		if status != "healthy" {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		_ = json.NewEncoder(w).Encode(map[string]any{
			"status":    status,
			"service":   "live-coding-trip",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
			"checks":    checks,
		})
	}
}
