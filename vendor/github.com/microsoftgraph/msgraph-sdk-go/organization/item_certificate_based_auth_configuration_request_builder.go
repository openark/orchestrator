package organization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCertificateBasedAuthConfigurationRequestBuilder provides operations to manage the certificateBasedAuthConfiguration property of the microsoft.graph.organization entity.
type ItemCertificateBasedAuthConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemCertificateBasedAuthConfigurationRequestBuilderGetQueryParameters get a list of certificateBasedAuthConfiguration objects.
type ItemCertificateBasedAuthConfigurationRequestBuilderGetQueryParameters struct {
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
// ItemCertificateBasedAuthConfigurationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCertificateBasedAuthConfigurationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCertificateBasedAuthConfigurationRequestBuilderGetQueryParameters
}
// NewItemCertificateBasedAuthConfigurationRequestBuilderInternal instantiates a new CertificateBasedAuthConfigurationRequestBuilder and sets the default values.
func NewItemCertificateBasedAuthConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCertificateBasedAuthConfigurationRequestBuilder) {
    m := &ItemCertificateBasedAuthConfigurationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/organization/{organization%2Did}/certificateBasedAuthConfiguration{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemCertificateBasedAuthConfigurationRequestBuilder instantiates a new CertificateBasedAuthConfigurationRequestBuilder and sets the default values.
func NewItemCertificateBasedAuthConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCertificateBasedAuthConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCertificateBasedAuthConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemCertificateBasedAuthConfigurationRequestBuilder) Count()(*ItemCertificateBasedAuthConfigurationCountRequestBuilder) {
    return NewItemCertificateBasedAuthConfigurationCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get get a list of certificateBasedAuthConfiguration objects.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/certificatebasedauthconfiguration-list?view=graph-rest-1.0
func (m *ItemCertificateBasedAuthConfigurationRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCertificateBasedAuthConfigurationRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CertificateBasedAuthConfigurationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCertificateBasedAuthConfigurationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CertificateBasedAuthConfigurationCollectionResponseable), nil
}
// ToGetRequestInformation get a list of certificateBasedAuthConfiguration objects.
func (m *ItemCertificateBasedAuthConfigurationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCertificateBasedAuthConfigurationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
