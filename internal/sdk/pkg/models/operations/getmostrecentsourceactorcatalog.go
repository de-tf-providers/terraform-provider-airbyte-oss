// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetMostRecentSourceActorCatalogResponse struct {
	// Successful operation
	ActorCatalogWithUpdatedAt *shared.ActorCatalogWithUpdatedAt
	ContentType               string
	// Input failed validation
	InvalidInputExceptionInfo *shared.InvalidInputExceptionInfo
	// Object with given id was not found.
	NotFoundKnownExceptionInfo *shared.NotFoundKnownExceptionInfo
	StatusCode                 int
	RawResponse                *http.Response
}
