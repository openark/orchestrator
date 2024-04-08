package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CallsItemMicrosoftGraphCancelMediaProcessingRequestBuilder provides operations to call the cancelMediaProcessing method.
type CallsItemMicrosoftGraphCancelMediaProcessingRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CallsItemMicrosoftGraphCancelMediaProcessingRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemMicrosoftGraphCancelMediaProcessingRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallsItemMicrosoftGraphCancelMediaProcessingRequestBuilderInternal instantiates a new MicrosoftGraphCancelMediaProcessingRequestBuilder and sets the default values.
func NewCallsItemMicrosoftGraphCancelMediaProcessingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemMicrosoftGraphCancelMediaProcessingRequestBuilder) {
    m := &CallsItemMicrosoftGraphCancelMediaProcessingRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/calls/{call%2Did}/microsoft.graph.cancelMediaProcessing";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCallsItemMicrosoftGraphCancelMediaProcessingRequestBuilder instantiates a new MicrosoftGraphCancelMediaProcessingRequestBuilder and sets the default values.
func NewCallsItemMicrosoftGraphCancelMediaProcessingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemMicrosoftGraphCancelMediaProcessingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsItemMicrosoftGraphCancelMediaProcessingRequestBuilderInternal(urlParams, requestAdapter)
}
// Post cancels processing for any in-progress media operations. Media operations refer to the IVR operations playPrompt and recordResponse, which are by default queued to process in order. The **cancelMediaProcessing** method cancels any operation that is in-process as well as operations that are queued. For example, this method can be used to clean up the IVR operation queue for a new media operation. However, it will not cancel a **subscribeToTone** operation because it operates independent of any operation queue.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/call-cancelmediaprocessing?view=graph-rest-1.0
func (m *CallsItemMicrosoftGraphCancelMediaProcessingRequestBuilder) Post(ctx context.Context, body CallsItemMicrosoftGraphCancelMediaProcessingCancelMediaProcessingPostRequestBodyable, requestConfiguration *CallsItemMicrosoftGraphCancelMediaProcessingRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CancelMediaProcessingOperationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCancelMediaProcessingOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CancelMediaProcessingOperationable), nil
}
// ToPostRequestInformation cancels processing for any in-progress media operations. Media operations refer to the IVR operations playPrompt and recordResponse, which are by default queued to process in order. The **cancelMediaProcessing** method cancels any operation that is in-process as well as operations that are queued. For example, this method can be used to clean up the IVR operation queue for a new media operation. However, it will not cancel a **subscribeToTone** operation because it operates independent of any operation queue.
func (m *CallsItemMicrosoftGraphCancelMediaProcessingRequestBuilder) ToPostRequestInformation(ctx context.Context, body CallsItemMicrosoftGraphCancelMediaProcessingCancelMediaProcessingPostRequestBodyable, requestConfiguration *CallsItemMicrosoftGraphCancelMediaProcessingRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
