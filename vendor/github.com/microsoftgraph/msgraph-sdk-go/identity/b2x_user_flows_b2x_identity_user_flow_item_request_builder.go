package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// B2xUserFlowsB2xIdentityUserFlowItemRequestBuilder provides operations to manage the b2xUserFlows property of the microsoft.graph.identityContainer entity.
type B2xUserFlowsB2xIdentityUserFlowItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// B2xUserFlowsB2xIdentityUserFlowItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xUserFlowsB2xIdentityUserFlowItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// B2xUserFlowsB2xIdentityUserFlowItemRequestBuilderGetQueryParameters represents entry point for B2X/self-service sign-up identity userflows.
type B2xUserFlowsB2xIdentityUserFlowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// B2xUserFlowsB2xIdentityUserFlowItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xUserFlowsB2xIdentityUserFlowItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2xUserFlowsB2xIdentityUserFlowItemRequestBuilderGetQueryParameters
}
// B2xUserFlowsB2xIdentityUserFlowItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xUserFlowsB2xIdentityUserFlowItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewB2xUserFlowsB2xIdentityUserFlowItemRequestBuilderInternal instantiates a new B2xIdentityUserFlowItemRequestBuilder and sets the default values.
func NewB2xUserFlowsB2xIdentityUserFlowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xUserFlowsB2xIdentityUserFlowItemRequestBuilder) {
    m := &B2xUserFlowsB2xIdentityUserFlowItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewB2xUserFlowsB2xIdentityUserFlowItemRequestBuilder instantiates a new B2xIdentityUserFlowItemRequestBuilder and sets the default values.
func NewB2xUserFlowsB2xIdentityUserFlowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xUserFlowsB2xIdentityUserFlowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xUserFlowsB2xIdentityUserFlowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property b2xUserFlows for identity
func (m *B2xUserFlowsB2xIdentityUserFlowItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *B2xUserFlowsB2xIdentityUserFlowItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents entry point for B2X/self-service sign-up identity userflows.
func (m *B2xUserFlowsB2xIdentityUserFlowItemRequestBuilder) Get(ctx context.Context, requestConfiguration *B2xUserFlowsB2xIdentityUserFlowItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.B2xIdentityUserFlowable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateB2xIdentityUserFlowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.B2xIdentityUserFlowable), nil
}
// IdentityProviders provides operations to manage the identityProviders property of the microsoft.graph.b2xIdentityUserFlow entity.
func (m *B2xUserFlowsB2xIdentityUserFlowItemRequestBuilder) IdentityProviders()(*B2xUserFlowsItemIdentityProvidersRequestBuilder) {
    return NewB2xUserFlowsItemIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IdentityProvidersById provides operations to manage the identityProviders property of the microsoft.graph.b2xIdentityUserFlow entity.
func (m *B2xUserFlowsB2xIdentityUserFlowItemRequestBuilder) IdentityProvidersById(id string)(*B2xUserFlowsItemIdentityProvidersIdentityProviderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProvider%2Did"] = id
    }
    return NewB2xUserFlowsItemIdentityProvidersIdentityProviderItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Languages provides operations to manage the languages property of the microsoft.graph.b2xIdentityUserFlow entity.
func (m *B2xUserFlowsB2xIdentityUserFlowItemRequestBuilder) Languages()(*B2xUserFlowsItemLanguagesRequestBuilder) {
    return NewB2xUserFlowsItemLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// LanguagesById provides operations to manage the languages property of the microsoft.graph.b2xIdentityUserFlow entity.
func (m *B2xUserFlowsB2xIdentityUserFlowItemRequestBuilder) LanguagesById(id string)(*B2xUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguageConfiguration%2Did"] = id
    }
    return NewB2xUserFlowsItemLanguagesUserFlowLanguageConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property b2xUserFlows in identity
func (m *B2xUserFlowsB2xIdentityUserFlowItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.B2xIdentityUserFlowable, requestConfiguration *B2xUserFlowsB2xIdentityUserFlowItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.B2xIdentityUserFlowable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateB2xIdentityUserFlowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.B2xIdentityUserFlowable), nil
}
// ToDeleteRequestInformation delete navigation property b2xUserFlows for identity
func (m *B2xUserFlowsB2xIdentityUserFlowItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *B2xUserFlowsB2xIdentityUserFlowItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation represents entry point for B2X/self-service sign-up identity userflows.
func (m *B2xUserFlowsB2xIdentityUserFlowItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2xUserFlowsB2xIdentityUserFlowItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property b2xUserFlows in identity
func (m *B2xUserFlowsB2xIdentityUserFlowItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.B2xIdentityUserFlowable, requestConfiguration *B2xUserFlowsB2xIdentityUserFlowItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserAttributeAssignments provides operations to manage the userAttributeAssignments property of the microsoft.graph.b2xIdentityUserFlow entity.
func (m *B2xUserFlowsB2xIdentityUserFlowItemRequestBuilder) UserAttributeAssignments()(*B2xUserFlowsItemUserAttributeAssignmentsRequestBuilder) {
    return NewB2xUserFlowsItemUserAttributeAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserAttributeAssignmentsById provides operations to manage the userAttributeAssignments property of the microsoft.graph.b2xIdentityUserFlow entity.
func (m *B2xUserFlowsB2xIdentityUserFlowItemRequestBuilder) UserAttributeAssignmentsById(id string)(*B2xUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityUserFlowAttributeAssignment%2Did"] = id
    }
    return NewB2xUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// UserFlowIdentityProviders provides operations to manage the userFlowIdentityProviders property of the microsoft.graph.b2xIdentityUserFlow entity.
func (m *B2xUserFlowsB2xIdentityUserFlowItemRequestBuilder) UserFlowIdentityProviders()(*B2xUserFlowsItemUserFlowIdentityProvidersRequestBuilder) {
    return NewB2xUserFlowsItemUserFlowIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserFlowIdentityProvidersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.b2xUserFlows.item.userFlowIdentityProviders.item collection
func (m *B2xUserFlowsB2xIdentityUserFlowItemRequestBuilder) UserFlowIdentityProvidersById(id string)(*B2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProviderBase%2Did"] = id
    }
    return NewB2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
