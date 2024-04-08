package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ImportedWindowsAutopilotDeviceIdentitiesMicrosoftGraphImportRequestBuilder provides operations to call the import method.
type ImportedWindowsAutopilotDeviceIdentitiesMicrosoftGraphImportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ImportedWindowsAutopilotDeviceIdentitiesMicrosoftGraphImportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImportedWindowsAutopilotDeviceIdentitiesMicrosoftGraphImportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewImportedWindowsAutopilotDeviceIdentitiesMicrosoftGraphImportRequestBuilderInternal instantiates a new MicrosoftGraphImportRequestBuilder and sets the default values.
func NewImportedWindowsAutopilotDeviceIdentitiesMicrosoftGraphImportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportedWindowsAutopilotDeviceIdentitiesMicrosoftGraphImportRequestBuilder) {
    m := &ImportedWindowsAutopilotDeviceIdentitiesMicrosoftGraphImportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/importedWindowsAutopilotDeviceIdentities/microsoft.graph.import";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewImportedWindowsAutopilotDeviceIdentitiesMicrosoftGraphImportRequestBuilder instantiates a new MicrosoftGraphImportRequestBuilder and sets the default values.
func NewImportedWindowsAutopilotDeviceIdentitiesMicrosoftGraphImportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportedWindowsAutopilotDeviceIdentitiesMicrosoftGraphImportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImportedWindowsAutopilotDeviceIdentitiesMicrosoftGraphImportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action import
func (m *ImportedWindowsAutopilotDeviceIdentitiesMicrosoftGraphImportRequestBuilder) Post(ctx context.Context, body ImportedWindowsAutopilotDeviceIdentitiesMicrosoftGraphImportImportPostRequestBodyable, requestConfiguration *ImportedWindowsAutopilotDeviceIdentitiesMicrosoftGraphImportRequestBuilderPostRequestConfiguration)(ImportedWindowsAutopilotDeviceIdentitiesMicrosoftGraphImportImportResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateImportedWindowsAutopilotDeviceIdentitiesMicrosoftGraphImportImportResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ImportedWindowsAutopilotDeviceIdentitiesMicrosoftGraphImportImportResponseable), nil
}
// ToPostRequestInformation invoke action import
func (m *ImportedWindowsAutopilotDeviceIdentitiesMicrosoftGraphImportRequestBuilder) ToPostRequestInformation(ctx context.Context, body ImportedWindowsAutopilotDeviceIdentitiesMicrosoftGraphImportImportPostRequestBodyable, requestConfiguration *ImportedWindowsAutopilotDeviceIdentitiesMicrosoftGraphImportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
