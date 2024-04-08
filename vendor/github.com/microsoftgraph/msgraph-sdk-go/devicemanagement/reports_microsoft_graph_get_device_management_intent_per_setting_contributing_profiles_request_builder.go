package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder provides operations to call the getDeviceManagementIntentPerSettingContributingProfiles method.
type ReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderInternal instantiates a new MicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder) {
    m := &ReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getDeviceManagementIntentPerSettingContributingProfiles";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder instantiates a new MicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getDeviceManagementIntentPerSettingContributingProfiles
func (m *ReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder) Post(ctx context.Context, body ReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesGetDeviceManagementIntentPerSettingContributingProfilesPostRequestBodyable, requestConfiguration *ReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToPostRequestInformation invoke action getDeviceManagementIntentPerSettingContributingProfiles
func (m *ReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesGetDeviceManagementIntentPerSettingContributingProfilesPostRequestBodyable, requestConfiguration *ReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
