package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilder provides operations to count the resources in the collection.
type ItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilderGetQueryParameters get the number of the resource
type ItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// ItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilderGetQueryParameters
}
// NewItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilder) {
    m := &ItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/termStores/{store%2Did}/groups/{group%2Did}/sets/{set%2Did}/children/{term%2Did}/children/{term%2Did1}/relations/$count{?%24search,%24filter}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *ItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
func (m *ItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTermStoresItemGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "text/plain")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
