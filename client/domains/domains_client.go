// Code generated by go-swagger; DO NOT EDIT.

package domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new domains API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for domains API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AssignTagsToExistingDomain(params *AssignTagsToExistingDomainParams, opts ...ClientOption) (*AssignTagsToExistingDomainOK, error)

	AssignableTagsToDomain(params *AssignableTagsToDomainParams, opts ...ClientOption) (*AssignableTagsToDomainOK, error)

	CreateDomain(params *CreateDomainParams, opts ...ClientOption) (*CreateDomainOK, *CreateDomainAccepted, error)

	DeleteDomain(params *DeleteDomainParams, opts ...ClientOption) (*DeleteDomainOK, *DeleteDomainAccepted, error)

	GetClusterCriteria(params *GetClusterCriteriaParams, opts ...ClientOption) (*GetClusterCriteriaOK, error)

	GetClusterCriterion(params *GetClusterCriterionParams, opts ...ClientOption) (*GetClusterCriterionOK, error)

	GetClusterQueryResponse(params *GetClusterQueryResponseParams, opts ...ClientOption) (*GetClusterQueryResponseOK, error)

	GetClustersQueryResponse(params *GetClustersQueryResponseParams, opts ...ClientOption) (*GetClustersQueryResponseOK, error)

	GetDatastoreCriterion1(params *GetDatastoreCriterion1Params, opts ...ClientOption) (*GetDatastoreCriterion1OK, error)

	GetDatastoreQueryResponse1(params *GetDatastoreQueryResponse1Params, opts ...ClientOption) (*GetDatastoreQueryResponse1OK, error)

	GetDatastoresCriteria1(params *GetDatastoresCriteria1Params, opts ...ClientOption) (*GetDatastoresCriteria1OK, error)

	GetDomain(params *GetDomainParams, opts ...ClientOption) (*GetDomainOK, error)

	GetDomainEndpoints(params *GetDomainEndpointsParams, opts ...ClientOption) (*GetDomainEndpointsOK, error)

	GetDomainTagManagerURL(params *GetDomainTagManagerURLParams, opts ...ClientOption) (*GetDomainTagManagerURLOK, error)

	GetDomains(params *GetDomainsParams, opts ...ClientOption) (*GetDomainsOK, error)

	GetTagsAssignedToDomain(params *GetTagsAssignedToDomainParams, opts ...ClientOption) (*GetTagsAssignedToDomainOK, error)

	GetTagsAssignedToDomains(params *GetTagsAssignedToDomainsParams, opts ...ClientOption) (*GetTagsAssignedToDomainsOK, error)

	PostClusterQuery(params *PostClusterQueryParams, opts ...ClientOption) (*PostClusterQueryOK, error)

	PostClustersQuery(params *PostClustersQueryParams, opts ...ClientOption) (*PostClustersQueryOK, error)

	PostDatastoreQuery1(params *PostDatastoreQuery1Params, opts ...ClientOption) (*PostDatastoreQuery1OK, error)

	RemoveTagsFromDomain(params *RemoveTagsFromDomainParams, opts ...ClientOption) (*RemoveTagsFromDomainOK, error)

	UpdateDomain(params *UpdateDomainParams, opts ...ClientOption) (*UpdateDomainOK, *UpdateDomainAccepted, error)

	ValidateDomainsOperations(params *ValidateDomainsOperationsParams, opts ...ClientOption) (*ValidateDomainsOperationsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AssignTagsToExistingDomain assigns tags to domain
*/
func (a *Client) AssignTagsToExistingDomain(params *AssignTagsToExistingDomainParams, opts ...ClientOption) (*AssignTagsToExistingDomainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssignTagsToExistingDomainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "assignTagsToExistingDomain",
		Method:             "PUT",
		PathPattern:        "/v1/domains/{id}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AssignTagsToExistingDomainReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AssignTagsToExistingDomainOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for assignTagsToExistingDomain: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AssignableTagsToDomain gets assignable tags to domain
*/
func (a *Client) AssignableTagsToDomain(params *AssignableTagsToDomainParams, opts ...ClientOption) (*AssignableTagsToDomainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssignableTagsToDomainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "assignableTagsToDomain",
		Method:             "GET",
		PathPattern:        "/v1/domains/{id}/tags/assignable-tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AssignableTagsToDomainReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AssignableTagsToDomainOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for assignableTagsToDomain: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateDomain creates a domain
*/
func (a *Client) CreateDomain(params *CreateDomainParams, opts ...ClientOption) (*CreateDomainOK, *CreateDomainAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDomainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createDomain",
		Method:             "POST",
		PathPattern:        "/v1/domains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDomainReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateDomainOK:
		return value, nil, nil
	case *CreateDomainAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for domains: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteDomain deletes a domain if it has been previously initialized for deletion
*/
func (a *Client) DeleteDomain(params *DeleteDomainParams, opts ...ClientOption) (*DeleteDomainOK, *DeleteDomainAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDomainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteDomain",
		Method:             "DELETE",
		PathPattern:        "/v1/domains/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDomainReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteDomainOK:
		return value, nil, nil
	case *DeleteDomainAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for domains: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetClusterCriteria gets all cluster criteria
*/
func (a *Client) GetClusterCriteria(params *GetClusterCriteriaParams, opts ...ClientOption) (*GetClusterCriteriaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterCriteriaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getClusterCriteria",
		Method:             "GET",
		PathPattern:        "/v1/domains/{domainId}/clusters/criteria",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterCriteriaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClusterCriteriaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getClusterCriteria: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetClusterCriterion gets a criterion to query for cluster
*/
func (a *Client) GetClusterCriterion(params *GetClusterCriterionParams, opts ...ClientOption) (*GetClusterCriterionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterCriterionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getClusterCriterion",
		Method:             "GET",
		PathPattern:        "/v1/domains/{domainId}/clusters/criteria/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterCriterionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClusterCriterionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getClusterCriterion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetClusterQueryResponse gets cluster query response
*/
func (a *Client) GetClusterQueryResponse(params *GetClusterQueryResponseParams, opts ...ClientOption) (*GetClusterQueryResponseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterQueryResponseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getClusterQueryResponse",
		Method:             "GET",
		PathPattern:        "/v1/domains/{domainId}/clusters/{clusterName}/queries/{queryId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterQueryResponseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClusterQueryResponseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getClusterQueryResponse: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetClustersQueryResponse gets clusters query response
*/
func (a *Client) GetClustersQueryResponse(params *GetClustersQueryResponseParams, opts ...ClientOption) (*GetClustersQueryResponseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClustersQueryResponseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getClustersQueryResponse",
		Method:             "GET",
		PathPattern:        "/v1/domains/{domainId}/clusters/queries/{queryId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClustersQueryResponseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClustersQueryResponseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getClustersQueryResponse: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDatastoreCriterion1 gets a criterion to query for datastore
*/
func (a *Client) GetDatastoreCriterion1(params *GetDatastoreCriterion1Params, opts ...ClientOption) (*GetDatastoreCriterion1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDatastoreCriterion1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDatastoreCriterion_1",
		Method:             "GET",
		PathPattern:        "/v1/domains/{domainId}/datastores/criteria/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDatastoreCriterion1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDatastoreCriterion1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDatastoreCriterion_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDatastoreQueryResponse1 gets datastore query response
*/
func (a *Client) GetDatastoreQueryResponse1(params *GetDatastoreQueryResponse1Params, opts ...ClientOption) (*GetDatastoreQueryResponse1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDatastoreQueryResponse1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDatastoreQueryResponse_1",
		Method:             "GET",
		PathPattern:        "/v1/domains/{domainId}/datastores/queries/{queryId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDatastoreQueryResponse1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDatastoreQueryResponse1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDatastoreQueryResponse_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDatastoresCriteria1 gets all datastore criteria
*/
func (a *Client) GetDatastoresCriteria1(params *GetDatastoresCriteria1Params, opts ...ClientOption) (*GetDatastoresCriteria1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDatastoresCriteria1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDatastoresCriteria_1",
		Method:             "GET",
		PathPattern:        "/v1/domains/{domainId}/datastores/criteria",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDatastoresCriteria1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDatastoresCriteria1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDatastoresCriteria_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDomain gets a domain
*/
func (a *Client) GetDomain(params *GetDomainParams, opts ...ClientOption) (*GetDomainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDomainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDomain",
		Method:             "GET",
		PathPattern:        "/v1/domains/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDomainReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDomainOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDomain: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDomainEndpoints gets endpoints of a domain
*/
func (a *Client) GetDomainEndpoints(params *GetDomainEndpointsParams, opts ...ClientOption) (*GetDomainEndpointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDomainEndpointsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDomainEndpoints",
		Method:             "GET",
		PathPattern:        "/v1/domains/{id}/endpoints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDomainEndpointsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDomainEndpointsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDomainEndpoints: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDomainTagManagerURL gets domain tag manager Url
*/
func (a *Client) GetDomainTagManagerURL(params *GetDomainTagManagerURLParams, opts ...ClientOption) (*GetDomainTagManagerURLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDomainTagManagerURLParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDomainTagManagerUrl",
		Method:             "GET",
		PathPattern:        "/v1/domains/{id}/tags/tag-manager",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDomainTagManagerURLReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDomainTagManagerURLOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDomainTagManagerUrl: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDomains gets the domains
*/
func (a *Client) GetDomains(params *GetDomainsParams, opts ...ClientOption) (*GetDomainsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDomainsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDomains",
		Method:             "GET",
		PathPattern:        "/v1/domains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDomainsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDomainsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDomains: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTagsAssignedToDomain gets tags assigned to domain
*/
func (a *Client) GetTagsAssignedToDomain(params *GetTagsAssignedToDomainParams, opts ...ClientOption) (*GetTagsAssignedToDomainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTagsAssignedToDomainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTagsAssignedToDomain",
		Method:             "GET",
		PathPattern:        "/v1/domains/{id}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTagsAssignedToDomainReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTagsAssignedToDomainOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTagsAssignedToDomain: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTagsAssignedToDomains gets tags assigned to domains
*/
func (a *Client) GetTagsAssignedToDomains(params *GetTagsAssignedToDomainsParams, opts ...ClientOption) (*GetTagsAssignedToDomainsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTagsAssignedToDomainsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTagsAssignedToDomains",
		Method:             "GET",
		PathPattern:        "/v1/domains/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTagsAssignedToDomainsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTagsAssignedToDomainsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTagsAssignedToDomains: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostClusterQuery posts a cluster query
*/
func (a *Client) PostClusterQuery(params *PostClusterQueryParams, opts ...ClientOption) (*PostClusterQueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostClusterQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postClusterQuery",
		Method:             "POST",
		PathPattern:        "/v1/domains/{domainId}/clusters/{clusterName}/queries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostClusterQueryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostClusterQueryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postClusterQuery: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostClustersQuery posts clusters query
*/
func (a *Client) PostClustersQuery(params *PostClustersQueryParams, opts ...ClientOption) (*PostClustersQueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostClustersQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postClustersQuery",
		Method:             "POST",
		PathPattern:        "/v1/domains/{domainId}/clusters/queries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostClustersQueryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostClustersQueryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postClustersQuery: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostDatastoreQuery1 posts a datastore query
*/
func (a *Client) PostDatastoreQuery1(params *PostDatastoreQuery1Params, opts ...ClientOption) (*PostDatastoreQuery1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostDatastoreQuery1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "postDatastoreQuery_1",
		Method:             "POST",
		PathPattern:        "/v1/domains/{domainId}/datastores/queries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostDatastoreQuery1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostDatastoreQuery1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postDatastoreQuery_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RemoveTagsFromDomain removes tags from domain
*/
func (a *Client) RemoveTagsFromDomain(params *RemoveTagsFromDomainParams, opts ...ClientOption) (*RemoveTagsFromDomainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveTagsFromDomainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeTagsFromDomain",
		Method:             "DELETE",
		PathPattern:        "/v1/domains/{id}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RemoveTagsFromDomainReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveTagsFromDomainOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeTagsFromDomain: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateDomain updates a domain
*/
func (a *Client) UpdateDomain(params *UpdateDomainParams, opts ...ClientOption) (*UpdateDomainOK, *UpdateDomainAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDomainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateDomain",
		Method:             "PATCH",
		PathPattern:        "/v1/domains/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDomainReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateDomainOK:
		return value, nil, nil
	case *UpdateDomainAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for domains: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ValidateDomainsOperations validates the input spec for domains operations
*/
func (a *Client) ValidateDomainsOperations(params *ValidateDomainsOperationsParams, opts ...ClientOption) (*ValidateDomainsOperationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateDomainsOperationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "validateDomainsOperations",
		Method:             "POST",
		PathPattern:        "/v1/domains/validations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateDomainsOperationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ValidateDomainsOperationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for validateDomainsOperations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}