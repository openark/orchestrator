package chats

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ChatItemRequestBuilder provides operations to manage the collection of chat entities.
type ChatItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ChatItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChatItemRequestBuilderGetQueryParameters retrieve a single chat (without its messages). This method supports federation. To access a chat, at least one chat member must belong to the tenant the request initiated from.
type ChatItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ChatItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChatItemRequestBuilderGetQueryParameters
}
// ChatItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChatItemRequestBuilderInternal instantiates a new ChatItemRequestBuilder and sets the default values.
func NewChatItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatItemRequestBuilder) {
    m := &ChatItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/chats/{chat%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewChatItemRequestBuilder instantiates a new ChatItemRequestBuilder and sets the default values.
func NewChatItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete entity from chats
func (m *ChatItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ChatItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve a single chat (without its messages). This method supports federation. To access a chat, at least one chat member must belong to the tenant the request initiated from.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/chat-get?view=graph-rest-1.0
func (m *ChatItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ChatItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable, error) {
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
func (m *ChatItemRequestBuilder) InstalledApps()(*ItemInstalledAppsRequestBuilder) {
    return NewItemInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InstalledAppsById provides operations to manage the installedApps property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) InstalledAppsById(id string)(*ItemInstalledAppsTeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return NewItemInstalledAppsTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// LastMessagePreview provides operations to manage the lastMessagePreview property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) LastMessagePreview()(*ItemLastMessagePreviewRequestBuilder) {
    return NewItemLastMessagePreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Members provides operations to manage the members property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) Members()(*ItemMembersRequestBuilder) {
    return NewItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MembersById provides operations to manage the members property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) MembersById(id string)(*ItemMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewItemMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Messages provides operations to manage the messages property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) Messages()(*ItemMessagesRequestBuilder) {
    return NewItemMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) MessagesById(id string)(*ItemMessagesChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return NewItemMessagesChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphHideForUser provides operations to call the hideForUser method.
func (m *ChatItemRequestBuilder) MicrosoftGraphHideForUser()(*ItemMicrosoftGraphHideForUserRequestBuilder) {
    return NewItemMicrosoftGraphHideForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMarkChatReadForUser provides operations to call the markChatReadForUser method.
func (m *ChatItemRequestBuilder) MicrosoftGraphMarkChatReadForUser()(*ItemMicrosoftGraphMarkChatReadForUserRequestBuilder) {
    return NewItemMicrosoftGraphMarkChatReadForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMarkChatUnreadForUser provides operations to call the markChatUnreadForUser method.
func (m *ChatItemRequestBuilder) MicrosoftGraphMarkChatUnreadForUser()(*ItemMicrosoftGraphMarkChatUnreadForUserRequestBuilder) {
    return NewItemMicrosoftGraphMarkChatUnreadForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSendActivityNotification provides operations to call the sendActivityNotification method.
func (m *ChatItemRequestBuilder) MicrosoftGraphSendActivityNotification()(*ItemMicrosoftGraphSendActivityNotificationRequestBuilder) {
    return NewItemMicrosoftGraphSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUnhideForUser provides operations to call the unhideForUser method.
func (m *ChatItemRequestBuilder) MicrosoftGraphUnhideForUser()(*ItemMicrosoftGraphUnhideForUserRequestBuilder) {
    return NewItemMicrosoftGraphUnhideForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the properties of a chat object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/chat-patch?view=graph-rest-1.0
func (m *ChatItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable, requestConfiguration *ChatItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable, error) {
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
func (m *ChatItemRequestBuilder) PinnedMessages()(*ItemPinnedMessagesRequestBuilder) {
    return NewItemPinnedMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PinnedMessagesById provides operations to manage the pinnedMessages property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) PinnedMessagesById(id string)(*ItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["pinnedChatMessageInfo%2Did"] = id
    }
    return NewItemPinnedMessagesPinnedChatMessageInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Tabs provides operations to manage the tabs property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) Tabs()(*ItemTabsRequestBuilder) {
    return NewItemTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TabsById provides operations to manage the tabs property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) TabsById(id string)(*ItemTabsTeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return NewItemTabsTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete entity from chats
func (m *ChatItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ChatItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation retrieve a single chat (without its messages). This method supports federation. To access a chat, at least one chat member must belong to the tenant the request initiated from.
func (m *ChatItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ChatItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a chat object.
func (m *ChatItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable, requestConfiguration *ChatItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
