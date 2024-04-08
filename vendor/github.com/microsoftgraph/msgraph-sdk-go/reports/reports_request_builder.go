package reports

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ReportsRequestBuilder provides operations to manage the reportRoot singleton.
type ReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReportsRequestBuilderGetQueryParameters get reports
type ReportsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ReportsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReportsRequestBuilderGetQueryParameters
}
// ReportsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsRequestBuilderInternal instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsRequestBuilder) {
    m := &ReportsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewReportsRequestBuilder instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// DailyPrintUsageByPrinter provides operations to manage the dailyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageByPrinter()(*DailyPrintUsageByPrinterRequestBuilder) {
    return NewDailyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DailyPrintUsageByPrinterById provides operations to manage the dailyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageByPrinterById(id string)(*DailyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return NewDailyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DailyPrintUsageByUser provides operations to manage the dailyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageByUser()(*DailyPrintUsageByUserRequestBuilder) {
    return NewDailyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DailyPrintUsageByUserById provides operations to manage the dailyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageByUserById(id string)(*DailyPrintUsageByUserPrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return NewDailyPrintUsageByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get get reports
func (m *ReportsRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReportRootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateReportRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReportRootable), nil
}
// MicrosoftGraphDeviceConfigurationDeviceActivity provides operations to call the deviceConfigurationDeviceActivity method.
func (m *ReportsRequestBuilder) MicrosoftGraphDeviceConfigurationDeviceActivity()(*MicrosoftGraphDeviceConfigurationDeviceActivityRequestBuilder) {
    return NewMicrosoftGraphDeviceConfigurationDeviceActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDeviceConfigurationUserActivity provides operations to call the deviceConfigurationUserActivity method.
func (m *ReportsRequestBuilder) MicrosoftGraphDeviceConfigurationUserActivity()(*MicrosoftGraphDeviceConfigurationUserActivityRequestBuilder) {
    return NewMicrosoftGraphDeviceConfigurationUserActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetEmailActivityCountsWithPeriod provides operations to call the getEmailActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailActivityCountsWithPeriod(period *string)(*MicrosoftGraphGetEmailActivityCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetEmailActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailActivityUserCountsWithPeriod provides operations to call the getEmailActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailActivityUserCountsWithPeriod(period *string)(*MicrosoftGraphGetEmailActivityUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetEmailActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailActivityUserDetailWithDate provides operations to call the getEmailActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetEmailActivityUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetEmailActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetEmailActivityUserDetailWithPeriod provides operations to call the getEmailActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailActivityUserDetailWithPeriod(period *string)(*MicrosoftGraphGetEmailActivityUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetEmailActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailAppUsageAppsUserCountsWithPeriod provides operations to call the getEmailAppUsageAppsUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailAppUsageAppsUserCountsWithPeriod(period *string)(*MicrosoftGraphGetEmailAppUsageAppsUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetEmailAppUsageAppsUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailAppUsageUserCountsWithPeriod provides operations to call the getEmailAppUsageUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailAppUsageUserCountsWithPeriod(period *string)(*MicrosoftGraphGetEmailAppUsageUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetEmailAppUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailAppUsageUserDetailWithDate provides operations to call the getEmailAppUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailAppUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetEmailAppUsageUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetEmailAppUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetEmailAppUsageUserDetailWithPeriod provides operations to call the getEmailAppUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailAppUsageUserDetailWithPeriod(period *string)(*MicrosoftGraphGetEmailAppUsageUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetEmailAppUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailAppUsageVersionsUserCountsWithPeriod provides operations to call the getEmailAppUsageVersionsUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailAppUsageVersionsUserCountsWithPeriod(period *string)(*MicrosoftGraphGetEmailAppUsageVersionsUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetEmailAppUsageVersionsUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTime provides operations to call the getGroupArchivedPrintJobs method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, groupId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*MicrosoftGraphGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewMicrosoftGraphGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, groupId, startDateTime)
}
// MicrosoftGraphGetM365AppPlatformUserCountsWithPeriod provides operations to call the getM365AppPlatformUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetM365AppPlatformUserCountsWithPeriod(period *string)(*MicrosoftGraphGetM365AppPlatformUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetM365AppPlatformUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetM365AppUserCountsWithPeriod provides operations to call the getM365AppUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetM365AppUserCountsWithPeriod(period *string)(*MicrosoftGraphGetM365AppUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetM365AppUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetM365AppUserDetailWithDate provides operations to call the getM365AppUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetM365AppUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetM365AppUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetM365AppUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetM365AppUserDetailWithPeriod provides operations to call the getM365AppUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetM365AppUserDetailWithPeriod(period *string)(*MicrosoftGraphGetM365AppUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetM365AppUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetMailboxUsageDetailWithPeriod provides operations to call the getMailboxUsageDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetMailboxUsageDetailWithPeriod(period *string)(*MicrosoftGraphGetMailboxUsageDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetMailboxUsageDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetMailboxUsageMailboxCountsWithPeriod provides operations to call the getMailboxUsageMailboxCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetMailboxUsageMailboxCountsWithPeriod(period *string)(*MicrosoftGraphGetMailboxUsageMailboxCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetMailboxUsageMailboxCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetMailboxUsageQuotaStatusMailboxCountsWithPeriod provides operations to call the getMailboxUsageQuotaStatusMailboxCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetMailboxUsageQuotaStatusMailboxCountsWithPeriod(period *string)(*MicrosoftGraphGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetMailboxUsageStorageWithPeriod provides operations to call the getMailboxUsageStorage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetMailboxUsageStorageWithPeriod(period *string)(*MicrosoftGraphGetMailboxUsageStorageWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetMailboxUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365ActivationCounts provides operations to call the getOffice365ActivationCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActivationCounts()(*MicrosoftGraphGetOffice365ActivationCountsRequestBuilder) {
    return NewMicrosoftGraphGetOffice365ActivationCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetOffice365ActivationsUserCounts provides operations to call the getOffice365ActivationsUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActivationsUserCounts()(*MicrosoftGraphGetOffice365ActivationsUserCountsRequestBuilder) {
    return NewMicrosoftGraphGetOffice365ActivationsUserCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetOffice365ActivationsUserDetail provides operations to call the getOffice365ActivationsUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActivationsUserDetail()(*MicrosoftGraphGetOffice365ActivationsUserDetailRequestBuilder) {
    return NewMicrosoftGraphGetOffice365ActivationsUserDetailRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetOffice365ActiveUserCountsWithPeriod provides operations to call the getOffice365ActiveUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActiveUserCountsWithPeriod(period *string)(*MicrosoftGraphGetOffice365ActiveUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOffice365ActiveUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365ActiveUserDetailWithDate provides operations to call the getOffice365ActiveUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActiveUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetOffice365ActiveUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetOffice365ActiveUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetOffice365ActiveUserDetailWithPeriod provides operations to call the getOffice365ActiveUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActiveUserDetailWithPeriod(period *string)(*MicrosoftGraphGetOffice365ActiveUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOffice365ActiveUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365GroupsActivityCountsWithPeriod provides operations to call the getOffice365GroupsActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityCountsWithPeriod(period *string)(*MicrosoftGraphGetOffice365GroupsActivityCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOffice365GroupsActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365GroupsActivityDetailWithDate provides operations to call the getOffice365GroupsActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetOffice365GroupsActivityDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetOffice365GroupsActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetOffice365GroupsActivityDetailWithPeriod provides operations to call the getOffice365GroupsActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityDetailWithPeriod(period *string)(*MicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriod provides operations to call the getOffice365GroupsActivityFileCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriod(period *string)(*MicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365GroupsActivityGroupCountsWithPeriod provides operations to call the getOffice365GroupsActivityGroupCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityGroupCountsWithPeriod(period *string)(*MicrosoftGraphGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365GroupsActivityStorageWithPeriod provides operations to call the getOffice365GroupsActivityStorage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityStorageWithPeriod(period *string)(*MicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365ServicesUserCountsWithPeriod provides operations to call the getOffice365ServicesUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ServicesUserCountsWithPeriod(period *string)(*MicrosoftGraphGetOffice365ServicesUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOffice365ServicesUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveActivityFileCountsWithPeriod provides operations to call the getOneDriveActivityFileCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveActivityFileCountsWithPeriod(period *string)(*MicrosoftGraphGetOneDriveActivityFileCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOneDriveActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveActivityUserCountsWithPeriod provides operations to call the getOneDriveActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveActivityUserCountsWithPeriod(period *string)(*MicrosoftGraphGetOneDriveActivityUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOneDriveActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveActivityUserDetailWithDate provides operations to call the getOneDriveActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetOneDriveActivityUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetOneDriveActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetOneDriveActivityUserDetailWithPeriod provides operations to call the getOneDriveActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveActivityUserDetailWithPeriod(period *string)(*MicrosoftGraphGetOneDriveActivityUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOneDriveActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveUsageAccountCountsWithPeriod provides operations to call the getOneDriveUsageAccountCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveUsageAccountCountsWithPeriod(period *string)(*MicrosoftGraphGetOneDriveUsageAccountCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOneDriveUsageAccountCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveUsageAccountDetailWithDate provides operations to call the getOneDriveUsageAccountDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveUsageAccountDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetOneDriveUsageAccountDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetOneDriveUsageAccountDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetOneDriveUsageAccountDetailWithPeriod provides operations to call the getOneDriveUsageAccountDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveUsageAccountDetailWithPeriod(period *string)(*MicrosoftGraphGetOneDriveUsageAccountDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOneDriveUsageAccountDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveUsageFileCountsWithPeriod provides operations to call the getOneDriveUsageFileCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveUsageFileCountsWithPeriod(period *string)(*MicrosoftGraphGetOneDriveUsageFileCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOneDriveUsageFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveUsageStorageWithPeriod provides operations to call the getOneDriveUsageStorage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveUsageStorageWithPeriod(period *string)(*MicrosoftGraphGetOneDriveUsageStorageWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOneDriveUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTime provides operations to call the getPrinterArchivedPrintJobs method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, printerId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*MicrosoftGraphGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewMicrosoftGraphGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, printerId, startDateTime)
}
// MicrosoftGraphGetSharePointActivityFileCountsWithPeriod provides operations to call the getSharePointActivityFileCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointActivityFileCountsWithPeriod(period *string)(*MicrosoftGraphGetSharePointActivityFileCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSharePointActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointActivityPagesWithPeriod provides operations to call the getSharePointActivityPages method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointActivityPagesWithPeriod(period *string)(*MicrosoftGraphGetSharePointActivityPagesWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSharePointActivityPagesWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointActivityUserCountsWithPeriod provides operations to call the getSharePointActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointActivityUserCountsWithPeriod(period *string)(*MicrosoftGraphGetSharePointActivityUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSharePointActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointActivityUserDetailWithDate provides operations to call the getSharePointActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetSharePointActivityUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetSharePointActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetSharePointActivityUserDetailWithPeriod provides operations to call the getSharePointActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointActivityUserDetailWithPeriod(period *string)(*MicrosoftGraphGetSharePointActivityUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSharePointActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointSiteUsageDetailWithDate provides operations to call the getSharePointSiteUsageDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsageDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetSharePointSiteUsageDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetSharePointSiteUsageDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetSharePointSiteUsageDetailWithPeriod provides operations to call the getSharePointSiteUsageDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsageDetailWithPeriod(period *string)(*MicrosoftGraphGetSharePointSiteUsageDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSharePointSiteUsageDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointSiteUsageFileCountsWithPeriod provides operations to call the getSharePointSiteUsageFileCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsageFileCountsWithPeriod(period *string)(*MicrosoftGraphGetSharePointSiteUsageFileCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSharePointSiteUsageFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointSiteUsagePagesWithPeriod provides operations to call the getSharePointSiteUsagePages method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsagePagesWithPeriod(period *string)(*MicrosoftGraphGetSharePointSiteUsagePagesWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSharePointSiteUsagePagesWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointSiteUsageSiteCountsWithPeriod provides operations to call the getSharePointSiteUsageSiteCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsageSiteCountsWithPeriod(period *string)(*MicrosoftGraphGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointSiteUsageStorageWithPeriod provides operations to call the getSharePointSiteUsageStorage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsageStorageWithPeriod(period *string)(*MicrosoftGraphGetSharePointSiteUsageStorageWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSharePointSiteUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessActivityCountsWithPeriod provides operations to call the getSkypeForBusinessActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessActivityCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessActivityCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessActivityUserCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessActivityUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessActivityUserDetailWithDate provides operations to call the getSkypeForBusinessActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetSkypeForBusinessActivityUserDetailWithPeriod provides operations to call the getSkypeForBusinessActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessActivityUserDetailWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessActivityUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageDistributionUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessDeviceUsageUserCountsWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessDeviceUsageUserCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithDate provides operations to call the getSkypeForBusinessDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessOrganizerActivityCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessOrganizerActivityCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityMinuteCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessOrganizerActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessOrganizerActivityUserCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityMinuteCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessParticipantActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessParticipantActivityUserCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessParticipantActivityUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessParticipantActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessPeerToPeerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityMinuteCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageDistributionUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageDistributionUserCountsWithPeriod(period *string)(*MicrosoftGraphGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsDeviceUsageUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageUserCountsWithPeriod(period *string)(*MicrosoftGraphGetTeamsDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsDeviceUsageUserDetailWithDate provides operations to call the getTeamsDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetTeamsDeviceUsageUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetTeamsDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetTeamsDeviceUsageUserDetailWithPeriod provides operations to call the getTeamsDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageUserDetailWithPeriod(period *string)(*MicrosoftGraphGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityCountsWithPeriod provides operations to call the getTeamsUserActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityCountsWithPeriod(period *string)(*MicrosoftGraphGetTeamsUserActivityCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsUserActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityUserCountsWithPeriod provides operations to call the getTeamsUserActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityUserCountsWithPeriod(period *string)(*MicrosoftGraphGetTeamsUserActivityUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsUserActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityUserDetailWithDate provides operations to call the getTeamsUserActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetTeamsUserActivityUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetTeamsUserActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetTeamsUserActivityUserDetailWithPeriod provides operations to call the getTeamsUserActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityUserDetailWithPeriod(period *string)(*MicrosoftGraphGetTeamsUserActivityUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsUserActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime provides operations to call the getUserArchivedPrintJobs method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, userId *string)(*MicrosoftGraphGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewMicrosoftGraphGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime, userId)
}
// MicrosoftGraphGetYammerActivityCountsWithPeriod provides operations to call the getYammerActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerActivityCountsWithPeriod(period *string)(*MicrosoftGraphGetYammerActivityCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetYammerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerActivityUserCountsWithPeriod provides operations to call the getYammerActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerActivityUserCountsWithPeriod(period *string)(*MicrosoftGraphGetYammerActivityUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetYammerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerActivityUserDetailWithDate provides operations to call the getYammerActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetYammerActivityUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetYammerActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetYammerActivityUserDetailWithPeriod provides operations to call the getYammerActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerActivityUserDetailWithPeriod(period *string)(*MicrosoftGraphGetYammerActivityUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetYammerActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getYammerDeviceUsageDistributionUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriod(period *string)(*MicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerDeviceUsageUserCountsWithPeriod provides operations to call the getYammerDeviceUsageUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerDeviceUsageUserCountsWithPeriod(period *string)(*MicrosoftGraphGetYammerDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetYammerDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerDeviceUsageUserDetailWithDate provides operations to call the getYammerDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetYammerDeviceUsageUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetYammerDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetYammerDeviceUsageUserDetailWithPeriod provides operations to call the getYammerDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerDeviceUsageUserDetailWithPeriod(period *string)(*MicrosoftGraphGetYammerDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetYammerDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerGroupsActivityCountsWithPeriod provides operations to call the getYammerGroupsActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerGroupsActivityCountsWithPeriod(period *string)(*MicrosoftGraphGetYammerGroupsActivityCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetYammerGroupsActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerGroupsActivityDetailWithDate provides operations to call the getYammerGroupsActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerGroupsActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetYammerGroupsActivityDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetYammerGroupsActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetYammerGroupsActivityDetailWithPeriod provides operations to call the getYammerGroupsActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerGroupsActivityDetailWithPeriod(period *string)(*MicrosoftGraphGetYammerGroupsActivityDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetYammerGroupsActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerGroupsActivityGroupCountsWithPeriod provides operations to call the getYammerGroupsActivityGroupCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerGroupsActivityGroupCountsWithPeriod(period *string)(*MicrosoftGraphGetYammerGroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetYammerGroupsActivityGroupCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphManagedDeviceEnrollmentFailureDetails provides operations to call the managedDeviceEnrollmentFailureDetails method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentFailureDetails()(*MicrosoftGraphManagedDeviceEnrollmentFailureDetailsRequestBuilder) {
    return NewMicrosoftGraphManagedDeviceEnrollmentFailureDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentFailureDetails method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*MicrosoftGraphManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return NewMicrosoftGraphManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top)
}
// MicrosoftGraphManagedDeviceEnrollmentTopFailures provides operations to call the managedDeviceEnrollmentTopFailures method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentTopFailures()(*MicrosoftGraphManagedDeviceEnrollmentTopFailuresRequestBuilder) {
    return NewMicrosoftGraphManagedDeviceEnrollmentTopFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphManagedDeviceEnrollmentTopFailuresWithPeriod provides operations to call the managedDeviceEnrollmentTopFailures method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentTopFailuresWithPeriod(period *string)(*MicrosoftGraphManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) {
    return NewMicrosoftGraphManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MonthlyPrintUsageByPrinter provides operations to manage the monthlyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageByPrinter()(*MonthlyPrintUsageByPrinterRequestBuilder) {
    return NewMonthlyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MonthlyPrintUsageByPrinterById provides operations to manage the monthlyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageByPrinterById(id string)(*MonthlyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return NewMonthlyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MonthlyPrintUsageByUser provides operations to manage the monthlyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageByUser()(*MonthlyPrintUsageByUserRequestBuilder) {
    return NewMonthlyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MonthlyPrintUsageByUserById provides operations to manage the monthlyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageByUserById(id string)(*MonthlyPrintUsageByUserPrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return NewMonthlyPrintUsageByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update reports
func (m *ReportsRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReportRootable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReportRootable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateReportRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReportRootable), nil
}
// Security provides operations to manage the security property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) Security()(*SecurityRequestBuilder) {
    return NewSecurityRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation get reports
func (m *ReportsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReportsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update reports
func (m *ReportsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReportRootable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
