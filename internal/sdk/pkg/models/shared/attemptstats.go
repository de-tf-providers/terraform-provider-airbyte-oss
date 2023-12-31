// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AttemptStats struct {
	BytesCommitted       *int64 `json:"bytesCommitted,omitempty"`
	BytesEmitted         *int64 `json:"bytesEmitted,omitempty"`
	EstimatedBytes       *int64 `json:"estimatedBytes,omitempty"`
	EstimatedRecords     *int64 `json:"estimatedRecords,omitempty"`
	RecordsCommitted     *int64 `json:"recordsCommitted,omitempty"`
	RecordsEmitted       *int64 `json:"recordsEmitted,omitempty"`
	StateMessagesEmitted *int64 `json:"stateMessagesEmitted,omitempty"`
}
