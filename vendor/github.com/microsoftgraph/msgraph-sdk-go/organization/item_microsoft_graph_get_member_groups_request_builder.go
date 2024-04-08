package organization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMicrosoftGraphGetMemberGroupsRequestBuilder provides operations to call the getMemberGroups method.
type ItemMicrosoftGraphGetMemberGroupsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMicrosoftGraphGetMemberGroupsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMicrosoftGraphGetMemberGroupsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMicrosoftGraphGetMemberGroupsRequestBuilderInternal instantiates a new MicrosoftGraphGetMemberGroupsRequestBuilder and sets the default values.
func NewItemMicrosoftGraphGetMemberGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphGetMemberGroupsRequestBuilder) {
    m := &ItemMicrosoftGraphGetMemberGroupsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/organization/{organization%2Did}/microsoft.graph.getMemberGroups";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMicrosoftGraphGetMemberGroupsRequestBuilder instantiates a new MicrosoftGraphGetMemberGroupsRequestBuilder and sets the default values.
func NewItemMicrosoftGraphGetMemberGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphGetMemberGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMicrosoftGraphGetMemberGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post return all the group IDs for the groups that the specified user, group, service principal, organizational contact, device, or directory object is a member of. This function is transitive.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/directoryobject-getmembergroups?view=graph-rest-1.0
func (m *ItemMicrosoftGraphGetMemberGroupsRequestBuilder) Post(ctx context.Context, body ItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBodyable, requestConfiguration *ItemMicrosoftGraphGetMemberGroupsRequestBuilderPostRequestConfiguration)(ItemMicrosoftGraphGetMemberGroupsGetMemberGroupsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemMicrosoftGraphGetMemberGroupsGetMemberGroupsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemMicrosoftGraphGetMemberGroupsGetMemberGroupsResponseable), nil
}
// ToPostRequestInformation return all the group IDs for the groups that the specified user, group, service principal, organizational contact, device, or directory object is a member of. This function is transitive.
func (m *ItemMicrosoftGraphGetMemberGroupsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBodyable, requestConfiguration *ItemMicrosoftGraphGetMemberGroupsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
