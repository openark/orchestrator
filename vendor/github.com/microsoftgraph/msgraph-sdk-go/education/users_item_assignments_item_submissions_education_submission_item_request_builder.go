package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder provides operations to manage the submissions property of the microsoft.graph.educationAssignment entity.
type UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderGetQueryParameters once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
type UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderGetQueryParameters
}
// UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderInternal instantiates a new EducationSubmissionItemRequestBuilder and sets the default values.
func NewUsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) {
    m := &UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/users/{educationUser%2Did}/assignments/{educationAssignment%2Did}/submissions/{educationSubmission%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewUsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder instantiates a new EducationSubmissionItemRequestBuilder and sets the default values.
func NewUsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property submissions for education
func (m *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionable, error) {
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
func (m *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) MicrosoftGraphReassign()(*UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphReassignRequestBuilder) {
    return NewUsersItemAssignmentsItemSubmissionsItemMicrosoftGraphReassignRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReturn provides operations to call the return method.
func (m *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) MicrosoftGraphReturn()(*UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphReturnRequestBuilder) {
    return NewUsersItemAssignmentsItemSubmissionsItemMicrosoftGraphReturnRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetUpResourcesFolder provides operations to call the setUpResourcesFolder method.
func (m *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) MicrosoftGraphSetUpResourcesFolder()(*UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphSetUpResourcesFolderRequestBuilder) {
    return NewUsersItemAssignmentsItemSubmissionsItemMicrosoftGraphSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSubmit provides operations to call the submit method.
func (m *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) MicrosoftGraphSubmit()(*UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphSubmitRequestBuilder) {
    return NewUsersItemAssignmentsItemSubmissionsItemMicrosoftGraphSubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUnsubmit provides operations to call the unsubmit method.
func (m *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) MicrosoftGraphUnsubmit()(*UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphUnsubmitRequestBuilder) {
    return NewUsersItemAssignmentsItemSubmissionsItemMicrosoftGraphUnsubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Outcomes provides operations to manage the outcomes property of the microsoft.graph.educationSubmission entity.
func (m *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) Outcomes()(*UsersItemAssignmentsItemSubmissionsItemOutcomesRequestBuilder) {
    return NewUsersItemAssignmentsItemSubmissionsItemOutcomesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OutcomesById provides operations to manage the outcomes property of the microsoft.graph.educationSubmission entity.
func (m *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) OutcomesById(id string)(*UsersItemAssignmentsItemSubmissionsItemOutcomesEducationOutcomeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationOutcome%2Did"] = id
    }
    return NewUsersItemAssignmentsItemSubmissionsItemOutcomesEducationOutcomeItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property submissions in education
func (m *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionable, requestConfiguration *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionable, error) {
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
func (m *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) Resources()(*UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilder) {
    return NewUsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ResourcesById provides operations to manage the resources property of the microsoft.graph.educationSubmission entity.
func (m *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) ResourcesById(id string)(*UsersItemAssignmentsItemSubmissionsItemResourcesEducationSubmissionResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource%2Did"] = id
    }
    return NewUsersItemAssignmentsItemSubmissionsItemResourcesEducationSubmissionResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SubmittedResources provides operations to manage the submittedResources property of the microsoft.graph.educationSubmission entity.
func (m *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) SubmittedResources()(*UsersItemAssignmentsItemSubmissionsItemSubmittedResourcesRequestBuilder) {
    return NewUsersItemAssignmentsItemSubmissionsItemSubmittedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SubmittedResourcesById provides operations to manage the submittedResources property of the microsoft.graph.educationSubmission entity.
func (m *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) SubmittedResourcesById(id string)(*UsersItemAssignmentsItemSubmissionsItemSubmittedResourcesEducationSubmissionResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource%2Did"] = id
    }
    return NewUsersItemAssignmentsItemSubmissionsItemSubmittedResourcesEducationSubmissionResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property submissions for education
func (m *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionable, requestConfiguration *UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
