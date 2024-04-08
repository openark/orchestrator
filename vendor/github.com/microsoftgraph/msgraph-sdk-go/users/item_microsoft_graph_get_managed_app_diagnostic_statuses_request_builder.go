package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder provides operations to call the getManagedAppDiagnosticStatuses method.
type ItemMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderGetQueryParameters gets diagnostics validation status for a given user.
type ItemMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderGetQueryParameters struct {
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
// ItemMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderGetQueryParameters
}
// NewItemMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderInternal instantiates a new MicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder and sets the default values.
func NewItemMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder) {
    m := &ItemMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/microsoft.graph.getManagedAppDiagnosticStatuses(){?%24top,%24skip,%24search,%24filter,%24count}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder instantiates a new MicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder and sets the default values.
func NewItemMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets diagnostics validation status for a given user.
func (m *ItemMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderGetRequestConfiguration)(ItemMicrosoftGraphGetManagedAppDiagnosticStatusesGetManagedAppDiagnosticStatusesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemMicrosoftGraphGetManagedAppDiagnosticStatusesGetManagedAppDiagnosticStatusesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemMicrosoftGraphGetManagedAppDiagnosticStatusesGetManagedAppDiagnosticStatusesResponseable), nil
}
// ToGetRequestInformation gets diagnostics validation status for a given user.
func (m *ItemMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
