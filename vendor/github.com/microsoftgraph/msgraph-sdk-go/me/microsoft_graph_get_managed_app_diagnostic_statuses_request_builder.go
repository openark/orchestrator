package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder provides operations to call the getManagedAppDiagnosticStatuses method.
type MicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderGetQueryParameters gets diagnostics validation status for a given user.
type MicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderGetQueryParameters struct {
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
// MicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderGetQueryParameters
}
// NewMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderInternal instantiates a new MicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder and sets the default values.
func NewMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder) {
    m := &MicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/microsoft.graph.getManagedAppDiagnosticStatuses(){?%24top,%24skip,%24search,%24filter,%24count}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder instantiates a new MicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder and sets the default values.
func NewMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets diagnostics validation status for a given user.
func (m *MicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderGetRequestConfiguration)(MicrosoftGraphGetManagedAppDiagnosticStatusesGetManagedAppDiagnosticStatusesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateMicrosoftGraphGetManagedAppDiagnosticStatusesGetManagedAppDiagnosticStatusesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MicrosoftGraphGetManagedAppDiagnosticStatusesGetManagedAppDiagnosticStatusesResponseable), nil
}
// ToGetRequestInformation gets diagnostics validation status for a given user.
func (m *MicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosoftGraphGetManagedAppDiagnosticStatusesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
