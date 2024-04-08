package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MeRequestBuilder provides operations to manage the user singleton.
type MeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeRequestBuilderGetQueryParameters retrieve the properties and relationships of user object.
type MeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeRequestBuilderGetQueryParameters
}
// MeRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities provides operations to manage the activities property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Activities()(*ActivitiesRequestBuilder) {
    return NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ActivitiesById provides operations to manage the activities property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ActivitiesById(id string)(*ActivitiesUserActivityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userActivity%2Did"] = id
    }
    return NewActivitiesUserActivityItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// AgreementAcceptances provides operations to manage the agreementAcceptances property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) AgreementAcceptances()(*AgreementAcceptancesRequestBuilder) {
    return NewAgreementAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AgreementAcceptancesById provides operations to manage the agreementAcceptances property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) AgreementAcceptancesById(id string)(*AgreementAcceptancesAgreementAcceptanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementAcceptance%2Did"] = id
    }
    return NewAgreementAcceptancesAgreementAcceptanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// AppRoleAssignments provides operations to manage the appRoleAssignments property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) AppRoleAssignments()(*AppRoleAssignmentsRequestBuilder) {
    return NewAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AppRoleAssignmentsById provides operations to manage the appRoleAssignments property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) AppRoleAssignmentsById(id string)(*AppRoleAssignmentsAppRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment%2Did"] = id
    }
    return NewAppRoleAssignmentsAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Authentication provides operations to manage the authentication property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Authentication()(*AuthenticationRequestBuilder) {
    return NewAuthenticationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Calendar()(*CalendarRequestBuilder) {
    return NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CalendarGroups provides operations to manage the calendarGroups property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) CalendarGroups()(*CalendarGroupsRequestBuilder) {
    return NewCalendarGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CalendarGroupsById provides operations to manage the calendarGroups property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) CalendarGroupsById(id string)(*CalendarGroupsCalendarGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarGroup%2Did"] = id
    }
    return NewCalendarGroupsCalendarGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Calendars provides operations to manage the calendars property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Calendars()(*CalendarsRequestBuilder) {
    return NewCalendarsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CalendarsById provides operations to manage the calendars property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) CalendarsById(id string)(*CalendarsCalendarItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendar%2Did"] = id
    }
    return NewCalendarsCalendarItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// CalendarView provides operations to manage the calendarView property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) CalendarView()(*CalendarViewRequestBuilder) {
    return NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CalendarViewById provides operations to manage the calendarView property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) CalendarViewById(id string)(*CalendarViewEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return NewCalendarViewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Chats provides operations to manage the chats property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Chats()(*ChatsRequestBuilder) {
    return NewChatsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ChatsById provides operations to manage the chats property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ChatsById(id string)(*ChatsChatItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chat%2Did"] = id
    }
    return NewChatsChatItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewMeRequestBuilderInternal instantiates a new MeRequestBuilder and sets the default values.
func NewMeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeRequestBuilder) {
    m := &MeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMeRequestBuilder instantiates a new MeRequestBuilder and sets the default values.
func NewMeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactFolders provides operations to manage the contactFolders property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ContactFolders()(*ContactFoldersRequestBuilder) {
    return NewContactFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ContactFoldersById provides operations to manage the contactFolders property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ContactFoldersById(id string)(*ContactFoldersContactFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contactFolder%2Did"] = id
    }
    return NewContactFoldersContactFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Contacts provides operations to manage the contacts property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Contacts()(*ContactsRequestBuilder) {
    return NewContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ContactsById provides operations to manage the contacts property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ContactsById(id string)(*ContactsContactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contact%2Did"] = id
    }
    return NewContactsContactItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// CreatedObjects provides operations to manage the createdObjects property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) CreatedObjects()(*CreatedObjectsRequestBuilder) {
    return NewCreatedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CreatedObjectsById provides operations to manage the createdObjects property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) CreatedObjectsById(id string)(*CreatedObjectsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewCreatedObjectsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DeviceManagementTroubleshootingEvents provides operations to manage the deviceManagementTroubleshootingEvents property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) DeviceManagementTroubleshootingEvents()(*DeviceManagementTroubleshootingEventsRequestBuilder) {
    return NewDeviceManagementTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceManagementTroubleshootingEventsById provides operations to manage the deviceManagementTroubleshootingEvents property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) DeviceManagementTroubleshootingEventsById(id string)(*DeviceManagementTroubleshootingEventsDeviceManagementTroubleshootingEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTroubleshootingEvent%2Did"] = id
    }
    return NewDeviceManagementTroubleshootingEventsDeviceManagementTroubleshootingEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DirectReports provides operations to manage the directReports property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) DirectReports()(*DirectReportsRequestBuilder) {
    return NewDirectReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DirectReportsById provides operations to manage the directReports property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) DirectReportsById(id string)(*DirectReportsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewDirectReportsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Drive provides operations to manage the drive property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Drive()(*DriveRequestBuilder) {
    return NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Drives provides operations to manage the drives property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Drives()(*DrivesRequestBuilder) {
    return NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DrivesById provides operations to manage the drives property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) DrivesById(id string)(*DrivesDriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return NewDrivesDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Events provides operations to manage the events property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Events()(*EventsRequestBuilder) {
    return NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// EventsById provides operations to manage the events property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) EventsById(id string)(*EventsEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return NewEventsEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Extensions()(*ExtensionsRequestBuilder) {
    return NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ExtensionsById(id string)(*ExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// FollowedSites provides operations to manage the followedSites property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) FollowedSites()(*FollowedSitesRequestBuilder) {
    return NewFollowedSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// FollowedSitesById provides operations to manage the followedSites property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) FollowedSitesById(id string)(*FollowedSitesSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did"] = id
    }
    return NewFollowedSitesSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get retrieve the properties and relationships of user object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/user-get?view=graph-rest-1.0
func (m *MeRequestBuilder) Get(ctx context.Context, requestConfiguration *MeRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable), nil
}
// InferenceClassification provides operations to manage the inferenceClassification property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) InferenceClassification()(*InferenceClassificationRequestBuilder) {
    return NewInferenceClassificationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Insights provides operations to manage the insights property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Insights()(*InsightsRequestBuilder) {
    return NewInsightsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// JoinedTeams provides operations to manage the joinedTeams property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) JoinedTeams()(*JoinedTeamsRequestBuilder) {
    return NewJoinedTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// JoinedTeamsById provides operations to manage the joinedTeams property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) JoinedTeamsById(id string)(*JoinedTeamsTeamItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["team%2Did"] = id
    }
    return NewJoinedTeamsTeamItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// LicenseDetails provides operations to manage the licenseDetails property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) LicenseDetails()(*LicenseDetailsRequestBuilder) {
    return NewLicenseDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// LicenseDetailsById provides operations to manage the licenseDetails property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) LicenseDetailsById(id string)(*LicenseDetailsLicenseDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["licenseDetails%2Did"] = id
    }
    return NewLicenseDetailsLicenseDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MailFolders provides operations to manage the mailFolders property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) MailFolders()(*MailFoldersRequestBuilder) {
    return NewMailFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MailFoldersById provides operations to manage the mailFolders property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) MailFoldersById(id string)(*MailFoldersMailFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mailFolder%2Did"] = id
    }
    return NewMailFoldersMailFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ManagedAppRegistrations provides operations to manage the managedAppRegistrations property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ManagedAppRegistrations()(*ManagedAppRegistrationsRequestBuilder) {
    return NewManagedAppRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ManagedAppRegistrationsById provides operations to manage the managedAppRegistrations property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ManagedAppRegistrationsById(id string)(*ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppRegistration%2Did"] = id
    }
    return NewManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ManagedDevices provides operations to manage the managedDevices property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ManagedDevices()(*ManagedDevicesRequestBuilder) {
    return NewManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ManagedDevicesById provides operations to manage the managedDevices property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ManagedDevicesById(id string)(*ManagedDevicesManagedDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDevice%2Did"] = id
    }
    return NewManagedDevicesManagedDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Manager provides operations to manage the manager property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Manager()(*ManagerRequestBuilder) {
    return NewManagerRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MemberOf provides operations to manage the memberOf property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) MemberOf()(*MemberOfRequestBuilder) {
    return NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MemberOfById provides operations to manage the memberOf property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) MemberOfById(id string)(*MemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Messages provides operations to manage the messages property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Messages()(*MessagesRequestBuilder) {
    return NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) MessagesById(id string)(*MessagesMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message%2Did"] = id
    }
    return NewMessagesMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphAssignLicense provides operations to call the assignLicense method.
func (m *MeRequestBuilder) MicrosoftGraphAssignLicense()(*MicrosoftGraphAssignLicenseRequestBuilder) {
    return NewMicrosoftGraphAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphChangePassword provides operations to call the changePassword method.
func (m *MeRequestBuilder) MicrosoftGraphChangePassword()(*MicrosoftGraphChangePasswordRequestBuilder) {
    return NewMicrosoftGraphChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *MeRequestBuilder) MicrosoftGraphCheckMemberGroups()(*MicrosoftGraphCheckMemberGroupsRequestBuilder) {
    return NewMicrosoftGraphCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *MeRequestBuilder) MicrosoftGraphCheckMemberObjects()(*MicrosoftGraphCheckMemberObjectsRequestBuilder) {
    return NewMicrosoftGraphCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphExportDeviceAndAppManagementData provides operations to call the exportDeviceAndAppManagementData method.
func (m *MeRequestBuilder) MicrosoftGraphExportDeviceAndAppManagementData()(*MicrosoftGraphExportDeviceAndAppManagementDataRequestBuilder) {
    return NewMicrosoftGraphExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *MeRequestBuilder) MicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*MicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return NewMicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top)
}
// MicrosoftGraphExportPersonalData provides operations to call the exportPersonalData method.
func (m *MeRequestBuilder) MicrosoftGraphExportPersonalData()(*MicrosoftGraphExportPersonalDataRequestBuilder) {
    return NewMicrosoftGraphExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphFindMeetingTimes provides operations to call the findMeetingTimes method.
func (m *MeRequestBuilder) MicrosoftGraphFindMeetingTimes()(*MicrosoftGraphFindMeetingTimesRequestBuilder) {
    return NewMicrosoftGraphFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetMailTips provides operations to call the getMailTips method.
func (m *MeRequestBuilder) MicrosoftGraphGetMailTips()(*MicrosoftGraphGetMailTipsRequestBuilder) {
    return NewMicrosoftGraphGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *MeRequestBuilder) MicrosoftGraphGetManagedAppDiagnosticStatuses()(*MicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder) {
    return NewMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *MeRequestBuilder) MicrosoftGraphGetManagedAppPolicies()(*MicrosoftGraphGetManagedAppPoliciesRequestBuilder) {
    return NewMicrosoftGraphGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *MeRequestBuilder) MicrosoftGraphGetManagedDevicesWithAppFailures()(*MicrosoftGraphGetManagedDevicesWithAppFailuresRequestBuilder) {
    return NewMicrosoftGraphGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetMemberGroups provides operations to call the getMemberGroups method.
func (m *MeRequestBuilder) MicrosoftGraphGetMemberGroups()(*MicrosoftGraphGetMemberGroupsRequestBuilder) {
    return NewMicrosoftGraphGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetMemberObjects provides operations to call the getMemberObjects method.
func (m *MeRequestBuilder) MicrosoftGraphGetMemberObjects()(*MicrosoftGraphGetMemberObjectsRequestBuilder) {
    return NewMicrosoftGraphGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *MeRequestBuilder) MicrosoftGraphReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*MicrosoftGraphReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewMicrosoftGraphReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime)
}
// MicrosoftGraphRemoveAllDevicesFromManagement provides operations to call the removeAllDevicesFromManagement method.
func (m *MeRequestBuilder) MicrosoftGraphRemoveAllDevicesFromManagement()(*MicrosoftGraphRemoveAllDevicesFromManagementRequestBuilder) {
    return NewMicrosoftGraphRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReprocessLicenseAssignment provides operations to call the reprocessLicenseAssignment method.
func (m *MeRequestBuilder) MicrosoftGraphReprocessLicenseAssignment()(*MicrosoftGraphReprocessLicenseAssignmentRequestBuilder) {
    return NewMicrosoftGraphReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRestore provides operations to call the restore method.
func (m *MeRequestBuilder) MicrosoftGraphRestore()(*MicrosoftGraphRestoreRequestBuilder) {
    return NewMicrosoftGraphRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRevokeSignInSessions provides operations to call the revokeSignInSessions method.
func (m *MeRequestBuilder) MicrosoftGraphRevokeSignInSessions()(*MicrosoftGraphRevokeSignInSessionsRequestBuilder) {
    return NewMicrosoftGraphRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSendMail provides operations to call the sendMail method.
func (m *MeRequestBuilder) MicrosoftGraphSendMail()(*MicrosoftGraphSendMailRequestBuilder) {
    return NewMicrosoftGraphSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTranslateExchangeIds provides operations to call the translateExchangeIds method.
func (m *MeRequestBuilder) MicrosoftGraphTranslateExchangeIds()(*MicrosoftGraphTranslateExchangeIdsRequestBuilder) {
    return NewMicrosoftGraphTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWipeManagedAppRegistrationsByDeviceTag provides operations to call the wipeManagedAppRegistrationsByDeviceTag method.
func (m *MeRequestBuilder) MicrosoftGraphWipeManagedAppRegistrationsByDeviceTag()(*MicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return NewMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Oauth2PermissionGrants provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Oauth2PermissionGrants()(*Oauth2PermissionGrantsRequestBuilder) {
    return NewOauth2PermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Oauth2PermissionGrantsById provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Oauth2PermissionGrantsById(id string)(*Oauth2PermissionGrantsOAuth2PermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["oAuth2PermissionGrant%2Did"] = id
    }
    return NewOauth2PermissionGrantsOAuth2PermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Onenote provides operations to manage the onenote property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Onenote()(*OnenoteRequestBuilder) {
    return NewOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OnlineMeetings provides operations to manage the onlineMeetings property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) OnlineMeetings()(*OnlineMeetingsRequestBuilder) {
    return NewOnlineMeetingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OnlineMeetingsById provides operations to manage the onlineMeetings property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) OnlineMeetingsById(id string)(*OnlineMeetingsOnlineMeetingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onlineMeeting%2Did"] = id
    }
    return NewOnlineMeetingsOnlineMeetingItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Outlook provides operations to manage the outlook property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Outlook()(*OutlookRequestBuilder) {
    return NewOutlookRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OwnedDevices provides operations to manage the ownedDevices property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) OwnedDevices()(*OwnedDevicesRequestBuilder) {
    return NewOwnedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OwnedDevicesById provides operations to manage the ownedDevices property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) OwnedDevicesById(id string)(*OwnedDevicesDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewOwnedDevicesDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// OwnedObjects provides operations to manage the ownedObjects property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) OwnedObjects()(*OwnedObjectsRequestBuilder) {
    return NewOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OwnedObjectsById provides operations to manage the ownedObjects property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) OwnedObjectsById(id string)(*OwnedObjectsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewOwnedObjectsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the properties of a user object. Not all properties can be updated by Member or Guest users with their default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/user-update?view=graph-rest-1.0
func (m *MeRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, requestConfiguration *MeRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable), nil
}
// People provides operations to manage the people property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) People()(*PeopleRequestBuilder) {
    return NewPeopleRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PeopleById provides operations to manage the people property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) PeopleById(id string)(*PeoplePersonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["person%2Did"] = id
    }
    return NewPeoplePersonItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Photo provides operations to manage the photo property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Photo()(*PhotoRequestBuilder) {
    return NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Photos provides operations to manage the photos property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Photos()(*PhotosRequestBuilder) {
    return NewPhotosRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PhotosById provides operations to manage the photos property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) PhotosById(id string)(*PhotosProfilePhotoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["profilePhoto%2Did"] = id
    }
    return NewPhotosProfilePhotoItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Planner provides operations to manage the planner property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Planner()(*PlannerRequestBuilder) {
    return NewPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Presence provides operations to manage the presence property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Presence()(*PresenceRequestBuilder) {
    return NewPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RegisteredDevices provides operations to manage the registeredDevices property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) RegisteredDevices()(*RegisteredDevicesRequestBuilder) {
    return NewRegisteredDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RegisteredDevicesById provides operations to manage the registeredDevices property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) RegisteredDevicesById(id string)(*RegisteredDevicesDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewRegisteredDevicesDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ScopedRoleMemberOf provides operations to manage the scopedRoleMemberOf property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ScopedRoleMemberOf()(*ScopedRoleMemberOfRequestBuilder) {
    return NewScopedRoleMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ScopedRoleMemberOfById provides operations to manage the scopedRoleMemberOf property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ScopedRoleMemberOfById(id string)(*ScopedRoleMemberOfScopedRoleMembershipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership%2Did"] = id
    }
    return NewScopedRoleMemberOfScopedRoleMembershipItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Settings provides operations to manage the settings property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Settings()(*SettingsRequestBuilder) {
    return NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Teamwork provides operations to manage the teamwork property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Teamwork()(*TeamworkRequestBuilder) {
    return NewTeamworkRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Todo provides operations to manage the todo property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Todo()(*TodoRequestBuilder) {
    return NewTodoRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation retrieve the properties and relationships of user object.
func (m *MeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a user object. Not all properties can be updated by Member or Guest users with their default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage.
func (m *MeRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, requestConfiguration *MeRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// TransitiveMemberOf provides operations to manage the transitiveMemberOf property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) TransitiveMemberOf()(*TransitiveMemberOfRequestBuilder) {
    return NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TransitiveMemberOfById provides operations to manage the transitiveMemberOf property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) TransitiveMemberOfById(id string)(*TransitiveMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
