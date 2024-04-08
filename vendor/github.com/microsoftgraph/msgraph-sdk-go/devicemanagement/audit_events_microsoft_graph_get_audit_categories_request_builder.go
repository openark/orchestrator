package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuditEventsMicrosoftGraphGetAuditCategoriesRequestBuilder provides operations to call the getAuditCategories method.
type AuditEventsMicrosoftGraphGetAuditCategoriesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuditEventsMicrosoftGraphGetAuditCategoriesRequestBuilderGetQueryParameters invoke function getAuditCategories
type AuditEventsMicrosoftGraphGetAuditCategoriesRequestBuilderGetQueryParameters struct {
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
// AuditEventsMicrosoftGraphGetAuditCategoriesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuditEventsMicrosoftGraphGetAuditCategoriesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuditEventsMicrosoftGraphGetAuditCategoriesRequestBuilderGetQueryParameters
}
// NewAuditEventsMicrosoftGraphGetAuditCategoriesRequestBuilderInternal instantiates a new MicrosoftGraphGetAuditCategoriesRequestBuilder and sets the default values.
func NewAuditEventsMicrosoftGraphGetAuditCategoriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuditEventsMicrosoftGraphGetAuditCategoriesRequestBuilder) {
    m := &AuditEventsMicrosoftGraphGetAuditCategoriesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/auditEvents/microsoft.graph.getAuditCategories(){?%24top,%24skip,%24search,%24filter,%24count}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAuditEventsMicrosoftGraphGetAuditCategoriesRequestBuilder instantiates a new MicrosoftGraphGetAuditCategoriesRequestBuilder and sets the default values.
func NewAuditEventsMicrosoftGraphGetAuditCategoriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuditEventsMicrosoftGraphGetAuditCategoriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuditEventsMicrosoftGraphGetAuditCategoriesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getAuditCategories
func (m *AuditEventsMicrosoftGraphGetAuditCategoriesRequestBuilder) Get(ctx context.Context, requestConfiguration *AuditEventsMicrosoftGraphGetAuditCategoriesRequestBuilderGetRequestConfiguration)(AuditEventsMicrosoftGraphGetAuditCategoriesGetAuditCategoriesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateAuditEventsMicrosoftGraphGetAuditCategoriesGetAuditCategoriesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AuditEventsMicrosoftGraphGetAuditCategoriesGetAuditCategoriesResponseable), nil
}
// ToGetRequestInformation invoke function getAuditCategories
func (m *AuditEventsMicrosoftGraphGetAuditCategoriesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuditEventsMicrosoftGraphGetAuditCategoriesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
