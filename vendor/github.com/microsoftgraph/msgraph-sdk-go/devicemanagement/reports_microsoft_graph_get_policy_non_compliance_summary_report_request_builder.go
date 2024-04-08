package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportRequestBuilder provides operations to call the getPolicyNonComplianceSummaryReport method.
type ReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportRequestBuilderInternal instantiates a new MicrosoftGraphGetPolicyNonComplianceSummaryReportRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportRequestBuilder) {
    m := &ReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getPolicyNonComplianceSummaryReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportRequestBuilder instantiates a new MicrosoftGraphGetPolicyNonComplianceSummaryReportRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getPolicyNonComplianceSummaryReport
func (m *ReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportRequestBuilder) Post(ctx context.Context, body ReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportGetPolicyNonComplianceSummaryReportPostRequestBodyable, requestConfiguration *ReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportRequestBuilderPostRequestConfiguration)([]byte, error) {
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
// ToPostRequestInformation invoke action getPolicyNonComplianceSummaryReport
func (m *ReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportGetPolicyNonComplianceSummaryReportPostRequestBodyable, requestConfiguration *ReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
