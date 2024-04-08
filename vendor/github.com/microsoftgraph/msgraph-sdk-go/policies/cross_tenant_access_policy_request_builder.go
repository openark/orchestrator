package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CrossTenantAccessPolicyRequestBuilder provides operations to manage the crossTenantAccessPolicy property of the microsoft.graph.policyRoot entity.
type CrossTenantAccessPolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CrossTenantAccessPolicyRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrossTenantAccessPolicyRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CrossTenantAccessPolicyRequestBuilderGetQueryParameters read the properties and relationships of a crossTenantAccessPolicy object.
type CrossTenantAccessPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CrossTenantAccessPolicyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrossTenantAccessPolicyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CrossTenantAccessPolicyRequestBuilderGetQueryParameters
}
// CrossTenantAccessPolicyRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrossTenantAccessPolicyRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCrossTenantAccessPolicyRequestBuilderInternal instantiates a new CrossTenantAccessPolicyRequestBuilder and sets the default values.
func NewCrossTenantAccessPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CrossTenantAccessPolicyRequestBuilder) {
    m := &CrossTenantAccessPolicyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/crossTenantAccessPolicy{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCrossTenantAccessPolicyRequestBuilder instantiates a new CrossTenantAccessPolicyRequestBuilder and sets the default values.
func NewCrossTenantAccessPolicyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CrossTenantAccessPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCrossTenantAccessPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// DefaultEscaped provides operations to manage the default property of the microsoft.graph.crossTenantAccessPolicy entity.
func (m *CrossTenantAccessPolicyRequestBuilder) DefaultEscaped()(*CrossTenantAccessPolicyDefaultRequestBuilder) {
    return NewCrossTenantAccessPolicyDefaultRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete navigation property crossTenantAccessPolicy for policies
func (m *CrossTenantAccessPolicyRequestBuilder) Delete(ctx context.Context, requestConfiguration *CrossTenantAccessPolicyRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a crossTenantAccessPolicy object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/crosstenantaccesspolicy-get?view=graph-rest-1.0
func (m *CrossTenantAccessPolicyRequestBuilder) Get(ctx context.Context, requestConfiguration *CrossTenantAccessPolicyRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCrossTenantAccessPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyable), nil
}
// Partners provides operations to manage the partners property of the microsoft.graph.crossTenantAccessPolicy entity.
func (m *CrossTenantAccessPolicyRequestBuilder) Partners()(*CrossTenantAccessPolicyPartnersRequestBuilder) {
    return NewCrossTenantAccessPolicyPartnersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PartnersById provides operations to manage the partners property of the microsoft.graph.crossTenantAccessPolicy entity.
func (m *CrossTenantAccessPolicyRequestBuilder) PartnersById(id string)(*CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["crossTenantAccessPolicyConfigurationPartner%2DtenantId"] = id
    }
    return NewCrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the properties of a cross-tenant access policy.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/crosstenantaccesspolicy-update?view=graph-rest-1.0
func (m *CrossTenantAccessPolicyRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyable, requestConfiguration *CrossTenantAccessPolicyRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCrossTenantAccessPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyable), nil
}
// ToDeleteRequestInformation delete navigation property crossTenantAccessPolicy for policies
func (m *CrossTenantAccessPolicyRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CrossTenantAccessPolicyRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation read the properties and relationships of a crossTenantAccessPolicy object.
func (m *CrossTenantAccessPolicyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CrossTenantAccessPolicyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a cross-tenant access policy.
func (m *CrossTenantAccessPolicyRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyable, requestConfiguration *CrossTenantAccessPolicyRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
