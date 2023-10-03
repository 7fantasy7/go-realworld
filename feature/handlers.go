package main

import (
	"context"
	"encoding/json"
	"feature/db"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

var (
	queries *db.Queries
) // todo global state

func init() { // todo env var
	url := "postgres://postgres:postgres@localhost:5432/feature"
	//url := "postgres://postgres:postgres@host.docker.internal:5432/feature"

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatal(err)
	}
	queries = db.New(conn)
}

func getOrganisationFeatures(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	paramStr := chi.URLParam(r, "organisationId")
	organisationId, err := strconv.Atoi(paramStr)
	if err != nil {
		http.Error(w, "Missing organisationId", http.StatusBadRequest)
		return
	}

	features, err := queries.GetFeaturesByOrganisation(ctx, int64(organisationId))
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(features)
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
} // todo simplify error handling?

func checkEnabledInOrganisation(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	organisationIdStr := chi.URLParam(r, "organisationId")
	organisationId, err := strconv.Atoi(organisationIdStr)
	if err != nil {
		http.Error(w, "Missing organisationId", http.StatusBadRequest)
		return
	}

	keyStr := chi.URLParam(r, "key")
	if keyStr == "" {
		http.Error(w, "Missing key", http.StatusBadRequest)
		return
	}

	params := &db.IsEnabledByKeyAndOrganisationParams{Key: keyStr, OrganisationID: int64(organisationId)}
	features, err := queries.IsEnabledByKeyAndOrganisation(ctx, *params)
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(features)
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
}

func getOrganisationsWithFeature(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	keyStr := chi.URLParam(r, "key")
	if keyStr == "" {
		http.Error(w, "Missing key", http.StatusBadRequest)
		return
	}

	organisationIds, err := queries.GetOrganisationsWithFeature(ctx, keyStr)
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(organisationIds)
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
}
func updateOrganisationFeatures(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	organisationIdStr := chi.URLParam(r, "organisationId")
	organisationId, err := strconv.Atoi(organisationIdStr)
	if err != nil {
		http.Error(w, "Missing organisationId", http.StatusBadRequest)
		return
	}

	keyStr := chi.URLParam(r, "key")
	if keyStr == "" {
		http.Error(w, "Missing key", http.StatusBadRequest)
		return
	}

	enabledStr := chi.URLParam(r, "enabled")
	if enabledStr == "" {
		http.Error(w, "Missing key", http.StatusBadRequest)
		return
	}
	enabled, err := strconv.ParseBool(enabledStr)
	if err != nil {
		http.Error(w, "Missing enabled", http.StatusBadRequest)
		return
	}

	params := &db.UpdateOrganisationFeatureParams{Key: keyStr, OrganisationID: int64(organisationId), Enabled: enabled}

	err = queries.UpdateOrganisationFeature(ctx, *params)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
