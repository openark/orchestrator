package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRequestBuilder provides operations to call the removeGroup method.
type ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRequestBuilderInternal instantiates a new MicrosoftGraphRemoveGroupRequestBuilder and sets the default values.
func NewItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRequestBuilder) {
    m := &ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/groupLifecyclePolicies/{groupLifecyclePolicy%2Did}/microsoft.graph.removeGroup";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRequestBuilder instantiates a new MicrosoftGraphRemoveGroupRequestBuilder and sets the default values.
func NewItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action removeGroup
func (m *ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRequestBuilder) Post(ctx context.Context, body ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRemoveGroupPostRequestBodyable, requestConfiguration *ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRequestBuilderPostRequestConfiguration)(ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRemoveGroupResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRemoveGroupResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRemoveGroupResponseable), nil
}
// ToPostRequestInformation invoke action removeGroup
func (m *ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRemoveGroupPostRequestBodyable, requestConfiguration *ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
