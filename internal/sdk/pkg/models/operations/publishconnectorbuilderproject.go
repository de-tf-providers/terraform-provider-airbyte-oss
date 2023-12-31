// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PublishConnectorBuilderProjectResponse struct {
	ContentType string
	// Successful operation
	SourceDefinitionIDBody *shared.SourceDefinitionIDBody
	StatusCode             int
	RawResponse            *http.Response
}
