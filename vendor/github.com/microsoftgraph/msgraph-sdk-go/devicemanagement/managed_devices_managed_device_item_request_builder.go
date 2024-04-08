package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ManagedDevicesManagedDeviceItemRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.deviceManagement entity.
type ManagedDevicesManagedDeviceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters the list of managed devices.
type ManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters
}
// ManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedDevicesManagedDeviceItemRequestBuilderInternal instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewManagedDevicesManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesManagedDeviceItemRequestBuilder) {
    m := &ManagedDevicesManagedDeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewManagedDevicesManagedDeviceItemRequestBuilder instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewManagedDevicesManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property managedDevices for deviceManagement
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceCategory provides operations to manage the deviceCategory property of the microsoft.graph.managedDevice entity.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) DeviceCategory()(*ManagedDevicesItemDeviceCategoryRequestBuilder) {
    return NewManagedDevicesItemDeviceCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceCompliancePolicyStates provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*ManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilder) {
    return NewManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceCompliancePolicyStatesById provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStatesById(id string)(*ManagedDevicesItemDeviceCompliancePolicyStatesDeviceCompliancePolicyStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyState%2Did"] = id
    }
    return NewManagedDevicesItemDeviceCompliancePolicyStatesDeviceCompliancePolicyStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DeviceConfigurationStates provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*ManagedDevicesItemDeviceConfigurationStatesRequestBuilder) {
    return NewManagedDevicesItemDeviceConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceConfigurationStatesById provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStatesById(id string)(*ManagedDevicesItemDeviceConfigurationStatesDeviceConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationState%2Did"] = id
    }
    return NewManagedDevicesItemDeviceConfigurationStatesDeviceConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the list of managed devices.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable), nil
}
// MicrosoftGraphBypassActivationLock provides operations to call the bypassActivationLock method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphBypassActivationLock()(*ManagedDevicesItemMicrosoftGraphBypassActivationLockRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCleanWindowsDevice provides operations to call the cleanWindowsDevice method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphCleanWindowsDevice()(*ManagedDevicesItemMicrosoftGraphCleanWindowsDeviceRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDeleteUserFromSharedAppleDevice provides operations to call the deleteUserFromSharedAppleDevice method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphDeleteUserFromSharedAppleDevice()(*ManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDisableLostMode provides operations to call the disableLostMode method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphDisableLostMode()(*ManagedDevicesItemMicrosoftGraphDisableLostModeRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLocateDevice provides operations to call the locateDevice method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphLocateDevice()(*ManagedDevicesItemMicrosoftGraphLocateDeviceRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphLogoutSharedAppleDeviceActiveUser()(*ManagedDevicesItemMicrosoftGraphLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRebootNow provides operations to call the rebootNow method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRebootNow()(*ManagedDevicesItemMicrosoftGraphRebootNowRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRecoverPasscode provides operations to call the recoverPasscode method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRecoverPasscode()(*ManagedDevicesItemMicrosoftGraphRecoverPasscodeRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRemoteLock provides operations to call the remoteLock method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRemoteLock()(*ManagedDevicesItemMicrosoftGraphRemoteLockRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRequestRemoteAssistance()(*ManagedDevicesItemMicrosoftGraphRequestRemoteAssistanceRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphResetPasscode provides operations to call the resetPasscode method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphResetPasscode()(*ManagedDevicesItemMicrosoftGraphResetPasscodeRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRetire provides operations to call the retire method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRetire()(*ManagedDevicesItemMicrosoftGraphRetireRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphShutDown provides operations to call the shutDown method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphShutDown()(*ManagedDevicesItemMicrosoftGraphShutDownRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSyncDevice provides operations to call the syncDevice method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphSyncDevice()(*ManagedDevicesItemMicrosoftGraphSyncDeviceRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUpdateWindowsDeviceAccount provides operations to call the updateWindowsDeviceAccount method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphUpdateWindowsDeviceAccount()(*ManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWindowsDefenderScan provides operations to call the windowsDefenderScan method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphWindowsDefenderScan()(*ManagedDevicesItemMicrosoftGraphWindowsDefenderScanRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphWindowsDefenderUpdateSignatures()(*ManagedDevicesItemMicrosoftGraphWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWipe provides operations to call the wipe method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphWipe()(*ManagedDevicesItemMicrosoftGraphWipeRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property managedDevices in deviceManagement
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable, requestConfiguration *ManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable), nil
}
// ToDeleteRequestInformation delete navigation property managedDevices for deviceManagement
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the list of managed devices.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property managedDevices in deviceManagement
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable, requestConfiguration *ManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Users provides operations to manage the users property of the microsoft.graph.managedDevice entity.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) Users()(*ManagedDevicesItemUsersRequestBuilder) {
    return NewManagedDevicesItemUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
