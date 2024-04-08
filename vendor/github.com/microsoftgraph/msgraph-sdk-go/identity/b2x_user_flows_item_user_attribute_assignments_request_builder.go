package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// B2xUserFlowsItemUserAttributeAssignmentsRequestBuilder provides operations to manage the userAttributeAssignments property of the microsoft.graph.b2xIdentityUserFlow entity.
type B2xUserFlowsItemUserAttributeAssignmentsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// B2xUserFlowsItemUserAttributeAssignmentsRequestBuilderGetQueryParameters get the identityUserFlowAttributeAssignment resources from the userAttributeAssignments navigation property in a b2xIdentityUserFlow.
type B2xUserFlowsItemUserAttributeAssignmentsRequestBuilderGetQueryParameters struct {
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
// B2xUserFlowsItemUserAttributeAssignmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xUserFlowsItemUserAttributeAssignmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2xUserFlowsItemUserAttributeAssignmentsRequestBuilderGetQueryParameters
}
// B2xUserFlowsItemUserAttributeAssignmentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xUserFlowsItemUserAttributeAssignmentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewB2xUserFlowsItemUserAttributeAssignmentsRequestBuilderInternal instantiates a new UserAttributeAssignmentsRequestBuilder and sets the default values.
func NewB2xUserFlowsItemUserAttributeAssignmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xUserFlowsItemUserAttributeAssignmentsRequestBuilder) {
    m := &B2xUserFlowsItemUserAttributeAssignmentsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/userAttributeAssignments{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewB2xUserFlowsItemUserAttributeAssignmentsRequestBuilder instantiates a new UserAttributeAssignmentsRequestBuilder and sets the default values.
func NewB2xUserFlowsItemUserAttributeAssignmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xUserFlowsItemUserAttributeAssignmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xUserFlowsItemUserAttributeAssignmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *B2xUserFlowsItemUserAttributeAssignmentsRequestBuilder) Count()(*B2xUserFlowsItemUserAttributeAssignmentsCountRequestBuilder) {
    return NewB2xUserFlowsItemUserAttributeAssignmentsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get get the identityUserFlowAttributeAssignment resources from the userAttributeAssignments navigation property in a b2xIdentityUserFlow.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/b2xidentityuserflow-list-userattributeassignments?view=graph-rest-1.0
func (m *B2xUserFlowsItemUserAttributeAssignmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *B2xUserFlowsItemUserAttributeAssignmentsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeAssignmentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIdentityUserFlowAttributeAssignmentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeAssignmentCollectionResponseable), nil
}
// MicrosoftGraphGetOrder provides operations to call the getOrder method.
func (m *B2xUserFlowsItemUserAttributeAssignmentsRequestBuilder) MicrosoftGraphGetOrder()(*B2xUserFlowsItemUserAttributeAssignmentsMicrosoftGraphGetOrderRequestBuilder) {
    return NewB2xUserFlowsItemUserAttributeAssignmentsMicrosoftGraphGetOrderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetOrder provides operations to call the setOrder method.
func (m *B2xUserFlowsItemUserAttributeAssignmentsRequestBuilder) MicrosoftGraphSetOrder()(*B2xUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderRequestBuilder) {
    return NewB2xUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Post create a new identityUserFlowAttributeAssignment object in a b2xIdentityUserFlow.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/b2xidentityuserflow-post-userattributeassignments?view=graph-rest-1.0
func (m *B2xUserFlowsItemUserAttributeAssignmentsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeAssignmentable, requestConfiguration *B2xUserFlowsItemUserAttributeAssignmentsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeAssignmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIdentityUserFlowAttributeAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeAssignmentable), nil
}
// ToGetRequestInformation get the identityUserFlowAttributeAssignment resources from the userAttributeAssignments navigation property in a b2xIdentityUserFlow.
func (m *B2xUserFlowsItemUserAttributeAssignmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2xUserFlowsItemUserAttributeAssignmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new identityUserFlowAttributeAssignment object in a b2xIdentityUserFlow.
func (m *B2xUserFlowsItemUserAttributeAssignmentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeAssignmentable, requestConfiguration *B2xUserFlowsItemUserAttributeAssignmentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
