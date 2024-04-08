package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MicrosoftGraphGetPresencesByUserIdRequestBuilder provides operations to call the getPresencesByUserId method.
type MicrosoftGraphGetPresencesByUserIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftGraphGetPresencesByUserIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphGetPresencesByUserIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftGraphGetPresencesByUserIdRequestBuilderInternal instantiates a new MicrosoftGraphGetPresencesByUserIdRequestBuilder and sets the default values.
func NewMicrosoftGraphGetPresencesByUserIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphGetPresencesByUserIdRequestBuilder) {
    m := &MicrosoftGraphGetPresencesByUserIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/microsoft.graph.getPresencesByUserId";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMicrosoftGraphGetPresencesByUserIdRequestBuilder instantiates a new MicrosoftGraphGetPresencesByUserIdRequestBuilder and sets the default values.
func NewMicrosoftGraphGetPresencesByUserIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphGetPresencesByUserIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphGetPresencesByUserIdRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get the presence information for multiple users.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/cloudcommunications-getpresencesbyuserid?view=graph-rest-1.0
func (m *MicrosoftGraphGetPresencesByUserIdRequestBuilder) Post(ctx context.Context, body MicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdPostRequestBodyable, requestConfiguration *MicrosoftGraphGetPresencesByUserIdRequestBuilderPostRequestConfiguration)(MicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateMicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdResponseable), nil
}
// ToPostRequestInformation get the presence information for multiple users.
func (m *MicrosoftGraphGetPresencesByUserIdRequestBuilder) ToPostRequestInformation(ctx context.Context, body MicrosoftGraphGetPresencesByUserIdGetPresencesByUserIdPostRequestBodyable, requestConfiguration *MicrosoftGraphGetPresencesByUserIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
