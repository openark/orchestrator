package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ClassesItemAssignmentsItemSubmissionsItemResourcesCountRequestBuilder provides operations to count the resources in the collection.
type ClassesItemAssignmentsItemSubmissionsItemResourcesCountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ClassesItemAssignmentsItemSubmissionsItemResourcesCountRequestBuilderGetQueryParameters get the number of the resource
type ClassesItemAssignmentsItemSubmissionsItemResourcesCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// ClassesItemAssignmentsItemSubmissionsItemResourcesCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsItemSubmissionsItemResourcesCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ClassesItemAssignmentsItemSubmissionsItemResourcesCountRequestBuilderGetQueryParameters
}
// NewClassesItemAssignmentsItemSubmissionsItemResourcesCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewClassesItemAssignmentsItemSubmissionsItemResourcesCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsItemSubmissionsItemResourcesCountRequestBuilder) {
    m := &ClassesItemAssignmentsItemSubmissionsItemResourcesCountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/classes/{educationClass%2Did}/assignments/{educationAssignment%2Did}/submissions/{educationSubmission%2Did}/resources/$count{?%24search,%24filter}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewClassesItemAssignmentsItemSubmissionsItemResourcesCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewClassesItemAssignmentsItemSubmissionsItemResourcesCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsItemSubmissionsItemResourcesCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClassesItemAssignmentsItemSubmissionsItemResourcesCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *ClassesItemAssignmentsItemSubmissionsItemResourcesCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ClassesItemAssignmentsItemSubmissionsItemResourcesCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
func (m *ClassesItemAssignmentsItemSubmissionsItemResourcesCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ClassesItemAssignmentsItemSubmissionsItemResourcesCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "text/plain")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
