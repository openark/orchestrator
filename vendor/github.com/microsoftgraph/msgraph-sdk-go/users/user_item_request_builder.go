package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserItemRequestBuilder provides operations to manage the collection of user entities.
type UserItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UserItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserItemRequestBuilderGetQueryParameters retrieve the properties and relationships of user object.
type UserItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserItemRequestBuilderGetQueryParameters
}
// UserItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities provides operations to manage the activities property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Activities()(*ItemActivitiesRequestBuilder) {
    return NewItemActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ActivitiesById provides operations to manage the activities property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) ActivitiesById(id string)(*ItemActivitiesUserActivityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userActivity%2Did"] = id
    }
    return NewItemActivitiesUserActivityItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// AgreementAcceptances provides operations to manage the agreementAcceptances property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) AgreementAcceptances()(*ItemAgreementAcceptancesRequestBuilder) {
    return NewItemAgreementAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AgreementAcceptancesById provides operations to manage the agreementAcceptances property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) AgreementAcceptancesById(id string)(*ItemAgreementAcceptancesAgreementAcceptanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementAcceptance%2Did"] = id
    }
    return NewItemAgreementAcceptancesAgreementAcceptanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// AppRoleAssignments provides operations to manage the appRoleAssignments property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) AppRoleAssignments()(*ItemAppRoleAssignmentsRequestBuilder) {
    return NewItemAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AppRoleAssignmentsById provides operations to manage the appRoleAssignments property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) AppRoleAssignmentsById(id string)(*ItemAppRoleAssignmentsAppRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment%2Did"] = id
    }
    return NewItemAppRoleAssignmentsAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Authentication provides operations to manage the authentication property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Authentication()(*ItemAuthenticationRequestBuilder) {
    return NewItemAuthenticationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Calendar()(*ItemCalendarRequestBuilder) {
    return NewItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CalendarGroups provides operations to manage the calendarGroups property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) CalendarGroups()(*ItemCalendarGroupsRequestBuilder) {
    return NewItemCalendarGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CalendarGroupsById provides operations to manage the calendarGroups property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) CalendarGroupsById(id string)(*ItemCalendarGroupsCalendarGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarGroup%2Did"] = id
    }
    return NewItemCalendarGroupsCalendarGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Calendars provides operations to manage the calendars property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Calendars()(*ItemCalendarsRequestBuilder) {
    return NewItemCalendarsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CalendarsById provides operations to manage the calendars property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) CalendarsById(id string)(*ItemCalendarsCalendarItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendar%2Did"] = id
    }
    return NewItemCalendarsCalendarItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// CalendarView provides operations to manage the calendarView property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) CalendarView()(*ItemCalendarViewRequestBuilder) {
    return NewItemCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CalendarViewById provides operations to manage the calendarView property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) CalendarViewById(id string)(*ItemCalendarViewEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return NewItemCalendarViewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Chats provides operations to manage the chats property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Chats()(*ItemChatsRequestBuilder) {
    return NewItemChatsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ChatsById provides operations to manage the chats property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) ChatsById(id string)(*ItemChatsChatItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chat%2Did"] = id
    }
    return NewItemChatsChatItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewUserItemRequestBuilderInternal instantiates a new UserItemRequestBuilder and sets the default values.
func NewUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserItemRequestBuilder) {
    m := &UserItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewUserItemRequestBuilder instantiates a new UserItemRequestBuilder and sets the default values.
func NewUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactFolders provides operations to manage the contactFolders property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) ContactFolders()(*ItemContactFoldersRequestBuilder) {
    return NewItemContactFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ContactFoldersById provides operations to manage the contactFolders property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) ContactFoldersById(id string)(*ItemContactFoldersContactFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contactFolder%2Did"] = id
    }
    return NewItemContactFoldersContactFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Contacts provides operations to manage the contacts property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Contacts()(*ItemContactsRequestBuilder) {
    return NewItemContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ContactsById provides operations to manage the contacts property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) ContactsById(id string)(*ItemContactsContactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contact%2Did"] = id
    }
    return NewItemContactsContactItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// CreatedObjects provides operations to manage the createdObjects property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) CreatedObjects()(*ItemCreatedObjectsRequestBuilder) {
    return NewItemCreatedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CreatedObjectsById provides operations to manage the createdObjects property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) CreatedObjectsById(id string)(*ItemCreatedObjectsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemCreatedObjectsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Delete delete user.   When deleted, user resources are moved to a temporary container and can be restored within 30 days.  After that time, they are permanently deleted.  To learn more, see deletedItems.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/user-delete?view=graph-rest-1.0
func (m *UserItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceManagementTroubleshootingEvents provides operations to manage the deviceManagementTroubleshootingEvents property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) DeviceManagementTroubleshootingEvents()(*ItemDeviceManagementTroubleshootingEventsRequestBuilder) {
    return NewItemDeviceManagementTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceManagementTroubleshootingEventsById provides operations to manage the deviceManagementTroubleshootingEvents property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) DeviceManagementTroubleshootingEventsById(id string)(*ItemDeviceManagementTroubleshootingEventsDeviceManagementTroubleshootingEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTroubleshootingEvent%2Did"] = id
    }
    return NewItemDeviceManagementTroubleshootingEventsDeviceManagementTroubleshootingEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DirectReports provides operations to manage the directReports property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) DirectReports()(*ItemDirectReportsRequestBuilder) {
    return NewItemDirectReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DirectReportsById provides operations to manage the directReports property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) DirectReportsById(id string)(*ItemDirectReportsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemDirectReportsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Drive provides operations to manage the drive property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Drive()(*ItemDriveRequestBuilder) {
    return NewItemDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Drives provides operations to manage the drives property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Drives()(*ItemDrivesRequestBuilder) {
    return NewItemDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DrivesById provides operations to manage the drives property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) DrivesById(id string)(*ItemDrivesDriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return NewItemDrivesDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Events provides operations to manage the events property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Events()(*ItemEventsRequestBuilder) {
    return NewItemEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// EventsById provides operations to manage the events property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) EventsById(id string)(*ItemEventsEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return NewItemEventsEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Extensions()(*ItemExtensionsRequestBuilder) {
    return NewItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) ExtensionsById(id string)(*ItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// FollowedSites provides operations to manage the followedSites property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) FollowedSites()(*ItemFollowedSitesRequestBuilder) {
    return NewItemFollowedSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// FollowedSitesById provides operations to manage the followedSites property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) FollowedSitesById(id string)(*ItemFollowedSitesSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did"] = id
    }
    return NewItemFollowedSitesSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get retrieve the properties and relationships of user object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/user-get?view=graph-rest-1.0
func (m *UserItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
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
func (m *UserItemRequestBuilder) InferenceClassification()(*ItemInferenceClassificationRequestBuilder) {
    return NewItemInferenceClassificationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Insights provides operations to manage the insights property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Insights()(*ItemInsightsRequestBuilder) {
    return NewItemInsightsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// JoinedTeams provides operations to manage the joinedTeams property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) JoinedTeams()(*ItemJoinedTeamsRequestBuilder) {
    return NewItemJoinedTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// JoinedTeamsById provides operations to manage the joinedTeams property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) JoinedTeamsById(id string)(*ItemJoinedTeamsTeamItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["team%2Did"] = id
    }
    return NewItemJoinedTeamsTeamItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// LicenseDetails provides operations to manage the licenseDetails property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) LicenseDetails()(*ItemLicenseDetailsRequestBuilder) {
    return NewItemLicenseDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// LicenseDetailsById provides operations to manage the licenseDetails property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) LicenseDetailsById(id string)(*ItemLicenseDetailsLicenseDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["licenseDetails%2Did"] = id
    }
    return NewItemLicenseDetailsLicenseDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MailFolders provides operations to manage the mailFolders property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) MailFolders()(*ItemMailFoldersRequestBuilder) {
    return NewItemMailFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MailFoldersById provides operations to manage the mailFolders property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) MailFoldersById(id string)(*ItemMailFoldersMailFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mailFolder%2Did"] = id
    }
    return NewItemMailFoldersMailFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ManagedAppRegistrations provides operations to manage the managedAppRegistrations property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) ManagedAppRegistrations()(*ItemManagedAppRegistrationsRequestBuilder) {
    return NewItemManagedAppRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ManagedAppRegistrationsById provides operations to manage the managedAppRegistrations property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) ManagedAppRegistrationsById(id string)(*ItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppRegistration%2Did"] = id
    }
    return NewItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ManagedDevices provides operations to manage the managedDevices property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) ManagedDevices()(*ItemManagedDevicesRequestBuilder) {
    return NewItemManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ManagedDevicesById provides operations to manage the managedDevices property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) ManagedDevicesById(id string)(*ItemManagedDevicesManagedDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDevice%2Did"] = id
    }
    return NewItemManagedDevicesManagedDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Manager provides operations to manage the manager property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Manager()(*ItemManagerRequestBuilder) {
    return NewItemManagerRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MemberOf provides operations to manage the memberOf property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) MemberOf()(*ItemMemberOfRequestBuilder) {
    return NewItemMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MemberOfById provides operations to manage the memberOf property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) MemberOfById(id string)(*ItemMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Messages provides operations to manage the messages property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Messages()(*ItemMessagesRequestBuilder) {
    return NewItemMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) MessagesById(id string)(*ItemMessagesMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message%2Did"] = id
    }
    return NewItemMessagesMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphAssignLicense provides operations to call the assignLicense method.
func (m *UserItemRequestBuilder) MicrosoftGraphAssignLicense()(*ItemMicrosoftGraphAssignLicenseRequestBuilder) {
    return NewItemMicrosoftGraphAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphChangePassword provides operations to call the changePassword method.
func (m *UserItemRequestBuilder) MicrosoftGraphChangePassword()(*ItemMicrosoftGraphChangePasswordRequestBuilder) {
    return NewItemMicrosoftGraphChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *UserItemRequestBuilder) MicrosoftGraphCheckMemberGroups()(*ItemMicrosoftGraphCheckMemberGroupsRequestBuilder) {
    return NewItemMicrosoftGraphCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *UserItemRequestBuilder) MicrosoftGraphCheckMemberObjects()(*ItemMicrosoftGraphCheckMemberObjectsRequestBuilder) {
    return NewItemMicrosoftGraphCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphExportDeviceAndAppManagementData provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserItemRequestBuilder) MicrosoftGraphExportDeviceAndAppManagementData()(*ItemMicrosoftGraphExportDeviceAndAppManagementDataRequestBuilder) {
    return NewItemMicrosoftGraphExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserItemRequestBuilder) MicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*ItemMicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return NewItemMicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top)
}
// MicrosoftGraphExportPersonalData provides operations to call the exportPersonalData method.
func (m *UserItemRequestBuilder) MicrosoftGraphExportPersonalData()(*ItemMicrosoftGraphExportPersonalDataRequestBuilder) {
    return NewItemMicrosoftGraphExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphFindMeetingTimes provides operations to call the findMeetingTimes method.
func (m *UserItemRequestBuilder) MicrosoftGraphFindMeetingTimes()(*ItemMicrosoftGraphFindMeetingTimesRequestBuilder) {
    return NewItemMicrosoftGraphFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetMailTips provides operations to call the getMailTips method.
func (m *UserItemRequestBuilder) MicrosoftGraphGetMailTips()(*ItemMicrosoftGraphGetMailTipsRequestBuilder) {
    return NewItemMicrosoftGraphGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserItemRequestBuilder) MicrosoftGraphGetManagedAppDiagnosticStatuses()(*ItemMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder) {
    return NewItemMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserItemRequestBuilder) MicrosoftGraphGetManagedAppPolicies()(*ItemMicrosoftGraphGetManagedAppPoliciesRequestBuilder) {
    return NewItemMicrosoftGraphGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserItemRequestBuilder) MicrosoftGraphGetManagedDevicesWithAppFailures()(*ItemMicrosoftGraphGetManagedDevicesWithAppFailuresRequestBuilder) {
    return NewItemMicrosoftGraphGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetMemberGroups provides operations to call the getMemberGroups method.
func (m *UserItemRequestBuilder) MicrosoftGraphGetMemberGroups()(*ItemMicrosoftGraphGetMemberGroupsRequestBuilder) {
    return NewItemMicrosoftGraphGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetMemberObjects provides operations to call the getMemberObjects method.
func (m *UserItemRequestBuilder) MicrosoftGraphGetMemberObjects()(*ItemMicrosoftGraphGetMemberObjectsRequestBuilder) {
    return NewItemMicrosoftGraphGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserItemRequestBuilder) MicrosoftGraphReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*ItemMicrosoftGraphReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewItemMicrosoftGraphReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime)
}
// MicrosoftGraphRemoveAllDevicesFromManagement provides operations to call the removeAllDevicesFromManagement method.
func (m *UserItemRequestBuilder) MicrosoftGraphRemoveAllDevicesFromManagement()(*ItemMicrosoftGraphRemoveAllDevicesFromManagementRequestBuilder) {
    return NewItemMicrosoftGraphRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReprocessLicenseAssignment provides operations to call the reprocessLicenseAssignment method.
func (m *UserItemRequestBuilder) MicrosoftGraphReprocessLicenseAssignment()(*ItemMicrosoftGraphReprocessLicenseAssignmentRequestBuilder) {
    return NewItemMicrosoftGraphReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRestore provides operations to call the restore method.
func (m *UserItemRequestBuilder) MicrosoftGraphRestore()(*ItemMicrosoftGraphRestoreRequestBuilder) {
    return NewItemMicrosoftGraphRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRevokeSignInSessions provides operations to call the revokeSignInSessions method.
func (m *UserItemRequestBuilder) MicrosoftGraphRevokeSignInSessions()(*ItemMicrosoftGraphRevokeSignInSessionsRequestBuilder) {
    return NewItemMicrosoftGraphRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSendMail provides operations to call the sendMail method.
func (m *UserItemRequestBuilder) MicrosoftGraphSendMail()(*ItemMicrosoftGraphSendMailRequestBuilder) {
    return NewItemMicrosoftGraphSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTranslateExchangeIds provides operations to call the translateExchangeIds method.
func (m *UserItemRequestBuilder) MicrosoftGraphTranslateExchangeIds()(*ItemMicrosoftGraphTranslateExchangeIdsRequestBuilder) {
    return NewItemMicrosoftGraphTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWipeManagedAppRegistrationsByDeviceTag provides operations to call the wipeManagedAppRegistrationsByDeviceTag method.
func (m *UserItemRequestBuilder) MicrosoftGraphWipeManagedAppRegistrationsByDeviceTag()(*ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return NewItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Oauth2PermissionGrants provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Oauth2PermissionGrants()(*ItemOauth2PermissionGrantsRequestBuilder) {
    return NewItemOauth2PermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Oauth2PermissionGrantsById provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Oauth2PermissionGrantsById(id string)(*ItemOauth2PermissionGrantsOAuth2PermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["oAuth2PermissionGrant%2Did"] = id
    }
    return NewItemOauth2PermissionGrantsOAuth2PermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Onenote provides operations to manage the onenote property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Onenote()(*ItemOnenoteRequestBuilder) {
    return NewItemOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OnlineMeetings provides operations to manage the onlineMeetings property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) OnlineMeetings()(*ItemOnlineMeetingsRequestBuilder) {
    return NewItemOnlineMeetingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OnlineMeetingsById provides operations to manage the onlineMeetings property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) OnlineMeetingsById(id string)(*ItemOnlineMeetingsOnlineMeetingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onlineMeeting%2Did"] = id
    }
    return NewItemOnlineMeetingsOnlineMeetingItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Outlook provides operations to manage the outlook property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Outlook()(*ItemOutlookRequestBuilder) {
    return NewItemOutlookRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OwnedDevices provides operations to manage the ownedDevices property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) OwnedDevices()(*ItemOwnedDevicesRequestBuilder) {
    return NewItemOwnedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OwnedDevicesById provides operations to manage the ownedDevices property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) OwnedDevicesById(id string)(*ItemOwnedDevicesDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemOwnedDevicesDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// OwnedObjects provides operations to manage the ownedObjects property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) OwnedObjects()(*ItemOwnedObjectsRequestBuilder) {
    return NewItemOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OwnedObjectsById provides operations to manage the ownedObjects property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) OwnedObjectsById(id string)(*ItemOwnedObjectsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemOwnedObjectsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the properties of a user object. Not all properties can be updated by Member or Guest users with their default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/user-update?view=graph-rest-1.0
func (m *UserItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, requestConfiguration *UserItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
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
func (m *UserItemRequestBuilder) People()(*ItemPeopleRequestBuilder) {
    return NewItemPeopleRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PeopleById provides operations to manage the people property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) PeopleById(id string)(*ItemPeoplePersonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["person%2Did"] = id
    }
    return NewItemPeoplePersonItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Photo provides operations to manage the photo property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Photo()(*ItemPhotoRequestBuilder) {
    return NewItemPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Photos provides operations to manage the photos property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Photos()(*ItemPhotosRequestBuilder) {
    return NewItemPhotosRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PhotosById provides operations to manage the photos property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) PhotosById(id string)(*ItemPhotosProfilePhotoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["profilePhoto%2Did"] = id
    }
    return NewItemPhotosProfilePhotoItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Planner provides operations to manage the planner property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Planner()(*ItemPlannerRequestBuilder) {
    return NewItemPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Presence provides operations to manage the presence property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Presence()(*ItemPresenceRequestBuilder) {
    return NewItemPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RegisteredDevices provides operations to manage the registeredDevices property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) RegisteredDevices()(*ItemRegisteredDevicesRequestBuilder) {
    return NewItemRegisteredDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RegisteredDevicesById provides operations to manage the registeredDevices property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) RegisteredDevicesById(id string)(*ItemRegisteredDevicesDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemRegisteredDevicesDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ScopedRoleMemberOf provides operations to manage the scopedRoleMemberOf property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) ScopedRoleMemberOf()(*ItemScopedRoleMemberOfRequestBuilder) {
    return NewItemScopedRoleMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ScopedRoleMemberOfById provides operations to manage the scopedRoleMemberOf property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) ScopedRoleMemberOfById(id string)(*ItemScopedRoleMemberOfScopedRoleMembershipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership%2Did"] = id
    }
    return NewItemScopedRoleMemberOfScopedRoleMembershipItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Settings provides operations to manage the settings property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Settings()(*ItemSettingsRequestBuilder) {
    return NewItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Teamwork provides operations to manage the teamwork property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Teamwork()(*ItemTeamworkRequestBuilder) {
    return NewItemTeamworkRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete user.   When deleted, user resources are moved to a temporary container and can be restored within 30 days.  After that time, they are permanently deleted.  To learn more, see deletedItems.
func (m *UserItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Todo provides operations to manage the todo property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) Todo()(*ItemTodoRequestBuilder) {
    return NewItemTodoRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation retrieve the properties and relationships of user object.
func (m *UserItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UserItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, requestConfiguration *UserItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UserItemRequestBuilder) TransitiveMemberOf()(*ItemTransitiveMemberOfRequestBuilder) {
    return NewItemTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TransitiveMemberOfById provides operations to manage the transitiveMemberOf property of the microsoft.graph.user entity.
func (m *UserItemRequestBuilder) TransitiveMemberOfById(id string)(*ItemTransitiveMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
