package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilder provides operations to call the softDelete method.
type ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilderInternal instantiates a new MicrosoftGraphSoftDeleteRequestBuilder and sets the default values.
func NewItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilder) {
    m := &ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/team/primaryChannel/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/microsoft.graph.softDelete";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilder instantiates a new MicrosoftGraphSoftDeleteRequestBuilder and sets the default values.
func NewItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action softDelete
func (m *ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation invoke action softDelete
func (m *ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
