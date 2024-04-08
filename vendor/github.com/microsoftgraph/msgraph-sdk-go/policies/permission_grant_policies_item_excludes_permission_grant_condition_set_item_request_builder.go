package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder provides operations to manage the excludes property of the microsoft.graph.permissionGrantPolicy entity.
type PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderGetQueryParameters condition sets which are excluded in this permission grant policy. Automatically expanded on GET.
type PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderGetQueryParameters
}
// PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderInternal instantiates a new PermissionGrantConditionSetItemRequestBuilder and sets the default values.
func NewPermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder) {
    m := &PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/permissionGrantPolicies/{permissionGrantPolicy%2Did}/excludes/{permissionGrantConditionSet%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewPermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder instantiates a new PermissionGrantConditionSetItemRequestBuilder and sets the default values.
func NewPermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property excludes for policies
func (m *PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get condition sets which are excluded in this permission grant policy. Automatically expanded on GET.
func (m *PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePermissionGrantConditionSetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetable), nil
}
// Patch update the navigation property excludes in policies
func (m *PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetable, requestConfiguration *PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePermissionGrantConditionSetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetable), nil
}
// ToDeleteRequestInformation delete navigation property excludes for policies
func (m *PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation condition sets which are excluded in this permission grant policy. Automatically expanded on GET.
func (m *PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property excludes in policies
func (m *PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetable, requestConfiguration *PermissionGrantPoliciesItemExcludesPermissionGrantConditionSetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
