package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ChatsItemMicrosoftGraphUnhideForUserRequestBuilder provides operations to call the unhideForUser method.
type ChatsItemMicrosoftGraphUnhideForUserRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ChatsItemMicrosoftGraphUnhideForUserRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatsItemMicrosoftGraphUnhideForUserRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChatsItemMicrosoftGraphUnhideForUserRequestBuilderInternal instantiates a new MicrosoftGraphUnhideForUserRequestBuilder and sets the default values.
func NewChatsItemMicrosoftGraphUnhideForUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatsItemMicrosoftGraphUnhideForUserRequestBuilder) {
    m := &ChatsItemMicrosoftGraphUnhideForUserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/chats/{chat%2Did}/microsoft.graph.unhideForUser";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewChatsItemMicrosoftGraphUnhideForUserRequestBuilder instantiates a new MicrosoftGraphUnhideForUserRequestBuilder and sets the default values.
func NewChatsItemMicrosoftGraphUnhideForUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatsItemMicrosoftGraphUnhideForUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatsItemMicrosoftGraphUnhideForUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Post unhide a chat for a user.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/chat-unhideforuser?view=graph-rest-1.0
func (m *ChatsItemMicrosoftGraphUnhideForUserRequestBuilder) Post(ctx context.Context, body ChatsItemMicrosoftGraphUnhideForUserUnhideForUserPostRequestBodyable, requestConfiguration *ChatsItemMicrosoftGraphUnhideForUserRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation unhide a chat for a user.
func (m *ChatsItemMicrosoftGraphUnhideForUserRequestBuilder) ToPostRequestInformation(ctx context.Context, body ChatsItemMicrosoftGraphUnhideForUserUnhideForUserPostRequestBodyable, requestConfiguration *ChatsItemMicrosoftGraphUnhideForUserRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
