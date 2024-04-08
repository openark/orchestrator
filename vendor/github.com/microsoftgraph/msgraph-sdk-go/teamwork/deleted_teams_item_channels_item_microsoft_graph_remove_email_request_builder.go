package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeletedTeamsItemChannelsItemMicrosoftGraphRemoveEmailRequestBuilder provides operations to call the removeEmail method.
type DeletedTeamsItemChannelsItemMicrosoftGraphRemoveEmailRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeletedTeamsItemChannelsItemMicrosoftGraphRemoveEmailRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedTeamsItemChannelsItemMicrosoftGraphRemoveEmailRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeletedTeamsItemChannelsItemMicrosoftGraphRemoveEmailRequestBuilderInternal instantiates a new MicrosoftGraphRemoveEmailRequestBuilder and sets the default values.
func NewDeletedTeamsItemChannelsItemMicrosoftGraphRemoveEmailRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedTeamsItemChannelsItemMicrosoftGraphRemoveEmailRequestBuilder) {
    m := &DeletedTeamsItemChannelsItemMicrosoftGraphRemoveEmailRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}/microsoft.graph.removeEmail";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDeletedTeamsItemChannelsItemMicrosoftGraphRemoveEmailRequestBuilder instantiates a new MicrosoftGraphRemoveEmailRequestBuilder and sets the default values.
func NewDeletedTeamsItemChannelsItemMicrosoftGraphRemoveEmailRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedTeamsItemChannelsItemMicrosoftGraphRemoveEmailRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedTeamsItemChannelsItemMicrosoftGraphRemoveEmailRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove the email address of a channel. You can remove an email address only if it was provisioned using the provisionEmail method or through the Microsoft Teams client.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/channel-removeemail?view=graph-rest-1.0
func (m *DeletedTeamsItemChannelsItemMicrosoftGraphRemoveEmailRequestBuilder) Post(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemMicrosoftGraphRemoveEmailRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation remove the email address of a channel. You can remove an email address only if it was provisioned using the provisionEmail method or through the Microsoft Teams client.
func (m *DeletedTeamsItemChannelsItemMicrosoftGraphRemoveEmailRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemMicrosoftGraphRemoveEmailRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
