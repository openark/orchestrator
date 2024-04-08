package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// OwnedDevicesMicrosoftGraphEndpointRequestBuilder casts the previous resource to endpoint.
type OwnedDevicesMicrosoftGraphEndpointRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OwnedDevicesMicrosoftGraphEndpointRequestBuilderGetQueryParameters get the items of type microsoft.graph.endpoint in the microsoft.graph.directoryObject collection
type OwnedDevicesMicrosoftGraphEndpointRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// OwnedDevicesMicrosoftGraphEndpointRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OwnedDevicesMicrosoftGraphEndpointRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OwnedDevicesMicrosoftGraphEndpointRequestBuilderGetQueryParameters
}
// NewOwnedDevicesMicrosoftGraphEndpointRequestBuilderInternal instantiates a new MicrosoftGraphEndpointRequestBuilder and sets the default values.
func NewOwnedDevicesMicrosoftGraphEndpointRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OwnedDevicesMicrosoftGraphEndpointRequestBuilder) {
    m := &OwnedDevicesMicrosoftGraphEndpointRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/ownedDevices/microsoft.graph.endpoint{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewOwnedDevicesMicrosoftGraphEndpointRequestBuilder instantiates a new MicrosoftGraphEndpointRequestBuilder and sets the default values.
func NewOwnedDevicesMicrosoftGraphEndpointRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OwnedDevicesMicrosoftGraphEndpointRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOwnedDevicesMicrosoftGraphEndpointRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *OwnedDevicesMicrosoftGraphEndpointRequestBuilder) Count()(*OwnedDevicesMicrosoftGraphEndpointCountRequestBuilder) {
    return NewOwnedDevicesMicrosoftGraphEndpointCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get get the items of type microsoft.graph.endpoint in the microsoft.graph.directoryObject collection
func (m *OwnedDevicesMicrosoftGraphEndpointRequestBuilder) Get(ctx context.Context, requestConfiguration *OwnedDevicesMicrosoftGraphEndpointRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EndpointCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEndpointCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EndpointCollectionResponseable), nil
}
// ToGetRequestInformation get the items of type microsoft.graph.endpoint in the microsoft.graph.directoryObject collection
func (m *OwnedDevicesMicrosoftGraphEndpointRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OwnedDevicesMicrosoftGraphEndpointRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
