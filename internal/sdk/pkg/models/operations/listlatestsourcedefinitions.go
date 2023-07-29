// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListLatestSourceDefinitionsResponse struct {
	ContentType string
	// Successful operation
	SourceDefinitionReadList *shared.SourceDefinitionReadList
	StatusCode               int
	RawResponse              *http.Response
}
