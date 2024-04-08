package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MicrosoftGraphExportPersonalDataRequestBuilder provides operations to call the exportPersonalData method.
type MicrosoftGraphExportPersonalDataRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftGraphExportPersonalDataRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphExportPersonalDataRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftGraphExportPersonalDataRequestBuilderInternal instantiates a new MicrosoftGraphExportPersonalDataRequestBuilder and sets the default values.
func NewMicrosoftGraphExportPersonalDataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphExportPersonalDataRequestBuilder) {
    m := &MicrosoftGraphExportPersonalDataRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/microsoft.graph.exportPersonalData";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMicrosoftGraphExportPersonalDataRequestBuilder instantiates a new MicrosoftGraphExportPersonalDataRequestBuilder and sets the default values.
func NewMicrosoftGraphExportPersonalDataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphExportPersonalDataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphExportPersonalDataRequestBuilderInternal(urlParams, requestAdapter)
}
// Post submit a data policy operation request from a company administrator or an application to export an organizational user's data. This data includes the user's data stored in OneDrive and their activity reports. For more guidance about exporting data while complying with regulations, see Data Subject Requests and the GDPR and CCPA.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/user-exportpersonaldata?view=graph-rest-1.0
func (m *MicrosoftGraphExportPersonalDataRequestBuilder) Post(ctx context.Context, body MicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBodyable, requestConfiguration *MicrosoftGraphExportPersonalDataRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation submit a data policy operation request from a company administrator or an application to export an organizational user's data. This data includes the user's data stored in OneDrive and their activity reports. For more guidance about exporting data while complying with regulations, see Data Subject Requests and the GDPR and CCPA.
func (m *MicrosoftGraphExportPersonalDataRequestBuilder) ToPostRequestInformation(ctx context.Context, body MicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBodyable, requestConfiguration *MicrosoftGraphExportPersonalDataRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
