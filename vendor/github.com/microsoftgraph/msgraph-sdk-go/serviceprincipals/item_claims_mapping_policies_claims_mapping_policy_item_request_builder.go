package serviceprincipals

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\claimsMappingPolicies\{claimsMappingPolicy-id}
type ItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilderInternal instantiates a new ClaimsMappingPolicyItemRequestBuilder and sets the default values.
func NewItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilder) {
    m := &ItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/claimsMappingPolicies/{claimsMappingPolicy%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilder instantiates a new ClaimsMappingPolicyItemRequestBuilder and sets the default values.
func NewItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of servicePrincipal entities.
func (m *ItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilder) Ref()(*ItemClaimsMappingPoliciesItemRefRequestBuilder) {
    return NewItemClaimsMappingPoliciesItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
