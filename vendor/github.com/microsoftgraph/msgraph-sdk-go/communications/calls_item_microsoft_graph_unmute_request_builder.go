package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CallsItemMicrosoftGraphUnmuteRequestBuilder provides operations to call the unmute method.
type CallsItemMicrosoftGraphUnmuteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CallsItemMicrosoftGraphUnmuteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemMicrosoftGraphUnmuteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallsItemMicrosoftGraphUnmuteRequestBuilderInternal instantiates a new MicrosoftGraphUnmuteRequestBuilder and sets the default values.
func NewCallsItemMicrosoftGraphUnmuteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemMicrosoftGraphUnmuteRequestBuilder) {
    m := &CallsItemMicrosoftGraphUnmuteRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/calls/{call%2Did}/microsoft.graph.unmute";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCallsItemMicrosoftGraphUnmuteRequestBuilder instantiates a new MicrosoftGraphUnmuteRequestBuilder and sets the default values.
func NewCallsItemMicrosoftGraphUnmuteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemMicrosoftGraphUnmuteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsItemMicrosoftGraphUnmuteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post allow the application to unmute itself. This is a server unmute, meaning that the server will start sending audio packets for this participant to other participants again. For more information about how to handle unmute operations, see unmuteParticipantOperation.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/call-unmute?view=graph-rest-1.0
func (m *CallsItemMicrosoftGraphUnmuteRequestBuilder) Post(ctx context.Context, body CallsItemMicrosoftGraphUnmuteUnmutePostRequestBodyable, requestConfiguration *CallsItemMicrosoftGraphUnmuteRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnmuteParticipantOperationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnmuteParticipantOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnmuteParticipantOperationable), nil
}
// ToPostRequestInformation allow the application to unmute itself. This is a server unmute, meaning that the server will start sending audio packets for this participant to other participants again. For more information about how to handle unmute operations, see unmuteParticipantOperation.
func (m *CallsItemMicrosoftGraphUnmuteRequestBuilder) ToPostRequestInformation(ctx context.Context, body CallsItemMicrosoftGraphUnmuteUnmutePostRequestBodyable, requestConfiguration *CallsItemMicrosoftGraphUnmuteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
