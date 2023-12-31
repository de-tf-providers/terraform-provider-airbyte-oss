// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DeclarativeSourceManifest struct {
	Description string `json:"description"`
	// Low code CDK manifest JSON object
	Manifest DeclarativeManifest `json:"manifest"`
	// The specification for what values are required to configure the sourceDefinition.
	Spec    SourceDefinitionSpecification `json:"spec"`
	Version int64                         `json:"version"`
}
