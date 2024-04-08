package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendarRequestBuilder provides operations to manage the calendar property of the microsoft.graph.group entity.
type ItemCalendarRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemCalendarRequestBuilderGetQueryParameters the group's calendar. Read-only.
type ItemCalendarRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCalendarRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarRequestBuilderGetQueryParameters
}
// CalendarPermissions provides operations to manage the calendarPermissions property of the microsoft.graph.calendar entity.
func (m *ItemCalendarRequestBuilder) CalendarPermissions()(*ItemCalendarCalendarPermissionsRequestBuilder) {
    return NewItemCalendarCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CalendarPermissionsById provides operations to manage the calendarPermissions property of the microsoft.graph.calendar entity.
func (m *ItemCalendarRequestBuilder) CalendarPermissionsById(id string)(*ItemCalendarCalendarPermissionsCalendarPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission%2Did"] = id
    }
    return NewItemCalendarCalendarPermissionsCalendarPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// CalendarView provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
func (m *ItemCalendarRequestBuilder) CalendarView()(*ItemCalendarCalendarViewRequestBuilder) {
    return NewItemCalendarCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CalendarViewById provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
func (m *ItemCalendarRequestBuilder) CalendarViewById(id string)(*ItemCalendarCalendarViewEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return NewItemCalendarCalendarViewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemCalendarRequestBuilderInternal instantiates a new CalendarRequestBuilder and sets the default values.
func NewItemCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarRequestBuilder) {
    m := &ItemCalendarRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendar{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemCalendarRequestBuilder instantiates a new CalendarRequestBuilder and sets the default values.
func NewItemCalendarRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarRequestBuilderInternal(urlParams, requestAdapter)
}
// Events provides operations to manage the events property of the microsoft.graph.calendar entity.
func (m *ItemCalendarRequestBuilder) Events()(*ItemCalendarEventsRequestBuilder) {
    return NewItemCalendarEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// EventsById provides operations to manage the events property of the microsoft.graph.calendar entity.
func (m *ItemCalendarRequestBuilder) EventsById(id string)(*ItemCalendarEventsEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return NewItemCalendarEventsEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the group's calendar. Read-only.
func (m *ItemCalendarRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable, error) {
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
func (m *ItemCalendarRequestBuilder) MicrosoftGraphAllowedCalendarSharingRolesWithUser(user *string)(*ItemCalendarMicrosoftGraphAllowedCalendarSharingRolesWithUserRequestBuilder) {
    return NewItemCalendarMicrosoftGraphAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user)
}
// MicrosoftGraphGetSchedule provides operations to call the getSchedule method.
func (m *ItemCalendarRequestBuilder) MicrosoftGraphGetSchedule()(*ItemCalendarMicrosoftGraphGetScheduleRequestBuilder) {
    return NewItemCalendarMicrosoftGraphGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.calendar entity.
func (m *ItemCalendarRequestBuilder) MultiValueExtendedProperties()(*ItemCalendarMultiValueExtendedPropertiesRequestBuilder) {
    return NewItemCalendarMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.calendar entity.
func (m *ItemCalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ItemCalendarMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemCalendarMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.calendar entity.
func (m *ItemCalendarRequestBuilder) SingleValueExtendedProperties()(*ItemCalendarSingleValueExtendedPropertiesRequestBuilder) {
    return NewItemCalendarSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.calendar entity.
func (m *ItemCalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ItemCalendarSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemCalendarSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation the group's calendar. Read-only.
func (m *ItemCalendarRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
