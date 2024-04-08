package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UsersItemAssignmentsEducationAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.educationUser entity.
type UsersItemAssignmentsEducationAssignmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemAssignmentsEducationAssignmentItemRequestBuilderGetQueryParameters assignments belonging to the user.
type UsersItemAssignmentsEducationAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemAssignmentsEducationAssignmentItemRequestBuilderGetQueryParameters
}
// UsersItemAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Categories provides operations to manage the categories property of the microsoft.graph.educationAssignment entity.
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) Categories()(*UsersItemAssignmentsItemCategoriesRequestBuilder) {
    return NewUsersItemAssignmentsItemCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CategoriesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.users.item.assignments.item.categories.item collection
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) CategoriesById(id string)(*UsersItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationCategory%2Did"] = id
    }
    return NewUsersItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewUsersItemAssignmentsEducationAssignmentItemRequestBuilderInternal instantiates a new EducationAssignmentItemRequestBuilder and sets the default values.
func NewUsersItemAssignmentsEducationAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAssignmentsEducationAssignmentItemRequestBuilder) {
    m := &UsersItemAssignmentsEducationAssignmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/users/{educationUser%2Did}/assignments/{educationAssignment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewUsersItemAssignmentsEducationAssignmentItemRequestBuilder instantiates a new EducationAssignmentItemRequestBuilder and sets the default values.
func NewUsersItemAssignmentsEducationAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAssignmentsEducationAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemAssignmentsEducationAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignments for education
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable, error) {
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
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) MicrosoftGraphPublish()(*UsersItemAssignmentsItemMicrosoftGraphPublishRequestBuilder) {
    return NewUsersItemAssignmentsItemMicrosoftGraphPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetUpFeedbackResourcesFolder provides operations to call the setUpFeedbackResourcesFolder method.
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) MicrosoftGraphSetUpFeedbackResourcesFolder()(*UsersItemAssignmentsItemMicrosoftGraphSetUpFeedbackResourcesFolderRequestBuilder) {
    return NewUsersItemAssignmentsItemMicrosoftGraphSetUpFeedbackResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetUpResourcesFolder provides operations to call the setUpResourcesFolder method.
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) MicrosoftGraphSetUpResourcesFolder()(*UsersItemAssignmentsItemMicrosoftGraphSetUpResourcesFolderRequestBuilder) {
    return NewUsersItemAssignmentsItemMicrosoftGraphSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property assignments in education
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable, requestConfiguration *UsersItemAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable, error) {
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
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) Resources()(*UsersItemAssignmentsItemResourcesRequestBuilder) {
    return NewUsersItemAssignmentsItemResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ResourcesById provides operations to manage the resources property of the microsoft.graph.educationAssignment entity.
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) ResourcesById(id string)(*UsersItemAssignmentsItemResourcesEducationAssignmentResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignmentResource%2Did"] = id
    }
    return NewUsersItemAssignmentsItemResourcesEducationAssignmentResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Rubric provides operations to manage the rubric property of the microsoft.graph.educationAssignment entity.
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) Rubric()(*UsersItemAssignmentsItemRubricRequestBuilder) {
    return NewUsersItemAssignmentsItemRubricRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Submissions provides operations to manage the submissions property of the microsoft.graph.educationAssignment entity.
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) Submissions()(*UsersItemAssignmentsItemSubmissionsRequestBuilder) {
    return NewUsersItemAssignmentsItemSubmissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SubmissionsById provides operations to manage the submissions property of the microsoft.graph.educationAssignment entity.
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) SubmissionsById(id string)(*UsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmission%2Did"] = id
    }
    return NewUsersItemAssignmentsItemSubmissionsEducationSubmissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property assignments for education
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemAssignmentsEducationAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemAssignmentsEducationAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemAssignmentsEducationAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable, requestConfiguration *UsersItemAssignmentsEducationAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
