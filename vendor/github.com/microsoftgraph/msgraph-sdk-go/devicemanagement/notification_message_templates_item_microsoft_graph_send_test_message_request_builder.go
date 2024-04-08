package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// NotificationMessageTemplatesItemMicrosoftGraphSendTestMessageRequestBuilder provides operations to call the sendTestMessage method.
type NotificationMessageTemplatesItemMicrosoftGraphSendTestMessageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NotificationMessageTemplatesItemMicrosoftGraphSendTestMessageRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NotificationMessageTemplatesItemMicrosoftGraphSendTestMessageRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewNotificationMessageTemplatesItemMicrosoftGraphSendTestMessageRequestBuilderInternal instantiates a new MicrosoftGraphSendTestMessageRequestBuilder and sets the default values.
func NewNotificationMessageTemplatesItemMicrosoftGraphSendTestMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NotificationMessageTemplatesItemMicrosoftGraphSendTestMessageRequestBuilder) {
    m := &NotificationMessageTemplatesItemMicrosoftGraphSendTestMessageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/notificationMessageTemplates/{notificationMessageTemplate%2Did}/microsoft.graph.sendTestMessage";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewNotificationMessageTemplatesItemMicrosoftGraphSendTestMessageRequestBuilder instantiates a new MicrosoftGraphSendTestMessageRequestBuilder and sets the default values.
func NewNotificationMessageTemplatesItemMicrosoftGraphSendTestMessageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NotificationMessageTemplatesItemMicrosoftGraphSendTestMessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNotificationMessageTemplatesItemMicrosoftGraphSendTestMessageRequestBuilderInternal(urlParams, requestAdapter)
}
// Post sends test message using the specified notificationMessageTemplate in the default locale
func (m *NotificationMessageTemplatesItemMicrosoftGraphSendTestMessageRequestBuilder) Post(ctx context.Context, requestConfiguration *NotificationMessageTemplatesItemMicrosoftGraphSendTestMessageRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation sends test message using the specified notificationMessageTemplate in the default locale
func (m *NotificationMessageTemplatesItemMicrosoftGraphSendTestMessageRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *NotificationMessageTemplatesItemMicrosoftGraphSendTestMessageRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
