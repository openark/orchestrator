package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesRequestBuilder provides operations to call the updateDeviceProperties method.
type WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesRequestBuilderInternal instantiates a new MicrosoftGraphUpdateDevicePropertiesRequestBuilder and sets the default values.
func NewWindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesRequestBuilder) {
    m := &WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsAutopilotDeviceIdentities/{windowsAutopilotDeviceIdentity%2Did}/microsoft.graph.updateDeviceProperties";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewWindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesRequestBuilder instantiates a new MicrosoftGraphUpdateDevicePropertiesRequestBuilder and sets the default values.
func NewWindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post updates properties on Autopilot devices.
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesRequestBuilder) Post(ctx context.Context, body WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBodyable, requestConfiguration *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation updates properties on Autopilot devices.
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBodyable, requestConfiguration *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
