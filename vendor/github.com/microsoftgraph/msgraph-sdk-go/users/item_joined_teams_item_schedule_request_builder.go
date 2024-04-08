package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedTeamsItemScheduleRequestBuilder provides operations to manage the schedule property of the microsoft.graph.team entity.
type ItemJoinedTeamsItemScheduleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemJoinedTeamsItemScheduleRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemScheduleRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemJoinedTeamsItemScheduleRequestBuilderGetQueryParameters retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the `provisionStatus` property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the `provisionStatusCode` property. Clients can also inspect the configuration of the schedule.
type ItemJoinedTeamsItemScheduleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemJoinedTeamsItemScheduleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemScheduleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemJoinedTeamsItemScheduleRequestBuilderGetQueryParameters
}
// ItemJoinedTeamsItemScheduleRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemScheduleRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemJoinedTeamsItemScheduleRequestBuilderInternal instantiates a new ScheduleRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemScheduleRequestBuilder) {
    m := &ItemJoinedTeamsItemScheduleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/schedule{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemJoinedTeamsItemScheduleRequestBuilder instantiates a new ScheduleRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemScheduleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedTeamsItemScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property schedule for users
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemScheduleRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemScheduleRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, error) {
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
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) MicrosoftGraphShare()(*ItemJoinedTeamsItemScheduleMicrosoftGraphShareRequestBuilder) {
    return NewItemJoinedTeamsItemScheduleMicrosoftGraphShareRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OfferShiftRequests provides operations to manage the offerShiftRequests property of the microsoft.graph.schedule entity.
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) OfferShiftRequests()(*ItemJoinedTeamsItemScheduleOfferShiftRequestsRequestBuilder) {
    return NewItemJoinedTeamsItemScheduleOfferShiftRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OfferShiftRequestsById provides operations to manage the offerShiftRequests property of the microsoft.graph.schedule entity.
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) OfferShiftRequestsById(id string)(*ItemJoinedTeamsItemScheduleOfferShiftRequestsOfferShiftRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["offerShiftRequest%2Did"] = id
    }
    return NewItemJoinedTeamsItemScheduleOfferShiftRequestsOfferShiftRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// OpenShiftChangeRequests provides operations to manage the openShiftChangeRequests property of the microsoft.graph.schedule entity.
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) OpenShiftChangeRequests()(*ItemJoinedTeamsItemScheduleOpenShiftChangeRequestsRequestBuilder) {
    return NewItemJoinedTeamsItemScheduleOpenShiftChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OpenShiftChangeRequestsById provides operations to manage the openShiftChangeRequests property of the microsoft.graph.schedule entity.
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) OpenShiftChangeRequestsById(id string)(*ItemJoinedTeamsItemScheduleOpenShiftChangeRequestsOpenShiftChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShiftChangeRequest%2Did"] = id
    }
    return NewItemJoinedTeamsItemScheduleOpenShiftChangeRequestsOpenShiftChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// OpenShifts provides operations to manage the openShifts property of the microsoft.graph.schedule entity.
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) OpenShifts()(*ItemJoinedTeamsItemScheduleOpenShiftsRequestBuilder) {
    return NewItemJoinedTeamsItemScheduleOpenShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OpenShiftsById provides operations to manage the openShifts property of the microsoft.graph.schedule entity.
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) OpenShiftsById(id string)(*ItemJoinedTeamsItemScheduleOpenShiftsOpenShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShift%2Did"] = id
    }
    return NewItemJoinedTeamsItemScheduleOpenShiftsOpenShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Put update the navigation property schedule in users
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) Put(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, requestConfiguration *ItemJoinedTeamsItemScheduleRequestBuilderPutRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, error) {
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
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) SchedulingGroups()(*ItemJoinedTeamsItemScheduleSchedulingGroupsRequestBuilder) {
    return NewItemJoinedTeamsItemScheduleSchedulingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SchedulingGroupsById provides operations to manage the schedulingGroups property of the microsoft.graph.schedule entity.
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) SchedulingGroupsById(id string)(*ItemJoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["schedulingGroup%2Did"] = id
    }
    return NewItemJoinedTeamsItemScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Shifts provides operations to manage the shifts property of the microsoft.graph.schedule entity.
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) Shifts()(*ItemJoinedTeamsItemScheduleShiftsRequestBuilder) {
    return NewItemJoinedTeamsItemScheduleShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ShiftsById provides operations to manage the shifts property of the microsoft.graph.schedule entity.
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) ShiftsById(id string)(*ItemJoinedTeamsItemScheduleShiftsShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shift%2Did"] = id
    }
    return NewItemJoinedTeamsItemScheduleShiftsShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SwapShiftsChangeRequests provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) SwapShiftsChangeRequests()(*ItemJoinedTeamsItemScheduleSwapShiftsChangeRequestsRequestBuilder) {
    return NewItemJoinedTeamsItemScheduleSwapShiftsChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SwapShiftsChangeRequestsById provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) SwapShiftsChangeRequestsById(id string)(*ItemJoinedTeamsItemScheduleSwapShiftsChangeRequestsSwapShiftsChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["swapShiftsChangeRequest%2Did"] = id
    }
    return NewItemJoinedTeamsItemScheduleSwapShiftsChangeRequestsSwapShiftsChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// TimeOffReasons provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) TimeOffReasons()(*ItemJoinedTeamsItemScheduleTimeOffReasonsRequestBuilder) {
    return NewItemJoinedTeamsItemScheduleTimeOffReasonsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TimeOffReasonsById provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) TimeOffReasonsById(id string)(*ItemJoinedTeamsItemScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffReason%2Did"] = id
    }
    return NewItemJoinedTeamsItemScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// TimeOffRequests provides operations to manage the timeOffRequests property of the microsoft.graph.schedule entity.
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) TimeOffRequests()(*ItemJoinedTeamsItemScheduleTimeOffRequestsRequestBuilder) {
    return NewItemJoinedTeamsItemScheduleTimeOffRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TimeOffRequestsById provides operations to manage the timeOffRequests property of the microsoft.graph.schedule entity.
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) TimeOffRequestsById(id string)(*ItemJoinedTeamsItemScheduleTimeOffRequestsTimeOffRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffRequest%2Did"] = id
    }
    return NewItemJoinedTeamsItemScheduleTimeOffRequestsTimeOffRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// TimesOff provides operations to manage the timesOff property of the microsoft.graph.schedule entity.
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) TimesOff()(*ItemJoinedTeamsItemScheduleTimesOffRequestBuilder) {
    return NewItemJoinedTeamsItemScheduleTimesOffRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TimesOffById provides operations to manage the timesOff property of the microsoft.graph.schedule entity.
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) TimesOffById(id string)(*ItemJoinedTeamsItemScheduleTimesOffTimeOffItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOff%2Did"] = id
    }
    return NewItemJoinedTeamsItemScheduleTimesOffTimeOffItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property schedule for users
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemScheduleRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemScheduleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation update the navigation property schedule in users
func (m *ItemJoinedTeamsItemScheduleRequestBuilder) ToPutRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, requestConfiguration *ItemJoinedTeamsItemScheduleRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
