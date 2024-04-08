package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// IdentityProvidersMicrosoftGraphAvailableProviderTypesRequestBuilder provides operations to call the availableProviderTypes method.
type IdentityProvidersMicrosoftGraphAvailableProviderTypesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityProvidersMicrosoftGraphAvailableProviderTypesRequestBuilderGetQueryParameters invoke function availableProviderTypes
type IdentityProvidersMicrosoftGraphAvailableProviderTypesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// IdentityProvidersMicrosoftGraphAvailableProviderTypesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityProvidersMicrosoftGraphAvailableProviderTypesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityProvidersMicrosoftGraphAvailableProviderTypesRequestBuilderGetQueryParameters
}
// NewIdentityProvidersMicrosoftGraphAvailableProviderTypesRequestBuilderInternal instantiates a new MicrosoftGraphAvailableProviderTypesRequestBuilder and sets the default values.
func NewIdentityProvidersMicrosoftGraphAvailableProviderTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityProvidersMicrosoftGraphAvailableProviderTypesRequestBuilder) {
    m := &IdentityProvidersMicrosoftGraphAvailableProviderTypesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/identityProviders/microsoft.graph.availableProviderTypes(){?%24top,%24skip,%24search,%24filter,%24count}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewIdentityProvidersMicrosoftGraphAvailableProviderTypesRequestBuilder instantiates a new MicrosoftGraphAvailableProviderTypesRequestBuilder and sets the default values.
func NewIdentityProvidersMicrosoftGraphAvailableProviderTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityProvidersMicrosoftGraphAvailableProviderTypesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityProvidersMicrosoftGraphAvailableProviderTypesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function availableProviderTypes
func (m *IdentityProvidersMicrosoftGraphAvailableProviderTypesRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityProvidersMicrosoftGraphAvailableProviderTypesRequestBuilderGetRequestConfiguration)(IdentityProvidersMicrosoftGraphAvailableProviderTypesAvailableProviderTypesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateIdentityProvidersMicrosoftGraphAvailableProviderTypesAvailableProviderTypesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(IdentityProvidersMicrosoftGraphAvailableProviderTypesAvailableProviderTypesResponseable), nil
}
// ToGetRequestInformation invoke function availableProviderTypes
func (m *IdentityProvidersMicrosoftGraphAvailableProviderTypesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IdentityProvidersMicrosoftGraphAvailableProviderTypesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
