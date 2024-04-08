package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilder provides operations to call the getCompatibleHubContentTypes method.
type ItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilderGetQueryParameters invoke function getCompatibleHubContentTypes
type ItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilderGetQueryParameters struct {
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
// ItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilderGetQueryParameters
}
// NewItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilderInternal instantiates a new MicrosoftGraphGetCompatibleHubContentTypesRequestBuilder and sets the default values.
func NewItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilder) {
    m := &ItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/contentTypes/microsoft.graph.getCompatibleHubContentTypes(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilder instantiates a new MicrosoftGraphGetCompatibleHubContentTypesRequestBuilder and sets the default values.
func NewItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getCompatibleHubContentTypes
func (m *ItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration)(ItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesGetCompatibleHubContentTypesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesGetCompatibleHubContentTypesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesGetCompatibleHubContentTypesResponseable), nil
}
// ToGetRequestInformation invoke function getCompatibleHubContentTypes
func (m *ItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemContentTypesMicrosoftGraphGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
