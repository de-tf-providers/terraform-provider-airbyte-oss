// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type SaveSyncConfigResponse struct {
	ContentType string
	// Successful Operation
	InternalOperationResult *shared.InternalOperationResult
	StatusCode              int
	RawResponse             *http.Response
}