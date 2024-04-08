package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilder provides operations to call the getAttackSimulationTrainingUserCoverage method.
type SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilderGetQueryParameters invoke function getAttackSimulationTrainingUserCoverage
type SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilderGetQueryParameters struct {
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
// SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilderGetQueryParameters
}
// NewSecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilderInternal instantiates a new MicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilder and sets the default values.
func NewSecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilder) {
    m := &SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/security/microsoft.graph.getAttackSimulationTrainingUserCoverage(){?%24top,%24skip,%24search,%24filter,%24count}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewSecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilder instantiates a new MicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilder and sets the default values.
func NewSecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getAttackSimulationTrainingUserCoverage
func (m *SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilder) Get(ctx context.Context, requestConfiguration *SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilderGetRequestConfiguration)(SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageGetAttackSimulationTrainingUserCoverageResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateSecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageGetAttackSimulationTrainingUserCoverageResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageGetAttackSimulationTrainingUserCoverageResponseable), nil
}
// ToGetRequestInformation invoke function getAttackSimulationTrainingUserCoverage
func (m *SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SecurityMicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
