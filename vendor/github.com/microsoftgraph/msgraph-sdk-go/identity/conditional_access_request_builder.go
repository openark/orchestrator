package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ConditionalAccessRequestBuilder provides operations to manage the conditionalAccess property of the microsoft.graph.identityContainer entity.
type ConditionalAccessRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ConditionalAccessRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConditionalAccessRequestBuilderGetQueryParameters the entry point for the Conditional Access (CA) object model.
type ConditionalAccessRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConditionalAccessRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConditionalAccessRequestBuilderGetQueryParameters
}
// ConditionalAccessRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationContextClassReferences provides operations to manage the authenticationContextClassReferences property of the microsoft.graph.conditionalAccessRoot entity.
func (m *ConditionalAccessRequestBuilder) AuthenticationContextClassReferences()(*ConditionalAccessAuthenticationContextClassReferencesRequestBuilder) {
    return NewConditionalAccessAuthenticationContextClassReferencesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AuthenticationContextClassReferencesById provides operations to manage the authenticationContextClassReferences property of the microsoft.graph.conditionalAccessRoot entity.
func (m *ConditionalAccessRequestBuilder) AuthenticationContextClassReferencesById(id string)(*ConditionalAccessAuthenticationContextClassReferencesAuthenticationContextClassReferenceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationContextClassReference%2Did"] = id
    }
    return NewConditionalAccessAuthenticationContextClassReferencesAuthenticationContextClassReferenceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewConditionalAccessRequestBuilderInternal instantiates a new ConditionalAccessRequestBuilder and sets the default values.
func NewConditionalAccessRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessRequestBuilder) {
    m := &ConditionalAccessRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/conditionalAccess{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewConditionalAccessRequestBuilder instantiates a new ConditionalAccessRequestBuilder and sets the default values.
func NewConditionalAccessRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalAccessRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property conditionalAccess for identity
func (m *ConditionalAccessRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConditionalAccessRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the entry point for the Conditional Access (CA) object model.
func (m *ConditionalAccessRequestBuilder) Get(ctx context.Context, requestConfiguration *ConditionalAccessRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessRootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateConditionalAccessRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessRootable), nil
}
// NamedLocations provides operations to manage the namedLocations property of the microsoft.graph.conditionalAccessRoot entity.
func (m *ConditionalAccessRequestBuilder) NamedLocations()(*ConditionalAccessNamedLocationsRequestBuilder) {
    return NewConditionalAccessNamedLocationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NamedLocationsById provides operations to manage the namedLocations property of the microsoft.graph.conditionalAccessRoot entity.
func (m *ConditionalAccessRequestBuilder) NamedLocationsById(id string)(*ConditionalAccessNamedLocationsNamedLocationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["namedLocation%2Did"] = id
    }
    return NewConditionalAccessNamedLocationsNamedLocationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property conditionalAccess in identity
func (m *ConditionalAccessRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessRootable, requestConfiguration *ConditionalAccessRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessRootable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateConditionalAccessRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessRootable), nil
}
// Policies provides operations to manage the policies property of the microsoft.graph.conditionalAccessRoot entity.
func (m *ConditionalAccessRequestBuilder) Policies()(*ConditionalAccessPoliciesRequestBuilder) {
    return NewConditionalAccessPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PoliciesById provides operations to manage the policies property of the microsoft.graph.conditionalAccessRoot entity.
func (m *ConditionalAccessRequestBuilder) PoliciesById(id string)(*ConditionalAccessPoliciesConditionalAccessPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conditionalAccessPolicy%2Did"] = id
    }
    return NewConditionalAccessPoliciesConditionalAccessPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Templates provides operations to manage the templates property of the microsoft.graph.conditionalAccessRoot entity.
func (m *ConditionalAccessRequestBuilder) Templates()(*ConditionalAccessTemplatesRequestBuilder) {
    return NewConditionalAccessTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TemplatesById provides operations to manage the templates property of the microsoft.graph.conditionalAccessRoot entity.
func (m *ConditionalAccessRequestBuilder) TemplatesById(id string)(*ConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conditionalAccessTemplate%2Did"] = id
    }
    return NewConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property conditionalAccess for identity
func (m *ConditionalAccessRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConditionalAccessRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the entry point for the Conditional Access (CA) object model.
func (m *ConditionalAccessRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConditionalAccessRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property conditionalAccess in identity
func (m *ConditionalAccessRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessRootable, requestConfiguration *ConditionalAccessRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
