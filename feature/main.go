package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/profile"
	"net/http"
)

func main() {
	defer profile.Start().Stop()

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Get("/organisations/{organisationId}/features", getOrganisationFeatures)
	router.Get("/organisations/{organisationId}/features/{key}/enabled", checkEnabledInOrganisation)
	router.Put("/organisations/{organisationId}/features/{key}/{enabled}", updateOrganisationFeatures)
	router.Get("/features/{key}/organisations", getOrganisationsWithFeature)

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		panic(err)
	}
}
