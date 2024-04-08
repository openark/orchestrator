package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DirectoryRoleEligibilityScheduleRequestsRequestBuilder provides operations to manage the roleEligibilityScheduleRequests property of the microsoft.graph.rbacApplication entity.
type DirectoryRoleEligibilityScheduleRequestsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectoryRoleEligibilityScheduleRequestsRequestBuilderGetQueryParameters in PIM, retrieve the requests for role eligibilities for principals made through the unifiedRoleEligibilityScheduleRequest object.
type DirectoryRoleEligibilityScheduleRequestsRequestBuilderGetQueryParameters struct {
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
// DirectoryRoleEligibilityScheduleRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleEligibilityScheduleRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleEligibilityScheduleRequestsRequestBuilderGetQueryParameters
}
// DirectoryRoleEligibilityScheduleRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleEligibilityScheduleRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDirectoryRoleEligibilityScheduleRequestsRequestBuilderInternal instantiates a new RoleEligibilityScheduleRequestsRequestBuilder and sets the default values.
func NewDirectoryRoleEligibilityScheduleRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleEligibilityScheduleRequestsRequestBuilder) {
    m := &DirectoryRoleEligibilityScheduleRequestsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/directory/roleEligibilityScheduleRequests{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDirectoryRoleEligibilityScheduleRequestsRequestBuilder instantiates a new RoleEligibilityScheduleRequestsRequestBuilder and sets the default values.
func NewDirectoryRoleEligibilityScheduleRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleEligibilityScheduleRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleEligibilityScheduleRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *DirectoryRoleEligibilityScheduleRequestsRequestBuilder) Count()(*DirectoryRoleEligibilityScheduleRequestsCountRequestBuilder) {
    return NewDirectoryRoleEligibilityScheduleRequestsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get in PIM, retrieve the requests for role eligibilities for principals made through the unifiedRoleEligibilityScheduleRequest object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/rbacapplication-list-roleeligibilityschedulerequests?view=graph-rest-1.0
func (m *DirectoryRoleEligibilityScheduleRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleEligibilityScheduleRequestsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleEligibilityScheduleRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleRequestCollectionResponseable), nil
}
// MicrosoftGraphFilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
func (m *DirectoryRoleEligibilityScheduleRequestsRequestBuilder) MicrosoftGraphFilterByCurrentUserWithOn(on *string)(*DirectoryRoleEligibilityScheduleRequestsMicrosoftGraphFilterByCurrentUserWithOnRequestBuilder) {
    return NewDirectoryRoleEligibilityScheduleRequestsMicrosoftGraphFilterByCurrentUserWithOnRequestBuilderInternal(m.pathParameters, m.requestAdapter, on)
}
// Post in PIM, request for a role eligibility for a principal through the unifiedRoleEligibilityScheduleRequest object. This operation allows both admins and eligible users to add, revoke, or extend eligible assignments.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/rbacapplication-post-roleeligibilityschedulerequests?view=graph-rest-1.0
func (m *DirectoryRoleEligibilityScheduleRequestsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleRequestable, requestConfiguration *DirectoryRoleEligibilityScheduleRequestsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleEligibilityScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleRequestable), nil
}
// ToGetRequestInformation in PIM, retrieve the requests for role eligibilities for principals made through the unifiedRoleEligibilityScheduleRequest object.
func (m *DirectoryRoleEligibilityScheduleRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleEligibilityScheduleRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation in PIM, request for a role eligibility for a principal through the unifiedRoleEligibilityScheduleRequest object. This operation allows both admins and eligible users to add, revoke, or extend eligible assignments.
func (m *DirectoryRoleEligibilityScheduleRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleRequestable, requestConfiguration *DirectoryRoleEligibilityScheduleRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
