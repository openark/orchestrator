package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDismissReminderRequestBuilder provides operations to call the dismissReminder method.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDismissReminderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDismissReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDismissReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDismissReminderRequestBuilderInternal instantiates a new MicrosoftGraphDismissReminderRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDismissReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDismissReminderRequestBuilder) {
    m := &ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDismissReminderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/microsoft.graph.dismissReminder";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDismissReminderRequestBuilder instantiates a new MicrosoftGraphDismissReminderRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDismissReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDismissReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDismissReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post dismiss a reminder that has been triggered for an event in a user calendar.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/event-dismissreminder?view=graph-rest-1.0
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDismissReminderRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDismissReminderRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation dismiss a reminder that has been triggered for an event in a user calendar.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDismissReminderRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDismissReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
