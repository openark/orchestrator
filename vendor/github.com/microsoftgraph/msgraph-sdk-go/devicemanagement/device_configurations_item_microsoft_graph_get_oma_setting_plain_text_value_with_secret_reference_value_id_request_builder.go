package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeviceConfigurationsItemMicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder provides operations to call the getOmaSettingPlainTextValue method.
type DeviceConfigurationsItemMicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceConfigurationsItemMicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceConfigurationsItemMicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceConfigurationsItemMicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal instantiates a new MicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder and sets the default values.
func NewDeviceConfigurationsItemMicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, secretReferenceValueId *string)(*DeviceConfigurationsItemMicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) {
    m := &DeviceConfigurationsItemMicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration%2Did}/microsoft.graph.getOmaSettingPlainTextValue(secretReferenceValueId='{secretReferenceValueId}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if secretReferenceValueId != nil {
        urlTplParams["secretReferenceValueId"] = *secretReferenceValueId
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDeviceConfigurationsItemMicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder instantiates a new MicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder and sets the default values.
func NewDeviceConfigurationsItemMicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceConfigurationsItemMicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceConfigurationsItemMicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getOmaSettingPlainTextValue
func (m *DeviceConfigurationsItemMicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceConfigurationsItemMicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderGetRequestConfiguration)(DeviceConfigurationsItemMicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdGetOmaSettingPlainTextValueWithSecretReferenceValueIdResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateDeviceConfigurationsItemMicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdGetOmaSettingPlainTextValueWithSecretReferenceValueIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceConfigurationsItemMicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdGetOmaSettingPlainTextValueWithSecretReferenceValueIdResponseable), nil
}
// ToGetRequestInformation invoke function getOmaSettingPlainTextValue
func (m *DeviceConfigurationsItemMicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceConfigurationsItemMicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
