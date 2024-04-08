package identity

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// B2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder builds and executes requests for operations under \identity\b2xUserFlows\{b2xIdentityUserFlow-id}\userFlowIdentityProviders\{identityProviderBase-id}
type B2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewB2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilderInternal instantiates a new IdentityProviderBaseItemRequestBuilder and sets the default values.
func NewB2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder) {
    m := &B2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/userFlowIdentityProviders/{identityProviderBase%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewB2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder instantiates a new IdentityProviderBaseItemRequestBuilder and sets the default values.
func NewB2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of identityContainer entities.
func (m *B2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder) Ref()(*B2xUserFlowsItemUserFlowIdentityProvidersItemRefRequestBuilder) {
    return NewB2xUserFlowsItemUserFlowIdentityProvidersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
