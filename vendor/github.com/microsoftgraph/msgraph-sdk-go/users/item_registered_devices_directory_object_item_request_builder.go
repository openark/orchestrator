package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemRegisteredDevicesDirectoryObjectItemRequestBuilder provides operations to manage the registeredDevices property of the microsoft.graph.user entity.
type ItemRegisteredDevicesDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemRegisteredDevicesDirectoryObjectItemRequestBuilderGetQueryParameters devices that are registered for the user. Read-only. Nullable. Supports $expand.
type ItemRegisteredDevicesDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemRegisteredDevicesDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRegisteredDevicesDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemRegisteredDevicesDirectoryObjectItemRequestBuilderGetQueryParameters
}
// NewItemRegisteredDevicesDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemRegisteredDevicesDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredDevicesDirectoryObjectItemRequestBuilder) {
    m := &ItemRegisteredDevicesDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/registeredDevices/{directoryObject%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemRegisteredDevicesDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemRegisteredDevicesDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredDevicesDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRegisteredDevicesDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get devices that are registered for the user. Read-only. Nullable. Supports $expand.
func (m *ItemRegisteredDevicesDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemRegisteredDevicesDirectoryObjectItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable), nil
}
// MicrosoftGraphAppRoleAssignment casts the previous resource to appRoleAssignment.
func (m *ItemRegisteredDevicesDirectoryObjectItemRequestBuilder) MicrosoftGraphAppRoleAssignment()(*ItemRegisteredDevicesItemMicrosoftGraphAppRoleAssignmentRequestBuilder) {
    return NewItemRegisteredDevicesItemMicrosoftGraphAppRoleAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDevice casts the previous resource to device.
func (m *ItemRegisteredDevicesDirectoryObjectItemRequestBuilder) MicrosoftGraphDevice()(*ItemRegisteredDevicesItemMicrosoftGraphDeviceRequestBuilder) {
    return NewItemRegisteredDevicesItemMicrosoftGraphDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEndpoint casts the previous resource to endpoint.
func (m *ItemRegisteredDevicesDirectoryObjectItemRequestBuilder) MicrosoftGraphEndpoint()(*ItemRegisteredDevicesItemMicrosoftGraphEndpointRequestBuilder) {
    return NewItemRegisteredDevicesItemMicrosoftGraphEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation devices that are registered for the user. Read-only. Nullable. Supports $expand.
func (m *ItemRegisteredDevicesDirectoryObjectItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemRegisteredDevicesDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
