package education

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ClassesItemTeachersEducationUserItemRequestBuilder builds and executes requests for operations under \education\classes\{educationClass-id}\teachers\{educationUser-id}
type ClassesItemTeachersEducationUserItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewClassesItemTeachersEducationUserItemRequestBuilderInternal instantiates a new EducationUserItemRequestBuilder and sets the default values.
func NewClassesItemTeachersEducationUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemTeachersEducationUserItemRequestBuilder) {
    m := &ClassesItemTeachersEducationUserItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/classes/{educationClass%2Did}/teachers/{educationUser%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewClassesItemTeachersEducationUserItemRequestBuilder instantiates a new EducationUserItemRequestBuilder and sets the default values.
func NewClassesItemTeachersEducationUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemTeachersEducationUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClassesItemTeachersEducationUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of educationRoot entities.
func (m *ClassesItemTeachersEducationUserItemRequestBuilder) Ref()(*ClassesItemTeachersItemRefRequestBuilder) {
    return NewClassesItemTeachersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
