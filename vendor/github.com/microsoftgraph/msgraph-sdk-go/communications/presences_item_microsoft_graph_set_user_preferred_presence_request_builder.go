package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PresencesItemMicrosoftGraphSetUserPreferredPresenceRequestBuilder provides operations to call the setUserPreferredPresence method.
type PresencesItemMicrosoftGraphSetUserPreferredPresenceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PresencesItemMicrosoftGraphSetUserPreferredPresenceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PresencesItemMicrosoftGraphSetUserPreferredPresenceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPresencesItemMicrosoftGraphSetUserPreferredPresenceRequestBuilderInternal instantiates a new MicrosoftGraphSetUserPreferredPresenceRequestBuilder and sets the default values.
func NewPresencesItemMicrosoftGraphSetUserPreferredPresenceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresencesItemMicrosoftGraphSetUserPreferredPresenceRequestBuilder) {
    m := &PresencesItemMicrosoftGraphSetUserPreferredPresenceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/presences/{presence%2Did}/microsoft.graph.setUserPreferredPresence";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewPresencesItemMicrosoftGraphSetUserPreferredPresenceRequestBuilder instantiates a new MicrosoftGraphSetUserPreferredPresenceRequestBuilder and sets the default values.
func NewPresencesItemMicrosoftGraphSetUserPreferredPresenceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresencesItemMicrosoftGraphSetUserPreferredPresenceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPresencesItemMicrosoftGraphSetUserPreferredPresenceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set the preferred availability and activity status for a user. If the preferred presence of a user is set, the user's presence shows as the preferred status. Preferred presence takes effect only when at least one presence session exists for the user. Otherwise, the user's presence shows as `Offline`. A presence session is created as a result of a successful setPresence operation, or if the user is signed in on a Microsoft Teams client. For more details, see presence sessions and time-out and expiration.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/presence-setuserpreferredpresence?view=graph-rest-1.0
func (m *PresencesItemMicrosoftGraphSetUserPreferredPresenceRequestBuilder) Post(ctx context.Context, body PresencesItemMicrosoftGraphSetUserPreferredPresenceSetUserPreferredPresencePostRequestBodyable, requestConfiguration *PresencesItemMicrosoftGraphSetUserPreferredPresenceRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation set the preferred availability and activity status for a user. If the preferred presence of a user is set, the user's presence shows as the preferred status. Preferred presence takes effect only when at least one presence session exists for the user. Otherwise, the user's presence shows as `Offline`. A presence session is created as a result of a successful setPresence operation, or if the user is signed in on a Microsoft Teams client. For more details, see presence sessions and time-out and expiration.
func (m *PresencesItemMicrosoftGraphSetUserPreferredPresenceRequestBuilder) ToPostRequestInformation(ctx context.Context, body PresencesItemMicrosoftGraphSetUserPreferredPresenceSetUserPreferredPresencePostRequestBodyable, requestConfiguration *PresencesItemMicrosoftGraphSetUserPreferredPresenceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
