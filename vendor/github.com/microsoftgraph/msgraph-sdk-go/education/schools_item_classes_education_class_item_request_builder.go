package education

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SchoolsItemClassesEducationClassItemRequestBuilder builds and executes requests for operations under \education\schools\{educationSchool-id}\classes\{educationClass-id}
type SchoolsItemClassesEducationClassItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewSchoolsItemClassesEducationClassItemRequestBuilderInternal instantiates a new EducationClassItemRequestBuilder and sets the default values.
func NewSchoolsItemClassesEducationClassItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchoolsItemClassesEducationClassItemRequestBuilder) {
    m := &SchoolsItemClassesEducationClassItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/schools/{educationSchool%2Did}/classes/{educationClass%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewSchoolsItemClassesEducationClassItemRequestBuilder instantiates a new EducationClassItemRequestBuilder and sets the default values.
func NewSchoolsItemClassesEducationClassItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchoolsItemClassesEducationClassItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchoolsItemClassesEducationClassItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of educationRoot entities.
func (m *SchoolsItemClassesEducationClassItemRequestBuilder) Ref()(*SchoolsItemClassesItemRefRequestBuilder) {
    return NewSchoolsItemClassesItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
