package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder provides operations to manage the tasks property of the microsoft.graph.plannerBucket entity.
type PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderGetQueryParameters read-only. Nullable. The collection of tasks in the bucket.
type PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderGetQueryParameters
}
// PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssignedToTaskBoardFormat provides operations to manage the assignedToTaskBoardFormat property of the microsoft.graph.plannerTask entity.
func (m *PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) AssignedToTaskBoardFormat()(*PlannerPlansItemBucketsItemTasksItemAssignedToTaskBoardFormatRequestBuilder) {
    return NewPlannerPlansItemBucketsItemTasksItemAssignedToTaskBoardFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// BucketTaskBoardFormat provides operations to manage the bucketTaskBoardFormat property of the microsoft.graph.plannerTask entity.
func (m *PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) BucketTaskBoardFormat()(*PlannerPlansItemBucketsItemTasksItemBucketTaskBoardFormatRequestBuilder) {
    return NewPlannerPlansItemBucketsItemTasksItemBucketTaskBoardFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderInternal instantiates a new PlannerTaskItemRequestBuilder and sets the default values.
func NewPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) {
    m := &PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/planner/plans/{plannerPlan%2Did}/buckets/{plannerBucket%2Did}/tasks/{plannerTask%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder instantiates a new PlannerTaskItemRequestBuilder and sets the default values.
func NewPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property tasks for me
func (m *PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Details provides operations to manage the details property of the microsoft.graph.plannerTask entity.
func (m *PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) Details()(*PlannerPlansItemBucketsItemTasksItemDetailsRequestBuilder) {
    return NewPlannerPlansItemBucketsItemTasksItemDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get read-only. Nullable. The collection of tasks in the bucket.
func (m *PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlannerTaskable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePlannerTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlannerTaskable), nil
}
// Patch update the navigation property tasks in me
func (m *PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlannerTaskable, requestConfiguration *PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlannerTaskable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePlannerTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlannerTaskable), nil
}
// ProgressTaskBoardFormat provides operations to manage the progressTaskBoardFormat property of the microsoft.graph.plannerTask entity.
func (m *PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) ProgressTaskBoardFormat()(*PlannerPlansItemBucketsItemTasksItemProgressTaskBoardFormatRequestBuilder) {
    return NewPlannerPlansItemBucketsItemTasksItemProgressTaskBoardFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property tasks for me
func (m *PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation read-only. Nullable. The collection of tasks in the bucket.
func (m *PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property tasks in me
func (m *PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlannerTaskable, requestConfiguration *PlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
