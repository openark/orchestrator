package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ServiceAnnouncementHealthOverviewsItemIssuesItemMicrosoftGraphIncidentReportRequestBuilder provides operations to call the incidentReport method.
type ServiceAnnouncementHealthOverviewsItemIssuesItemMicrosoftGraphIncidentReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ServiceAnnouncementHealthOverviewsItemIssuesItemMicrosoftGraphIncidentReportRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceAnnouncementHealthOverviewsItemIssuesItemMicrosoftGraphIncidentReportRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServiceAnnouncementHealthOverviewsItemIssuesItemMicrosoftGraphIncidentReportRequestBuilderInternal instantiates a new MicrosoftGraphIncidentReportRequestBuilder and sets the default values.
func NewServiceAnnouncementHealthOverviewsItemIssuesItemMicrosoftGraphIncidentReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAnnouncementHealthOverviewsItemIssuesItemMicrosoftGraphIncidentReportRequestBuilder) {
    m := &ServiceAnnouncementHealthOverviewsItemIssuesItemMicrosoftGraphIncidentReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/serviceAnnouncement/healthOverviews/{serviceHealth%2Did}/issues/{serviceHealthIssue%2Did}/microsoft.graph.incidentReport()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewServiceAnnouncementHealthOverviewsItemIssuesItemMicrosoftGraphIncidentReportRequestBuilder instantiates a new MicrosoftGraphIncidentReportRequestBuilder and sets the default values.
func NewServiceAnnouncementHealthOverviewsItemIssuesItemMicrosoftGraphIncidentReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAnnouncementHealthOverviewsItemIssuesItemMicrosoftGraphIncidentReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceAnnouncementHealthOverviewsItemIssuesItemMicrosoftGraphIncidentReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function incidentReport
func (m *ServiceAnnouncementHealthOverviewsItemIssuesItemMicrosoftGraphIncidentReportRequestBuilder) Get(ctx context.Context, requestConfiguration *ServiceAnnouncementHealthOverviewsItemIssuesItemMicrosoftGraphIncidentReportRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// ToGetRequestInformation invoke function incidentReport
func (m *ServiceAnnouncementHealthOverviewsItemIssuesItemMicrosoftGraphIncidentReportRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServiceAnnouncementHealthOverviewsItemIssuesItemMicrosoftGraphIncidentReportRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
