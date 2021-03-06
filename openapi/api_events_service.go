/*
 * eventsite
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/umemak/eventsite_go/db"
	"github.com/umemak/eventsite_go/sqlc"
)

// EventsApiService is a service that implements the logic for the EventsApiServicer
// This service should implement the business logic for every endpoint for the EventsApi API.
// Include any external packages or services that will be required by this service.
type EventsApiService struct {
}

// NewEventsApiService creates a default api service
func NewEventsApiService() EventsApiServicer {
	return &EventsApiService{}
}

// EventsGet - Get all events.
func (s *EventsApiService) EventsGet(ctx context.Context) (ImplResponse, error) {
	db, err := db.Open()
	if err != nil {
		return Response(500, nil), fmt.Errorf("db.Open: %w", err)
	}
	defer db.Close()
	queries := sqlc.New(db)
	events, err := queries.ListEvents(ctx)
	if err != nil {
		return Response(500, nil), fmt.Errorf("queries.ListEvents: %w", err)
	}
	return Response(http.StatusOK, events), nil
}
