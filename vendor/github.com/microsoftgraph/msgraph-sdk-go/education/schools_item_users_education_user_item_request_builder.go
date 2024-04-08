package education

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SchoolsItemUsersEducationUserItemRequestBuilder builds and executes requests for operations under \education\schools\{educationSchool-id}\users\{educationUser-id}
type SchoolsItemUsersEducationUserItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewSchoolsItemUsersEducationUserItemRequestBuilderInternal instantiates a new EducationUserItemRequestBuilder and sets the default values.
func NewSchoolsItemUsersEducationUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchoolsItemUsersEducationUserItemRequestBuilder) {
    m := &SchoolsItemUsersEducationUserItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/schools/{educationSchool%2Did}/users/{educationUser%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewSchoolsItemUsersEducationUserItemRequestBuilder instantiates a new EducationUserItemRequestBuilder and sets the default values.
func NewSchoolsItemUsersEducationUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchoolsItemUsersEducationUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchoolsItemUsersEducationUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of educationRoot entities.
func (m *SchoolsItemUsersEducationUserItemRequestBuilder) Ref()(*SchoolsItemUsersItemRefRequestBuilder) {
    return NewSchoolsItemUsersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
