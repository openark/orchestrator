package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemListsItemContentTypesItemMicrosoftGraphIsPublishedRequestBuilder provides operations to call the isPublished method.
type ItemSitesItemListsItemContentTypesItemMicrosoftGraphIsPublishedRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemListsItemContentTypesItemMicrosoftGraphIsPublishedRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemContentTypesItemMicrosoftGraphIsPublishedRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemListsItemContentTypesItemMicrosoftGraphIsPublishedRequestBuilderInternal instantiates a new MicrosoftGraphIsPublishedRequestBuilder and sets the default values.
func NewItemSitesItemListsItemContentTypesItemMicrosoftGraphIsPublishedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemContentTypesItemMicrosoftGraphIsPublishedRequestBuilder) {
    m := &ItemSitesItemListsItemContentTypesItemMicrosoftGraphIsPublishedRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/contentTypes/{contentType%2Did}/microsoft.graph.isPublished()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesItemListsItemContentTypesItemMicrosoftGraphIsPublishedRequestBuilder instantiates a new MicrosoftGraphIsPublishedRequestBuilder and sets the default values.
func NewItemSitesItemListsItemContentTypesItemMicrosoftGraphIsPublishedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemContentTypesItemMicrosoftGraphIsPublishedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemListsItemContentTypesItemMicrosoftGraphIsPublishedRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function isPublished
func (m *ItemSitesItemListsItemContentTypesItemMicrosoftGraphIsPublishedRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemListsItemContentTypesItemMicrosoftGraphIsPublishedRequestBuilderGetRequestConfiguration)(ItemSitesItemListsItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemSitesItemListsItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemListsItemContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponseable), nil
}
// ToGetRequestInformation invoke function isPublished
func (m *ItemSitesItemListsItemContentTypesItemMicrosoftGraphIsPublishedRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemListsItemContentTypesItemMicrosoftGraphIsPublishedRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
