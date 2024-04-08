package policies

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// FeatureRolloutPoliciesItemAppliesToDirectoryObjectItemRequestBuilder builds and executes requests for operations under \policies\featureRolloutPolicies\{featureRolloutPolicy-id}\appliesTo\{directoryObject-id}
type FeatureRolloutPoliciesItemAppliesToDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewFeatureRolloutPoliciesItemAppliesToDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewFeatureRolloutPoliciesItemAppliesToDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FeatureRolloutPoliciesItemAppliesToDirectoryObjectItemRequestBuilder) {
    m := &FeatureRolloutPoliciesItemAppliesToDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/featureRolloutPolicies/{featureRolloutPolicy%2Did}/appliesTo/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewFeatureRolloutPoliciesItemAppliesToDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewFeatureRolloutPoliciesItemAppliesToDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FeatureRolloutPoliciesItemAppliesToDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFeatureRolloutPoliciesItemAppliesToDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of policyRoot entities.
func (m *FeatureRolloutPoliciesItemAppliesToDirectoryObjectItemRequestBuilder) Ref()(*FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder) {
    return NewFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
