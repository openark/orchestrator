package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder provides operations to manage the collection of policyRoot entities.
type FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderDeleteQueryParameters delete ref of navigation property appliesTo for policies
type FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderDeleteQueryParameters struct {
    // Delete Uri
    Id *string `uriparametername:"%40id"`
}
// FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderDeleteQueryParameters
}
// NewFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderInternal instantiates a new RefRequestBuilder and sets the default values.
func NewFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder) {
    m := &FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/featureRolloutPolicies/{featureRolloutPolicy%2Did}/appliesTo/{directoryObject%2Did}/$ref{?%40id*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder instantiates a new RefRequestBuilder and sets the default values.
func NewFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete ref of navigation property appliesTo for policies
func (m *FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderDeleteRequestConfiguration)(error) {
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
// ToDeleteRequestInformation delete ref of navigation property appliesTo for policies
func (m *FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
