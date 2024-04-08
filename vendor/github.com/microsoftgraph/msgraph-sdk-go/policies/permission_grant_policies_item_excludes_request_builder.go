package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PermissionGrantPoliciesItemExcludesRequestBuilder provides operations to manage the excludes property of the microsoft.graph.permissionGrantPolicy entity.
type PermissionGrantPoliciesItemExcludesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PermissionGrantPoliciesItemExcludesRequestBuilderGetQueryParameters retrieve the condition sets which are *excluded* in a permissionGrantPolicy.
type PermissionGrantPoliciesItemExcludesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// PermissionGrantPoliciesItemExcludesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionGrantPoliciesItemExcludesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissionGrantPoliciesItemExcludesRequestBuilderGetQueryParameters
}
// PermissionGrantPoliciesItemExcludesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionGrantPoliciesItemExcludesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPermissionGrantPoliciesItemExcludesRequestBuilderInternal instantiates a new ExcludesRequestBuilder and sets the default values.
func NewPermissionGrantPoliciesItemExcludesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionGrantPoliciesItemExcludesRequestBuilder) {
    m := &PermissionGrantPoliciesItemExcludesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/permissionGrantPolicies/{permissionGrantPolicy%2Did}/excludes{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewPermissionGrantPoliciesItemExcludesRequestBuilder instantiates a new ExcludesRequestBuilder and sets the default values.
func NewPermissionGrantPoliciesItemExcludesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionGrantPoliciesItemExcludesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionGrantPoliciesItemExcludesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *PermissionGrantPoliciesItemExcludesRequestBuilder) Count()(*PermissionGrantPoliciesItemExcludesCountRequestBuilder) {
    return NewPermissionGrantPoliciesItemExcludesCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get retrieve the condition sets which are *excluded* in a permissionGrantPolicy.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/permissiongrantpolicy-list-excludes?view=graph-rest-1.0
func (m *PermissionGrantPoliciesItemExcludesRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissionGrantPoliciesItemExcludesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePermissionGrantConditionSetCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetCollectionResponseable), nil
}
// Post add conditions under which a permission grant event is *excluded* in a permission grant policy. You do this by adding a permissionGrantConditionSet to the **excludes** collection of a  permissionGrantPolicy.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/permissiongrantpolicy-post-excludes?view=graph-rest-1.0
func (m *PermissionGrantPoliciesItemExcludesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetable, requestConfiguration *PermissionGrantPoliciesItemExcludesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation retrieve the condition sets which are *excluded* in a permissionGrantPolicy.
func (m *PermissionGrantPoliciesItemExcludesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissionGrantPoliciesItemExcludesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation add conditions under which a permission grant event is *excluded* in a permission grant policy. You do this by adding a permissionGrantConditionSet to the **excludes** collection of a  permissionGrantPolicy.
func (m *PermissionGrantPoliciesItemExcludesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetable, requestConfiguration *PermissionGrantPoliciesItemExcludesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
