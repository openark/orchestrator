package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MeAssignmentsEducationAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.educationUser entity.
type MeAssignmentsEducationAssignmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MeAssignmentsEducationAssignmentItemRequestBuilderGetQueryParameters assignments belonging to the user.
type MeAssignmentsEducationAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeAssignmentsEducationAssignmentItemRequestBuilderGetQueryParameters
}
// MeAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Categories provides operations to manage the categories property of the microsoft.graph.educationAssignment entity.
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Categories()(*MeAssignmentsItemCategoriesRequestBuilder) {
    return NewMeAssignmentsItemCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CategoriesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.me.assignments.item.categories.item collection
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) CategoriesById(id string)(*MeAssignmentsItemCategoriesEducationCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationCategory%2Did"] = id
    }
    return NewMeAssignmentsItemCategoriesEducationCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewMeAssignmentsEducationAssignmentItemRequestBuilderInternal instantiates a new EducationAssignmentItemRequestBuilder and sets the default values.
func NewMeAssignmentsEducationAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsEducationAssignmentItemRequestBuilder) {
    m := &MeAssignmentsEducationAssignmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/me/assignments/{educationAssignment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMeAssignmentsEducationAssignmentItemRequestBuilder instantiates a new EducationAssignmentItemRequestBuilder and sets the default values.
func NewMeAssignmentsEducationAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsEducationAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAssignmentsEducationAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignments for education
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MeAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get assignments belonging to the user.
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable, error) {
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
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) MicrosoftGraphPublish()(*MeAssignmentsItemMicrosoftGraphPublishRequestBuilder) {
    return NewMeAssignmentsItemMicrosoftGraphPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetUpFeedbackResourcesFolder provides operations to call the setUpFeedbackResourcesFolder method.
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) MicrosoftGraphSetUpFeedbackResourcesFolder()(*MeAssignmentsItemMicrosoftGraphSetUpFeedbackResourcesFolderRequestBuilder) {
    return NewMeAssignmentsItemMicrosoftGraphSetUpFeedbackResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetUpResourcesFolder provides operations to call the setUpResourcesFolder method.
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) MicrosoftGraphSetUpResourcesFolder()(*MeAssignmentsItemMicrosoftGraphSetUpResourcesFolderRequestBuilder) {
    return NewMeAssignmentsItemMicrosoftGraphSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property assignments in education
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable, requestConfiguration *MeAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable, error) {
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
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Resources()(*MeAssignmentsItemResourcesRequestBuilder) {
    return NewMeAssignmentsItemResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ResourcesById provides operations to manage the resources property of the microsoft.graph.educationAssignment entity.
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) ResourcesById(id string)(*MeAssignmentsItemResourcesEducationAssignmentResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignmentResource%2Did"] = id
    }
    return NewMeAssignmentsItemResourcesEducationAssignmentResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Rubric provides operations to manage the rubric property of the microsoft.graph.educationAssignment entity.
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Rubric()(*MeAssignmentsItemRubricRequestBuilder) {
    return NewMeAssignmentsItemRubricRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Submissions provides operations to manage the submissions property of the microsoft.graph.educationAssignment entity.
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) Submissions()(*MeAssignmentsItemSubmissionsRequestBuilder) {
    return NewMeAssignmentsItemSubmissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SubmissionsById provides operations to manage the submissions property of the microsoft.graph.educationAssignment entity.
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) SubmissionsById(id string)(*MeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmission%2Did"] = id
    }
    return NewMeAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property assignments for education
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MeAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation assignments belonging to the user.
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MeAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MeAssignmentsEducationAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable, requestConfiguration *MeAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
