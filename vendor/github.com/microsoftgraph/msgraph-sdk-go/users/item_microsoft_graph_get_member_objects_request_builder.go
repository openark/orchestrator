package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMicrosoftGraphGetMemberObjectsRequestBuilder provides operations to call the getMemberObjects method.
type ItemMicrosoftGraphGetMemberObjectsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMicrosoftGraphGetMemberObjectsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMicrosoftGraphGetMemberObjectsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMicrosoftGraphGetMemberObjectsRequestBuilderInternal instantiates a new MicrosoftGraphGetMemberObjectsRequestBuilder and sets the default values.
func NewItemMicrosoftGraphGetMemberObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphGetMemberObjectsRequestBuilder) {
    m := &ItemMicrosoftGraphGetMemberObjectsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/microsoft.graph.getMemberObjects";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMicrosoftGraphGetMemberObjectsRequestBuilder instantiates a new MicrosoftGraphGetMemberObjectsRequestBuilder and sets the default values.
func NewItemMicrosoftGraphGetMemberObjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphGetMemberObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMicrosoftGraphGetMemberObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post return all IDs for the groups, administrative units, and directory roles that a user, group, service principal, organizational contact, device, or directory object is a member of. This function is transitive. **Note:** Only users and role-enabled groups can be members of directory roles.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/directoryobject-getmemberobjects?view=graph-rest-1.0
func (m *ItemMicrosoftGraphGetMemberObjectsRequestBuilder) Post(ctx context.Context, body ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostRequestBodyable, requestConfiguration *ItemMicrosoftGraphGetMemberObjectsRequestBuilderPostRequestConfiguration)(ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemMicrosoftGraphGetMemberObjectsGetMemberObjectsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsResponseable), nil
}
// ToPostRequestInformation return all IDs for the groups, administrative units, and directory roles that a user, group, service principal, organizational contact, device, or directory object is a member of. This function is transitive. **Note:** Only users and role-enabled groups can be members of directory roles.
func (m *ItemMicrosoftGraphGetMemberObjectsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostRequestBodyable, requestConfiguration *ItemMicrosoftGraphGetMemberObjectsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
