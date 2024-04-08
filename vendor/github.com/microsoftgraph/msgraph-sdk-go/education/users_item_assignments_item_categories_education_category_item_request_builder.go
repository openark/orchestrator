package education

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UsersItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilder builds and executes requests for operations under \education\users\{educationUser-id}\assignments\{educationAssignment-id}\categories\{educationCategory-id}
type UsersItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewUsersItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilderInternal instantiates a new EducationCategoryItemRequestBuilder and sets the default values.
func NewUsersItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilder) {
    m := &UsersItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/users/{educationUser%2Did}/assignments/{educationAssignment%2Did}/categories/{educationCategory%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewUsersItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilder instantiates a new EducationCategoryItemRequestBuilder and sets the default values.
func NewUsersItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of educationRoot entities.
func (m *UsersItemAssignmentsItemCategoriesEducationCategoryItemRequestBuilder) Ref()(*UsersItemAssignmentsItemCategoriesItemRefRequestBuilder) {
    return NewUsersItemAssignmentsItemCategoriesItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
