package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilder provides operations to manage the schedulingGroups property of the microsoft.graph.schedule entity.
type JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderGetQueryParameters the logical grouping of users in the schedule (usually by role).
type JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderGetQueryParameters
}
// JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewJoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderInternal instantiates a new SchedulingGroupItemRequestBuilder and sets the default values.
func NewJoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilder) {
    m := &JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedTeams/{team%2Did}/schedule/schedulingGroups/{schedulingGroup%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewJoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilder instantiates a new SchedulingGroupItemRequestBuilder and sets the default values.
func NewJoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewJoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property schedulingGroups for me
func (m *JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the logical grouping of users in the schedule (usually by role).
func (m *JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilder) Get(ctx context.Context, requestConfiguration *JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SchedulingGroupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSchedulingGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SchedulingGroupable), nil
}
// Patch update the navigation property schedulingGroups in me
func (m *JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SchedulingGroupable, requestConfiguration *JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SchedulingGroupable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSchedulingGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SchedulingGroupable), nil
}
// ToDeleteRequestInformation delete navigation property schedulingGroups for me
func (m *JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the logical grouping of users in the schedule (usually by role).
func (m *JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property schedulingGroups in me
func (m *JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SchedulingGroupable, requestConfiguration *JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
