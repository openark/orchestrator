package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemManagedDevicesManagedDeviceItemRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.user entity.
type ItemManagedDevicesManagedDeviceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters the managed devices associated with the user.
type ItemManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters
}
// ItemManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManagedDevicesManagedDeviceItemRequestBuilderInternal instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewItemManagedDevicesManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesManagedDeviceItemRequestBuilder) {
    m := &ItemManagedDevicesManagedDeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemManagedDevicesManagedDeviceItemRequestBuilder instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewItemManagedDevicesManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property managedDevices for users
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceCategory()(*ItemManagedDevicesItemDeviceCategoryRequestBuilder) {
    return NewItemManagedDevicesItemDeviceCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceCompliancePolicyStates provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*ItemManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilder) {
    return NewItemManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceCompliancePolicyStatesById provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStatesById(id string)(*ItemManagedDevicesItemDeviceCompliancePolicyStatesDeviceCompliancePolicyStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyState%2Did"] = id
    }
    return NewItemManagedDevicesItemDeviceCompliancePolicyStatesDeviceCompliancePolicyStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DeviceConfigurationStates provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*ItemManagedDevicesItemDeviceConfigurationStatesRequestBuilder) {
    return NewItemManagedDevicesItemDeviceConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceConfigurationStatesById provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStatesById(id string)(*ItemManagedDevicesItemDeviceConfigurationStatesDeviceConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationState%2Did"] = id
    }
    return NewItemManagedDevicesItemDeviceConfigurationStatesDeviceConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the managed devices associated with the user.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable, error) {
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
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphBypassActivationLock()(*ItemManagedDevicesItemMicrosoftGraphBypassActivationLockRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCleanWindowsDevice provides operations to call the cleanWindowsDevice method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphCleanWindowsDevice()(*ItemManagedDevicesItemMicrosoftGraphCleanWindowsDeviceRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDeleteUserFromSharedAppleDevice provides operations to call the deleteUserFromSharedAppleDevice method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphDeleteUserFromSharedAppleDevice()(*ItemManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDisableLostMode provides operations to call the disableLostMode method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphDisableLostMode()(*ItemManagedDevicesItemMicrosoftGraphDisableLostModeRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLocateDevice provides operations to call the locateDevice method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphLocateDevice()(*ItemManagedDevicesItemMicrosoftGraphLocateDeviceRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphLogoutSharedAppleDeviceActiveUser()(*ItemManagedDevicesItemMicrosoftGraphLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRebootNow provides operations to call the rebootNow method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRebootNow()(*ItemManagedDevicesItemMicrosoftGraphRebootNowRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRecoverPasscode provides operations to call the recoverPasscode method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRecoverPasscode()(*ItemManagedDevicesItemMicrosoftGraphRecoverPasscodeRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRemoteLock provides operations to call the remoteLock method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRemoteLock()(*ItemManagedDevicesItemMicrosoftGraphRemoteLockRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRequestRemoteAssistance()(*ItemManagedDevicesItemMicrosoftGraphRequestRemoteAssistanceRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphResetPasscode provides operations to call the resetPasscode method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphResetPasscode()(*ItemManagedDevicesItemMicrosoftGraphResetPasscodeRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRetire provides operations to call the retire method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRetire()(*ItemManagedDevicesItemMicrosoftGraphRetireRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphShutDown provides operations to call the shutDown method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphShutDown()(*ItemManagedDevicesItemMicrosoftGraphShutDownRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSyncDevice provides operations to call the syncDevice method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphSyncDevice()(*ItemManagedDevicesItemMicrosoftGraphSyncDeviceRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUpdateWindowsDeviceAccount provides operations to call the updateWindowsDeviceAccount method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphUpdateWindowsDeviceAccount()(*ItemManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWindowsDefenderScan provides operations to call the windowsDefenderScan method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphWindowsDefenderScan()(*ItemManagedDevicesItemMicrosoftGraphWindowsDefenderScanRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphWindowsDefenderUpdateSignatures()(*ItemManagedDevicesItemMicrosoftGraphWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWipe provides operations to call the wipe method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphWipe()(*ItemManagedDevicesItemMicrosoftGraphWipeRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property managedDevices in users
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable, error) {
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
// ToDeleteRequestInformation delete navigation property managedDevices for users
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the managed devices associated with the user.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property managedDevices in users
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceable, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Users()(*ItemManagedDevicesItemUsersRequestBuilder) {
    return NewItemManagedDevicesItemUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
