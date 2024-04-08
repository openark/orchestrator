package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendarsItemCalendarViewEventItemRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
type ItemCalendarsItemCalendarViewEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemCalendarsItemCalendarViewEventItemRequestBuilderGetQueryParameters the calendar view for the calendar. Navigation property. Read-only.
type ItemCalendarsItemCalendarViewEventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string
}
// ItemCalendarsItemCalendarViewEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarsItemCalendarViewEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarsItemCalendarViewEventItemRequestBuilderGetQueryParameters
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewEventItemRequestBuilder) Attachments()(*ItemCalendarsItemCalendarViewItemAttachmentsRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewEventItemRequestBuilder) AttachmentsById(id string)(*ItemCalendarsItemCalendarViewItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewItemCalendarsItemCalendarViewItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewEventItemRequestBuilder) Calendar()(*ItemCalendarsItemCalendarViewItemCalendarRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemCalendarsItemCalendarViewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewItemCalendarsItemCalendarViewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemCalendarViewEventItemRequestBuilder) {
    m := &ItemCalendarsItemCalendarViewEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}{?startDateTime*,endDateTime*,%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemCalendarsItemCalendarViewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewItemCalendarsItemCalendarViewEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemCalendarViewEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarsItemCalendarViewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewEventItemRequestBuilder) Extensions()(*ItemCalendarsItemCalendarViewItemExtensionsRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewEventItemRequestBuilder) ExtensionsById(id string)(*ItemCalendarsItemCalendarViewItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewItemCalendarsItemCalendarViewItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the calendar view for the calendar. Navigation property. Read-only.
func (m *ItemCalendarsItemCalendarViewEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarsItemCalendarViewEventItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, error) {
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
func (m *ItemCalendarsItemCalendarViewEventItemRequestBuilder) Instances()(*ItemCalendarsItemCalendarViewItemInstancesRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InstancesById provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewEventItemRequestBuilder) InstancesById(id string)(*ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return NewItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphAccept provides operations to call the accept method.
func (m *ItemCalendarsItemCalendarViewEventItemRequestBuilder) MicrosoftGraphAccept()(*ItemCalendarsItemCalendarViewItemMicrosoftGraphAcceptRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemMicrosoftGraphAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCancel provides operations to call the cancel method.
func (m *ItemCalendarsItemCalendarViewEventItemRequestBuilder) MicrosoftGraphCancel()(*ItemCalendarsItemCalendarViewItemMicrosoftGraphCancelRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemMicrosoftGraphCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDecline provides operations to call the decline method.
func (m *ItemCalendarsItemCalendarViewEventItemRequestBuilder) MicrosoftGraphDecline()(*ItemCalendarsItemCalendarViewItemMicrosoftGraphDeclineRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemMicrosoftGraphDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDismissReminder provides operations to call the dismissReminder method.
func (m *ItemCalendarsItemCalendarViewEventItemRequestBuilder) MicrosoftGraphDismissReminder()(*ItemCalendarsItemCalendarViewItemMicrosoftGraphDismissReminderRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemMicrosoftGraphDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphForward provides operations to call the forward method.
func (m *ItemCalendarsItemCalendarViewEventItemRequestBuilder) MicrosoftGraphForward()(*ItemCalendarsItemCalendarViewItemMicrosoftGraphForwardRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemMicrosoftGraphForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSnoozeReminder provides operations to call the snoozeReminder method.
func (m *ItemCalendarsItemCalendarViewEventItemRequestBuilder) MicrosoftGraphSnoozeReminder()(*ItemCalendarsItemCalendarViewItemMicrosoftGraphSnoozeReminderRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemMicrosoftGraphSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTentativelyAccept provides operations to call the tentativelyAccept method.
func (m *ItemCalendarsItemCalendarViewEventItemRequestBuilder) MicrosoftGraphTentativelyAccept()(*ItemCalendarsItemCalendarViewItemMicrosoftGraphTentativelyAcceptRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemMicrosoftGraphTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewEventItemRequestBuilder) MultiValueExtendedProperties()(*ItemCalendarsItemCalendarViewItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ItemCalendarsItemCalendarViewItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemCalendarsItemCalendarViewItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewEventItemRequestBuilder) SingleValueExtendedProperties()(*ItemCalendarsItemCalendarViewItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ItemCalendarsItemCalendarViewItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemCalendarsItemCalendarViewItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation the calendar view for the calendar. Navigation property. Read-only.
func (m *ItemCalendarsItemCalendarViewEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarsItemCalendarViewEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
