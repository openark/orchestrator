package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedTeamsItemMembersMicrosoftGraphAddRequestBuilder provides operations to call the add method.
type ItemJoinedTeamsItemMembersMicrosoftGraphAddRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemJoinedTeamsItemMembersMicrosoftGraphAddRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemMembersMicrosoftGraphAddRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemJoinedTeamsItemMembersMicrosoftGraphAddRequestBuilderInternal instantiates a new MicrosoftGraphAddRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemMembersMicrosoftGraphAddRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemMembersMicrosoftGraphAddRequestBuilder) {
    m := &ItemJoinedTeamsItemMembersMicrosoftGraphAddRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/members/microsoft.graph.add";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemJoinedTeamsItemMembersMicrosoftGraphAddRequestBuilder instantiates a new MicrosoftGraphAddRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemMembersMicrosoftGraphAddRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemMembersMicrosoftGraphAddRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedTeamsItemMembersMicrosoftGraphAddRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add multiple members in a single request to a team. The response provides details about which memberships could and couldn't be created.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/conversationmembers-add?view=graph-rest-1.0
func (m *ItemJoinedTeamsItemMembersMicrosoftGraphAddRequestBuilder) Post(ctx context.Context, body ItemJoinedTeamsItemMembersMicrosoftGraphAddAddPostRequestBodyable, requestConfiguration *ItemJoinedTeamsItemMembersMicrosoftGraphAddRequestBuilderPostRequestConfiguration)(ItemJoinedTeamsItemMembersMicrosoftGraphAddAddResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemJoinedTeamsItemMembersMicrosoftGraphAddAddResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemJoinedTeamsItemMembersMicrosoftGraphAddAddResponseable), nil
}
// ToPostRequestInformation add multiple members in a single request to a team. The response provides details about which memberships could and couldn't be created.
func (m *ItemJoinedTeamsItemMembersMicrosoftGraphAddRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemJoinedTeamsItemMembersMicrosoftGraphAddAddPostRequestBodyable, requestConfiguration *ItemJoinedTeamsItemMembersMicrosoftGraphAddRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
