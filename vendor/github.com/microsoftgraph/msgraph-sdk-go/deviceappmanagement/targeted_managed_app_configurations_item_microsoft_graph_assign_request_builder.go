package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// TargetedManagedAppConfigurationsItemMicrosoftGraphAssignRequestBuilder provides operations to call the assign method.
type TargetedManagedAppConfigurationsItemMicrosoftGraphAssignRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TargetedManagedAppConfigurationsItemMicrosoftGraphAssignRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TargetedManagedAppConfigurationsItemMicrosoftGraphAssignRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTargetedManagedAppConfigurationsItemMicrosoftGraphAssignRequestBuilderInternal instantiates a new MicrosoftGraphAssignRequestBuilder and sets the default values.
func NewTargetedManagedAppConfigurationsItemMicrosoftGraphAssignRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedManagedAppConfigurationsItemMicrosoftGraphAssignRequestBuilder) {
    m := &TargetedManagedAppConfigurationsItemMicrosoftGraphAssignRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration%2Did}/microsoft.graph.assign";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewTargetedManagedAppConfigurationsItemMicrosoftGraphAssignRequestBuilder instantiates a new MicrosoftGraphAssignRequestBuilder and sets the default values.
func NewTargetedManagedAppConfigurationsItemMicrosoftGraphAssignRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedManagedAppConfigurationsItemMicrosoftGraphAssignRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetedManagedAppConfigurationsItemMicrosoftGraphAssignRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action assign
func (m *TargetedManagedAppConfigurationsItemMicrosoftGraphAssignRequestBuilder) Post(ctx context.Context, body TargetedManagedAppConfigurationsItemMicrosoftGraphAssignAssignPostRequestBodyable, requestConfiguration *TargetedManagedAppConfigurationsItemMicrosoftGraphAssignRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action assign
func (m *TargetedManagedAppConfigurationsItemMicrosoftGraphAssignRequestBuilder) ToPostRequestInformation(ctx context.Context, body TargetedManagedAppConfigurationsItemMicrosoftGraphAssignAssignPostRequestBodyable, requestConfiguration *TargetedManagedAppConfigurationsItemMicrosoftGraphAssignRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
