package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UsersEducationUserItemRequestBuilder provides operations to manage the users property of the microsoft.graph.educationRoot entity.
type UsersEducationUserItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersEducationUserItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersEducationUserItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersEducationUserItemRequestBuilderGetQueryParameters get users from education
type UsersEducationUserItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersEducationUserItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersEducationUserItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersEducationUserItemRequestBuilderGetQueryParameters
}
// UsersEducationUserItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersEducationUserItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.educationUser entity.
func (m *UsersEducationUserItemRequestBuilder) Assignments()(*UsersItemAssignmentsRequestBuilder) {
    return NewUsersItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.educationUser entity.
func (m *UsersEducationUserItemRequestBuilder) AssignmentsById(id string)(*UsersItemAssignmentsEducationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignment%2Did"] = id
    }
    return NewUsersItemAssignmentsEducationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Classes provides operations to manage the classes property of the microsoft.graph.educationUser entity.
func (m *UsersEducationUserItemRequestBuilder) Classes()(*UsersItemClassesRequestBuilder) {
    return NewUsersItemClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ClassesById provides operations to manage the classes property of the microsoft.graph.educationUser entity.
func (m *UsersEducationUserItemRequestBuilder) ClassesById(id string)(*UsersItemClassesEducationClassItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationClass%2Did"] = id
    }
    return NewUsersItemClassesEducationClassItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewUsersEducationUserItemRequestBuilderInternal instantiates a new EducationUserItemRequestBuilder and sets the default values.
func NewUsersEducationUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersEducationUserItemRequestBuilder) {
    m := &UsersEducationUserItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/users/{educationUser%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewUsersEducationUserItemRequestBuilder instantiates a new EducationUserItemRequestBuilder and sets the default values.
func NewUsersEducationUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersEducationUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersEducationUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property users for education
func (m *UsersEducationUserItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersEducationUserItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get users from education
func (m *UsersEducationUserItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersEducationUserItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationUserable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationUserable), nil
}
// Patch update the navigation property users in education
func (m *UsersEducationUserItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationUserable, requestConfiguration *UsersEducationUserItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationUserable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationUserable), nil
}
// Rubrics provides operations to manage the rubrics property of the microsoft.graph.educationUser entity.
func (m *UsersEducationUserItemRequestBuilder) Rubrics()(*UsersItemRubricsRequestBuilder) {
    return NewUsersItemRubricsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RubricsById provides operations to manage the rubrics property of the microsoft.graph.educationUser entity.
func (m *UsersEducationUserItemRequestBuilder) RubricsById(id string)(*UsersItemRubricsEducationRubricItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationRubric%2Did"] = id
    }
    return NewUsersItemRubricsEducationRubricItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Schools provides operations to manage the schools property of the microsoft.graph.educationUser entity.
func (m *UsersEducationUserItemRequestBuilder) Schools()(*UsersItemSchoolsRequestBuilder) {
    return NewUsersItemSchoolsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SchoolsById provides operations to manage the schools property of the microsoft.graph.educationUser entity.
func (m *UsersEducationUserItemRequestBuilder) SchoolsById(id string)(*UsersItemSchoolsEducationSchoolItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSchool%2Did"] = id
    }
    return NewUsersItemSchoolsEducationSchoolItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// TaughtClasses provides operations to manage the taughtClasses property of the microsoft.graph.educationUser entity.
func (m *UsersEducationUserItemRequestBuilder) TaughtClasses()(*UsersItemTaughtClassesRequestBuilder) {
    return NewUsersItemTaughtClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TaughtClassesById provides operations to manage the taughtClasses property of the microsoft.graph.educationUser entity.
func (m *UsersEducationUserItemRequestBuilder) TaughtClassesById(id string)(*UsersItemTaughtClassesEducationClassItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationClass%2Did"] = id
    }
    return NewUsersItemTaughtClassesEducationClassItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property users for education
func (m *UsersEducationUserItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersEducationUserItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get users from education
func (m *UsersEducationUserItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UsersEducationUserItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property users in education
func (m *UsersEducationUserItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationUserable, requestConfiguration *UsersEducationUserItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// User provides operations to manage the user property of the microsoft.graph.educationUser entity.
func (m *UsersEducationUserItemRequestBuilder) User()(*UsersItemUserRequestBuilder) {
    return NewUsersItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
