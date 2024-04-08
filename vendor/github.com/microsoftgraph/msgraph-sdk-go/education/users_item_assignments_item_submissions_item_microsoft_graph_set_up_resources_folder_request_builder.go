package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphSetUpResourcesFolderRequestBuilder provides operations to call the setUpResourcesFolder method.
type UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphSetUpResourcesFolderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphSetUpResourcesFolderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphSetUpResourcesFolderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersItemAssignmentsItemSubmissionsItemMicrosoftGraphSetUpResourcesFolderRequestBuilderInternal instantiates a new MicrosoftGraphSetUpResourcesFolderRequestBuilder and sets the default values.
func NewUsersItemAssignmentsItemSubmissionsItemMicrosoftGraphSetUpResourcesFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphSetUpResourcesFolderRequestBuilder) {
    m := &UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphSetUpResourcesFolderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/users/{educationUser%2Did}/assignments/{educationAssignment%2Did}/submissions/{educationSubmission%2Did}/microsoft.graph.setUpResourcesFolder";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewUsersItemAssignmentsItemSubmissionsItemMicrosoftGraphSetUpResourcesFolderRequestBuilder instantiates a new MicrosoftGraphSetUpResourcesFolderRequestBuilder and sets the default values.
func NewUsersItemAssignmentsItemSubmissionsItemMicrosoftGraphSetUpResourcesFolderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphSetUpResourcesFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemAssignmentsItemSubmissionsItemMicrosoftGraphSetUpResourcesFolderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post trigger the creation of the SharePoint resource folder where all file-based resources (Word, Excel, and so on) should be uploaded for a given submission. Only teachers and students can perform this operation. Note that files must be located in this folder in order to be added as resources. Only a student in the class can determine what files to upload in a given submission-level resource folder. 
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/educationsubmission-setupresourcesfolder?view=graph-rest-1.0
func (m *UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphSetUpResourcesFolderRequestBuilder) Post(ctx context.Context, requestConfiguration *UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphSetUpResourcesFolderRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionable, error) {
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
// ToPostRequestInformation trigger the creation of the SharePoint resource folder where all file-based resources (Word, Excel, and so on) should be uploaded for a given submission. Only teachers and students can perform this operation. Note that files must be located in this folder in order to be added as resources. Only a student in the class can determine what files to upload in a given submission-level resource folder. 
func (m *UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphSetUpResourcesFolderRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *UsersItemAssignmentsItemSubmissionsItemMicrosoftGraphSetUpResourcesFolderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
