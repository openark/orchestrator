package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemChatsItemMicrosoftGraphSendActivityNotificationRequestBuilder provides operations to call the sendActivityNotification method.
type ItemChatsItemMicrosoftGraphSendActivityNotificationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemChatsItemMicrosoftGraphSendActivityNotificationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemMicrosoftGraphSendActivityNotificationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemChatsItemMicrosoftGraphSendActivityNotificationRequestBuilderInternal instantiates a new MicrosoftGraphSendActivityNotificationRequestBuilder and sets the default values.
func NewItemChatsItemMicrosoftGraphSendActivityNotificationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemMicrosoftGraphSendActivityNotificationRequestBuilder) {
    m := &ItemChatsItemMicrosoftGraphSendActivityNotificationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/chats/{chat%2Did}/microsoft.graph.sendActivityNotification";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemChatsItemMicrosoftGraphSendActivityNotificationRequestBuilder instantiates a new MicrosoftGraphSendActivityNotificationRequestBuilder and sets the default values.
func NewItemChatsItemMicrosoftGraphSendActivityNotificationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemMicrosoftGraphSendActivityNotificationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChatsItemMicrosoftGraphSendActivityNotificationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post send an activity feed notification in scope of a chat. For more details about sending notifications and the requirements for doing so, see sending Teams activity notifications.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/chat-sendactivitynotification?view=graph-rest-1.0
func (m *ItemChatsItemMicrosoftGraphSendActivityNotificationRequestBuilder) Post(ctx context.Context, body ItemChatsItemMicrosoftGraphSendActivityNotificationSendActivityNotificationPostRequestBodyable, requestConfiguration *ItemChatsItemMicrosoftGraphSendActivityNotificationRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation send an activity feed notification in scope of a chat. For more details about sending notifications and the requirements for doing so, see sending Teams activity notifications.
func (m *ItemChatsItemMicrosoftGraphSendActivityNotificationRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemChatsItemMicrosoftGraphSendActivityNotificationSendActivityNotificationPostRequestBodyable, requestConfiguration *ItemChatsItemMicrosoftGraphSendActivityNotificationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
