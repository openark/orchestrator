package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphReassignRequestBuilder provides operations to call the reassign method.
type UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphReassignRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphReassignRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphReassignRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersItemAssignmentsItemSubmissionsItemMicrosoftGraphReassignRequestBuilderInternal instantiates a new MicrosoftGraphReassignRequestBuilder and sets the default values.
func NewUsersItemAssignmentsItemSubmissionsItemMicrosoftGraphReassignRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphReassignRequestBuilder) {
    m := &UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphReassignRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/users/{educationUser%2Did}/assignments/{educationAssignment%2Did}/submissions/{educationSubmission%2Did}/microsoft.graph.reassign";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewUsersItemAssignmentsItemSubmissionsItemMicrosoftGraphReassignRequestBuilder instantiates a new MicrosoftGraphReassignRequestBuilder and sets the default values.
func NewUsersItemAssignmentsItemSubmissionsItemMicrosoftGraphReassignRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphReassignRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemAssignmentsItemSubmissionsItemMicrosoftGraphReassignRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reassign the submission to the student with feedback for review. Only teachers can perform this action.  Include the `Prefer: include-unknown-enum-members` header when you call this method; otherwise, a reassigned submission will be treated as a returned submission. This means that the `reassigned` status will be mapped to the `returned` status, and **reassignedDateTime** and **reassignedBy** properties will be mapped to **returnedDateTime** and **returnedBy** respectively. If the header `Prefer: include-unknown-enum-members` is provided, a reassigned submission retains the `reassigned` status. For details, see the examples section.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/educationsubmission-reassign?view=graph-rest-1.0
func (m *UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphReassignRequestBuilder) Post(ctx context.Context, requestConfiguration *UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphReassignRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation reassign the submission to the student with feedback for review. Only teachers can perform this action.  Include the `Prefer: include-unknown-enum-members` header when you call this method; otherwise, a reassigned submission will be treated as a returned submission. This means that the `reassigned` status will be mapped to the `returned` status, and **reassignedDateTime** and **reassignedBy** properties will be mapped to **returnedDateTime** and **returnedBy** respectively. If the header `Prefer: include-unknown-enum-members` is provided, a reassigned submission retains the `reassigned` status. For details, see the examples section.
func (m *UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphReassignRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphReassignRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
