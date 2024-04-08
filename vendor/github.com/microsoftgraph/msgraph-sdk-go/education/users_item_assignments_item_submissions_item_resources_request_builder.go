package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilder provides operations to manage the resources property of the microsoft.graph.educationSubmission entity.
type UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilderGetQueryParameters list the resources associated with a submission. Only teachers, students, and applications with application permissions can perform this operation. The educationSubmissionResource object is a wrapper around the actual resource object the student is working on. The wrapper also includes a pointer to the resources on the assignment if this was copied from the assignment during the assign process. These resources are the working copy of the assignment. The **submittedResources** are the resources that have officially been submitted to be graded.
type UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilderGetQueryParameters
}
// UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilderInternal instantiates a new ResourcesRequestBuilder and sets the default values.
func NewUsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilder) {
    m := &UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/users/{educationUser%2Did}/assignments/{educationAssignment%2Did}/submissions/{educationSubmission%2Did}/resources{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewUsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilder instantiates a new ResourcesRequestBuilder and sets the default values.
func NewUsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilder) Count()(*UsersItemAssignmentsItemSubmissionsItemResourcesCountRequestBuilder) {
    return NewUsersItemAssignmentsItemSubmissionsItemResourcesCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get list the resources associated with a submission. Only teachers, students, and applications with application permissions can perform this operation. The educationSubmissionResource object is a wrapper around the actual resource object the student is working on. The wrapper also includes a pointer to the resources on the assignment if this was copied from the assignment during the assign process. These resources are the working copy of the assignment. The **submittedResources** are the resources that have officially been submitted to be graded.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/educationsubmission-list-resources?view=graph-rest-1.0
func (m *UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionResourceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationSubmissionResourceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionResourceCollectionResponseable), nil
}
// Post add an educationSubmissionResource to a submission resource list. Only teachers and students can perform this operation. The operation will not succeed if the **allowStudentsToAddResources** flag is not set to `true`. To create a new file-based resource, upload the file to the resources folder associated with the submission. If the file doesn't exist or is not in that folder, the POST request will fail.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/educationsubmission-post-resources?view=graph-rest-1.0
func (m *UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionResourceable, requestConfiguration *UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionResourceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationSubmissionResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionResourceable), nil
}
// ToGetRequestInformation list the resources associated with a submission. Only teachers, students, and applications with application permissions can perform this operation. The educationSubmissionResource object is a wrapper around the actual resource object the student is working on. The wrapper also includes a pointer to the resources on the assignment if this was copied from the assignment during the assign process. These resources are the working copy of the assignment. The **submittedResources** are the resources that have officially been submitted to be graded.
func (m *UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation add an educationSubmissionResource to a submission resource list. Only teachers and students can perform this operation. The operation will not succeed if the **allowStudentsToAddResources** flag is not set to `true`. To create a new file-based resource, upload the file to the resources folder associated with the submission. If the file doesn't exist or is not in that folder, the POST request will fail.
func (m *UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionResourceable, requestConfiguration *UsersItemAssignmentsItemSubmissionsItemResourcesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
