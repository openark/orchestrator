package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
type CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderGetQueryParameters the calendar view for the calendar. Navigation property. Read-only.
type CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderGetQueryParameters
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Attachments()(*CalendarGroupsItemCalendarsItemCalendarViewItemAttachmentsRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) AttachmentsById(id string)(*CalendarGroupsItemCalendarsItemCalendarViewItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Calendar()(*CalendarGroupsItemCalendarsItemCalendarViewItemCalendarRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) {
    m := &CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Extensions()(*CalendarGroupsItemCalendarsItemCalendarViewItemExtensionsRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) ExtensionsById(id string)(*CalendarGroupsItemCalendarsItemCalendarViewItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the calendar view for the calendar. Navigation property. Read-only.
func (m *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable), nil
}
// Instances provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Instances()(*CalendarGroupsItemCalendarsItemCalendarViewItemInstancesRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InstancesById provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) InstancesById(id string)(*CalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphAccept provides operations to call the accept method.
func (m *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) MicrosoftGraphAccept()(*CalendarGroupsItemCalendarsItemCalendarViewItemMicrosoftGraphAcceptRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemMicrosoftGraphAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCancel provides operations to call the cancel method.
func (m *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) MicrosoftGraphCancel()(*CalendarGroupsItemCalendarsItemCalendarViewItemMicrosoftGraphCancelRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemMicrosoftGraphCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDecline provides operations to call the decline method.
func (m *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) MicrosoftGraphDecline()(*CalendarGroupsItemCalendarsItemCalendarViewItemMicrosoftGraphDeclineRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemMicrosoftGraphDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDismissReminder provides operations to call the dismissReminder method.
func (m *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) MicrosoftGraphDismissReminder()(*CalendarGroupsItemCalendarsItemCalendarViewItemMicrosoftGraphDismissReminderRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemMicrosoftGraphDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphForward provides operations to call the forward method.
func (m *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) MicrosoftGraphForward()(*CalendarGroupsItemCalendarsItemCalendarViewItemMicrosoftGraphForwardRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemMicrosoftGraphForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSnoozeReminder provides operations to call the snoozeReminder method.
func (m *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) MicrosoftGraphSnoozeReminder()(*CalendarGroupsItemCalendarsItemCalendarViewItemMicrosoftGraphSnoozeReminderRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemMicrosoftGraphSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTentativelyAccept provides operations to call the tentativelyAccept method.
func (m *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) MicrosoftGraphTentativelyAccept()(*CalendarGroupsItemCalendarsItemCalendarViewItemMicrosoftGraphTentativelyAcceptRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemMicrosoftGraphTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) MultiValueExtendedProperties()(*CalendarGroupsItemCalendarsItemCalendarViewItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*CalendarGroupsItemCalendarsItemCalendarViewItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) SingleValueExtendedProperties()(*CalendarGroupsItemCalendarsItemCalendarViewItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*CalendarGroupsItemCalendarsItemCalendarViewItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation the calendar view for the calendar. Navigation property. Read-only.
func (m *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
