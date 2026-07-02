package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	daysessionservice "day_session/application/services"
	daysessionrepository "day_session/infrastructure/driven/postgres/repository"
	daysessionrest "day_session/infrastructure/driving/rest"
	planservice "plan/application/services"
	planrepository "plan/infrastructure/driven/postgres/repository"
	planstoprest "plan/infrastructure/driving/rest"
	planversionrest "plan/infrastructure/driving/rest"
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
	mux.Handle("/api/day-sessions/", http.StripPrefix("/api", initDaySessionHandler(db)))
	mux.Handle("/api/plan-version/", http.StripPrefix("/api", initPlanVersionHandler(db)))
	mux.Handle("/api/plan-stop/", http.StripPrefix("/api", initPlanStopHandler(db)))
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

func initDaySessionHandler(db *pgxpool.Pool) http.Handler {
	daySessionRepo := daysessionrepository.NewDaySessionRepository(db)
	createDaySessionSvc := daysessionservice.NewDaySessionService(daySessionRepo)
	listDaySessionSvc := daysessionservice.NewListDaySessionService(daySessionRepo)
	getByTripIDAndDateDaySessionSvc := daysessionservice.NewDaySessionListService(daySessionRepo)
	return daysessionrest.NewHandler(createDaySessionSvc, listDaySessionSvc, getByTripIDAndDateDaySessionSvc)
}

func initPlanVersionHandler(db *pgxpool.Pool) http.Handler {
	planVersionRepo := planrepository.NewPlanVersionRepository(db)
	createPlanVersionSvc := planservice.NewCreatePlanVersionService(planVersionRepo)
	getPlanVersionSvc := planservice.NewGetByIDPlanVersionService(planVersionRepo)
	listPlanVersionSvc := planservice.NewListPlanVersionService(planVersionRepo)
	return planversionrest.NewHandler(createPlanVersionSvc, getPlanVersionSvc, listPlanVersionSvc)
}

func initPlanStopHandler(db *pgxpool.Pool) http.Handler {
	planRepo := planrepository.NewPlanStopRepository(db)
	createPlanStopSvc := planservice.NewCreatePlanStopService(planRepo)
	getPlanStopSvc := planservice.NewGetStopByIDService(planRepo)
	listPlanStopSvc := planservice.NewListPlanStopService(planRepo)
	return planstoprest.NewHandlers(createPlanStopSvc, getPlanStopSvc, listPlanStopSvc)
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
