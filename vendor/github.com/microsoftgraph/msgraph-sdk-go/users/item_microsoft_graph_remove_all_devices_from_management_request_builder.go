package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMicrosoftGraphRemoveAllDevicesFromManagementRequestBuilder provides operations to call the removeAllDevicesFromManagement method.
type ItemMicrosoftGraphRemoveAllDevicesFromManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMicrosoftGraphRemoveAllDevicesFromManagementRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMicrosoftGraphRemoveAllDevicesFromManagementRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMicrosoftGraphRemoveAllDevicesFromManagementRequestBuilderInternal instantiates a new MicrosoftGraphRemoveAllDevicesFromManagementRequestBuilder and sets the default values.
func NewItemMicrosoftGraphRemoveAllDevicesFromManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphRemoveAllDevicesFromManagementRequestBuilder) {
    m := &ItemMicrosoftGraphRemoveAllDevicesFromManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/microsoft.graph.removeAllDevicesFromManagement";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMicrosoftGraphRemoveAllDevicesFromManagementRequestBuilder instantiates a new MicrosoftGraphRemoveAllDevicesFromManagementRequestBuilder and sets the default values.
func NewItemMicrosoftGraphRemoveAllDevicesFromManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphRemoveAllDevicesFromManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMicrosoftGraphRemoveAllDevicesFromManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// Post retire all devices from management for this user
func (m *ItemMicrosoftGraphRemoveAllDevicesFromManagementRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemMicrosoftGraphRemoveAllDevicesFromManagementRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation retire all devices from management for this user
func (m *ItemMicrosoftGraphRemoveAllDevicesFromManagementRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemMicrosoftGraphRemoveAllDevicesFromManagementRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
