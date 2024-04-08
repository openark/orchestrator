package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CalendarsItemCalendarViewEventItemRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
type CalendarsItemCalendarViewEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarsItemCalendarViewEventItemRequestBuilderGetQueryParameters the calendar view for the calendar. Navigation property. Read-only.
type CalendarsItemCalendarViewEventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string
}
// CalendarsItemCalendarViewEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarsItemCalendarViewEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CalendarsItemCalendarViewEventItemRequestBuilderGetQueryParameters
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *CalendarsItemCalendarViewEventItemRequestBuilder) Attachments()(*CalendarsItemCalendarViewItemAttachmentsRequestBuilder) {
    return NewCalendarsItemCalendarViewItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *CalendarsItemCalendarViewEventItemRequestBuilder) AttachmentsById(id string)(*CalendarsItemCalendarViewItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewCalendarsItemCalendarViewItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *CalendarsItemCalendarViewEventItemRequestBuilder) Calendar()(*CalendarsItemCalendarViewItemCalendarRequestBuilder) {
    return NewCalendarsItemCalendarViewItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewCalendarsItemCalendarViewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewCalendarsItemCalendarViewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarsItemCalendarViewEventItemRequestBuilder) {
    m := &CalendarsItemCalendarViewEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar%2Did}/calendarView/{event%2Did}{?startDateTime*,endDateTime*,%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCalendarsItemCalendarViewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewCalendarsItemCalendarViewEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarsItemCalendarViewEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarsItemCalendarViewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *CalendarsItemCalendarViewEventItemRequestBuilder) Extensions()(*CalendarsItemCalendarViewItemExtensionsRequestBuilder) {
    return NewCalendarsItemCalendarViewItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *CalendarsItemCalendarViewEventItemRequestBuilder) ExtensionsById(id string)(*CalendarsItemCalendarViewItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewCalendarsItemCalendarViewItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the calendar view for the calendar. Navigation property. Read-only.
func (m *CalendarsItemCalendarViewEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CalendarsItemCalendarViewEventItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, error) {
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
func (m *CalendarsItemCalendarViewEventItemRequestBuilder) Instances()(*CalendarsItemCalendarViewItemInstancesRequestBuilder) {
    return NewCalendarsItemCalendarViewItemInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InstancesById provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *CalendarsItemCalendarViewEventItemRequestBuilder) InstancesById(id string)(*CalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return NewCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphAccept provides operations to call the accept method.
func (m *CalendarsItemCalendarViewEventItemRequestBuilder) MicrosoftGraphAccept()(*CalendarsItemCalendarViewItemMicrosoftGraphAcceptRequestBuilder) {
    return NewCalendarsItemCalendarViewItemMicrosoftGraphAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCancel provides operations to call the cancel method.
func (m *CalendarsItemCalendarViewEventItemRequestBuilder) MicrosoftGraphCancel()(*CalendarsItemCalendarViewItemMicrosoftGraphCancelRequestBuilder) {
    return NewCalendarsItemCalendarViewItemMicrosoftGraphCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDecline provides operations to call the decline method.
func (m *CalendarsItemCalendarViewEventItemRequestBuilder) MicrosoftGraphDecline()(*CalendarsItemCalendarViewItemMicrosoftGraphDeclineRequestBuilder) {
    return NewCalendarsItemCalendarViewItemMicrosoftGraphDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDismissReminder provides operations to call the dismissReminder method.
func (m *CalendarsItemCalendarViewEventItemRequestBuilder) MicrosoftGraphDismissReminder()(*CalendarsItemCalendarViewItemMicrosoftGraphDismissReminderRequestBuilder) {
    return NewCalendarsItemCalendarViewItemMicrosoftGraphDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphForward provides operations to call the forward method.
func (m *CalendarsItemCalendarViewEventItemRequestBuilder) MicrosoftGraphForward()(*CalendarsItemCalendarViewItemMicrosoftGraphForwardRequestBuilder) {
    return NewCalendarsItemCalendarViewItemMicrosoftGraphForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSnoozeReminder provides operations to call the snoozeReminder method.
func (m *CalendarsItemCalendarViewEventItemRequestBuilder) MicrosoftGraphSnoozeReminder()(*CalendarsItemCalendarViewItemMicrosoftGraphSnoozeReminderRequestBuilder) {
    return NewCalendarsItemCalendarViewItemMicrosoftGraphSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTentativelyAccept provides operations to call the tentativelyAccept method.
func (m *CalendarsItemCalendarViewEventItemRequestBuilder) MicrosoftGraphTentativelyAccept()(*CalendarsItemCalendarViewItemMicrosoftGraphTentativelyAcceptRequestBuilder) {
    return NewCalendarsItemCalendarViewItemMicrosoftGraphTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarsItemCalendarViewEventItemRequestBuilder) MultiValueExtendedProperties()(*CalendarsItemCalendarViewItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewCalendarsItemCalendarViewItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarsItemCalendarViewEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*CalendarsItemCalendarViewItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewCalendarsItemCalendarViewItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarsItemCalendarViewEventItemRequestBuilder) SingleValueExtendedProperties()(*CalendarsItemCalendarViewItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewCalendarsItemCalendarViewItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarsItemCalendarViewEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*CalendarsItemCalendarViewItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewCalendarsItemCalendarViewItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation the calendar view for the calendar. Navigation property. Read-only.
func (m *CalendarsItemCalendarViewEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CalendarsItemCalendarViewEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
