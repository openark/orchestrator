package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ChatsChatItemRequestBuilder provides operations to manage the chats property of the microsoft.graph.user entity.
type ChatsChatItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ChatsChatItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatsChatItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChatsChatItemRequestBuilderGetQueryParameters get chats from me
type ChatsChatItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ChatsChatItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatsChatItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChatsChatItemRequestBuilderGetQueryParameters
}
// ChatsChatItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatsChatItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChatsChatItemRequestBuilderInternal instantiates a new ChatItemRequestBuilder and sets the default values.
func NewChatsChatItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatsChatItemRequestBuilder) {
    m := &ChatsChatItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/chats/{chat%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewChatsChatItemRequestBuilder instantiates a new ChatItemRequestBuilder and sets the default values.
func NewChatsChatItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatsChatItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatsChatItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property chats for me
func (m *ChatsChatItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ChatsChatItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get chats from me
func (m *ChatsChatItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ChatsChatItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChatFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable), nil
}
// InstalledApps provides operations to manage the installedApps property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) InstalledApps()(*ChatsItemInstalledAppsRequestBuilder) {
    return NewChatsItemInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InstalledAppsById provides operations to manage the installedApps property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) InstalledAppsById(id string)(*ChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return NewChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// LastMessagePreview provides operations to manage the lastMessagePreview property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) LastMessagePreview()(*ChatsItemLastMessagePreviewRequestBuilder) {
    return NewChatsItemLastMessagePreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Members provides operations to manage the members property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) Members()(*ChatsItemMembersRequestBuilder) {
    return NewChatsItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MembersById provides operations to manage the members property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) MembersById(id string)(*ChatsItemMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewChatsItemMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Messages provides operations to manage the messages property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) Messages()(*ChatsItemMessagesRequestBuilder) {
    return NewChatsItemMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) MessagesById(id string)(*ChatsItemMessagesChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return NewChatsItemMessagesChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphHideForUser provides operations to call the hideForUser method.
func (m *ChatsChatItemRequestBuilder) MicrosoftGraphHideForUser()(*ChatsItemMicrosoftGraphHideForUserRequestBuilder) {
    return NewChatsItemMicrosoftGraphHideForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMarkChatReadForUser provides operations to call the markChatReadForUser method.
func (m *ChatsChatItemRequestBuilder) MicrosoftGraphMarkChatReadForUser()(*ChatsItemMicrosoftGraphMarkChatReadForUserRequestBuilder) {
    return NewChatsItemMicrosoftGraphMarkChatReadForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMarkChatUnreadForUser provides operations to call the markChatUnreadForUser method.
func (m *ChatsChatItemRequestBuilder) MicrosoftGraphMarkChatUnreadForUser()(*ChatsItemMicrosoftGraphMarkChatUnreadForUserRequestBuilder) {
    return NewChatsItemMicrosoftGraphMarkChatUnreadForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSendActivityNotification provides operations to call the sendActivityNotification method.
func (m *ChatsChatItemRequestBuilder) MicrosoftGraphSendActivityNotification()(*ChatsItemMicrosoftGraphSendActivityNotificationRequestBuilder) {
    return NewChatsItemMicrosoftGraphSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUnhideForUser provides operations to call the unhideForUser method.
func (m *ChatsChatItemRequestBuilder) MicrosoftGraphUnhideForUser()(*ChatsItemMicrosoftGraphUnhideForUserRequestBuilder) {
    return NewChatsItemMicrosoftGraphUnhideForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property chats in me
func (m *ChatsChatItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable, requestConfiguration *ChatsChatItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChatFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable), nil
}
// PinnedMessages provides operations to manage the pinnedMessages property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) PinnedMessages()(*ChatsItemPinnedMessagesRequestBuilder) {
    return NewChatsItemPinnedMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PinnedMessagesById provides operations to manage the pinnedMessages property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) PinnedMessagesById(id string)(*ChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["pinnedChatMessageInfo%2Did"] = id
    }
    return NewChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Tabs provides operations to manage the tabs property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) Tabs()(*ChatsItemTabsRequestBuilder) {
    return NewChatsItemTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TabsById provides operations to manage the tabs property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) TabsById(id string)(*ChatsItemTabsTeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return NewChatsItemTabsTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property chats for me
func (m *ChatsChatItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ChatsChatItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get chats from me
func (m *ChatsChatItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ChatsChatItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property chats in me
func (m *ChatsChatItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable, requestConfiguration *ChatsChatItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
