package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder provides operations to manage the submissions property of the microsoft.graph.educationAssignment entity.
type MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderGetQueryParameters once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
type MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderGetQueryParameters
}
// MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderInternal instantiates a new EducationSubmissionItemRequestBuilder and sets the default values.
func NewMeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) {
    m := &MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/me/assignments/{educationAssignment%2Did}/submissions/{educationSubmission%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder instantiates a new EducationSubmissionItemRequestBuilder and sets the default values.
func NewMeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property submissions for education
func (m *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationSubmissionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionable), nil
}
// MicrosoftGraphReassign provides operations to call the reassign method.
func (m *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) MicrosoftGraphReassign()(*MeAssignmentsItemSubmissionsItemMicrosoftGraphReassignRequestBuilder) {
    return NewMeAssignmentsItemSubmissionsItemMicrosoftGraphReassignRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReturn provides operations to call the return method.
func (m *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) MicrosoftGraphReturn()(*MeAssignmentsItemSubmissionsItemMicrosoftGraphReturnRequestBuilder) {
    return NewMeAssignmentsItemSubmissionsItemMicrosoftGraphReturnRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetUpResourcesFolder provides operations to call the setUpResourcesFolder method.
func (m *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) MicrosoftGraphSetUpResourcesFolder()(*MeAssignmentsItemSubmissionsItemMicrosoftGraphSetUpResourcesFolderRequestBuilder) {
    return NewMeAssignmentsItemSubmissionsItemMicrosoftGraphSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSubmit provides operations to call the submit method.
func (m *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) MicrosoftGraphSubmit()(*MeAssignmentsItemSubmissionsItemMicrosoftGraphSubmitRequestBuilder) {
    return NewMeAssignmentsItemSubmissionsItemMicrosoftGraphSubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUnsubmit provides operations to call the unsubmit method.
func (m *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) MicrosoftGraphUnsubmit()(*MeAssignmentsItemSubmissionsItemMicrosoftGraphUnsubmitRequestBuilder) {
    return NewMeAssignmentsItemSubmissionsItemMicrosoftGraphUnsubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Outcomes provides operations to manage the outcomes property of the microsoft.graph.educationSubmission entity.
func (m *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) Outcomes()(*MeAssignmentsItemSubmissionsItemOutcomesRequestBuilder) {
    return NewMeAssignmentsItemSubmissionsItemOutcomesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OutcomesById provides operations to manage the outcomes property of the microsoft.graph.educationSubmission entity.
func (m *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) OutcomesById(id string)(*MeAssignmentsItemSubmissionsItemOutcomesEducationOutcomeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationOutcome%2Did"] = id
    }
    return NewMeAssignmentsItemSubmissionsItemOutcomesEducationOutcomeItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property submissions in education
func (m *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionable, requestConfiguration *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationSubmissionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionable), nil
}
// Resources provides operations to manage the resources property of the microsoft.graph.educationSubmission entity.
func (m *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) Resources()(*MeAssignmentsItemSubmissionsItemResourcesRequestBuilder) {
    return NewMeAssignmentsItemSubmissionsItemResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ResourcesById provides operations to manage the resources property of the microsoft.graph.educationSubmission entity.
func (m *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) ResourcesById(id string)(*MeAssignmentsItemSubmissionsItemResourcesEducationSubmissionResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource%2Did"] = id
    }
    return NewMeAssignmentsItemSubmissionsItemResourcesEducationSubmissionResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SubmittedResources provides operations to manage the submittedResources property of the microsoft.graph.educationSubmission entity.
func (m *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) SubmittedResources()(*MeAssignmentsItemSubmissionsItemSubmittedResourcesRequestBuilder) {
    return NewMeAssignmentsItemSubmissionsItemSubmittedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SubmittedResourcesById provides operations to manage the submittedResources property of the microsoft.graph.educationSubmission entity.
func (m *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) SubmittedResourcesById(id string)(*MeAssignmentsItemSubmissionsItemSubmittedResourcesEducationSubmissionResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource%2Did"] = id
    }
    return NewMeAssignmentsItemSubmissionsItemSubmittedResourcesEducationSubmissionResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property submissions for education
func (m *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property submissions in education
func (m *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionable, requestConfiguration *MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
