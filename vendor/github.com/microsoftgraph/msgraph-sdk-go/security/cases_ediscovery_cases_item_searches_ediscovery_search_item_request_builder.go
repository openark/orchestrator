package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder provides operations to manage the searches property of the microsoft.graph.security.ediscoveryCase entity.
type CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilderGetQueryParameters returns a list of eDiscoverySearch objects associated with this case.
type CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilderGetQueryParameters
}
// CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdditionalSources provides operations to manage the additionalSources property of the microsoft.graph.security.ediscoverySearch entity.
func (m *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder) AdditionalSources()(*CasesEdiscoveryCasesItemSearchesItemAdditionalSourcesRequestBuilder) {
    return NewCasesEdiscoveryCasesItemSearchesItemAdditionalSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AdditionalSourcesById provides operations to manage the additionalSources property of the microsoft.graph.security.ediscoverySearch entity.
func (m *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder) AdditionalSourcesById(id string)(*CasesEdiscoveryCasesItemSearchesItemAdditionalSourcesDataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataSource%2Did"] = id
    }
    return NewCasesEdiscoveryCasesItemSearchesItemAdditionalSourcesDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// AddToReviewSetOperation provides operations to manage the addToReviewSetOperation property of the microsoft.graph.security.ediscoverySearch entity.
func (m *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder) AddToReviewSetOperation()(*CasesEdiscoveryCasesItemSearchesItemAddToReviewSetOperationRequestBuilder) {
    return NewCasesEdiscoveryCasesItemSearchesItemAddToReviewSetOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewCasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilderInternal instantiates a new EdiscoverySearchItemRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder) {
    m := &CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/searches/{ediscoverySearch%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder instantiates a new EdiscoverySearchItemRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CustodianSources provides operations to manage the custodianSources property of the microsoft.graph.security.ediscoverySearch entity.
func (m *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder) CustodianSources()(*CasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilder) {
    return NewCasesEdiscoveryCasesItemSearchesItemCustodianSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CustodianSourcesById provides operations to manage the custodianSources property of the microsoft.graph.security.ediscoverySearch entity.
func (m *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder) CustodianSourcesById(id string)(*CasesEdiscoveryCasesItemSearchesItemCustodianSourcesDataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataSource%2Did"] = id
    }
    return NewCasesEdiscoveryCasesItemSearchesItemCustodianSourcesDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Delete delete navigation property searches for security
func (m *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns a list of eDiscoverySearch objects associated with this case.
func (m *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoverySearchable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoverySearchFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoverySearchable), nil
}
// LastEstimateStatisticsOperation provides operations to manage the lastEstimateStatisticsOperation property of the microsoft.graph.security.ediscoverySearch entity.
func (m *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder) LastEstimateStatisticsOperation()(*CasesEdiscoveryCasesItemSearchesItemLastEstimateStatisticsOperationRequestBuilder) {
    return NewCasesEdiscoveryCasesItemSearchesItemLastEstimateStatisticsOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSecurityEstimateStatistics provides operations to call the estimateStatistics method.
func (m *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder) MicrosoftGraphSecurityEstimateStatistics()(*CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsRequestBuilder) {
    return NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSecurityPurgeData provides operations to call the purgeData method.
func (m *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder) MicrosoftGraphSecurityPurgeData()(*CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataRequestBuilder) {
    return NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NoncustodialSources provides operations to manage the noncustodialSources property of the microsoft.graph.security.ediscoverySearch entity.
func (m *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder) NoncustodialSources()(*CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesRequestBuilder) {
    return NewCasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NoncustodialSourcesById provides operations to manage the noncustodialSources property of the microsoft.graph.security.ediscoverySearch entity.
func (m *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder) NoncustodialSourcesById(id string)(*CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoveryNoncustodialDataSource%2Did"] = id
    }
    return NewCasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property searches in security
func (m *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoverySearchable, requestConfiguration *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoverySearchable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoverySearchFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoverySearchable), nil
}
// ToDeleteRequestInformation delete navigation property searches for security
func (m *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation returns a list of eDiscoverySearch objects associated with this case.
func (m *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property searches in security
func (m *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoverySearchable, requestConfiguration *CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
