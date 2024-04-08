package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder provides operations to manage the mobileAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
type MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderGetQueryParameters the Managed Device Mobile Application Configurations.
type MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderGetQueryParameters
}
// MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) Assignments()(*MobileAppConfigurationsItemAssignmentsRequestBuilder) {
    return NewMobileAppConfigurationsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) AssignmentsById(id string)(*MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationAssignment%2Did"] = id
    }
    return NewMobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewMobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderInternal instantiates a new ManagedDeviceMobileAppConfigurationItemRequestBuilder and sets the default values.
func NewMobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) {
    m := &MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileAppConfigurations/{managedDeviceMobileAppConfiguration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder instantiates a new ManagedDeviceMobileAppConfigurationItemRequestBuilder and sets the default values.
func NewMobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property mobileAppConfigurations for deviceAppManagement
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceStatuses provides operations to manage the deviceStatuses property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) DeviceStatuses()(*MobileAppConfigurationsItemDeviceStatusesRequestBuilder) {
    return NewMobileAppConfigurationsItemDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceStatusesById provides operations to manage the deviceStatuses property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) DeviceStatusesById(id string)(*MobileAppConfigurationsItemDeviceStatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationDeviceStatus%2Did"] = id
    }
    return NewMobileAppConfigurationsItemDeviceStatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DeviceStatusSummary provides operations to manage the deviceStatusSummary property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) DeviceStatusSummary()(*MobileAppConfigurationsItemDeviceStatusSummaryRequestBuilder) {
    return NewMobileAppConfigurationsItemDeviceStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get the Managed Device Mobile Application Configurations.
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedDeviceMobileAppConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationable), nil
}
// MicrosoftGraphAssign provides operations to call the assign method.
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) MicrosoftGraphAssign()(*MobileAppConfigurationsItemMicrosoftGraphAssignRequestBuilder) {
    return NewMobileAppConfigurationsItemMicrosoftGraphAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property mobileAppConfigurations in deviceAppManagement
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationable, requestConfiguration *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedDeviceMobileAppConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationable), nil
}
// ToDeleteRequestInformation delete navigation property mobileAppConfigurations for deviceAppManagement
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the Managed Device Mobile Application Configurations.
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property mobileAppConfigurations in deviceAppManagement
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationable, requestConfiguration *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserStatuses provides operations to manage the userStatuses property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) UserStatuses()(*MobileAppConfigurationsItemUserStatusesRequestBuilder) {
    return NewMobileAppConfigurationsItemUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserStatusesById provides operations to manage the userStatuses property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) UserStatusesById(id string)(*MobileAppConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationUserStatus%2Did"] = id
    }
    return NewMobileAppConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// UserStatusSummary provides operations to manage the userStatusSummary property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) UserStatusSummary()(*MobileAppConfigurationsItemUserStatusSummaryRequestBuilder) {
    return NewMobileAppConfigurationsItemUserStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
