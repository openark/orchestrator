package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesSiteItemRequestBuilder provides operations to manage the sites property of the microsoft.graph.group entity.
type ItemSitesSiteItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesSiteItemRequestBuilderGetQueryParameters the list of SharePoint sites in this group. Access the default site with /sites/root.
type ItemSitesSiteItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesSiteItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesSiteItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesSiteItemRequestBuilderGetQueryParameters
}
// ItemSitesSiteItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesSiteItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Analytics()(*ItemSitesItemAnalyticsRequestBuilder) {
    return NewItemSitesItemAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Columns()(*ItemSitesItemColumnsRequestBuilder) {
    return NewItemSitesItemColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ColumnsById provides operations to manage the columns property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) ColumnsById(id string)(*ItemSitesItemColumnsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return NewItemSitesItemColumnsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemSitesSiteItemRequestBuilderInternal instantiates a new SiteItemRequestBuilder and sets the default values.
func NewItemSitesSiteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesSiteItemRequestBuilder) {
    m := &ItemSitesSiteItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesSiteItemRequestBuilder instantiates a new SiteItemRequestBuilder and sets the default values.
func NewItemSitesSiteItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesSiteItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesSiteItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentTypes provides operations to manage the contentTypes property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) ContentTypes()(*ItemSitesItemContentTypesRequestBuilder) {
    return NewItemSitesItemContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ContentTypesById provides operations to manage the contentTypes property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) ContentTypesById(id string)(*ItemSitesItemContentTypesContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did"] = id
    }
    return NewItemSitesItemContentTypesContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Drive provides operations to manage the drive property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Drive()(*ItemSitesItemDriveRequestBuilder) {
    return NewItemSitesItemDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Drives provides operations to manage the drives property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Drives()(*ItemSitesItemDrivesRequestBuilder) {
    return NewItemSitesItemDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DrivesById provides operations to manage the drives property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) DrivesById(id string)(*ItemSitesItemDrivesDriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return NewItemSitesItemDrivesDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ExternalColumns provides operations to manage the externalColumns property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) ExternalColumns()(*ItemSitesItemExternalColumnsRequestBuilder) {
    return NewItemSitesItemExternalColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExternalColumnsById provides operations to manage the externalColumns property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) ExternalColumnsById(id string)(*ItemSitesItemExternalColumnsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return NewItemSitesItemExternalColumnsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *ItemSitesSiteItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesSiteItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSiteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable), nil
}
// Items provides operations to manage the items property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Items()(*ItemSitesItemItemsRequestBuilder) {
    return NewItemSitesItemItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ItemsById provides operations to manage the items property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) ItemsById(id string)(*ItemSitesItemItemsBaseItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["baseItem%2Did"] = id
    }
    return NewItemSitesItemItemsBaseItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Lists provides operations to manage the lists property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Lists()(*ItemSitesItemListsRequestBuilder) {
    return NewItemSitesItemListsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ListsById provides operations to manage the lists property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) ListsById(id string)(*ItemSitesItemListsListItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["list%2Did"] = id
    }
    return NewItemSitesItemListsListItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphGetActivitiesByInterval provides operations to call the getActivitiesByInterval method.
func (m *ItemSitesSiteItemRequestBuilder) MicrosoftGraphGetActivitiesByInterval()(*ItemSitesItemMicrosoftGraphGetActivitiesByIntervalRequestBuilder) {
    return NewItemSitesItemMicrosoftGraphGetActivitiesByIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *ItemSitesSiteItemRequestBuilder) MicrosoftGraphGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*ItemSitesItemMicrosoftGraphGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewItemSitesItemMicrosoftGraphGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime)
}
// MicrosoftGraphGetApplicableContentTypesForListWithListId provides operations to call the getApplicableContentTypesForList method.
func (m *ItemSitesSiteItemRequestBuilder) MicrosoftGraphGetApplicableContentTypesForListWithListId(listId *string)(*ItemSitesItemMicrosoftGraphGetApplicableContentTypesForListWithListIdRequestBuilder) {
    return NewItemSitesItemMicrosoftGraphGetApplicableContentTypesForListWithListIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, listId)
}
// MicrosoftGraphGetByPathWithPath provides operations to call the getByPath method.
func (m *ItemSitesSiteItemRequestBuilder) MicrosoftGraphGetByPathWithPath(path *string)(*ItemSitesItemMicrosoftGraphGetByPathWithPathRequestBuilder) {
    return NewItemSitesItemMicrosoftGraphGetByPathWithPathRequestBuilderInternal(m.pathParameters, m.requestAdapter, path)
}
// Onenote provides operations to manage the onenote property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Onenote()(*ItemSitesItemOnenoteRequestBuilder) {
    return NewItemSitesItemOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Operations()(*ItemSitesItemOperationsRequestBuilder) {
    return NewItemSitesItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) OperationsById(id string)(*ItemSitesItemOperationsRichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation%2Did"] = id
    }
    return NewItemSitesItemOperationsRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property sites in groups
func (m *ItemSitesSiteItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable, requestConfiguration *ItemSitesSiteItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSiteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable), nil
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Permissions()(*ItemSitesItemPermissionsRequestBuilder) {
    return NewItemSitesItemPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PermissionsById provides operations to manage the permissions property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) PermissionsById(id string)(*ItemSitesItemPermissionsPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return NewItemSitesItemPermissionsPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Sites provides operations to manage the sites property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) Sites()(*ItemSitesItemSitesRequestBuilder) {
    return NewItemSitesItemSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SitesById provides operations to manage the sites property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) SitesById(id string)(*ItemSitesItemSitesSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did1"] = id
    }
    return NewItemSitesItemSitesSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// TermStore provides operations to manage the termStore property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) TermStore()(*ItemSitesItemTermStoreRequestBuilder) {
    return NewItemSitesItemTermStoreRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TermStores provides operations to manage the termStores property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) TermStores()(*ItemSitesItemTermStoresRequestBuilder) {
    return NewItemSitesItemTermStoresRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TermStoresById provides operations to manage the termStores property of the microsoft.graph.site entity.
func (m *ItemSitesSiteItemRequestBuilder) TermStoresById(id string)(*ItemSitesItemTermStoresStoreItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["store%2Did"] = id
    }
    return NewItemSitesItemTermStoresStoreItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation the list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *ItemSitesSiteItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesSiteItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property sites in groups
func (m *ItemSitesSiteItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable, requestConfiguration *ItemSitesSiteItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
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
