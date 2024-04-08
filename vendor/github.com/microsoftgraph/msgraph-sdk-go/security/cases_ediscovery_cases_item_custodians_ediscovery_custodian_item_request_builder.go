package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder provides operations to manage the custodians property of the microsoft.graph.security.ediscoveryCase entity.
type CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderGetQueryParameters returns a list of case ediscoveryCustodian objects for this case.
type CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderGetQueryParameters
}
// CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderInternal instantiates a new EdiscoveryCustodianItemRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) {
    m := &CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/custodians/{ediscoveryCustodian%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder instantiates a new EdiscoveryCustodianItemRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property custodians for security
func (m *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns a list of case ediscoveryCustodian objects for this case.
func (m *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCustodianable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryCustodianFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCustodianable), nil
}
// LastIndexOperation provides operations to manage the lastIndexOperation property of the microsoft.graph.security.ediscoveryCustodian entity.
func (m *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) LastIndexOperation()(*CasesEdiscoveryCasesItemCustodiansItemLastIndexOperationRequestBuilder) {
    return NewCasesEdiscoveryCasesItemCustodiansItemLastIndexOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSecurityActivate provides operations to call the activate method.
func (m *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) MicrosoftGraphSecurityActivate()(*CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilder) {
    return NewCasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityActivateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSecurityApplyHold provides operations to call the applyHold method.
func (m *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) MicrosoftGraphSecurityApplyHold()(*CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityApplyHoldRequestBuilder) {
    return NewCasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityApplyHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSecurityRelease provides operations to call the release method.
func (m *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) MicrosoftGraphSecurityRelease()(*CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityReleaseRequestBuilder) {
    return NewCasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityReleaseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSecurityRemoveHold provides operations to call the removeHold method.
func (m *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) MicrosoftGraphSecurityRemoveHold()(*CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityRemoveHoldRequestBuilder) {
    return NewCasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityRemoveHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSecurityUpdateIndex provides operations to call the updateIndex method.
func (m *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) MicrosoftGraphSecurityUpdateIndex()(*CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityUpdateIndexRequestBuilder) {
    return NewCasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityUpdateIndexRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property custodians in security
func (m *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCustodianable, requestConfiguration *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCustodianable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryCustodianFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCustodianable), nil
}
// SiteSources provides operations to manage the siteSources property of the microsoft.graph.security.ediscoveryCustodian entity.
func (m *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) SiteSources()(*CasesEdiscoveryCasesItemCustodiansItemSiteSourcesRequestBuilder) {
    return NewCasesEdiscoveryCasesItemCustodiansItemSiteSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SiteSourcesById provides operations to manage the siteSources property of the microsoft.graph.security.ediscoveryCustodian entity.
func (m *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) SiteSourcesById(id string)(*CasesEdiscoveryCasesItemCustodiansItemSiteSourcesSiteSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["siteSource%2Did"] = id
    }
    return NewCasesEdiscoveryCasesItemCustodiansItemSiteSourcesSiteSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property custodians for security
func (m *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation returns a list of case ediscoveryCustodian objects for this case.
func (m *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property custodians in security
func (m *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCustodianable, requestConfiguration *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UnifiedGroupSources provides operations to manage the unifiedGroupSources property of the microsoft.graph.security.ediscoveryCustodian entity.
func (m *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) UnifiedGroupSources()(*CasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesRequestBuilder) {
    return NewCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UnifiedGroupSourcesById provides operations to manage the unifiedGroupSources property of the microsoft.graph.security.ediscoveryCustodian entity.
func (m *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) UnifiedGroupSourcesById(id string)(*CasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedGroupSource%2Did"] = id
    }
    return NewCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// UserSources provides operations to manage the userSources property of the microsoft.graph.security.ediscoveryCustodian entity.
func (m *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) UserSources()(*CasesEdiscoveryCasesItemCustodiansItemUserSourcesRequestBuilder) {
    return NewCasesEdiscoveryCasesItemCustodiansItemUserSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserSourcesById provides operations to manage the userSources property of the microsoft.graph.security.ediscoveryCustodian entity.
func (m *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) UserSourcesById(id string)(*CasesEdiscoveryCasesItemCustodiansItemUserSourcesUserSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userSource%2Did"] = id
    }
    return NewCasesEdiscoveryCasesItemCustodiansItemUserSourcesUserSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
