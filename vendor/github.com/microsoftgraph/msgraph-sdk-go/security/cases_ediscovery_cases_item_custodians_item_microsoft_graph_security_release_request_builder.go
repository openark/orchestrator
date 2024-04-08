package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityReleaseRequestBuilder provides operations to call the release method.
type CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityReleaseRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityReleaseRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityReleaseRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityReleaseRequestBuilderInternal instantiates a new MicrosoftGraphSecurityReleaseRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityReleaseRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityReleaseRequestBuilder) {
    m := &CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityReleaseRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/custodians/{ediscoveryCustodian%2Did}/microsoft.graph.security.release";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityReleaseRequestBuilder instantiates a new MicrosoftGraphSecurityReleaseRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityReleaseRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityReleaseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityReleaseRequestBuilderInternal(urlParams, requestAdapter)
}
// Post release a custodian from a case. For details, see Release a custodian from a case.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/security-ediscoverycustodian-release?view=graph-rest-1.0
func (m *CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityReleaseRequestBuilder) Post(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityReleaseRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation release a custodian from a case. For details, see Release a custodian from a case.
func (m *CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityReleaseRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemCustodiansItemMicrosoftGraphSecurityReleaseRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
