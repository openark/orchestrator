package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsRequestBuilder provides operations to call the estimateStatistics method.
type CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsRequestBuilderInternal instantiates a new MicrosoftGraphSecurityEstimateStatisticsRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsRequestBuilder) {
    m := &CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/searches/{ediscoverySearch%2Did}/microsoft.graph.security.estimateStatistics";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsRequestBuilder instantiates a new MicrosoftGraphSecurityEstimateStatisticsRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post run an estimate of the number of emails and documents in the eDiscovery search. To learn more about searches in eDiscovery, see Collect data for a case in eDiscovery (Premium).
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/security-ediscoverysearch-estimatestatistics?view=graph-rest-1.0
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsRequestBuilder) Post(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation run an estimate of the number of emails and documents in the eDiscovery search. To learn more about searches in eDiscovery, see Collect data for a case in eDiscovery (Premium).
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
