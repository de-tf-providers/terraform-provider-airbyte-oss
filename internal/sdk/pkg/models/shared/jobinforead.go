// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// JobInfoRead - Successful operation
type JobInfoRead struct {
	Attempts []AttemptInfoRead `json:"attempts"`
	Job      JobRead           `json:"job"`
}
