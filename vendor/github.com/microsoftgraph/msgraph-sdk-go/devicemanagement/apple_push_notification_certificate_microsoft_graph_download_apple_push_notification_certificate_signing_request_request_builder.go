package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ApplePushNotificationCertificateMicrosoftGraphDownloadApplePushNotificationCertificateSigningRequestRequestBuilder provides operations to call the downloadApplePushNotificationCertificateSigningRequest method.
type ApplePushNotificationCertificateMicrosoftGraphDownloadApplePushNotificationCertificateSigningRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ApplePushNotificationCertificateMicrosoftGraphDownloadApplePushNotificationCertificateSigningRequestRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplePushNotificationCertificateMicrosoftGraphDownloadApplePushNotificationCertificateSigningRequestRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewApplePushNotificationCertificateMicrosoftGraphDownloadApplePushNotificationCertificateSigningRequestRequestBuilderInternal instantiates a new MicrosoftGraphDownloadApplePushNotificationCertificateSigningRequestRequestBuilder and sets the default values.
func NewApplePushNotificationCertificateMicrosoftGraphDownloadApplePushNotificationCertificateSigningRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplePushNotificationCertificateMicrosoftGraphDownloadApplePushNotificationCertificateSigningRequestRequestBuilder) {
    m := &ApplePushNotificationCertificateMicrosoftGraphDownloadApplePushNotificationCertificateSigningRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/applePushNotificationCertificate/microsoft.graph.downloadApplePushNotificationCertificateSigningRequest()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewApplePushNotificationCertificateMicrosoftGraphDownloadApplePushNotificationCertificateSigningRequestRequestBuilder instantiates a new MicrosoftGraphDownloadApplePushNotificationCertificateSigningRequestRequestBuilder and sets the default values.
func NewApplePushNotificationCertificateMicrosoftGraphDownloadApplePushNotificationCertificateSigningRequestRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplePushNotificationCertificateMicrosoftGraphDownloadApplePushNotificationCertificateSigningRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApplePushNotificationCertificateMicrosoftGraphDownloadApplePushNotificationCertificateSigningRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// Get download Apple push notification certificate signing request
func (m *ApplePushNotificationCertificateMicrosoftGraphDownloadApplePushNotificationCertificateSigningRequestRequestBuilder) Get(ctx context.Context, requestConfiguration *ApplePushNotificationCertificateMicrosoftGraphDownloadApplePushNotificationCertificateSigningRequestRequestBuilderGetRequestConfiguration)(ApplePushNotificationCertificateMicrosoftGraphDownloadApplePushNotificationCertificateSigningRequestDownloadApplePushNotificationCertificateSigningRequestResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateApplePushNotificationCertificateMicrosoftGraphDownloadApplePushNotificationCertificateSigningRequestDownloadApplePushNotificationCertificateSigningRequestResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ApplePushNotificationCertificateMicrosoftGraphDownloadApplePushNotificationCertificateSigningRequestDownloadApplePushNotificationCertificateSigningRequestResponseable), nil
}
// ToGetRequestInformation download Apple push notification certificate signing request
func (m *ApplePushNotificationCertificateMicrosoftGraphDownloadApplePushNotificationCertificateSigningRequestRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ApplePushNotificationCertificateMicrosoftGraphDownloadApplePushNotificationCertificateSigningRequestRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
