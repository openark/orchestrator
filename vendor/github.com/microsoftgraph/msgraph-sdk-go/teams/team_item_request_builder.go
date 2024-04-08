package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// TeamItemRequestBuilder provides operations to manage the collection of team entities.
type TeamItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamItemRequestBuilderGetQueryParameters retrieve the properties and relationships of the specified team.
type TeamItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamItemRequestBuilderGetQueryParameters
}
// TeamItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllChannels provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *TeamItemRequestBuilder) AllChannels()(*ItemAllChannelsRequestBuilder) {
    return NewItemAllChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AllChannelsById provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *TeamItemRequestBuilder) AllChannelsById(id string)(*ItemAllChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewItemAllChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Channels provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *TeamItemRequestBuilder) Channels()(*ItemChannelsRequestBuilder) {
    return NewItemChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ChannelsById provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *TeamItemRequestBuilder) ChannelsById(id string)(*ItemChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewItemChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewTeamItemRequestBuilderInternal instantiates a new TeamItemRequestBuilder and sets the default values.
func NewTeamItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamItemRequestBuilder) {
    m := &TeamItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewTeamItemRequestBuilder instantiates a new TeamItemRequestBuilder and sets the default values.
func NewTeamItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete entity from teams
func (m *TeamItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the properties and relationships of the specified team.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/team-get?view=graph-rest-1.0
func (m *TeamItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, error) {
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
func (m *TeamItemRequestBuilder) Group()(*ItemGroupRequestBuilder) {
    return NewItemGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IncomingChannels provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *TeamItemRequestBuilder) IncomingChannels()(*ItemIncomingChannelsRequestBuilder) {
    return NewItemIncomingChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IncomingChannelsById provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *TeamItemRequestBuilder) IncomingChannelsById(id string)(*ItemIncomingChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewItemIncomingChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// InstalledApps provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *TeamItemRequestBuilder) InstalledApps()(*ItemInstalledAppsRequestBuilder) {
    return NewItemInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InstalledAppsById provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *TeamItemRequestBuilder) InstalledAppsById(id string)(*ItemInstalledAppsTeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return NewItemInstalledAppsTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Members provides operations to manage the members property of the microsoft.graph.team entity.
func (m *TeamItemRequestBuilder) Members()(*ItemMembersRequestBuilder) {
    return NewItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MembersById provides operations to manage the members property of the microsoft.graph.team entity.
func (m *TeamItemRequestBuilder) MembersById(id string)(*ItemMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewItemMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphArchive provides operations to call the archive method.
func (m *TeamItemRequestBuilder) MicrosoftGraphArchive()(*ItemMicrosoftGraphArchiveRequestBuilder) {
    return NewItemMicrosoftGraphArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphClone provides operations to call the clone method.
func (m *TeamItemRequestBuilder) MicrosoftGraphClone()(*ItemMicrosoftGraphCloneRequestBuilder) {
    return NewItemMicrosoftGraphCloneRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCompleteMigration provides operations to call the completeMigration method.
func (m *TeamItemRequestBuilder) MicrosoftGraphCompleteMigration()(*ItemMicrosoftGraphCompleteMigrationRequestBuilder) {
    return NewItemMicrosoftGraphCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSendActivityNotification provides operations to call the sendActivityNotification method.
func (m *TeamItemRequestBuilder) MicrosoftGraphSendActivityNotification()(*ItemMicrosoftGraphSendActivityNotificationRequestBuilder) {
    return NewItemMicrosoftGraphSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUnarchive provides operations to call the unarchive method.
func (m *TeamItemRequestBuilder) MicrosoftGraphUnarchive()(*ItemMicrosoftGraphUnarchiveRequestBuilder) {
    return NewItemMicrosoftGraphUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *TeamItemRequestBuilder) Operations()(*ItemOperationsRequestBuilder) {
    return NewItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *TeamItemRequestBuilder) OperationsById(id string)(*ItemOperationsTeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation%2Did"] = id
    }
    return NewItemOperationsTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the properties of the specified team.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/team-update?view=graph-rest-1.0
func (m *TeamItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, requestConfiguration *TeamItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, error) {
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
func (m *TeamItemRequestBuilder) Photo()(*ItemPhotoRequestBuilder) {
    return NewItemPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PrimaryChannel provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
func (m *TeamItemRequestBuilder) PrimaryChannel()(*ItemPrimaryChannelRequestBuilder) {
    return NewItemPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Schedule provides operations to manage the schedule property of the microsoft.graph.team entity.
func (m *TeamItemRequestBuilder) Schedule()(*ItemScheduleRequestBuilder) {
    return NewItemScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Tags provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *TeamItemRequestBuilder) Tags()(*ItemTagsRequestBuilder) {
    return NewItemTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TagsById provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *TeamItemRequestBuilder) TagsById(id string)(*ItemTagsTeamworkTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkTag%2Did"] = id
    }
    return NewItemTagsTeamworkTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Template provides operations to manage the template property of the microsoft.graph.team entity.
func (m *TeamItemRequestBuilder) Template()(*ItemTemplateRequestBuilder) {
    return NewItemTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete entity from teams
func (m *TeamItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation retrieve the properties and relationships of the specified team.
func (m *TeamItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of the specified team.
func (m *TeamItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, requestConfiguration *TeamItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
