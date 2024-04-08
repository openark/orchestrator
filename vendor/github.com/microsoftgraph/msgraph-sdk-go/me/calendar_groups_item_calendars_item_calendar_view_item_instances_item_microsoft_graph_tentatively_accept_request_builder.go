package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilder provides operations to call the tentativelyAccept method.
type CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilderInternal instantiates a new MicrosoftGraphTentativelyAcceptRequestBuilder and sets the default values.
func NewCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilder) {
    m := &CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/microsoft.graph.tentativelyAccept";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilder instantiates a new MicrosoftGraphTentativelyAcceptRequestBuilder and sets the default values.
func NewCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilderInternal(urlParams, requestAdapter)
}
// Post tentatively accept the specified event in a user calendar. If the event allows proposals for new times, on responding tentative to the event, an invitee can choose to suggest an alternative time by including the **proposedNewTime** parameter. For more information on how to propose a time, and how to receive and accept a new time proposal, see Propose new meeting times.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/event-tentativelyaccept?view=graph-rest-1.0
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilder) Post(ctx context.Context, body CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBodyable, requestConfiguration *CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation tentatively accept the specified event in a user calendar. If the event allows proposals for new times, on responding tentative to the event, an invitee can choose to suggest an alternative time by including the **proposedNewTime** parameter. For more information on how to propose a time, and how to receive and accept a new time proposal, see Propose new meeting times.
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilder) ToPostRequestInformation(ctx context.Context, body CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBodyable, requestConfiguration *CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
