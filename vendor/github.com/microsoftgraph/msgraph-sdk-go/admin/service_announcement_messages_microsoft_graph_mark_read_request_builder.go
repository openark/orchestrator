package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ServiceAnnouncementMessagesMicrosoftGraphMarkReadRequestBuilder provides operations to call the markRead method.
type ServiceAnnouncementMessagesMicrosoftGraphMarkReadRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ServiceAnnouncementMessagesMicrosoftGraphMarkReadRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceAnnouncementMessagesMicrosoftGraphMarkReadRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServiceAnnouncementMessagesMicrosoftGraphMarkReadRequestBuilderInternal instantiates a new MicrosoftGraphMarkReadRequestBuilder and sets the default values.
func NewServiceAnnouncementMessagesMicrosoftGraphMarkReadRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAnnouncementMessagesMicrosoftGraphMarkReadRequestBuilder) {
    m := &ServiceAnnouncementMessagesMicrosoftGraphMarkReadRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/serviceAnnouncement/messages/microsoft.graph.markRead";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewServiceAnnouncementMessagesMicrosoftGraphMarkReadRequestBuilder instantiates a new MicrosoftGraphMarkReadRequestBuilder and sets the default values.
func NewServiceAnnouncementMessagesMicrosoftGraphMarkReadRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAnnouncementMessagesMicrosoftGraphMarkReadRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceAnnouncementMessagesMicrosoftGraphMarkReadRequestBuilderInternal(urlParams, requestAdapter)
}
// Post mark a list of serviceUpdateMessages as **read** for the signed in user.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/serviceupdatemessage-markread?view=graph-rest-1.0
func (m *ServiceAnnouncementMessagesMicrosoftGraphMarkReadRequestBuilder) Post(ctx context.Context, body ServiceAnnouncementMessagesMicrosoftGraphMarkReadMarkReadPostRequestBodyable, requestConfiguration *ServiceAnnouncementMessagesMicrosoftGraphMarkReadRequestBuilderPostRequestConfiguration)(ServiceAnnouncementMessagesMicrosoftGraphMarkReadMarkReadResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateServiceAnnouncementMessagesMicrosoftGraphMarkReadMarkReadResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceAnnouncementMessagesMicrosoftGraphMarkReadMarkReadResponseable), nil
}
// ToPostRequestInformation mark a list of serviceUpdateMessages as **read** for the signed in user.
func (m *ServiceAnnouncementMessagesMicrosoftGraphMarkReadRequestBuilder) ToPostRequestInformation(ctx context.Context, body ServiceAnnouncementMessagesMicrosoftGraphMarkReadMarkReadPostRequestBodyable, requestConfiguration *ServiceAnnouncementMessagesMicrosoftGraphMarkReadRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
