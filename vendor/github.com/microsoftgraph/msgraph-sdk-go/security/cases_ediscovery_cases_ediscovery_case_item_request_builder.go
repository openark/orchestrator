package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder provides operations to manage the ediscoveryCases property of the microsoft.graph.security.casesRoot entity.
type CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderGetQueryParameters get ediscoveryCases from security
type CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderGetQueryParameters
}
// CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderInternal instantiates a new EdiscoveryCaseItemRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) {
    m := &CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder instantiates a new EdiscoveryCaseItemRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Custodians provides operations to manage the custodians property of the microsoft.graph.security.ediscoveryCase entity.
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Custodians()(*CasesEdiscoveryCasesItemCustodiansRequestBuilder) {
    return NewCasesEdiscoveryCasesItemCustodiansRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CustodiansById provides operations to manage the custodians property of the microsoft.graph.security.ediscoveryCase entity.
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) CustodiansById(id string)(*CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoveryCustodian%2Did"] = id
    }
    return NewCasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Delete delete navigation property ediscoveryCases for security
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get ediscoveryCases from security
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCaseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryCaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCaseable), nil
}
// MicrosoftGraphSecurityClose provides operations to call the close method.
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) MicrosoftGraphSecurityClose()(*CasesEdiscoveryCasesItemMicrosoftGraphSecurityCloseRequestBuilder) {
    return NewCasesEdiscoveryCasesItemMicrosoftGraphSecurityCloseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSecurityReopen provides operations to call the reopen method.
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) MicrosoftGraphSecurityReopen()(*CasesEdiscoveryCasesItemMicrosoftGraphSecurityReopenRequestBuilder) {
    return NewCasesEdiscoveryCasesItemMicrosoftGraphSecurityReopenRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NoncustodialDataSources provides operations to manage the noncustodialDataSources property of the microsoft.graph.security.ediscoveryCase entity.
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) NoncustodialDataSources()(*CasesEdiscoveryCasesItemNoncustodialDataSourcesRequestBuilder) {
    return NewCasesEdiscoveryCasesItemNoncustodialDataSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NoncustodialDataSourcesById provides operations to manage the noncustodialDataSources property of the microsoft.graph.security.ediscoveryCase entity.
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) NoncustodialDataSourcesById(id string)(*CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoveryNoncustodialDataSource%2Did"] = id
    }
    return NewCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.security.ediscoveryCase entity.
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Operations()(*CasesEdiscoveryCasesItemOperationsRequestBuilder) {
    return NewCasesEdiscoveryCasesItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.security.ediscoveryCase entity.
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) OperationsById(id string)(*CasesEdiscoveryCasesItemOperationsCaseOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["caseOperation%2Did"] = id
    }
    return NewCasesEdiscoveryCasesItemOperationsCaseOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property ediscoveryCases in security
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCaseable, requestConfiguration *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCaseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryCaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCaseable), nil
}
// ReviewSets provides operations to manage the reviewSets property of the microsoft.graph.security.ediscoveryCase entity.
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) ReviewSets()(*CasesEdiscoveryCasesItemReviewSetsRequestBuilder) {
    return NewCasesEdiscoveryCasesItemReviewSetsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ReviewSetsById provides operations to manage the reviewSets property of the microsoft.graph.security.ediscoveryCase entity.
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) ReviewSetsById(id string)(*CasesEdiscoveryCasesItemReviewSetsEdiscoveryReviewSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoveryReviewSet%2Did"] = id
    }
    return NewCasesEdiscoveryCasesItemReviewSetsEdiscoveryReviewSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Searches provides operations to manage the searches property of the microsoft.graph.security.ediscoveryCase entity.
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Searches()(*CasesEdiscoveryCasesItemSearchesRequestBuilder) {
    return NewCasesEdiscoveryCasesItemSearchesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SearchesById provides operations to manage the searches property of the microsoft.graph.security.ediscoveryCase entity.
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) SearchesById(id string)(*CasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoverySearch%2Did"] = id
    }
    return NewCasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Settings provides operations to manage the settings property of the microsoft.graph.security.ediscoveryCase entity.
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Settings()(*CasesEdiscoveryCasesItemSettingsRequestBuilder) {
    return NewCasesEdiscoveryCasesItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Tags provides operations to manage the tags property of the microsoft.graph.security.ediscoveryCase entity.
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Tags()(*CasesEdiscoveryCasesItemTagsRequestBuilder) {
    return NewCasesEdiscoveryCasesItemTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TagsById provides operations to manage the tags property of the microsoft.graph.security.ediscoveryCase entity.
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) TagsById(id string)(*CasesEdiscoveryCasesItemTagsEdiscoveryReviewTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoveryReviewTag%2Did"] = id
    }
    return NewCasesEdiscoveryCasesItemTagsEdiscoveryReviewTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property ediscoveryCases for security
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get ediscoveryCases from security
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property ediscoveryCases in security
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCaseable, requestConfiguration *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
