package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilder provides operations to manage the pinnedMessages property of the microsoft.graph.chat entity.
type ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilderGetQueryParameters a collection of all the pinned messages in the chat. Nullable.
type ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilderGetQueryParameters
}
// ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilderInternal instantiates a new PinnedChatMessageInfoItemRequestBuilder and sets the default values.
func NewChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilder) {
    m := &ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/chats/{chat%2Did}/pinnedMessages/{pinnedChatMessageInfo%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilder instantiates a new PinnedChatMessageInfoItemRequestBuilder and sets the default values.
func NewChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property pinnedMessages for me
func (m *ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get a collection of all the pinned messages in the chat. Nullable.
func (m *ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PinnedChatMessageInfoable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePinnedChatMessageInfoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PinnedChatMessageInfoable), nil
}
// Message provides operations to manage the message property of the microsoft.graph.pinnedChatMessageInfo entity.
func (m *ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilder) Message()(*ChatsItemPinnedMessagesItemMessageRequestBuilder) {
    return NewChatsItemPinnedMessagesItemMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property pinnedMessages in me
func (m *ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PinnedChatMessageInfoable, requestConfiguration *ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PinnedChatMessageInfoable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePinnedChatMessageInfoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PinnedChatMessageInfoable), nil
}
// ToDeleteRequestInformation delete navigation property pinnedMessages for me
func (m *ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation a collection of all the pinned messages in the chat. Nullable.
func (m *ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property pinnedMessages in me
func (m *ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PinnedChatMessageInfoable, requestConfiguration *ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
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
