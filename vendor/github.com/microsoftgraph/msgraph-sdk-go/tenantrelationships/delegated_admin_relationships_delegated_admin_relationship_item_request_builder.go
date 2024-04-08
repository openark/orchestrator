package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder provides operations to manage the delegatedAdminRelationships property of the microsoft.graph.tenantRelationship entity.
type DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderGetQueryParameters the details of the delegated administrative privileges that a Microsoft partner has in a customer tenant.
type DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderGetQueryParameters
}
// DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessAssignments provides operations to manage the accessAssignments property of the microsoft.graph.delegatedAdminRelationship entity.
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) AccessAssignments()(*DelegatedAdminRelationshipsItemAccessAssignmentsRequestBuilder) {
    return NewDelegatedAdminRelationshipsItemAccessAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AccessAssignmentsById provides operations to manage the accessAssignments property of the microsoft.graph.delegatedAdminRelationship entity.
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) AccessAssignmentsById(id string)(*DelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedAdminAccessAssignment%2Did"] = id
    }
    return NewDelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderInternal instantiates a new DelegatedAdminRelationshipItemRequestBuilder and sets the default values.
func NewDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) {
    m := &DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/delegatedAdminRelationships/{delegatedAdminRelationship%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder instantiates a new DelegatedAdminRelationshipItemRequestBuilder and sets the default values.
func NewDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property delegatedAdminRelationships for tenantRelationships
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the details of the delegated administrative privileges that a Microsoft partner has in a customer tenant.
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDelegatedAdminRelationshipFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipable), nil
}
// Operations provides operations to manage the operations property of the microsoft.graph.delegatedAdminRelationship entity.
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) Operations()(*DelegatedAdminRelationshipsItemOperationsRequestBuilder) {
    return NewDelegatedAdminRelationshipsItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.delegatedAdminRelationship entity.
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) OperationsById(id string)(*DelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedAdminRelationshipOperation%2Did"] = id
    }
    return NewDelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property delegatedAdminRelationships in tenantRelationships
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipable, requestConfiguration *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDelegatedAdminRelationshipFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipable), nil
}
// Requests provides operations to manage the requests property of the microsoft.graph.delegatedAdminRelationship entity.
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) Requests()(*DelegatedAdminRelationshipsItemRequestsRequestBuilder) {
    return NewDelegatedAdminRelationshipsItemRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RequestsById provides operations to manage the requests property of the microsoft.graph.delegatedAdminRelationship entity.
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) RequestsById(id string)(*DelegatedAdminRelationshipsItemRequestsDelegatedAdminRelationshipRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedAdminRelationshipRequest%2Did"] = id
    }
    return NewDelegatedAdminRelationshipsItemRequestsDelegatedAdminRelationshipRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property delegatedAdminRelationships for tenantRelationships
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the details of the delegated administrative privileges that a Microsoft partner has in a customer tenant.
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property delegatedAdminRelationships in tenantRelationships
func (m *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipable, requestConfiguration *DelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
