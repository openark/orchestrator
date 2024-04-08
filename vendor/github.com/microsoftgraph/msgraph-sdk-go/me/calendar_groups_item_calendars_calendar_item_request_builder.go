package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CalendarGroupsItemCalendarsCalendarItemRequestBuilder provides operations to manage the calendars property of the microsoft.graph.calendarGroup entity.
type CalendarGroupsItemCalendarsCalendarItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarGroupsItemCalendarsCalendarItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarGroupsItemCalendarsCalendarItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CalendarGroupsItemCalendarsCalendarItemRequestBuilderGetQueryParameters the calendars in the calendar group. Navigation property. Read-only. Nullable.
type CalendarGroupsItemCalendarsCalendarItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CalendarGroupsItemCalendarsCalendarItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarGroupsItemCalendarsCalendarItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CalendarGroupsItemCalendarsCalendarItemRequestBuilderGetQueryParameters
}
// CalendarGroupsItemCalendarsCalendarItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarGroupsItemCalendarsCalendarItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CalendarPermissions provides operations to manage the calendarPermissions property of the microsoft.graph.calendar entity.
func (m *CalendarGroupsItemCalendarsCalendarItemRequestBuilder) CalendarPermissions()(*CalendarGroupsItemCalendarsItemCalendarPermissionsRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CalendarPermissionsById provides operations to manage the calendarPermissions property of the microsoft.graph.calendar entity.
func (m *CalendarGroupsItemCalendarsCalendarItemRequestBuilder) CalendarPermissionsById(id string)(*CalendarGroupsItemCalendarsItemCalendarPermissionsCalendarPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission%2Did"] = id
    }
    return NewCalendarGroupsItemCalendarsItemCalendarPermissionsCalendarPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// CalendarView provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
func (m *CalendarGroupsItemCalendarsCalendarItemRequestBuilder) CalendarView()(*CalendarGroupsItemCalendarsItemCalendarViewRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CalendarViewById provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
func (m *CalendarGroupsItemCalendarsCalendarItemRequestBuilder) CalendarViewById(id string)(*CalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return NewCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewCalendarGroupsItemCalendarsCalendarItemRequestBuilderInternal instantiates a new CalendarItemRequestBuilder and sets the default values.
func NewCalendarGroupsItemCalendarsCalendarItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarGroupsItemCalendarsCalendarItemRequestBuilder) {
    m := &CalendarGroupsItemCalendarsCalendarItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCalendarGroupsItemCalendarsCalendarItemRequestBuilder instantiates a new CalendarItemRequestBuilder and sets the default values.
func NewCalendarGroupsItemCalendarsCalendarItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarGroupsItemCalendarsCalendarItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarGroupsItemCalendarsCalendarItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property calendars for me
func (m *CalendarGroupsItemCalendarsCalendarItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CalendarGroupsItemCalendarsCalendarItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Events provides operations to manage the events property of the microsoft.graph.calendar entity.
func (m *CalendarGroupsItemCalendarsCalendarItemRequestBuilder) Events()(*CalendarGroupsItemCalendarsItemEventsRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// EventsById provides operations to manage the events property of the microsoft.graph.calendar entity.
func (m *CalendarGroupsItemCalendarsCalendarItemRequestBuilder) EventsById(id string)(*CalendarGroupsItemCalendarsItemEventsEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return NewCalendarGroupsItemCalendarsItemEventsEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the calendars in the calendar group. Navigation property. Read-only. Nullable.
func (m *CalendarGroupsItemCalendarsCalendarItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CalendarGroupsItemCalendarsCalendarItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCalendarFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable), nil
}
// MicrosoftGraphAllowedCalendarSharingRolesWithUser provides operations to call the allowedCalendarSharingRoles method.
func (m *CalendarGroupsItemCalendarsCalendarItemRequestBuilder) MicrosoftGraphAllowedCalendarSharingRolesWithUser(user *string)(*CalendarGroupsItemCalendarsItemMicrosoftGraphAllowedCalendarSharingRolesWithUserRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemMicrosoftGraphAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user)
}
// MicrosoftGraphGetSchedule provides operations to call the getSchedule method.
func (m *CalendarGroupsItemCalendarsCalendarItemRequestBuilder) MicrosoftGraphGetSchedule()(*CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.calendar entity.
func (m *CalendarGroupsItemCalendarsCalendarItemRequestBuilder) MultiValueExtendedProperties()(*CalendarGroupsItemCalendarsItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.calendar entity.
func (m *CalendarGroupsItemCalendarsCalendarItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*CalendarGroupsItemCalendarsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewCalendarGroupsItemCalendarsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property calendars in me
func (m *CalendarGroupsItemCalendarsCalendarItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable, requestConfiguration *CalendarGroupsItemCalendarsCalendarItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCalendarFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable), nil
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.calendar entity.
func (m *CalendarGroupsItemCalendarsCalendarItemRequestBuilder) SingleValueExtendedProperties()(*CalendarGroupsItemCalendarsItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.calendar entity.
func (m *CalendarGroupsItemCalendarsCalendarItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*CalendarGroupsItemCalendarsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewCalendarGroupsItemCalendarsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property calendars for me
func (m *CalendarGroupsItemCalendarsCalendarItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CalendarGroupsItemCalendarsCalendarItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the calendars in the calendar group. Navigation property. Read-only. Nullable.
func (m *CalendarGroupsItemCalendarsCalendarItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CalendarGroupsItemCalendarsCalendarItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property calendars in me
func (m *CalendarGroupsItemCalendarsCalendarItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable, requestConfiguration *CalendarGroupsItemCalendarsCalendarItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
