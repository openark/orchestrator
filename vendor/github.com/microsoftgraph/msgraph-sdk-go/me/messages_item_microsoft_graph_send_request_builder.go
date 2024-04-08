package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MessagesItemMicrosoftGraphSendRequestBuilder provides operations to call the send method.
type MessagesItemMicrosoftGraphSendRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MessagesItemMicrosoftGraphSendRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MessagesItemMicrosoftGraphSendRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMessagesItemMicrosoftGraphSendRequestBuilderInternal instantiates a new MicrosoftGraphSendRequestBuilder and sets the default values.
func NewMessagesItemMicrosoftGraphSendRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MessagesItemMicrosoftGraphSendRequestBuilder) {
    m := &MessagesItemMicrosoftGraphSendRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/messages/{message%2Did}/microsoft.graph.send";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMessagesItemMicrosoftGraphSendRequestBuilder instantiates a new MicrosoftGraphSendRequestBuilder and sets the default values.
func NewMessagesItemMicrosoftGraphSendRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MessagesItemMicrosoftGraphSendRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMessagesItemMicrosoftGraphSendRequestBuilderInternal(urlParams, requestAdapter)
}
// Post send an existing draft message.  The draft message can be a new message draft, reply draft, reply-all draft, or a forward draft. This method saves the message in the **Sent Items** folder. Alternatively, send a new message in a single operation.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/message-send?view=graph-rest-1.0
func (m *MessagesItemMicrosoftGraphSendRequestBuilder) Post(ctx context.Context, requestConfiguration *MessagesItemMicrosoftGraphSendRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation send an existing draft message.  The draft message can be a new message draft, reply draft, reply-all draft, or a forward draft. This method saves the message in the **Sent Items** folder. Alternatively, send a new message in a single operation.
func (m *MessagesItemMicrosoftGraphSendRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MessagesItemMicrosoftGraphSendRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
