package organization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMicrosoftGraphSetMobileDeviceManagementAuthorityRequestBuilder provides operations to call the setMobileDeviceManagementAuthority method.
type ItemMicrosoftGraphSetMobileDeviceManagementAuthorityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMicrosoftGraphSetMobileDeviceManagementAuthorityRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMicrosoftGraphSetMobileDeviceManagementAuthorityRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMicrosoftGraphSetMobileDeviceManagementAuthorityRequestBuilderInternal instantiates a new MicrosoftGraphSetMobileDeviceManagementAuthorityRequestBuilder and sets the default values.
func NewItemMicrosoftGraphSetMobileDeviceManagementAuthorityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphSetMobileDeviceManagementAuthorityRequestBuilder) {
    m := &ItemMicrosoftGraphSetMobileDeviceManagementAuthorityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/organization/{organization%2Did}/microsoft.graph.setMobileDeviceManagementAuthority";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMicrosoftGraphSetMobileDeviceManagementAuthorityRequestBuilder instantiates a new MicrosoftGraphSetMobileDeviceManagementAuthorityRequestBuilder and sets the default values.
func NewItemMicrosoftGraphSetMobileDeviceManagementAuthorityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphSetMobileDeviceManagementAuthorityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMicrosoftGraphSetMobileDeviceManagementAuthorityRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set mobile device management authority
func (m *ItemMicrosoftGraphSetMobileDeviceManagementAuthorityRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemMicrosoftGraphSetMobileDeviceManagementAuthorityRequestBuilderPostRequestConfiguration)(ItemMicrosoftGraphSetMobileDeviceManagementAuthoritySetMobileDeviceManagementAuthorityResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemMicrosoftGraphSetMobileDeviceManagementAuthoritySetMobileDeviceManagementAuthorityResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemMicrosoftGraphSetMobileDeviceManagementAuthoritySetMobileDeviceManagementAuthorityResponseable), nil
}
// ToPostRequestInformation set mobile device management authority
func (m *ItemMicrosoftGraphSetMobileDeviceManagementAuthorityRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemMicrosoftGraphSetMobileDeviceManagementAuthorityRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
