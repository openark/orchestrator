package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphResetDecisionsRequestBuilder provides operations to call the resetDecisions method.
type AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphResetDecisionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphResetDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphResetDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphResetDecisionsRequestBuilderInternal instantiates a new MicrosoftGraphResetDecisionsRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphResetDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphResetDecisionsRequestBuilder) {
    m := &AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphResetDecisionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/microsoft.graph.resetDecisions";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphResetDecisionsRequestBuilder instantiates a new MicrosoftGraphResetDecisionsRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphResetDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphResetDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphResetDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post resets all accessReviewInstanceDecisionItem objects on an accessReviewInstance to `notReviewed`.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/accessreviewinstance-resetdecisions?view=graph-rest-1.0
func (m *AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphResetDecisionsRequestBuilder) Post(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphResetDecisionsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation resets all accessReviewInstanceDecisionItem objects on an accessReviewInstance to `notReviewed`.
func (m *AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphResetDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphResetDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
