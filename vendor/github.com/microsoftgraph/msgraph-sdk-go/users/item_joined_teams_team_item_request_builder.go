package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedTeamsTeamItemRequestBuilder provides operations to manage the joinedTeams property of the microsoft.graph.user entity.
type ItemJoinedTeamsTeamItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemJoinedTeamsTeamItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsTeamItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemJoinedTeamsTeamItemRequestBuilderGetQueryParameters get joinedTeams from users
type ItemJoinedTeamsTeamItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemJoinedTeamsTeamItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsTeamItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemJoinedTeamsTeamItemRequestBuilderGetQueryParameters
}
// ItemJoinedTeamsTeamItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsTeamItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllChannels provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) AllChannels()(*ItemJoinedTeamsItemAllChannelsRequestBuilder) {
    return NewItemJoinedTeamsItemAllChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AllChannelsById provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) AllChannelsById(id string)(*ItemJoinedTeamsItemAllChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewItemJoinedTeamsItemAllChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Channels provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) Channels()(*ItemJoinedTeamsItemChannelsRequestBuilder) {
    return NewItemJoinedTeamsItemChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ChannelsById provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) ChannelsById(id string)(*ItemJoinedTeamsItemChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewItemJoinedTeamsItemChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemJoinedTeamsTeamItemRequestBuilderInternal instantiates a new TeamItemRequestBuilder and sets the default values.
func NewItemJoinedTeamsTeamItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsTeamItemRequestBuilder) {
    m := &ItemJoinedTeamsTeamItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemJoinedTeamsTeamItemRequestBuilder instantiates a new TeamItemRequestBuilder and sets the default values.
func NewItemJoinedTeamsTeamItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsTeamItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedTeamsTeamItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property joinedTeams for users
func (m *ItemJoinedTeamsTeamItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemJoinedTeamsTeamItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get joinedTeams from users
func (m *ItemJoinedTeamsTeamItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedTeamsTeamItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable), nil
}
// Group provides operations to manage the group property of the microsoft.graph.team entity.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) Group()(*ItemJoinedTeamsItemGroupRequestBuilder) {
    return NewItemJoinedTeamsItemGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IncomingChannels provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) IncomingChannels()(*ItemJoinedTeamsItemIncomingChannelsRequestBuilder) {
    return NewItemJoinedTeamsItemIncomingChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IncomingChannelsById provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) IncomingChannelsById(id string)(*ItemJoinedTeamsItemIncomingChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewItemJoinedTeamsItemIncomingChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// InstalledApps provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) InstalledApps()(*ItemJoinedTeamsItemInstalledAppsRequestBuilder) {
    return NewItemJoinedTeamsItemInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InstalledAppsById provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) InstalledAppsById(id string)(*ItemJoinedTeamsItemInstalledAppsTeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return NewItemJoinedTeamsItemInstalledAppsTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Members provides operations to manage the members property of the microsoft.graph.team entity.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) Members()(*ItemJoinedTeamsItemMembersRequestBuilder) {
    return NewItemJoinedTeamsItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MembersById provides operations to manage the members property of the microsoft.graph.team entity.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) MembersById(id string)(*ItemJoinedTeamsItemMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewItemJoinedTeamsItemMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphArchive provides operations to call the archive method.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) MicrosoftGraphArchive()(*ItemJoinedTeamsItemMicrosoftGraphArchiveRequestBuilder) {
    return NewItemJoinedTeamsItemMicrosoftGraphArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphClone provides operations to call the clone method.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) MicrosoftGraphClone()(*ItemJoinedTeamsItemMicrosoftGraphCloneRequestBuilder) {
    return NewItemJoinedTeamsItemMicrosoftGraphCloneRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCompleteMigration provides operations to call the completeMigration method.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) MicrosoftGraphCompleteMigration()(*ItemJoinedTeamsItemMicrosoftGraphCompleteMigrationRequestBuilder) {
    return NewItemJoinedTeamsItemMicrosoftGraphCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSendActivityNotification provides operations to call the sendActivityNotification method.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) MicrosoftGraphSendActivityNotification()(*ItemJoinedTeamsItemMicrosoftGraphSendActivityNotificationRequestBuilder) {
    return NewItemJoinedTeamsItemMicrosoftGraphSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUnarchive provides operations to call the unarchive method.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) MicrosoftGraphUnarchive()(*ItemJoinedTeamsItemMicrosoftGraphUnarchiveRequestBuilder) {
    return NewItemJoinedTeamsItemMicrosoftGraphUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) Operations()(*ItemJoinedTeamsItemOperationsRequestBuilder) {
    return NewItemJoinedTeamsItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) OperationsById(id string)(*ItemJoinedTeamsItemOperationsTeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation%2Did"] = id
    }
    return NewItemJoinedTeamsItemOperationsTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property joinedTeams in users
func (m *ItemJoinedTeamsTeamItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, requestConfiguration *ItemJoinedTeamsTeamItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable), nil
}
// Photo provides operations to manage the photo property of the microsoft.graph.team entity.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) Photo()(*ItemJoinedTeamsItemPhotoRequestBuilder) {
    return NewItemJoinedTeamsItemPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PrimaryChannel provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) PrimaryChannel()(*ItemJoinedTeamsItemPrimaryChannelRequestBuilder) {
    return NewItemJoinedTeamsItemPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Schedule provides operations to manage the schedule property of the microsoft.graph.team entity.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) Schedule()(*ItemJoinedTeamsItemScheduleRequestBuilder) {
    return NewItemJoinedTeamsItemScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Tags provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) Tags()(*ItemJoinedTeamsItemTagsRequestBuilder) {
    return NewItemJoinedTeamsItemTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TagsById provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) TagsById(id string)(*ItemJoinedTeamsItemTagsTeamworkTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkTag%2Did"] = id
    }
    return NewItemJoinedTeamsItemTagsTeamworkTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Template provides operations to manage the template property of the microsoft.graph.team entity.
func (m *ItemJoinedTeamsTeamItemRequestBuilder) Template()(*ItemJoinedTeamsItemTemplateRequestBuilder) {
    return NewItemJoinedTeamsItemTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property joinedTeams for users
func (m *ItemJoinedTeamsTeamItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedTeamsTeamItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get joinedTeams from users
func (m *ItemJoinedTeamsTeamItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedTeamsTeamItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property joinedTeams in users
func (m *ItemJoinedTeamsTeamItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, requestConfiguration *ItemJoinedTeamsTeamItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
