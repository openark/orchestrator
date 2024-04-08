package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilder provides operations to manage the permissionGrantPolicies property of the microsoft.graph.policyRoot entity.
type PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilderGetQueryParameters the policy that specifies the conditions under which consent can be granted.
type PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilderGetQueryParameters
}
// PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilderInternal instantiates a new PermissionGrantPolicyItemRequestBuilder and sets the default values.
func NewPermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilder) {
    m := &PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/permissionGrantPolicies/{permissionGrantPolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewPermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilder instantiates a new PermissionGrantPolicyItemRequestBuilder and sets the default values.
func NewPermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property permissionGrantPolicies for policies
func (m *PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Excludes provides operations to manage the excludes property of the microsoft.graph.permissionGrantPolicy entity.
func (m *PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilder) Excludes()(*PermissionGrantPoliciesItemExcludesRequestBuilder) {
    return NewPermissionGrantPoliciesItemExcludesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExcludesById provides operations to manage the excludes property of the microsoft.graph.permissionGrantPolicy entity.
func (m *PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilder) ExcludesById(id string)(*PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permissionGrantConditionSet%2Did"] = id
    }
    return NewPermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the policy that specifies the conditions under which consent can be granted.
func (m *PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePermissionGrantPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantPolicyable), nil
}
// Includes provides operations to manage the includes property of the microsoft.graph.permissionGrantPolicy entity.
func (m *PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilder) Includes()(*PermissionGrantPoliciesItemIncludesRequestBuilder) {
    return NewPermissionGrantPoliciesItemIncludesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IncludesById provides operations to manage the includes property of the microsoft.graph.permissionGrantPolicy entity.
func (m *PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilder) IncludesById(id string)(*PermissionGrantPoliciesItemIncludesPermissionGrantConditionSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permissionGrantConditionSet%2Did"] = id
    }
    return NewPermissionGrantPoliciesItemIncludesPermissionGrantConditionSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property permissionGrantPolicies in policies
func (m *PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantPolicyable, requestConfiguration *PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePermissionGrantPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantPolicyable), nil
}
// ToDeleteRequestInformation delete navigation property permissionGrantPolicies for policies
func (m *PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the policy that specifies the conditions under which consent can be granted.
func (m *PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property permissionGrantPolicies in policies
func (m *PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantPolicyable, requestConfiguration *PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
