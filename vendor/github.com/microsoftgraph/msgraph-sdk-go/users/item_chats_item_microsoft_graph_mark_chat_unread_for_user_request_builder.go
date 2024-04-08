package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemChatsItemMicrosoftGraphMarkChatUnreadForUserRequestBuilder provides operations to call the markChatUnreadForUser method.
type ItemChatsItemMicrosoftGraphMarkChatUnreadForUserRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemChatsItemMicrosoftGraphMarkChatUnreadForUserRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemMicrosoftGraphMarkChatUnreadForUserRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemChatsItemMicrosoftGraphMarkChatUnreadForUserRequestBuilderInternal instantiates a new MicrosoftGraphMarkChatUnreadForUserRequestBuilder and sets the default values.
func NewItemChatsItemMicrosoftGraphMarkChatUnreadForUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemMicrosoftGraphMarkChatUnreadForUserRequestBuilder) {
    m := &ItemChatsItemMicrosoftGraphMarkChatUnreadForUserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/chats/{chat%2Did}/microsoft.graph.markChatUnreadForUser";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemChatsItemMicrosoftGraphMarkChatUnreadForUserRequestBuilder instantiates a new MicrosoftGraphMarkChatUnreadForUserRequestBuilder and sets the default values.
func NewItemChatsItemMicrosoftGraphMarkChatUnreadForUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemMicrosoftGraphMarkChatUnreadForUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChatsItemMicrosoftGraphMarkChatUnreadForUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Post mark a chat as unread for a user.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/chat-markchatunreadforuser?view=graph-rest-1.0
func (m *ItemChatsItemMicrosoftGraphMarkChatUnreadForUserRequestBuilder) Post(ctx context.Context, body ItemChatsItemMicrosoftGraphMarkChatUnreadForUserMarkChatUnreadForUserPostRequestBodyable, requestConfiguration *ItemChatsItemMicrosoftGraphMarkChatUnreadForUserRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation mark a chat as unread for a user.
func (m *ItemChatsItemMicrosoftGraphMarkChatUnreadForUserRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemChatsItemMicrosoftGraphMarkChatUnreadForUserMarkChatUnreadForUserPostRequestBodyable, requestConfiguration *ItemChatsItemMicrosoftGraphMarkChatUnreadForUserRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
