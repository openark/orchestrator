package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// JoinedTeamsTeamItemRequestBuilder provides operations to manage the joinedTeams property of the microsoft.graph.user entity.
type JoinedTeamsTeamItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// JoinedTeamsTeamItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type JoinedTeamsTeamItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// JoinedTeamsTeamItemRequestBuilderGetQueryParameters get joinedTeams from me
type JoinedTeamsTeamItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// JoinedTeamsTeamItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type JoinedTeamsTeamItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *JoinedTeamsTeamItemRequestBuilderGetQueryParameters
}
// JoinedTeamsTeamItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type JoinedTeamsTeamItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllChannels provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *JoinedTeamsTeamItemRequestBuilder) AllChannels()(*JoinedTeamsItemAllChannelsRequestBuilder) {
    return NewJoinedTeamsItemAllChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AllChannelsById provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *JoinedTeamsTeamItemRequestBuilder) AllChannelsById(id string)(*JoinedTeamsItemAllChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewJoinedTeamsItemAllChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Channels provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *JoinedTeamsTeamItemRequestBuilder) Channels()(*JoinedTeamsItemChannelsRequestBuilder) {
    return NewJoinedTeamsItemChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ChannelsById provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *JoinedTeamsTeamItemRequestBuilder) ChannelsById(id string)(*JoinedTeamsItemChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewJoinedTeamsItemChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewJoinedTeamsTeamItemRequestBuilderInternal instantiates a new TeamItemRequestBuilder and sets the default values.
func NewJoinedTeamsTeamItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*JoinedTeamsTeamItemRequestBuilder) {
    m := &JoinedTeamsTeamItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedTeams/{team%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewJoinedTeamsTeamItemRequestBuilder instantiates a new TeamItemRequestBuilder and sets the default values.
func NewJoinedTeamsTeamItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*JoinedTeamsTeamItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewJoinedTeamsTeamItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property joinedTeams for me
func (m *JoinedTeamsTeamItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *JoinedTeamsTeamItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get joinedTeams from me
func (m *JoinedTeamsTeamItemRequestBuilder) Get(ctx context.Context, requestConfiguration *JoinedTeamsTeamItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, error) {
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
func (m *JoinedTeamsTeamItemRequestBuilder) Group()(*JoinedTeamsItemGroupRequestBuilder) {
    return NewJoinedTeamsItemGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IncomingChannels provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *JoinedTeamsTeamItemRequestBuilder) IncomingChannels()(*JoinedTeamsItemIncomingChannelsRequestBuilder) {
    return NewJoinedTeamsItemIncomingChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IncomingChannelsById provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *JoinedTeamsTeamItemRequestBuilder) IncomingChannelsById(id string)(*JoinedTeamsItemIncomingChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewJoinedTeamsItemIncomingChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// InstalledApps provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *JoinedTeamsTeamItemRequestBuilder) InstalledApps()(*JoinedTeamsItemInstalledAppsRequestBuilder) {
    return NewJoinedTeamsItemInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InstalledAppsById provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *JoinedTeamsTeamItemRequestBuilder) InstalledAppsById(id string)(*JoinedTeamsItemInstalledAppsTeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return NewJoinedTeamsItemInstalledAppsTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Members provides operations to manage the members property of the microsoft.graph.team entity.
func (m *JoinedTeamsTeamItemRequestBuilder) Members()(*JoinedTeamsItemMembersRequestBuilder) {
    return NewJoinedTeamsItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MembersById provides operations to manage the members property of the microsoft.graph.team entity.
func (m *JoinedTeamsTeamItemRequestBuilder) MembersById(id string)(*JoinedTeamsItemMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewJoinedTeamsItemMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphArchive provides operations to call the archive method.
func (m *JoinedTeamsTeamItemRequestBuilder) MicrosoftGraphArchive()(*JoinedTeamsItemMicrosoftGraphArchiveRequestBuilder) {
    return NewJoinedTeamsItemMicrosoftGraphArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphClone provides operations to call the clone method.
func (m *JoinedTeamsTeamItemRequestBuilder) MicrosoftGraphClone()(*JoinedTeamsItemMicrosoftGraphCloneRequestBuilder) {
    return NewJoinedTeamsItemMicrosoftGraphCloneRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCompleteMigration provides operations to call the completeMigration method.
func (m *JoinedTeamsTeamItemRequestBuilder) MicrosoftGraphCompleteMigration()(*JoinedTeamsItemMicrosoftGraphCompleteMigrationRequestBuilder) {
    return NewJoinedTeamsItemMicrosoftGraphCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSendActivityNotification provides operations to call the sendActivityNotification method.
func (m *JoinedTeamsTeamItemRequestBuilder) MicrosoftGraphSendActivityNotification()(*JoinedTeamsItemMicrosoftGraphSendActivityNotificationRequestBuilder) {
    return NewJoinedTeamsItemMicrosoftGraphSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUnarchive provides operations to call the unarchive method.
func (m *JoinedTeamsTeamItemRequestBuilder) MicrosoftGraphUnarchive()(*JoinedTeamsItemMicrosoftGraphUnarchiveRequestBuilder) {
    return NewJoinedTeamsItemMicrosoftGraphUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *JoinedTeamsTeamItemRequestBuilder) Operations()(*JoinedTeamsItemOperationsRequestBuilder) {
    return NewJoinedTeamsItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *JoinedTeamsTeamItemRequestBuilder) OperationsById(id string)(*JoinedTeamsItemOperationsTeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation%2Did"] = id
    }
    return NewJoinedTeamsItemOperationsTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property joinedTeams in me
func (m *JoinedTeamsTeamItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, requestConfiguration *JoinedTeamsTeamItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, error) {
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
func (m *JoinedTeamsTeamItemRequestBuilder) Photo()(*JoinedTeamsItemPhotoRequestBuilder) {
    return NewJoinedTeamsItemPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PrimaryChannel provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
func (m *JoinedTeamsTeamItemRequestBuilder) PrimaryChannel()(*JoinedTeamsItemPrimaryChannelRequestBuilder) {
    return NewJoinedTeamsItemPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Schedule provides operations to manage the schedule property of the microsoft.graph.team entity.
func (m *JoinedTeamsTeamItemRequestBuilder) Schedule()(*JoinedTeamsItemScheduleRequestBuilder) {
    return NewJoinedTeamsItemScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Tags provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *JoinedTeamsTeamItemRequestBuilder) Tags()(*JoinedTeamsItemTagsRequestBuilder) {
    return NewJoinedTeamsItemTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TagsById provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *JoinedTeamsTeamItemRequestBuilder) TagsById(id string)(*JoinedTeamsItemTagsTeamworkTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkTag%2Did"] = id
    }
    return NewJoinedTeamsItemTagsTeamworkTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Template provides operations to manage the template property of the microsoft.graph.team entity.
func (m *JoinedTeamsTeamItemRequestBuilder) Template()(*JoinedTeamsItemTemplateRequestBuilder) {
    return NewJoinedTeamsItemTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property joinedTeams for me
func (m *JoinedTeamsTeamItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *JoinedTeamsTeamItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get joinedTeams from me
func (m *JoinedTeamsTeamItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *JoinedTeamsTeamItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property joinedTeams in me
func (m *JoinedTeamsTeamItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, requestConfiguration *JoinedTeamsTeamItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
