package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PresencesItemMicrosoftGraphClearPresenceRequestBuilder provides operations to call the clearPresence method.
type PresencesItemMicrosoftGraphClearPresenceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PresencesItemMicrosoftGraphClearPresenceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PresencesItemMicrosoftGraphClearPresenceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPresencesItemMicrosoftGraphClearPresenceRequestBuilderInternal instantiates a new MicrosoftGraphClearPresenceRequestBuilder and sets the default values.
func NewPresencesItemMicrosoftGraphClearPresenceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresencesItemMicrosoftGraphClearPresenceRequestBuilder) {
    m := &PresencesItemMicrosoftGraphClearPresenceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/presences/{presence%2Did}/microsoft.graph.clearPresence";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewPresencesItemMicrosoftGraphClearPresenceRequestBuilder instantiates a new MicrosoftGraphClearPresenceRequestBuilder and sets the default values.
func NewPresencesItemMicrosoftGraphClearPresenceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresencesItemMicrosoftGraphClearPresenceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPresencesItemMicrosoftGraphClearPresenceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post clear the application's presence session for a user. If it is the user's only presence session, the user's presence will change to `Offline/Offline`. For details about presences sessions, see presence: setPresence.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/presence-clearpresence?view=graph-rest-1.0
func (m *PresencesItemMicrosoftGraphClearPresenceRequestBuilder) Post(ctx context.Context, body PresencesItemMicrosoftGraphClearPresenceClearPresencePostRequestBodyable, requestConfiguration *PresencesItemMicrosoftGraphClearPresenceRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation clear the application's presence session for a user. If it is the user's only presence session, the user's presence will change to `Offline/Offline`. For details about presences sessions, see presence: setPresence.
func (m *PresencesItemMicrosoftGraphClearPresenceRequestBuilder) ToPostRequestInformation(ctx context.Context, body PresencesItemMicrosoftGraphClearPresenceClearPresencePostRequestBodyable, requestConfiguration *PresencesItemMicrosoftGraphClearPresenceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
