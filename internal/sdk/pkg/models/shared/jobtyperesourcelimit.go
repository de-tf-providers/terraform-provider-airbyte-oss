// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// JobTypeResourceLimit - sets resource requirements for a specific job type for an actor definition. these values override the default, if both are set.
type JobTypeResourceLimit struct {
	// enum that describes the different types of jobs that the platform runs.
	JobType JobType `json:"jobType"`
	// optional resource requirements to run workers (blank for unbounded allocations)
	ResourceRequirements ResourceRequirements `json:"resourceRequirements"`
}