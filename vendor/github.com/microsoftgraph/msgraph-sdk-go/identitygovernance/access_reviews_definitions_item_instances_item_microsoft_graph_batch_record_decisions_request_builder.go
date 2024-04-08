package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphBatchRecordDecisionsRequestBuilder provides operations to call the batchRecordDecisions method.
type AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphBatchRecordDecisionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphBatchRecordDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphBatchRecordDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphBatchRecordDecisionsRequestBuilderInternal instantiates a new MicrosoftGraphBatchRecordDecisionsRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphBatchRecordDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphBatchRecordDecisionsRequestBuilder) {
    m := &AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphBatchRecordDecisionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/microsoft.graph.batchRecordDecisions";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphBatchRecordDecisionsRequestBuilder instantiates a new MicrosoftGraphBatchRecordDecisionsRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphBatchRecordDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphBatchRecordDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphBatchRecordDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post enables reviewers to review all accessReviewInstanceDecisionItem objects in batches by using **principalId**, **resourceId**, or neither.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/accessreviewinstance-batchrecorddecisions?view=graph-rest-1.0
func (m *AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphBatchRecordDecisionsRequestBuilder) Post(ctx context.Context, body AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBodyable, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphBatchRecordDecisionsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation enables reviewers to review all accessReviewInstanceDecisionItem objects in batches by using **principalId**, **resourceId**, or neither.
func (m *AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphBatchRecordDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBodyable, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphBatchRecordDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
