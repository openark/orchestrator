package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilder provides operations to manage the windowsHelloForBusinessMethods property of the microsoft.graph.authentication entity.
type ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderGetQueryParameters represents the Windows Hello for Business authentication method registered to a user for authentication.
type ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderGetQueryParameters
}
// NewItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderInternal instantiates a new WindowsHelloForBusinessAuthenticationMethodItemRequestBuilder and sets the default values.
func NewItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilder) {
    m := &ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilder instantiates a new WindowsHelloForBusinessAuthenticationMethodItemRequestBuilder and sets the default values.
func NewItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property windowsHelloForBusinessMethods for users
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Device provides operations to manage the device property of the microsoft.graph.windowsHelloForBusinessAuthenticationMethod entity.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilder) Device()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get represents the Windows Hello for Business authentication method registered to a user for authentication.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsHelloForBusinessAuthenticationMethodable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsHelloForBusinessAuthenticationMethodFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsHelloForBusinessAuthenticationMethodable), nil
}
// ToDeleteRequestInformation delete navigation property windowsHelloForBusinessMethods for users
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation represents the Windows Hello for Business authentication method registered to a user for authentication.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
