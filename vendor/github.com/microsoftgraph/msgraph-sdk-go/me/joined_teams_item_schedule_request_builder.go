package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// JoinedTeamsItemScheduleRequestBuilder provides operations to manage the schedule property of the microsoft.graph.team entity.
type JoinedTeamsItemScheduleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// JoinedTeamsItemScheduleRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type JoinedTeamsItemScheduleRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// JoinedTeamsItemScheduleRequestBuilderGetQueryParameters retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the `provisionStatus` property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the `provisionStatusCode` property. Clients can also inspect the configuration of the schedule.
type JoinedTeamsItemScheduleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// JoinedTeamsItemScheduleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type JoinedTeamsItemScheduleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *JoinedTeamsItemScheduleRequestBuilderGetQueryParameters
}
// JoinedTeamsItemScheduleRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type JoinedTeamsItemScheduleRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewJoinedTeamsItemScheduleRequestBuilderInternal instantiates a new ScheduleRequestBuilder and sets the default values.
func NewJoinedTeamsItemScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*JoinedTeamsItemScheduleRequestBuilder) {
    m := &JoinedTeamsItemScheduleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedTeams/{team%2Did}/schedule{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewJoinedTeamsItemScheduleRequestBuilder instantiates a new ScheduleRequestBuilder and sets the default values.
func NewJoinedTeamsItemScheduleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*JoinedTeamsItemScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewJoinedTeamsItemScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property schedule for me
func (m *JoinedTeamsItemScheduleRequestBuilder) Delete(ctx context.Context, requestConfiguration *JoinedTeamsItemScheduleRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the `provisionStatus` property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the `provisionStatusCode` property. Clients can also inspect the configuration of the schedule.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/schedule-get?view=graph-rest-1.0
func (m *JoinedTeamsItemScheduleRequestBuilder) Get(ctx context.Context, requestConfiguration *JoinedTeamsItemScheduleRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable), nil
}
// MicrosoftGraphShare provides operations to call the share method.
func (m *JoinedTeamsItemScheduleRequestBuilder) MicrosoftGraphShare()(*JoinedTeamsItemScheduleMicrosoftGraphShareRequestBuilder) {
    return NewJoinedTeamsItemScheduleMicrosoftGraphShareRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OfferShiftRequests provides operations to manage the offerShiftRequests property of the microsoft.graph.schedule entity.
func (m *JoinedTeamsItemScheduleRequestBuilder) OfferShiftRequests()(*JoinedTeamsItemScheduleOfferShiftRequestsRequestBuilder) {
    return NewJoinedTeamsItemScheduleOfferShiftRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OfferShiftRequestsById provides operations to manage the offerShiftRequests property of the microsoft.graph.schedule entity.
func (m *JoinedTeamsItemScheduleRequestBuilder) OfferShiftRequestsById(id string)(*JoinedTeamsItemScheduleOfferShiftRequestsOfferShiftRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["offerShiftRequest%2Did"] = id
    }
    return NewJoinedTeamsItemScheduleOfferShiftRequestsOfferShiftRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// OpenShiftChangeRequests provides operations to manage the openShiftChangeRequests property of the microsoft.graph.schedule entity.
func (m *JoinedTeamsItemScheduleRequestBuilder) OpenShiftChangeRequests()(*JoinedTeamsItemScheduleOpenShiftChangeRequestsRequestBuilder) {
    return NewJoinedTeamsItemScheduleOpenShiftChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OpenShiftChangeRequestsById provides operations to manage the openShiftChangeRequests property of the microsoft.graph.schedule entity.
func (m *JoinedTeamsItemScheduleRequestBuilder) OpenShiftChangeRequestsById(id string)(*JoinedTeamsItemScheduleOpenShiftChangeRequestsOpenShiftChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShiftChangeRequest%2Did"] = id
    }
    return NewJoinedTeamsItemScheduleOpenShiftChangeRequestsOpenShiftChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// OpenShifts provides operations to manage the openShifts property of the microsoft.graph.schedule entity.
func (m *JoinedTeamsItemScheduleRequestBuilder) OpenShifts()(*JoinedTeamsItemScheduleOpenShiftsRequestBuilder) {
    return NewJoinedTeamsItemScheduleOpenShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OpenShiftsById provides operations to manage the openShifts property of the microsoft.graph.schedule entity.
func (m *JoinedTeamsItemScheduleRequestBuilder) OpenShiftsById(id string)(*JoinedTeamsItemScheduleOpenShiftsOpenShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShift%2Did"] = id
    }
    return NewJoinedTeamsItemScheduleOpenShiftsOpenShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Put update the navigation property schedule in me
func (m *JoinedTeamsItemScheduleRequestBuilder) Put(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, requestConfiguration *JoinedTeamsItemScheduleRequestBuilderPutRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable), nil
}
// SchedulingGroups provides operations to manage the schedulingGroups property of the microsoft.graph.schedule entity.
func (m *JoinedTeamsItemScheduleRequestBuilder) SchedulingGroups()(*JoinedTeamsItemScheduleSchedulingGroupsRequestBuilder) {
    return NewJoinedTeamsItemScheduleSchedulingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SchedulingGroupsById provides operations to manage the schedulingGroups property of the microsoft.graph.schedule entity.
func (m *JoinedTeamsItemScheduleRequestBuilder) SchedulingGroupsById(id string)(*JoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["schedulingGroup%2Did"] = id
    }
    return NewJoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Shifts provides operations to manage the shifts property of the microsoft.graph.schedule entity.
func (m *JoinedTeamsItemScheduleRequestBuilder) Shifts()(*JoinedTeamsItemScheduleShiftsRequestBuilder) {
    return NewJoinedTeamsItemScheduleShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ShiftsById provides operations to manage the shifts property of the microsoft.graph.schedule entity.
func (m *JoinedTeamsItemScheduleRequestBuilder) ShiftsById(id string)(*JoinedTeamsItemScheduleShiftsShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shift%2Did"] = id
    }
    return NewJoinedTeamsItemScheduleShiftsShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SwapShiftsChangeRequests provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
func (m *JoinedTeamsItemScheduleRequestBuilder) SwapShiftsChangeRequests()(*JoinedTeamsItemScheduleSwapShiftsChangeRequestsRequestBuilder) {
    return NewJoinedTeamsItemScheduleSwapShiftsChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SwapShiftsChangeRequestsById provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
func (m *JoinedTeamsItemScheduleRequestBuilder) SwapShiftsChangeRequestsById(id string)(*JoinedTeamsItemScheduleSwapShiftsChangeRequestsSwapShiftsChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["swapShiftsChangeRequest%2Did"] = id
    }
    return NewJoinedTeamsItemScheduleSwapShiftsChangeRequestsSwapShiftsChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// TimeOffReasons provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
func (m *JoinedTeamsItemScheduleRequestBuilder) TimeOffReasons()(*JoinedTeamsItemScheduleTimeOffReasonsRequestBuilder) {
    return NewJoinedTeamsItemScheduleTimeOffReasonsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TimeOffReasonsById provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
func (m *JoinedTeamsItemScheduleRequestBuilder) TimeOffReasonsById(id string)(*JoinedTeamsItemScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffReason%2Did"] = id
    }
    return NewJoinedTeamsItemScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// TimeOffRequests provides operations to manage the timeOffRequests property of the microsoft.graph.schedule entity.
func (m *JoinedTeamsItemScheduleRequestBuilder) TimeOffRequests()(*JoinedTeamsItemScheduleTimeOffRequestsRequestBuilder) {
    return NewJoinedTeamsItemScheduleTimeOffRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TimeOffRequestsById provides operations to manage the timeOffRequests property of the microsoft.graph.schedule entity.
func (m *JoinedTeamsItemScheduleRequestBuilder) TimeOffRequestsById(id string)(*JoinedTeamsItemScheduleTimeOffRequestsTimeOffRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffRequest%2Did"] = id
    }
    return NewJoinedTeamsItemScheduleTimeOffRequestsTimeOffRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// TimesOff provides operations to manage the timesOff property of the microsoft.graph.schedule entity.
func (m *JoinedTeamsItemScheduleRequestBuilder) TimesOff()(*JoinedTeamsItemScheduleTimesOffRequestBuilder) {
    return NewJoinedTeamsItemScheduleTimesOffRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TimesOffById provides operations to manage the timesOff property of the microsoft.graph.schedule entity.
func (m *JoinedTeamsItemScheduleRequestBuilder) TimesOffById(id string)(*JoinedTeamsItemScheduleTimesOffTimeOffItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOff%2Did"] = id
    }
    return NewJoinedTeamsItemScheduleTimesOffTimeOffItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property schedule for me
func (m *JoinedTeamsItemScheduleRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *JoinedTeamsItemScheduleRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the `provisionStatus` property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the `provisionStatusCode` property. Clients can also inspect the configuration of the schedule.
func (m *JoinedTeamsItemScheduleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *JoinedTeamsItemScheduleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation update the navigation property schedule in me
func (m *JoinedTeamsItemScheduleRequestBuilder) ToPutRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, requestConfiguration *JoinedTeamsItemScheduleRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
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
