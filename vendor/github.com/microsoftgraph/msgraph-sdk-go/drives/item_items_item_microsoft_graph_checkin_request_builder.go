package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemMicrosoftGraphCheckinRequestBuilder provides operations to call the checkin method.
type ItemItemsItemMicrosoftGraphCheckinRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemMicrosoftGraphCheckinRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemMicrosoftGraphCheckinRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemMicrosoftGraphCheckinRequestBuilderInternal instantiates a new MicrosoftGraphCheckinRequestBuilder and sets the default values.
func NewItemItemsItemMicrosoftGraphCheckinRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemMicrosoftGraphCheckinRequestBuilder) {
    m := &ItemItemsItemMicrosoftGraphCheckinRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/microsoft.graph.checkin";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemMicrosoftGraphCheckinRequestBuilder instantiates a new MicrosoftGraphCheckinRequestBuilder and sets the default values.
func NewItemItemsItemMicrosoftGraphCheckinRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemMicrosoftGraphCheckinRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemMicrosoftGraphCheckinRequestBuilderInternal(urlParams, requestAdapter)
}
// Post check in a checked out **driveItem** resource, which makes the version of the document available to others.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/driveitem-checkin?view=graph-rest-1.0
func (m *ItemItemsItemMicrosoftGraphCheckinRequestBuilder) Post(ctx context.Context, body ItemItemsItemMicrosoftGraphCheckinCheckinPostRequestBodyable, requestConfiguration *ItemItemsItemMicrosoftGraphCheckinRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation check in a checked out **driveItem** resource, which makes the version of the document available to others.
func (m *ItemItemsItemMicrosoftGraphCheckinRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemMicrosoftGraphCheckinCheckinPostRequestBodyable, requestConfiguration *ItemItemsItemMicrosoftGraphCheckinRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
