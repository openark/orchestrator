package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphApplyDecisionsRequestBuilder provides operations to call the applyDecisions method.
type AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphApplyDecisionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphApplyDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphApplyDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphApplyDecisionsRequestBuilderInternal instantiates a new MicrosoftGraphApplyDecisionsRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphApplyDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphApplyDecisionsRequestBuilder) {
    m := &AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphApplyDecisionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/microsoft.graph.applyDecisions";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphApplyDecisionsRequestBuilder instantiates a new MicrosoftGraphApplyDecisionsRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphApplyDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphApplyDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphApplyDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post apply review decisions on an accessReviewInstance if the decisions were not applied automatically because the autoApplyDecisionsEnabled property is `false` in the review's accessReviewScheduleSettings. The status of the accessReviewInstance must be `Completed` to call this method.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/accessreviewinstance-applydecisions?view=graph-rest-1.0
func (m *AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphApplyDecisionsRequestBuilder) Post(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphApplyDecisionsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation apply review decisions on an accessReviewInstance if the decisions were not applied automatically because the autoApplyDecisionsEnabled property is `false` in the review's accessReviewScheduleSettings. The status of the accessReviewInstance must be `Completed` to call this method.
func (m *AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphApplyDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphApplyDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
