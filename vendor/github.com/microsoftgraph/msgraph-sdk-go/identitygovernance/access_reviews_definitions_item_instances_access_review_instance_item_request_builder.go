package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.accessReviewScheduleDefinition entity.
type AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderGetQueryParameters if the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review that does not recur will have exactly one instance. Instances also represent each unique resource under review in the accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource will have a unique instance for each recurrence.
type AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderGetQueryParameters
}
// AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderInternal instantiates a new AccessReviewInstanceItemRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) {
    m := &AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder instantiates a new AccessReviewInstanceItemRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactedReviewers provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
func (m *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) ContactedReviewers()(*AccessReviewsDefinitionsItemInstancesItemContactedReviewersRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemContactedReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ContactedReviewersById provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
func (m *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) ContactedReviewersById(id string)(*AccessReviewsDefinitionsItemInstancesItemContactedReviewersAccessReviewReviewerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer%2Did"] = id
    }
    return NewAccessReviewsDefinitionsItemInstancesItemContactedReviewersAccessReviewReviewerItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Decisions provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
func (m *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) Decisions()(*AccessReviewsDefinitionsItemInstancesItemDecisionsRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DecisionsById provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
func (m *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) DecisionsById(id string)(*AccessReviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstanceDecisionItem%2Did"] = id
    }
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Delete delete navigation property instances for identityGovernance
func (m *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get if the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review that does not recur will have exactly one instance. Instances also represent each unique resource under review in the accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource will have a unique instance for each recurrence.
func (m *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewInstanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessReviewInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewInstanceable), nil
}
// MicrosoftGraphAcceptRecommendations provides operations to call the acceptRecommendations method.
func (m *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) MicrosoftGraphAcceptRecommendations()(*AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphAcceptRecommendationsRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyDecisions provides operations to call the applyDecisions method.
func (m *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) MicrosoftGraphApplyDecisions()(*AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphApplyDecisionsRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBatchRecordDecisions provides operations to call the batchRecordDecisions method.
func (m *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) MicrosoftGraphBatchRecordDecisions()(*AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphBatchRecordDecisionsRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphResetDecisions provides operations to call the resetDecisions method.
func (m *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) MicrosoftGraphResetDecisions()(*AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphResetDecisionsRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSendReminder provides operations to call the sendReminder method.
func (m *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) MicrosoftGraphSendReminder()(*AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphSendReminderRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphStop provides operations to call the stop method.
func (m *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) MicrosoftGraphStop()(*AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphStopRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphStopRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property instances in identityGovernance
func (m *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewInstanceable, requestConfiguration *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewInstanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessReviewInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewInstanceable), nil
}
// Stages provides operations to manage the stages property of the microsoft.graph.accessReviewInstance entity.
func (m *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) Stages()(*AccessReviewsDefinitionsItemInstancesItemStagesRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemStagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// StagesById provides operations to manage the stages property of the microsoft.graph.accessReviewInstance entity.
func (m *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) StagesById(id string)(*AccessReviewsDefinitionsItemInstancesItemStagesAccessReviewStageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewStage%2Did"] = id
    }
    return NewAccessReviewsDefinitionsItemInstancesItemStagesAccessReviewStageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property instances for identityGovernance
func (m *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation if the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review that does not recur will have exactly one instance. Instances also represent each unique resource under review in the accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource will have a unique instance for each recurrence.
func (m *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property instances in identityGovernance
func (m *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewInstanceable, requestConfiguration *AccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
