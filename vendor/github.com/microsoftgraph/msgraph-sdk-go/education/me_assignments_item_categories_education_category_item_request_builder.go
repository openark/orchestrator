package education

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MeAssignmentsItemCategoriesEducationCategoryItemRequestBuilder builds and executes requests for operations under \education\me\assignments\{educationAssignment-id}\categories\{educationCategory-id}
type MeAssignmentsItemCategoriesEducationCategoryItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewMeAssignmentsItemCategoriesEducationCategoryItemRequestBuilderInternal instantiates a new EducationCategoryItemRequestBuilder and sets the default values.
func NewMeAssignmentsItemCategoriesEducationCategoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemCategoriesEducationCategoryItemRequestBuilder) {
    m := &MeAssignmentsItemCategoriesEducationCategoryItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/me/assignments/{educationAssignment%2Did}/categories/{educationCategory%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMeAssignmentsItemCategoriesEducationCategoryItemRequestBuilder instantiates a new EducationCategoryItemRequestBuilder and sets the default values.
func NewMeAssignmentsItemCategoriesEducationCategoryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemCategoriesEducationCategoryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAssignmentsItemCategoriesEducationCategoryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of educationRoot entities.
func (m *MeAssignmentsItemCategoriesEducationCategoryItemRequestBuilder) Ref()(*MeAssignmentsItemCategoriesItemRefRequestBuilder) {
    return NewMeAssignmentsItemCategoriesItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
