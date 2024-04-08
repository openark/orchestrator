package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendarsCalendarItemRequestBuilder provides operations to manage the calendars property of the microsoft.graph.user entity.
type ItemCalendarsCalendarItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemCalendarsCalendarItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarsCalendarItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemCalendarsCalendarItemRequestBuilderGetQueryParameters the user's calendars. Read-only. Nullable.
type ItemCalendarsCalendarItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCalendarsCalendarItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarsCalendarItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarsCalendarItemRequestBuilderGetQueryParameters
}
// ItemCalendarsCalendarItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarsCalendarItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CalendarPermissions provides operations to manage the calendarPermissions property of the microsoft.graph.calendar entity.
func (m *ItemCalendarsCalendarItemRequestBuilder) CalendarPermissions()(*ItemCalendarsItemCalendarPermissionsRequestBuilder) {
    return NewItemCalendarsItemCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CalendarPermissionsById provides operations to manage the calendarPermissions property of the microsoft.graph.calendar entity.
func (m *ItemCalendarsCalendarItemRequestBuilder) CalendarPermissionsById(id string)(*ItemCalendarsItemCalendarPermissionsCalendarPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission%2Did"] = id
    }
    return NewItemCalendarsItemCalendarPermissionsCalendarPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// CalendarView provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
func (m *ItemCalendarsCalendarItemRequestBuilder) CalendarView()(*ItemCalendarsItemCalendarViewRequestBuilder) {
    return NewItemCalendarsItemCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CalendarViewById provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
func (m *ItemCalendarsCalendarItemRequestBuilder) CalendarViewById(id string)(*ItemCalendarsItemCalendarViewEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return NewItemCalendarsItemCalendarViewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemCalendarsCalendarItemRequestBuilderInternal instantiates a new CalendarItemRequestBuilder and sets the default values.
func NewItemCalendarsCalendarItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsCalendarItemRequestBuilder) {
    m := &ItemCalendarsCalendarItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemCalendarsCalendarItemRequestBuilder instantiates a new CalendarItemRequestBuilder and sets the default values.
func NewItemCalendarsCalendarItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsCalendarItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarsCalendarItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property calendars for users
func (m *ItemCalendarsCalendarItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemCalendarsCalendarItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ItemCalendarsCalendarItemRequestBuilder) Events()(*ItemCalendarsItemEventsRequestBuilder) {
    return NewItemCalendarsItemEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// EventsById provides operations to manage the events property of the microsoft.graph.calendar entity.
func (m *ItemCalendarsCalendarItemRequestBuilder) EventsById(id string)(*ItemCalendarsItemEventsEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return NewItemCalendarsItemEventsEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the user's calendars. Read-only. Nullable.
func (m *ItemCalendarsCalendarItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarsCalendarItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable, error) {
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
func (m *ItemCalendarsCalendarItemRequestBuilder) MicrosoftGraphAllowedCalendarSharingRolesWithUser(user *string)(*ItemCalendarsItemMicrosoftGraphAllowedCalendarSharingRolesWithUserRequestBuilder) {
    return NewItemCalendarsItemMicrosoftGraphAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user)
}
// MicrosoftGraphGetSchedule provides operations to call the getSchedule method.
func (m *ItemCalendarsCalendarItemRequestBuilder) MicrosoftGraphGetSchedule()(*ItemCalendarsItemMicrosoftGraphGetScheduleRequestBuilder) {
    return NewItemCalendarsItemMicrosoftGraphGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.calendar entity.
func (m *ItemCalendarsCalendarItemRequestBuilder) MultiValueExtendedProperties()(*ItemCalendarsItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewItemCalendarsItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.calendar entity.
func (m *ItemCalendarsCalendarItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ItemCalendarsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemCalendarsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property calendars in users
func (m *ItemCalendarsCalendarItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable, requestConfiguration *ItemCalendarsCalendarItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable, error) {
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
func (m *ItemCalendarsCalendarItemRequestBuilder) SingleValueExtendedProperties()(*ItemCalendarsItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewItemCalendarsItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.calendar entity.
func (m *ItemCalendarsCalendarItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ItemCalendarsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemCalendarsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property calendars for users
func (m *ItemCalendarsCalendarItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarsCalendarItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the user's calendars. Read-only. Nullable.
func (m *ItemCalendarsCalendarItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarsCalendarItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property calendars in users
func (m *ItemCalendarsCalendarItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable, requestConfiguration *ItemCalendarsCalendarItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
