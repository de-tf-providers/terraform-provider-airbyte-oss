// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SaveAttemptSyncConfigRequestBody struct {
	AttemptNumber int               `json:"attemptNumber"`
	JobID         int64             `json:"jobId"`
	SyncConfig    AttemptSyncConfig `json:"syncConfig"`
}
