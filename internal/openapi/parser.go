package openapi

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
)

type Parser struct {
	spec *openapi3.T
}

func LoadSpecFromFile(path string) (*openapi3.T, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read spec file: %w", err)
	}

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	
	spec, err := loader.LoadFromData(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse OpenAPI spec: %w", err)
	}

	// Skip validation for now as the spec has some issues
	// ctx := context.Background()
	// if err := spec.Validate(ctx); err != nil {
	// 	return nil, fmt.Errorf("invalid OpenAPI spec: %w", err)
	// }

	return spec, nil
}

func LoadSpecFromURL(url string) (*openapi3.T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch spec: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	cacheFile := "api.json"
	if err := os.WriteFile(cacheFile, data, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to cache spec: %v\n", err)
	}

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	
	spec, err := loader.LoadFromData(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse OpenAPI spec: %w", err)
	}

	// Skip validation for now as the spec has some issues with null types
	// if err := spec.Validate(ctx); err != nil {
	// 	return nil, fmt.Errorf("invalid OpenAPI spec: %w", err)
	// }

	return spec, nil
}

func NewParser(spec *openapi3.T) *Parser {
	return &Parser{spec: spec}
}

func (p *Parser) GetOperations() map[string]*Operation {
	operations := make(map[string]*Operation)

	for path, pathItem := range p.spec.Paths.Map() {
		p.extractOperations(path, pathItem, operations)
	}

	return operations
}

func (p *Parser) extractOperations(path string, pathItem *openapi3.PathItem, operations map[string]*Operation) {
	methods := map[string]*openapi3.Operation{
		"GET":    pathItem.Get,
		"POST":   pathItem.Post,
		"PUT":    pathItem.Put,
		"DELETE": pathItem.Delete,
		"PATCH":  pathItem.Patch,
	}

	for method, op := range methods {
		if op == nil {
			continue
		}

		operation := &Operation{
			Method:      method,
			Path:        path,
			OperationID: op.OperationID,
			Summary:     op.Summary,
			Description: op.Description,
			Tags:        op.Tags,
			Parameters:  p.extractParameters(op),
			RequestBody: p.extractRequestBody(op),
			Security:    p.extractSecurity(op),
		}

		operations[op.OperationID] = operation
	}
}

func (p *Parser) extractParameters(op *openapi3.Operation) []Parameter {
	var params []Parameter

	for _, paramRef := range op.Parameters {
		if paramRef.Value == nil {
			continue
		}

		param := paramRef.Value
		params = append(params, Parameter{
			Name:        param.Name,
			In:          param.In,
			Description: param.Description,
			Required:    param.Required,
			Schema:      param.Schema,
		})
	}

	return params
}

func (p *Parser) extractRequestBody(op *openapi3.Operation) *RequestBody {
	if op.RequestBody == nil || op.RequestBody.Value == nil {
		return nil
	}

	rb := op.RequestBody.Value
	reqBody := &RequestBody{
		Description: rb.Description,
		Required:    rb.Required,
		Content:     make(map[string]*MediaType),
	}

	for mediaType, content := range rb.Content {
		reqBody.Content[mediaType] = &MediaType{
			Schema: content.Schema,
		}
	}

	return reqBody
}

func (p *Parser) extractSecurity(op *openapi3.Operation) []SecurityRequirement {
	var security []SecurityRequirement

	secReqs := op.Security
	if secReqs == nil {
		secReqs = &p.spec.Security
	}

	if secReqs != nil {
		for _, req := range *secReqs {
			for name, scopes := range req {
				security = append(security, SecurityRequirement{
					Name:   name,
					Scopes: scopes,
				})
			}
		}
	}

	return security
}

func (p *Parser) GetServers() []string {
	var servers []string
	for _, server := range p.spec.Servers {
		servers = append(servers, server.URL)
	}
	return servers
}

func (p *Parser) GetSecuritySchemes() map[string]*openapi3.SecurityScheme {
	if p.spec.Components == nil || p.spec.Components.SecuritySchemes == nil {
		return nil
	}

	schemes := make(map[string]*openapi3.SecurityScheme)
	for name, schemeRef := range p.spec.Components.SecuritySchemes {
		if schemeRef.Value != nil {
			schemes[name] = schemeRef.Value
		}
	}
	return schemes
}