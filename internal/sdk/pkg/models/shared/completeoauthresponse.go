// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CompleteOAuthResponse - Successful operation
type CompleteOAuthResponse struct {
	AuthPayload      map[string]interface{} `json:"auth_payload"`
	RequestError     *string                `json:"request_error,omitempty"`
	RequestSucceeded bool                   `json:"request_succeeded"`
}
