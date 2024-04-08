package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemMicrosoftGraphInviteRequestBuilder provides operations to call the invite method.
type ItemItemsItemMicrosoftGraphInviteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemMicrosoftGraphInviteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemMicrosoftGraphInviteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemMicrosoftGraphInviteRequestBuilderInternal instantiates a new MicrosoftGraphInviteRequestBuilder and sets the default values.
func NewItemItemsItemMicrosoftGraphInviteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemMicrosoftGraphInviteRequestBuilder) {
    m := &ItemItemsItemMicrosoftGraphInviteRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/microsoft.graph.invite";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemMicrosoftGraphInviteRequestBuilder instantiates a new MicrosoftGraphInviteRequestBuilder and sets the default values.
func NewItemItemsItemMicrosoftGraphInviteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemMicrosoftGraphInviteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemMicrosoftGraphInviteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post sends a sharing invitation for a **driveItem**.A sharing invitation provides permissions to the recipients and optionally sends them an email with a [sharing link][].
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/driveitem-invite?view=graph-rest-1.0
func (m *ItemItemsItemMicrosoftGraphInviteRequestBuilder) Post(ctx context.Context, body ItemItemsItemMicrosoftGraphInviteInvitePostRequestBodyable, requestConfiguration *ItemItemsItemMicrosoftGraphInviteRequestBuilderPostRequestConfiguration)(ItemItemsItemMicrosoftGraphInviteInviteResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemItemsItemMicrosoftGraphInviteInviteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemsItemMicrosoftGraphInviteInviteResponseable), nil
}
// ToPostRequestInformation sends a sharing invitation for a **driveItem**.A sharing invitation provides permissions to the recipients and optionally sends them an email with a [sharing link][].
func (m *ItemItemsItemMicrosoftGraphInviteRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemMicrosoftGraphInviteInvitePostRequestBodyable, requestConfiguration *ItemItemsItemMicrosoftGraphInviteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
