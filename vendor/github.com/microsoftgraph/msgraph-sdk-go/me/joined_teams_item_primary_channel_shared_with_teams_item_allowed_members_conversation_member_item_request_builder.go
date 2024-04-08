package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// JoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder provides operations to manage the allowedMembers property of the microsoft.graph.sharedWithChannelTeamInfo entity.
type JoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// JoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderGetQueryParameters a collection of team members who have access to the shared channel.
type JoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// JoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type JoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *JoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderGetQueryParameters
}
// NewJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderInternal instantiates a new ConversationMemberItemRequestBuilder and sets the default values.
func NewJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*JoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder) {
    m := &JoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedTeams/{team%2Did}/primaryChannel/sharedWithTeams/{sharedWithChannelTeamInfo%2Did}/allowedMembers/{conversationMember%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder instantiates a new ConversationMemberItemRequestBuilder and sets the default values.
func NewJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*JoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get a collection of team members who have access to the shared channel.
func (m *JoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *JoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConversationMemberable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateConversationMemberFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConversationMemberable), nil
}
// ToGetRequestInformation a collection of team members who have access to the shared channel.
func (m *JoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *JoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
