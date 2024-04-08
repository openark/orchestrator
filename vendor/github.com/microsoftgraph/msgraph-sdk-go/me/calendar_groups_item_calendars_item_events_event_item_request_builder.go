package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder provides operations to manage the events property of the microsoft.graph.calendar entity.
type CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilderGetQueryParameters the events in the calendar. Navigation property. Read-only.
type CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilderGetQueryParameters
}
// CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) Attachments()(*CalendarGroupsItemCalendarsItemEventsItemAttachmentsRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) AttachmentsById(id string)(*CalendarGroupsItemCalendarsItemEventsItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewCalendarGroupsItemCalendarsItemEventsItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) Calendar()(*CalendarGroupsItemCalendarsItemEventsItemCalendarRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewCalendarGroupsItemCalendarsItemEventsEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewCalendarGroupsItemCalendarsItemEventsEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) {
    m := &CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewCalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarGroupsItemCalendarsItemEventsEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property events for me
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) Extensions()(*CalendarGroupsItemCalendarsItemEventsItemExtensionsRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) ExtensionsById(id string)(*CalendarGroupsItemCalendarsItemEventsItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewCalendarGroupsItemCalendarsItemEventsItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the events in the calendar. Navigation property. Read-only.
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, error) {
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
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) Instances()(*CalendarGroupsItemCalendarsItemEventsItemInstancesRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InstancesById provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) InstancesById(id string)(*CalendarGroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return NewCalendarGroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphAccept provides operations to call the accept method.
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) MicrosoftGraphAccept()(*CalendarGroupsItemCalendarsItemEventsItemMicrosoftGraphAcceptRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemMicrosoftGraphAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCancel provides operations to call the cancel method.
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) MicrosoftGraphCancel()(*CalendarGroupsItemCalendarsItemEventsItemMicrosoftGraphCancelRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemMicrosoftGraphCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDecline provides operations to call the decline method.
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) MicrosoftGraphDecline()(*CalendarGroupsItemCalendarsItemEventsItemMicrosoftGraphDeclineRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemMicrosoftGraphDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDismissReminder provides operations to call the dismissReminder method.
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) MicrosoftGraphDismissReminder()(*CalendarGroupsItemCalendarsItemEventsItemMicrosoftGraphDismissReminderRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemMicrosoftGraphDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphForward provides operations to call the forward method.
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) MicrosoftGraphForward()(*CalendarGroupsItemCalendarsItemEventsItemMicrosoftGraphForwardRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemMicrosoftGraphForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSnoozeReminder provides operations to call the snoozeReminder method.
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) MicrosoftGraphSnoozeReminder()(*CalendarGroupsItemCalendarsItemEventsItemMicrosoftGraphSnoozeReminderRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemMicrosoftGraphSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTentativelyAccept provides operations to call the tentativelyAccept method.
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) MicrosoftGraphTentativelyAccept()(*CalendarGroupsItemCalendarsItemEventsItemMicrosoftGraphTentativelyAcceptRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemMicrosoftGraphTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) MultiValueExtendedProperties()(*CalendarGroupsItemCalendarsItemEventsItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*CalendarGroupsItemCalendarsItemEventsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewCalendarGroupsItemCalendarsItemEventsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property events in me
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, requestConfiguration *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) SingleValueExtendedProperties()(*CalendarGroupsItemCalendarsItemEventsItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*CalendarGroupsItemCalendarsItemEventsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewCalendarGroupsItemCalendarsItemEventsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property events for me
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the events in the calendar. Navigation property. Read-only.
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property events in me
func (m *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, requestConfiguration *CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
