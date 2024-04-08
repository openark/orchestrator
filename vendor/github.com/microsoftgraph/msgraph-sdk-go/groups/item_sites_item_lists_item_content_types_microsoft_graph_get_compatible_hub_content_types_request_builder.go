package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilder provides operations to call the getCompatibleHubContentTypes method.
type ItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilderGetQueryParameters invoke function getCompatibleHubContentTypes
type ItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
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
// ItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilderGetQueryParameters
}
// NewItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilderInternal instantiates a new MicrosoftGraphGetCompatibleHubContentTypesRequestBuilder and sets the default values.
func NewItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilder) {
    m := &ItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/contentTypes/microsoft.graph.getCompatibleHubContentTypes(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilder instantiates a new MicrosoftGraphGetCompatibleHubContentTypesRequestBuilder and sets the default values.
func NewItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getCompatibleHubContentTypes
func (m *ItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration)(ItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesGetCompatibleHubContentTypesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesGetCompatibleHubContentTypesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesGetCompatibleHubContentTypesResponseable), nil
}
// ToGetRequestInformation invoke function getCompatibleHubContentTypes
func (m *ItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemListsItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
