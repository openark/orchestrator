package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ReportsRequestBuilder provides operations to manage the reports property of the microsoft.graph.deviceManagement entity.
type ReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReportsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ReportsRequestBuilderGetQueryParameters reports singleton
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
    m.urlTemplate = "{+baseurl}/deviceManagement/reports{?%24select,%24expand}";
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
// Delete delete navigation property reports for deviceManagement
func (m *ReportsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ReportsRequestBuilderDeleteRequestConfiguration)(error) {
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
// ExportJobs provides operations to manage the exportJobs property of the microsoft.graph.deviceManagementReports entity.
func (m *ReportsRequestBuilder) ExportJobs()(*ReportsExportJobsRequestBuilder) {
    return NewReportsExportJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExportJobsById provides operations to manage the exportJobs property of the microsoft.graph.deviceManagementReports entity.
func (m *ReportsRequestBuilder) ExportJobsById(id string)(*ReportsExportJobsDeviceManagementExportJobItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementExportJob%2Did"] = id
    }
    return NewReportsExportJobsDeviceManagementExportJobItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get reports singleton
func (m *ReportsRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementReportsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceManagementReportsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementReportsable), nil
}
// MicrosoftGraphGetCachedReport provides operations to call the getCachedReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetCachedReport()(*ReportsMicrosoftGraphGetCachedReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetCachedReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCompliancePolicyNonComplianceReport provides operations to call the getCompliancePolicyNonComplianceReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetCompliancePolicyNonComplianceReport()(*ReportsMicrosoftGraphGetCompliancePolicyNonComplianceReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetCompliancePolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCompliancePolicyNonComplianceSummaryReport provides operations to call the getCompliancePolicyNonComplianceSummaryReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetCompliancePolicyNonComplianceSummaryReport()(*ReportsMicrosoftGraphGetCompliancePolicyNonComplianceSummaryReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetCompliancePolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetComplianceSettingNonComplianceReport provides operations to call the getComplianceSettingNonComplianceReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetComplianceSettingNonComplianceReport()(*ReportsMicrosoftGraphGetComplianceSettingNonComplianceReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetComplianceSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetConfigurationPolicyNonComplianceReport provides operations to call the getConfigurationPolicyNonComplianceReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetConfigurationPolicyNonComplianceReport()(*ReportsMicrosoftGraphGetConfigurationPolicyNonComplianceReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetConfigurationPolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetConfigurationPolicyNonComplianceSummaryReport provides operations to call the getConfigurationPolicyNonComplianceSummaryReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetConfigurationPolicyNonComplianceSummaryReport()(*ReportsMicrosoftGraphGetConfigurationPolicyNonComplianceSummaryReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetConfigurationPolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetConfigurationSettingNonComplianceReport provides operations to call the getConfigurationSettingNonComplianceReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetConfigurationSettingNonComplianceReport()(*ReportsMicrosoftGraphGetConfigurationSettingNonComplianceReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetConfigurationSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfiles provides operations to call the getDeviceManagementIntentPerSettingContributingProfiles method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfiles()(*ReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder) {
    return NewReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetDeviceManagementIntentSettingsReport provides operations to call the getDeviceManagementIntentSettingsReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetDeviceManagementIntentSettingsReport()(*ReportsMicrosoftGraphGetDeviceManagementIntentSettingsReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetDeviceManagementIntentSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetDeviceNonComplianceReport provides operations to call the getDeviceNonComplianceReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetDeviceNonComplianceReport()(*ReportsMicrosoftGraphGetDeviceNonComplianceReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetDeviceNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetDevicesWithoutCompliancePolicyReport provides operations to call the getDevicesWithoutCompliancePolicyReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetDevicesWithoutCompliancePolicyReport()(*ReportsMicrosoftGraphGetDevicesWithoutCompliancePolicyReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetDevicesWithoutCompliancePolicyReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetHistoricalReport provides operations to call the getHistoricalReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetHistoricalReport()(*ReportsMicrosoftGraphGetHistoricalReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetHistoricalReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetNoncompliantDevicesAndSettingsReport provides operations to call the getNoncompliantDevicesAndSettingsReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetNoncompliantDevicesAndSettingsReport()(*ReportsMicrosoftGraphGetNoncompliantDevicesAndSettingsReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetNoncompliantDevicesAndSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetPolicyNonComplianceMetadata provides operations to call the getPolicyNonComplianceMetadata method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetPolicyNonComplianceMetadata()(*ReportsMicrosoftGraphGetPolicyNonComplianceMetadataRequestBuilder) {
    return NewReportsMicrosoftGraphGetPolicyNonComplianceMetadataRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetPolicyNonComplianceReport provides operations to call the getPolicyNonComplianceReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetPolicyNonComplianceReport()(*ReportsMicrosoftGraphGetPolicyNonComplianceReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetPolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetPolicyNonComplianceSummaryReport provides operations to call the getPolicyNonComplianceSummaryReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetPolicyNonComplianceSummaryReport()(*ReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetReportFilters provides operations to call the getReportFilters method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetReportFilters()(*ReportsMicrosoftGraphGetReportFiltersRequestBuilder) {
    return NewReportsMicrosoftGraphGetReportFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetSettingNonComplianceReport provides operations to call the getSettingNonComplianceReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSettingNonComplianceReport()(*ReportsMicrosoftGraphGetSettingNonComplianceReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property reports in deviceManagement
func (m *ReportsRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementReportsable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementReportsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceManagementReportsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementReportsable), nil
}
// ToDeleteRequestInformation delete navigation property reports for deviceManagement
func (m *ReportsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ReportsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation reports singleton
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
// ToPatchRequestInformation update the navigation property reports in deviceManagement
func (m *ReportsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementReportsable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
