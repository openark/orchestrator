package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// OutlookRequestBuilder provides operations to manage the outlook property of the microsoft.graph.user entity.
type OutlookRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OutlookRequestBuilderGetQueryParameters get outlook from me
type OutlookRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OutlookRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OutlookRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OutlookRequestBuilderGetQueryParameters
}
// NewOutlookRequestBuilderInternal instantiates a new OutlookRequestBuilder and sets the default values.
func NewOutlookRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutlookRequestBuilder) {
    m := &OutlookRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/outlook{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewOutlookRequestBuilder instantiates a new OutlookRequestBuilder and sets the default values.
func NewOutlookRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutlookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutlookRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get outlook from me
func (m *OutlookRequestBuilder) Get(ctx context.Context, requestConfiguration *OutlookRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OutlookUserable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOutlookUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OutlookUserable), nil
}
// MasterCategories provides operations to manage the masterCategories property of the microsoft.graph.outlookUser entity.
func (m *OutlookRequestBuilder) MasterCategories()(*OutlookMasterCategoriesRequestBuilder) {
    return NewOutlookMasterCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MasterCategoriesById provides operations to manage the masterCategories property of the microsoft.graph.outlookUser entity.
func (m *OutlookRequestBuilder) MasterCategoriesById(id string)(*OutlookMasterCategoriesOutlookCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookCategory%2Did"] = id
    }
    return NewOutlookMasterCategoriesOutlookCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphSupportedLanguages provides operations to call the supportedLanguages method.
func (m *OutlookRequestBuilder) MicrosoftGraphSupportedLanguages()(*OutlookMicrosoftGraphSupportedLanguagesRequestBuilder) {
    return NewOutlookMicrosoftGraphSupportedLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSupportedTimeZones provides operations to call the supportedTimeZones method.
func (m *OutlookRequestBuilder) MicrosoftGraphSupportedTimeZones()(*OutlookMicrosoftGraphSupportedTimeZonesRequestBuilder) {
    return NewOutlookMicrosoftGraphSupportedTimeZonesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSupportedTimeZonesWithTimeZoneStandard provides operations to call the supportedTimeZones method.
func (m *OutlookRequestBuilder) MicrosoftGraphSupportedTimeZonesWithTimeZoneStandard(timeZoneStandard *string)(*OutlookMicrosoftGraphSupportedTimeZonesWithTimeZoneStandardRequestBuilder) {
    return NewOutlookMicrosoftGraphSupportedTimeZonesWithTimeZoneStandardRequestBuilderInternal(m.pathParameters, m.requestAdapter, timeZoneStandard)
}
// ToGetRequestInformation get outlook from me
func (m *OutlookRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OutlookRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
