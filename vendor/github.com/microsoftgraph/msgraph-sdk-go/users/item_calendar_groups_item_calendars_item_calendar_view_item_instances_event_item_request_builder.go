package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderGetQueryParameters
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) Attachments()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemAttachmentsRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) AttachmentsById(id string)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) Calendar()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemCalendarRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) {
    m := &ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) Extensions()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExtensionsRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) ExtensionsById(id string)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, error) {
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
// MicrosoftGraphAccept provides operations to call the accept method.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) MicrosoftGraphAccept()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphAcceptRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCancel provides operations to call the cancel method.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) MicrosoftGraphCancel()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphCancelRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDecline provides operations to call the decline method.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) MicrosoftGraphDecline()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDeclineRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDismissReminder provides operations to call the dismissReminder method.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) MicrosoftGraphDismissReminder()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDismissReminderRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphForward provides operations to call the forward method.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) MicrosoftGraphForward()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphForwardRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSnoozeReminder provides operations to call the snoozeReminder method.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) MicrosoftGraphSnoozeReminder()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphSnoozeReminderRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTentativelyAccept provides operations to call the tentativelyAccept method.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) MicrosoftGraphTentativelyAccept()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) MultiValueExtendedProperties()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) SingleValueExtendedProperties()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
