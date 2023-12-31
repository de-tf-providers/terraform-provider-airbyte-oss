// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type WriteDiscoverCatalogResultResponse struct {
	ContentType string
	// Successful Operation
	DiscoverCatalogResult *shared.DiscoverCatalogResult
	StatusCode            int
	RawResponse           *http.Response
}
