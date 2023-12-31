// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetConnectorBuilderProjectResponse struct {
	// Successful operation
	ConnectorBuilderProjectRead *shared.ConnectorBuilderProjectRead
	ContentType                 string
	StatusCode                  int
	RawResponse                 *http.Response
}
