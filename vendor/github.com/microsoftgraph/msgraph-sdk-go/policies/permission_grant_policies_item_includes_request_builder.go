package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PermissionGrantPoliciesItemIncludesRequestBuilder provides operations to manage the includes property of the microsoft.graph.permissionGrantPolicy entity.
type PermissionGrantPoliciesItemIncludesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PermissionGrantPoliciesItemIncludesRequestBuilderGetQueryParameters retrieve the condition sets which are *included* in a permissionGrantPolicy.
type PermissionGrantPoliciesItemIncludesRequestBuilderGetQueryParameters struct {
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
// PermissionGrantPoliciesItemIncludesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionGrantPoliciesItemIncludesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissionGrantPoliciesItemIncludesRequestBuilderGetQueryParameters
}
// PermissionGrantPoliciesItemIncludesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionGrantPoliciesItemIncludesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPermissionGrantPoliciesItemIncludesRequestBuilderInternal instantiates a new IncludesRequestBuilder and sets the default values.
func NewPermissionGrantPoliciesItemIncludesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionGrantPoliciesItemIncludesRequestBuilder) {
    m := &PermissionGrantPoliciesItemIncludesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/permissionGrantPolicies/{permissionGrantPolicy%2Did}/includes{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewPermissionGrantPoliciesItemIncludesRequestBuilder instantiates a new IncludesRequestBuilder and sets the default values.
func NewPermissionGrantPoliciesItemIncludesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionGrantPoliciesItemIncludesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionGrantPoliciesItemIncludesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *PermissionGrantPoliciesItemIncludesRequestBuilder) Count()(*PermissionGrantPoliciesItemIncludesCountRequestBuilder) {
    return NewPermissionGrantPoliciesItemIncludesCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get retrieve the condition sets which are *included* in a permissionGrantPolicy.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/permissiongrantpolicy-list-includes?view=graph-rest-1.0
func (m *PermissionGrantPoliciesItemIncludesRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissionGrantPoliciesItemIncludesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetCollectionResponseable, error) {
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
// Post add conditions under which a permission grant event is *included* in a permission grant policy. You do this by adding a permissionGrantConditionSet to the **includes** collection of a  permissionGrantPolicy.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/permissiongrantpolicy-post-includes?view=graph-rest-1.0
func (m *PermissionGrantPoliciesItemIncludesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetable, requestConfiguration *PermissionGrantPoliciesItemIncludesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetable, error) {
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
// ToGetRequestInformation retrieve the condition sets which are *included* in a permissionGrantPolicy.
func (m *PermissionGrantPoliciesItemIncludesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissionGrantPoliciesItemIncludesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation add conditions under which a permission grant event is *included* in a permission grant policy. You do this by adding a permissionGrantConditionSet to the **includes** collection of a  permissionGrantPolicy.
func (m *PermissionGrantPoliciesItemIncludesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PermissionGrantConditionSetable, requestConfiguration *PermissionGrantPoliciesItemIncludesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
