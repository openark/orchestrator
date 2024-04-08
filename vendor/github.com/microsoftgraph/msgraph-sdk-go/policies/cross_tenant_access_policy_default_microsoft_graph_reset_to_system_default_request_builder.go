package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CrossTenantAccessPolicyDefaultMicrosoftGraphResetToSystemDefaultRequestBuilder provides operations to call the resetToSystemDefault method.
type CrossTenantAccessPolicyDefaultMicrosoftGraphResetToSystemDefaultRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CrossTenantAccessPolicyDefaultMicrosoftGraphResetToSystemDefaultRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrossTenantAccessPolicyDefaultMicrosoftGraphResetToSystemDefaultRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCrossTenantAccessPolicyDefaultMicrosoftGraphResetToSystemDefaultRequestBuilderInternal instantiates a new MicrosoftGraphResetToSystemDefaultRequestBuilder and sets the default values.
func NewCrossTenantAccessPolicyDefaultMicrosoftGraphResetToSystemDefaultRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CrossTenantAccessPolicyDefaultMicrosoftGraphResetToSystemDefaultRequestBuilder) {
    m := &CrossTenantAccessPolicyDefaultMicrosoftGraphResetToSystemDefaultRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/crossTenantAccessPolicy/default/microsoft.graph.resetToSystemDefault";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCrossTenantAccessPolicyDefaultMicrosoftGraphResetToSystemDefaultRequestBuilder instantiates a new MicrosoftGraphResetToSystemDefaultRequestBuilder and sets the default values.
func NewCrossTenantAccessPolicyDefaultMicrosoftGraphResetToSystemDefaultRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CrossTenantAccessPolicyDefaultMicrosoftGraphResetToSystemDefaultRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCrossTenantAccessPolicyDefaultMicrosoftGraphResetToSystemDefaultRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reset any changes made to the default configuration in a cross-tenant access policy back to the system default.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/crosstenantaccesspolicyconfigurationdefault-resettosystemdefault?view=graph-rest-1.0
func (m *CrossTenantAccessPolicyDefaultMicrosoftGraphResetToSystemDefaultRequestBuilder) Post(ctx context.Context, requestConfiguration *CrossTenantAccessPolicyDefaultMicrosoftGraphResetToSystemDefaultRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation reset any changes made to the default configuration in a cross-tenant access policy back to the system default.
func (m *CrossTenantAccessPolicyDefaultMicrosoftGraphResetToSystemDefaultRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CrossTenantAccessPolicyDefaultMicrosoftGraphResetToSystemDefaultRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
