// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceSearch struct {
	// The values required to configure the source. The schema for this must match the schema return by source_definition_specifications/get for the source.
	ConnectionConfiguration interface{} `json:"connectionConfiguration,omitempty"`
	Name                    *string     `json:"name,omitempty"`
	SourceDefinitionID      *string     `json:"sourceDefinitionId,omitempty"`
	SourceID                *string     `json:"sourceId,omitempty"`
	SourceName              *string     `json:"sourceName,omitempty"`
	WorkspaceID             *string     `json:"workspaceId,omitempty"`
}
