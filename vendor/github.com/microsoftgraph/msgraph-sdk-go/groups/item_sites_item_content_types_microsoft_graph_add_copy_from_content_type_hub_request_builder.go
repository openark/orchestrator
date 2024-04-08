package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemContentTypesMicrosoftGraphAddCopyFromContentTypeHubRequestBuilder provides operations to call the addCopyFromContentTypeHub method.
type ItemSitesItemContentTypesMicrosoftGraphAddCopyFromContentTypeHubRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemContentTypesMicrosoftGraphAddCopyFromContentTypeHubRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemContentTypesMicrosoftGraphAddCopyFromContentTypeHubRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemContentTypesMicrosoftGraphAddCopyFromContentTypeHubRequestBuilderInternal instantiates a new MicrosoftGraphAddCopyFromContentTypeHubRequestBuilder and sets the default values.
func NewItemSitesItemContentTypesMicrosoftGraphAddCopyFromContentTypeHubRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemContentTypesMicrosoftGraphAddCopyFromContentTypeHubRequestBuilder) {
    m := &ItemSitesItemContentTypesMicrosoftGraphAddCopyFromContentTypeHubRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/contentTypes/microsoft.graph.addCopyFromContentTypeHub";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesItemContentTypesMicrosoftGraphAddCopyFromContentTypeHubRequestBuilder instantiates a new MicrosoftGraphAddCopyFromContentTypeHubRequestBuilder and sets the default values.
func NewItemSitesItemContentTypesMicrosoftGraphAddCopyFromContentTypeHubRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemContentTypesMicrosoftGraphAddCopyFromContentTypeHubRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemContentTypesMicrosoftGraphAddCopyFromContentTypeHubRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add or sync a copy of a published content type from the content type hub to a target site or a list. This method is part of the content type publishing changes to optimize the syncing of published content types to sites and lists, effectively switching from a 'push everywhere' to 'pull as needed' approach. The method allows users to pull content types directly from the content type hub to a site or list. For more information, see contentType: getCompatibleHubContentTypes and the blog post Syntex Product Updates – August 2021.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/contenttype-addcopyfromcontenttypehub?view=graph-rest-1.0
func (m *ItemSitesItemContentTypesMicrosoftGraphAddCopyFromContentTypeHubRequestBuilder) Post(ctx context.Context, body ItemSitesItemContentTypesMicrosoftGraphAddCopyFromContentTypeHubAddCopyFromContentTypeHubPostRequestBodyable, requestConfiguration *ItemSitesItemContentTypesMicrosoftGraphAddCopyFromContentTypeHubRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateContentTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable), nil
}
// ToPostRequestInformation add or sync a copy of a published content type from the content type hub to a target site or a list. This method is part of the content type publishing changes to optimize the syncing of published content types to sites and lists, effectively switching from a 'push everywhere' to 'pull as needed' approach. The method allows users to pull content types directly from the content type hub to a site or list. For more information, see contentType: getCompatibleHubContentTypes and the blog post Syntex Product Updates – August 2021.
func (m *ItemSitesItemContentTypesMicrosoftGraphAddCopyFromContentTypeHubRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemContentTypesMicrosoftGraphAddCopyFromContentTypeHubAddCopyFromContentTypeHubPostRequestBodyable, requestConfiguration *ItemSitesItemContentTypesMicrosoftGraphAddCopyFromContentTypeHubRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
