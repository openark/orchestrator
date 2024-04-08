package applications

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder builds and executes requests for operations under \applications\{application-id}\tokenIssuancePolicies\{tokenIssuancePolicy-id}
type ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilderInternal instantiates a new TokenIssuancePolicyItemRequestBuilder and sets the default values.
func NewItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder) {
    m := &ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application%2Did}/tokenIssuancePolicies/{tokenIssuancePolicy%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder instantiates a new TokenIssuancePolicyItemRequestBuilder and sets the default values.
func NewItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of application entities.
func (m *ItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder) Ref()(*ItemTokenIssuancePoliciesItemRefRequestBuilder) {
    return NewItemTokenIssuancePoliciesItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
