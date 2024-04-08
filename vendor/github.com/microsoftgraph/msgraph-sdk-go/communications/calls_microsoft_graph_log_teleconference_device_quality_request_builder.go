package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CallsMicrosoftGraphLogTeleconferenceDeviceQualityRequestBuilder provides operations to call the logTeleconferenceDeviceQuality method.
type CallsMicrosoftGraphLogTeleconferenceDeviceQualityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CallsMicrosoftGraphLogTeleconferenceDeviceQualityRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsMicrosoftGraphLogTeleconferenceDeviceQualityRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallsMicrosoftGraphLogTeleconferenceDeviceQualityRequestBuilderInternal instantiates a new MicrosoftGraphLogTeleconferenceDeviceQualityRequestBuilder and sets the default values.
func NewCallsMicrosoftGraphLogTeleconferenceDeviceQualityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsMicrosoftGraphLogTeleconferenceDeviceQualityRequestBuilder) {
    m := &CallsMicrosoftGraphLogTeleconferenceDeviceQualityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/calls/microsoft.graph.logTeleconferenceDeviceQuality";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCallsMicrosoftGraphLogTeleconferenceDeviceQualityRequestBuilder instantiates a new MicrosoftGraphLogTeleconferenceDeviceQualityRequestBuilder and sets the default values.
func NewCallsMicrosoftGraphLogTeleconferenceDeviceQualityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsMicrosoftGraphLogTeleconferenceDeviceQualityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsMicrosoftGraphLogTeleconferenceDeviceQualityRequestBuilderInternal(urlParams, requestAdapter)
}
// Post log video teleconferencing device quality data. The Cloud Video Interop (CVI) bot represents video teleconferencing (VTC) devices and acts as a back-to-back agent for a VTC device in a conference call. Because a CVI bot is in the middle of the VTC and Microsoft Teams infrastructure as a VTC proxy, it has two media legs. One media leg is between the CVI bot and Teams infrastructure, such as Teams conference server or a Teams client. The other media leg is between the CVI bot and the VTC device.  The third-party partners own the VTC media leg and the Teams infrastructure cannot access the quality data of the third-party call leg.  This method is only for the CVI partners to provide their media quality data.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/call-logteleconferencedevicequality?view=graph-rest-1.0
func (m *CallsMicrosoftGraphLogTeleconferenceDeviceQualityRequestBuilder) Post(ctx context.Context, body CallsMicrosoftGraphLogTeleconferenceDeviceQualityLogTeleconferenceDeviceQualityPostRequestBodyable, requestConfiguration *CallsMicrosoftGraphLogTeleconferenceDeviceQualityRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation log video teleconferencing device quality data. The Cloud Video Interop (CVI) bot represents video teleconferencing (VTC) devices and acts as a back-to-back agent for a VTC device in a conference call. Because a CVI bot is in the middle of the VTC and Microsoft Teams infrastructure as a VTC proxy, it has two media legs. One media leg is between the CVI bot and Teams infrastructure, such as Teams conference server or a Teams client. The other media leg is between the CVI bot and the VTC device.  The third-party partners own the VTC media leg and the Teams infrastructure cannot access the quality data of the third-party call leg.  This method is only for the CVI partners to provide their media quality data.
func (m *CallsMicrosoftGraphLogTeleconferenceDeviceQualityRequestBuilder) ToPostRequestInformation(ctx context.Context, body CallsMicrosoftGraphLogTeleconferenceDeviceQualityLogTeleconferenceDeviceQualityPostRequestBodyable, requestConfiguration *CallsMicrosoftGraphLogTeleconferenceDeviceQualityRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
