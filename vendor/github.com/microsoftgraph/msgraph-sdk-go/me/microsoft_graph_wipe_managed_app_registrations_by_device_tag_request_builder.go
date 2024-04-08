package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilder provides operations to call the wipeManagedAppRegistrationsByDeviceTag method.
type MicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal instantiates a new MicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilder and sets the default values.
func NewMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    m := &MicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/microsoft.graph.wipeManagedAppRegistrationsByDeviceTag";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilder instantiates a new MicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilder and sets the default values.
func NewMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(urlParams, requestAdapter)
}
// Post issues a wipe operation on an app registration with specified device tag.
func (m *MicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilder) Post(ctx context.Context, body MicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBodyable, requestConfiguration *MicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation issues a wipe operation on an app registration with specified device tag.
func (m *MicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilder) ToPostRequestInformation(ctx context.Context, body MicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBodyable, requestConfiguration *MicrosoftGraphWipeManagedAppRegistrationsByDeviceTagRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
