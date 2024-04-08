package applications

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder builds and executes requests for operations under \applications\{application-id}\tokenLifetimePolicies\{tokenLifetimePolicy-id}
type ItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilderInternal instantiates a new TokenLifetimePolicyItemRequestBuilder and sets the default values.
func NewItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder) {
    m := &ItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application%2Did}/tokenLifetimePolicies/{tokenLifetimePolicy%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder instantiates a new TokenLifetimePolicyItemRequestBuilder and sets the default values.
func NewItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of application entities.
func (m *ItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder) Ref()(*ItemTokenLifetimePoliciesItemRefRequestBuilder) {
    return NewItemTokenLifetimePoliciesItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
