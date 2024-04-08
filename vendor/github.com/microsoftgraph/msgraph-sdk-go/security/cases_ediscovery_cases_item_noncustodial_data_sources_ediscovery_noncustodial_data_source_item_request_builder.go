package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder provides operations to manage the noncustodialDataSources property of the microsoft.graph.security.ediscoveryCase entity.
type CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetQueryParameters returns a list of case ediscoveryNoncustodialDataSource objects for this case.
type CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetQueryParameters
}
// CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderInternal instantiates a new EdiscoveryNoncustodialDataSourceItemRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) {
    m := &CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/noncustodialDataSources/{ediscoveryNoncustodialDataSource%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder instantiates a new EdiscoveryNoncustodialDataSourceItemRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// DataSource provides operations to manage the dataSource property of the microsoft.graph.security.ediscoveryNoncustodialDataSource entity.
func (m *CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) DataSource()(*CasesEdiscoveryCasesItemNoncustodialDataSourcesItemDataSourceRequestBuilder) {
    return NewCasesEdiscoveryCasesItemNoncustodialDataSourcesItemDataSourceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete navigation property noncustodialDataSources for security
func (m *CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns a list of case ediscoveryNoncustodialDataSource objects for this case.
func (m *CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryNoncustodialDataSourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryNoncustodialDataSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryNoncustodialDataSourceable), nil
}
// LastIndexOperation provides operations to manage the lastIndexOperation property of the microsoft.graph.security.ediscoveryNoncustodialDataSource entity.
func (m *CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) LastIndexOperation()(*CasesEdiscoveryCasesItemNoncustodialDataSourcesItemLastIndexOperationRequestBuilder) {
    return NewCasesEdiscoveryCasesItemNoncustodialDataSourcesItemLastIndexOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSecurityApplyHold provides operations to call the applyHold method.
func (m *CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) MicrosoftGraphSecurityApplyHold()(*CasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilder) {
    return NewCasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityApplyHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSecurityRelease provides operations to call the release method.
func (m *CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) MicrosoftGraphSecurityRelease()(*CasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityReleaseRequestBuilder) {
    return NewCasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityReleaseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSecurityRemoveHold provides operations to call the removeHold method.
func (m *CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) MicrosoftGraphSecurityRemoveHold()(*CasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityRemoveHoldRequestBuilder) {
    return NewCasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityRemoveHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSecurityUpdateIndex provides operations to call the updateIndex method.
func (m *CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) MicrosoftGraphSecurityUpdateIndex()(*CasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityUpdateIndexRequestBuilder) {
    return NewCasesEdiscoveryCasesItemNoncustodialDataSourcesItemMicrosoftGraphSecurityUpdateIndexRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property noncustodialDataSources in security
func (m *CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryNoncustodialDataSourceable, requestConfiguration *CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryNoncustodialDataSourceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryNoncustodialDataSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryNoncustodialDataSourceable), nil
}
// ToDeleteRequestInformation delete navigation property noncustodialDataSources for security
func (m *CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation returns a list of case ediscoveryNoncustodialDataSource objects for this case.
func (m *CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property noncustodialDataSources in security
func (m *CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryNoncustodialDataSourceable, requestConfiguration *CasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
