package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemChatsChatItemRequestBuilder provides operations to manage the chats property of the microsoft.graph.user entity.
type ItemChatsChatItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemChatsChatItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsChatItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemChatsChatItemRequestBuilderGetQueryParameters get chats from users
type ItemChatsChatItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemChatsChatItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsChatItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemChatsChatItemRequestBuilderGetQueryParameters
}
// ItemChatsChatItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsChatItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemChatsChatItemRequestBuilderInternal instantiates a new ChatItemRequestBuilder and sets the default values.
func NewItemChatsChatItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsChatItemRequestBuilder) {
    m := &ItemChatsChatItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/chats/{chat%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemChatsChatItemRequestBuilder instantiates a new ChatItemRequestBuilder and sets the default values.
func NewItemChatsChatItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsChatItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChatsChatItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property chats for users
func (m *ItemChatsChatItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemChatsChatItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get chats from users
func (m *ItemChatsChatItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemChatsChatItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable, error) {
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
func (m *ItemChatsChatItemRequestBuilder) InstalledApps()(*ItemChatsItemInstalledAppsRequestBuilder) {
    return NewItemChatsItemInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InstalledAppsById provides operations to manage the installedApps property of the microsoft.graph.chat entity.
func (m *ItemChatsChatItemRequestBuilder) InstalledAppsById(id string)(*ItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return NewItemChatsItemInstalledAppsTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// LastMessagePreview provides operations to manage the lastMessagePreview property of the microsoft.graph.chat entity.
func (m *ItemChatsChatItemRequestBuilder) LastMessagePreview()(*ItemChatsItemLastMessagePreviewRequestBuilder) {
    return NewItemChatsItemLastMessagePreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Members provides operations to manage the members property of the microsoft.graph.chat entity.
func (m *ItemChatsChatItemRequestBuilder) Members()(*ItemChatsItemMembersRequestBuilder) {
    return NewItemChatsItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MembersById provides operations to manage the members property of the microsoft.graph.chat entity.
func (m *ItemChatsChatItemRequestBuilder) MembersById(id string)(*ItemChatsItemMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewItemChatsItemMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Messages provides operations to manage the messages property of the microsoft.graph.chat entity.
func (m *ItemChatsChatItemRequestBuilder) Messages()(*ItemChatsItemMessagesRequestBuilder) {
    return NewItemChatsItemMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.chat entity.
func (m *ItemChatsChatItemRequestBuilder) MessagesById(id string)(*ItemChatsItemMessagesChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return NewItemChatsItemMessagesChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphHideForUser provides operations to call the hideForUser method.
func (m *ItemChatsChatItemRequestBuilder) MicrosoftGraphHideForUser()(*ItemChatsItemMicrosoftGraphHideForUserRequestBuilder) {
    return NewItemChatsItemMicrosoftGraphHideForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMarkChatReadForUser provides operations to call the markChatReadForUser method.
func (m *ItemChatsChatItemRequestBuilder) MicrosoftGraphMarkChatReadForUser()(*ItemChatsItemMicrosoftGraphMarkChatReadForUserRequestBuilder) {
    return NewItemChatsItemMicrosoftGraphMarkChatReadForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMarkChatUnreadForUser provides operations to call the markChatUnreadForUser method.
func (m *ItemChatsChatItemRequestBuilder) MicrosoftGraphMarkChatUnreadForUser()(*ItemChatsItemMicrosoftGraphMarkChatUnreadForUserRequestBuilder) {
    return NewItemChatsItemMicrosoftGraphMarkChatUnreadForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSendActivityNotification provides operations to call the sendActivityNotification method.
func (m *ItemChatsChatItemRequestBuilder) MicrosoftGraphSendActivityNotification()(*ItemChatsItemMicrosoftGraphSendActivityNotificationRequestBuilder) {
    return NewItemChatsItemMicrosoftGraphSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUnhideForUser provides operations to call the unhideForUser method.
func (m *ItemChatsChatItemRequestBuilder) MicrosoftGraphUnhideForUser()(*ItemChatsItemMicrosoftGraphUnhideForUserRequestBuilder) {
    return NewItemChatsItemMicrosoftGraphUnhideForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property chats in users
func (m *ItemChatsChatItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable, requestConfiguration *ItemChatsChatItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable, error) {
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
func (m *ItemChatsChatItemRequestBuilder) PinnedMessages()(*ItemChatsItemPinnedMessagesRequestBuilder) {
    return NewItemChatsItemPinnedMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PinnedMessagesById provides operations to manage the pinnedMessages property of the microsoft.graph.chat entity.
func (m *ItemChatsChatItemRequestBuilder) PinnedMessagesById(id string)(*ItemChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["pinnedChatMessageInfo%2Did"] = id
    }
    return NewItemChatsItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Tabs provides operations to manage the tabs property of the microsoft.graph.chat entity.
func (m *ItemChatsChatItemRequestBuilder) Tabs()(*ItemChatsItemTabsRequestBuilder) {
    return NewItemChatsItemTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TabsById provides operations to manage the tabs property of the microsoft.graph.chat entity.
func (m *ItemChatsChatItemRequestBuilder) TabsById(id string)(*ItemChatsItemTabsTeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return NewItemChatsItemTabsTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property chats for users
func (m *ItemChatsChatItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemChatsChatItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get chats from users
func (m *ItemChatsChatItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemChatsChatItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property chats in users
func (m *ItemChatsChatItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable, requestConfiguration *ItemChatsChatItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
