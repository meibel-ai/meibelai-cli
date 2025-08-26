package openapi

import "github.com/getkin/kin-openapi/openapi3"

type Operation struct {
	Method      string
	Path        string
	OperationID string
	Summary     string
	Description string
	Tags        []string
	Parameters  []Parameter
	RequestBody *RequestBody
	Security    []SecurityRequirement
}

type Parameter struct {
	Name        string
	In          string
	Description string
	Required    bool
	Schema      *openapi3.SchemaRef
}

type RequestBody struct {
	Description string
	Required    bool
	Content     map[string]*MediaType
}

type MediaType struct {
	Schema *openapi3.SchemaRef
}

type SecurityRequirement struct {
	Name   string
	Scopes []string
}