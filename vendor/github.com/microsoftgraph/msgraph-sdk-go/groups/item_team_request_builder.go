package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTeamRequestBuilder provides operations to manage the team property of the microsoft.graph.group entity.
type ItemTeamRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTeamRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTeamRequestBuilderGetQueryParameters the team associated with this group.
type ItemTeamRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTeamRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamRequestBuilderGetQueryParameters
}
// ItemTeamRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllChannels provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *ItemTeamRequestBuilder) AllChannels()(*ItemTeamAllChannelsRequestBuilder) {
    return NewItemTeamAllChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AllChannelsById provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *ItemTeamRequestBuilder) AllChannelsById(id string)(*ItemTeamAllChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewItemTeamAllChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Channels provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *ItemTeamRequestBuilder) Channels()(*ItemTeamChannelsRequestBuilder) {
    return NewItemTeamChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ChannelsById provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *ItemTeamRequestBuilder) ChannelsById(id string)(*ItemTeamChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewItemTeamChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemTeamRequestBuilderInternal instantiates a new TeamRequestBuilder and sets the default values.
func NewItemTeamRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamRequestBuilder) {
    m := &ItemTeamRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/team{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemTeamRequestBuilder instantiates a new TeamRequestBuilder and sets the default values.
func NewItemTeamRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property team for groups
func (m *ItemTeamRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTeamRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the team associated with this group.
func (m *ItemTeamRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, error) {
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
func (m *ItemTeamRequestBuilder) Group()(*ItemTeamGroupRequestBuilder) {
    return NewItemTeamGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IncomingChannels provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *ItemTeamRequestBuilder) IncomingChannels()(*ItemTeamIncomingChannelsRequestBuilder) {
    return NewItemTeamIncomingChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IncomingChannelsById provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *ItemTeamRequestBuilder) IncomingChannelsById(id string)(*ItemTeamIncomingChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewItemTeamIncomingChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// InstalledApps provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *ItemTeamRequestBuilder) InstalledApps()(*ItemTeamInstalledAppsRequestBuilder) {
    return NewItemTeamInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InstalledAppsById provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *ItemTeamRequestBuilder) InstalledAppsById(id string)(*ItemTeamInstalledAppsTeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return NewItemTeamInstalledAppsTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Members provides operations to manage the members property of the microsoft.graph.team entity.
func (m *ItemTeamRequestBuilder) Members()(*ItemTeamMembersRequestBuilder) {
    return NewItemTeamMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MembersById provides operations to manage the members property of the microsoft.graph.team entity.
func (m *ItemTeamRequestBuilder) MembersById(id string)(*ItemTeamMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewItemTeamMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphArchive provides operations to call the archive method.
func (m *ItemTeamRequestBuilder) MicrosoftGraphArchive()(*ItemTeamMicrosoftGraphArchiveRequestBuilder) {
    return NewItemTeamMicrosoftGraphArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphClone provides operations to call the clone method.
func (m *ItemTeamRequestBuilder) MicrosoftGraphClone()(*ItemTeamMicrosoftGraphCloneRequestBuilder) {
    return NewItemTeamMicrosoftGraphCloneRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCompleteMigration provides operations to call the completeMigration method.
func (m *ItemTeamRequestBuilder) MicrosoftGraphCompleteMigration()(*ItemTeamMicrosoftGraphCompleteMigrationRequestBuilder) {
    return NewItemTeamMicrosoftGraphCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSendActivityNotification provides operations to call the sendActivityNotification method.
func (m *ItemTeamRequestBuilder) MicrosoftGraphSendActivityNotification()(*ItemTeamMicrosoftGraphSendActivityNotificationRequestBuilder) {
    return NewItemTeamMicrosoftGraphSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUnarchive provides operations to call the unarchive method.
func (m *ItemTeamRequestBuilder) MicrosoftGraphUnarchive()(*ItemTeamMicrosoftGraphUnarchiveRequestBuilder) {
    return NewItemTeamMicrosoftGraphUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *ItemTeamRequestBuilder) Operations()(*ItemTeamOperationsRequestBuilder) {
    return NewItemTeamOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *ItemTeamRequestBuilder) OperationsById(id string)(*ItemTeamOperationsTeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation%2Did"] = id
    }
    return NewItemTeamOperationsTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch create a new team under a group. In order to create a team, the group must have a least one owner. If the group was created less than 15 minutes ago, it's possible for the Create team call to fail with a 404 error code due to replication delays. The recommended pattern is to retry the Create team call three times, with a 10 second delay between calls.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/team-put-teams?view=graph-rest-1.0
func (m *ItemTeamRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, requestConfiguration *ItemTeamRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, error) {
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
func (m *ItemTeamRequestBuilder) Photo()(*ItemTeamPhotoRequestBuilder) {
    return NewItemTeamPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PrimaryChannel provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
func (m *ItemTeamRequestBuilder) PrimaryChannel()(*ItemTeamPrimaryChannelRequestBuilder) {
    return NewItemTeamPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Schedule provides operations to manage the schedule property of the microsoft.graph.team entity.
func (m *ItemTeamRequestBuilder) Schedule()(*ItemTeamScheduleRequestBuilder) {
    return NewItemTeamScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Tags provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *ItemTeamRequestBuilder) Tags()(*ItemTeamTagsRequestBuilder) {
    return NewItemTeamTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TagsById provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *ItemTeamRequestBuilder) TagsById(id string)(*ItemTeamTagsTeamworkTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkTag%2Did"] = id
    }
    return NewItemTeamTagsTeamworkTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Template provides operations to manage the template property of the microsoft.graph.team entity.
func (m *ItemTeamRequestBuilder) Template()(*ItemTeamTemplateRequestBuilder) {
    return NewItemTeamTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property team for groups
func (m *ItemTeamRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTeamRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the team associated with this group.
func (m *ItemTeamRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation create a new team under a group. In order to create a team, the group must have a least one owner. If the group was created less than 15 minutes ago, it's possible for the Create team call to fail with a 404 error code due to replication delays. The recommended pattern is to retry the Create team call three times, with a 10 second delay between calls.
func (m *ItemTeamRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, requestConfiguration *ItemTeamRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
