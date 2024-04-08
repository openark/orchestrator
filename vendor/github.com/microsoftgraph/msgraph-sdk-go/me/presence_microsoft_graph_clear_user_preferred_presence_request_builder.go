package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PresenceMicrosoftGraphClearUserPreferredPresenceRequestBuilder provides operations to call the clearUserPreferredPresence method.
type PresenceMicrosoftGraphClearUserPreferredPresenceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PresenceMicrosoftGraphClearUserPreferredPresenceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PresenceMicrosoftGraphClearUserPreferredPresenceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPresenceMicrosoftGraphClearUserPreferredPresenceRequestBuilderInternal instantiates a new MicrosoftGraphClearUserPreferredPresenceRequestBuilder and sets the default values.
func NewPresenceMicrosoftGraphClearUserPreferredPresenceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresenceMicrosoftGraphClearUserPreferredPresenceRequestBuilder) {
    m := &PresenceMicrosoftGraphClearUserPreferredPresenceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/presence/microsoft.graph.clearUserPreferredPresence";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewPresenceMicrosoftGraphClearUserPreferredPresenceRequestBuilder instantiates a new MicrosoftGraphClearUserPreferredPresenceRequestBuilder and sets the default values.
func NewPresenceMicrosoftGraphClearUserPreferredPresenceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresenceMicrosoftGraphClearUserPreferredPresenceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPresenceMicrosoftGraphClearUserPreferredPresenceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post clear the preferred availability and activity status for a user.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/presence-clearuserpreferredpresence?view=graph-rest-1.0
func (m *PresenceMicrosoftGraphClearUserPreferredPresenceRequestBuilder) Post(ctx context.Context, requestConfiguration *PresenceMicrosoftGraphClearUserPreferredPresenceRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation clear the preferred availability and activity status for a user.
func (m *PresenceMicrosoftGraphClearUserPreferredPresenceRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *PresenceMicrosoftGraphClearUserPreferredPresenceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
