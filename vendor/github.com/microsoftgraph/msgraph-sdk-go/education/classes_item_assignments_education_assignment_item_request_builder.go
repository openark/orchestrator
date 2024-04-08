package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ClassesItemAssignmentsEducationAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.educationClass entity.
type ClassesItemAssignmentsEducationAssignmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ClassesItemAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ClassesItemAssignmentsEducationAssignmentItemRequestBuilderGetQueryParameters all assignments associated with this class. Nullable.
type ClassesItemAssignmentsEducationAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ClassesItemAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ClassesItemAssignmentsEducationAssignmentItemRequestBuilderGetQueryParameters
}
// ClassesItemAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Categories provides operations to manage the categories property of the microsoft.graph.educationAssignment entity.
func (m *ClassesItemAssignmentsEducationAssignmentItemRequestBuilder) Categories()(*ClassesItemAssignmentsItemCategoriesRequestBuilder) {
    return NewClassesItemAssignmentsItemCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CategoriesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.classes.item.assignments.item.categories.item collection
func (m *ClassesItemAssignmentsEducationAssignmentItemRequestBuilder) CategoriesById(id string)(*ClassesItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationCategory%2Did"] = id
    }
    return NewClassesItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewClassesItemAssignmentsEducationAssignmentItemRequestBuilderInternal instantiates a new EducationAssignmentItemRequestBuilder and sets the default values.
func NewClassesItemAssignmentsEducationAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsEducationAssignmentItemRequestBuilder) {
    m := &ClassesItemAssignmentsEducationAssignmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/classes/{educationClass%2Did}/assignments/{educationAssignment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewClassesItemAssignmentsEducationAssignmentItemRequestBuilder instantiates a new EducationAssignmentItemRequestBuilder and sets the default values.
func NewClassesItemAssignmentsEducationAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsEducationAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClassesItemAssignmentsEducationAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignments for education
func (m *ClassesItemAssignmentsEducationAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ClassesItemAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get all assignments associated with this class. Nullable.
func (m *ClassesItemAssignmentsEducationAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ClassesItemAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable), nil
}
// MicrosoftGraphPublish provides operations to call the publish method.
func (m *ClassesItemAssignmentsEducationAssignmentItemRequestBuilder) MicrosoftGraphPublish()(*ClassesItemAssignmentsItemMicrosoftGraphPublishRequestBuilder) {
    return NewClassesItemAssignmentsItemMicrosoftGraphPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetUpFeedbackResourcesFolder provides operations to call the setUpFeedbackResourcesFolder method.
func (m *ClassesItemAssignmentsEducationAssignmentItemRequestBuilder) MicrosoftGraphSetUpFeedbackResourcesFolder()(*ClassesItemAssignmentsItemMicrosoftGraphSetUpFeedbackResourcesFolderRequestBuilder) {
    return NewClassesItemAssignmentsItemMicrosoftGraphSetUpFeedbackResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetUpResourcesFolder provides operations to call the setUpResourcesFolder method.
func (m *ClassesItemAssignmentsEducationAssignmentItemRequestBuilder) MicrosoftGraphSetUpResourcesFolder()(*ClassesItemAssignmentsItemMicrosoftGraphSetUpResourcesFolderRequestBuilder) {
    return NewClassesItemAssignmentsItemMicrosoftGraphSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property assignments in education
func (m *ClassesItemAssignmentsEducationAssignmentItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable, requestConfiguration *ClassesItemAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable), nil
}
// Resources provides operations to manage the resources property of the microsoft.graph.educationAssignment entity.
func (m *ClassesItemAssignmentsEducationAssignmentItemRequestBuilder) Resources()(*ClassesItemAssignmentsItemResourcesRequestBuilder) {
    return NewClassesItemAssignmentsItemResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ResourcesById provides operations to manage the resources property of the microsoft.graph.educationAssignment entity.
func (m *ClassesItemAssignmentsEducationAssignmentItemRequestBuilder) ResourcesById(id string)(*ClassesItemAssignmentsItemResourcesEducationAssignmentResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignmentResource%2Did"] = id
    }
    return NewClassesItemAssignmentsItemResourcesEducationAssignmentResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Rubric provides operations to manage the rubric property of the microsoft.graph.educationAssignment entity.
func (m *ClassesItemAssignmentsEducationAssignmentItemRequestBuilder) Rubric()(*ClassesItemAssignmentsItemRubricRequestBuilder) {
    return NewClassesItemAssignmentsItemRubricRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Submissions provides operations to manage the submissions property of the microsoft.graph.educationAssignment entity.
func (m *ClassesItemAssignmentsEducationAssignmentItemRequestBuilder) Submissions()(*ClassesItemAssignmentsItemSubmissionsRequestBuilder) {
    return NewClassesItemAssignmentsItemSubmissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SubmissionsById provides operations to manage the submissions property of the microsoft.graph.educationAssignment entity.
func (m *ClassesItemAssignmentsEducationAssignmentItemRequestBuilder) SubmissionsById(id string)(*ClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmission%2Did"] = id
    }
    return NewClassesItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property assignments for education
func (m *ClassesItemAssignmentsEducationAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ClassesItemAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation all assignments associated with this class. Nullable.
func (m *ClassesItemAssignmentsEducationAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ClassesItemAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignments in education
func (m *ClassesItemAssignmentsEducationAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable, requestConfiguration *ClassesItemAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
