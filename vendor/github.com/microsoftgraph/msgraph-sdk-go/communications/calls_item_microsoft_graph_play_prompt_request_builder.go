package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CallsItemMicrosoftGraphPlayPromptRequestBuilder provides operations to call the playPrompt method.
type CallsItemMicrosoftGraphPlayPromptRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CallsItemMicrosoftGraphPlayPromptRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemMicrosoftGraphPlayPromptRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallsItemMicrosoftGraphPlayPromptRequestBuilderInternal instantiates a new MicrosoftGraphPlayPromptRequestBuilder and sets the default values.
func NewCallsItemMicrosoftGraphPlayPromptRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemMicrosoftGraphPlayPromptRequestBuilder) {
    m := &CallsItemMicrosoftGraphPlayPromptRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/calls/{call%2Did}/microsoft.graph.playPrompt";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCallsItemMicrosoftGraphPlayPromptRequestBuilder instantiates a new MicrosoftGraphPlayPromptRequestBuilder and sets the default values.
func NewCallsItemMicrosoftGraphPlayPromptRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemMicrosoftGraphPlayPromptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsItemMicrosoftGraphPlayPromptRequestBuilderInternal(urlParams, requestAdapter)
}
// Post play a prompt in the call. For more information about how to handle operations, see commsOperation
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/call-playprompt?view=graph-rest-1.0
func (m *CallsItemMicrosoftGraphPlayPromptRequestBuilder) Post(ctx context.Context, body CallsItemMicrosoftGraphPlayPromptPlayPromptPostRequestBodyable, requestConfiguration *CallsItemMicrosoftGraphPlayPromptRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlayPromptOperationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePlayPromptOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlayPromptOperationable), nil
}
// ToPostRequestInformation play a prompt in the call. For more information about how to handle operations, see commsOperation
func (m *CallsItemMicrosoftGraphPlayPromptRequestBuilder) ToPostRequestInformation(ctx context.Context, body CallsItemMicrosoftGraphPlayPromptPlayPromptPostRequestBodyable, requestConfiguration *CallsItemMicrosoftGraphPlayPromptRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
