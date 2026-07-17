package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	assistantsuggservice "assistant_suggestions/application/services"
	assistantsuggrepository "assistant_suggestions/infrastructure/driven/postgres/repository"
	assistantsuggrest "assistant_suggestions/infrastructure/driving/rest"
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
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
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

	if err := Run(ctx, db); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	mux.Handle("/api/trips", http.StripPrefix("/api", initTripHandler(db)))
	mux.Handle("/api/day-sessions", http.StripPrefix("/api", initDaySessionHandler(db)))
	mux.Handle("/api/day-sessions/{trip_id}", http.StripPrefix("/api", initDaySessionHandler(db)))
	mux.Handle("/api/day-sessions/{trip_id}/{date}", http.StripPrefix("/api", initDaySessionHandler(db)))
	mux.Handle("/api/day-sessions/{id}/plan-versions/", http.StripPrefix("/api", initPlanVersionHandler(db)))
	mux.Handle("/api/day-sessions/{id}/active-plan", http.StripPrefix("/api", initPlanStopHandler(db)))
	mux.Handle("/api/day-sessions/{id}/suggestions", http.StripPrefix("/api", initAssistantSuggestionHandler(db)))
	mux.Handle("/api/assistant-suggestions/{id}", http.StripPrefix("/api", initAssistantSuggestionHandler(db)))

	mux.HandleFunc("/health", healthHandler(db))

	log.Printf("Trips API: http://localhost:%s/api", port)
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

func Run(ctx context.Context, pool *pgxpool.Pool) error {

	db := stdlib.OpenDBFromPool(pool)
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("set goose dialect: %w", err)
	}
	wd, _ := os.Getwd()
	fmt.Println("Working directory:", wd)

	if err := goose.Up(db, "../../internal/platform/migration"); err != nil {
		return fmt.Errorf("run migrations: %w", err)
	}

	return nil
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
	listPlanStopSvc := planservice.NewListPlanStopService(planRepo)
	return planstoprest.NewHandlers(createPlanStopSvc, listPlanStopSvc)
}

func initAssistantSuggestionHandler(db *pgxpool.Pool) http.Handler {
	assistantsuggestionrepo := assistantsuggrepository.NewAssistantSuggestionRepository(db)
	createassistantsuggsvc := assistantsuggservice.NewAssistantSuggestionService(assistantsuggestionrepo)
	getassistantsuggsvc := assistantsuggservice.NewGetAssistantSuggestionService(assistantsuggestionrepo)
	editassistantsuggsvc := assistantsuggservice.NewEditAssistantSuggestionService(assistantsuggestionrepo)
	return assistantsuggrest.NewHandler(createassistantsuggsvc, getassistantsuggsvc, editassistantsuggsvc)

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
