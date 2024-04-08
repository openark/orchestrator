package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemContentTypesContentTypeItemRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.site entity.
type ItemSitesItemContentTypesContentTypeItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemContentTypesContentTypeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemContentTypesContentTypeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemContentTypesContentTypeItemRequestBuilderGetQueryParameters the collection of content types defined for this site.
type ItemSitesItemContentTypesContentTypeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemContentTypesContentTypeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemContentTypesContentTypeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemContentTypesContentTypeItemRequestBuilderGetQueryParameters
}
// ItemSitesItemContentTypesContentTypeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemContentTypesContentTypeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Base provides operations to manage the base property of the microsoft.graph.contentType entity.
func (m *ItemSitesItemContentTypesContentTypeItemRequestBuilder) Base()(*ItemSitesItemContentTypesItemBaseRequestBuilder) {
    return NewItemSitesItemContentTypesItemBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// BaseTypes provides operations to manage the baseTypes property of the microsoft.graph.contentType entity.
func (m *ItemSitesItemContentTypesContentTypeItemRequestBuilder) BaseTypes()(*ItemSitesItemContentTypesItemBaseTypesRequestBuilder) {
    return NewItemSitesItemContentTypesItemBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// BaseTypesById provides operations to manage the baseTypes property of the microsoft.graph.contentType entity.
func (m *ItemSitesItemContentTypesContentTypeItemRequestBuilder) BaseTypesById(id string)(*ItemSitesItemContentTypesItemBaseTypesContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did1"] = id
    }
    return NewItemSitesItemContentTypesItemBaseTypesContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ColumnLinks provides operations to manage the columnLinks property of the microsoft.graph.contentType entity.
func (m *ItemSitesItemContentTypesContentTypeItemRequestBuilder) ColumnLinks()(*ItemSitesItemContentTypesItemColumnLinksRequestBuilder) {
    return NewItemSitesItemContentTypesItemColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ColumnLinksById provides operations to manage the columnLinks property of the microsoft.graph.contentType entity.
func (m *ItemSitesItemContentTypesContentTypeItemRequestBuilder) ColumnLinksById(id string)(*ItemSitesItemContentTypesItemColumnLinksColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink%2Did"] = id
    }
    return NewItemSitesItemContentTypesItemColumnLinksColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ColumnPositions provides operations to manage the columnPositions property of the microsoft.graph.contentType entity.
func (m *ItemSitesItemContentTypesContentTypeItemRequestBuilder) ColumnPositions()(*ItemSitesItemContentTypesItemColumnPositionsRequestBuilder) {
    return NewItemSitesItemContentTypesItemColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ColumnPositionsById provides operations to manage the columnPositions property of the microsoft.graph.contentType entity.
func (m *ItemSitesItemContentTypesContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*ItemSitesItemContentTypesItemColumnPositionsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return NewItemSitesItemContentTypesItemColumnPositionsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.contentType entity.
func (m *ItemSitesItemContentTypesContentTypeItemRequestBuilder) Columns()(*ItemSitesItemContentTypesItemColumnsRequestBuilder) {
    return NewItemSitesItemContentTypesItemColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ColumnsById provides operations to manage the columns property of the microsoft.graph.contentType entity.
func (m *ItemSitesItemContentTypesContentTypeItemRequestBuilder) ColumnsById(id string)(*ItemSitesItemContentTypesItemColumnsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return NewItemSitesItemContentTypesItemColumnsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemSitesItemContentTypesContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewItemSitesItemContentTypesContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemContentTypesContentTypeItemRequestBuilder) {
    m := &ItemSitesItemContentTypesContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/contentTypes/{contentType%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesItemContentTypesContentTypeItemRequestBuilder instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewItemSitesItemContentTypesContentTypeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemContentTypesContentTypeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemContentTypesContentTypeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property contentTypes for groups
func (m *ItemSitesItemContentTypesContentTypeItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemContentTypesContentTypeItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get the collection of content types defined for this site.
func (m *ItemSitesItemContentTypesContentTypeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemContentTypesContentTypeItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateContentTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable), nil
}
// MicrosoftGraphAssociateWithHubSites provides operations to call the associateWithHubSites method.
func (m *ItemSitesItemContentTypesContentTypeItemRequestBuilder) MicrosoftGraphAssociateWithHubSites()(*ItemSitesItemContentTypesItemMicrosoftGraphAssociateWithHubSitesRequestBuilder) {
    return NewItemSitesItemContentTypesItemMicrosoftGraphAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCopyToDefaultContentLocation provides operations to call the copyToDefaultContentLocation method.
func (m *ItemSitesItemContentTypesContentTypeItemRequestBuilder) MicrosoftGraphCopyToDefaultContentLocation()(*ItemSitesItemContentTypesItemMicrosoftGraphCopyToDefaultContentLocationRequestBuilder) {
    return NewItemSitesItemContentTypesItemMicrosoftGraphCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphIsPublished provides operations to call the isPublished method.
func (m *ItemSitesItemContentTypesContentTypeItemRequestBuilder) MicrosoftGraphIsPublished()(*ItemSitesItemContentTypesItemMicrosoftGraphIsPublishedRequestBuilder) {
    return NewItemSitesItemContentTypesItemMicrosoftGraphIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPublish provides operations to call the publish method.
func (m *ItemSitesItemContentTypesContentTypeItemRequestBuilder) MicrosoftGraphPublish()(*ItemSitesItemContentTypesItemMicrosoftGraphPublishRequestBuilder) {
    return NewItemSitesItemContentTypesItemMicrosoftGraphPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUnpublish provides operations to call the unpublish method.
func (m *ItemSitesItemContentTypesContentTypeItemRequestBuilder) MicrosoftGraphUnpublish()(*ItemSitesItemContentTypesItemMicrosoftGraphUnpublishRequestBuilder) {
    return NewItemSitesItemContentTypesItemMicrosoftGraphUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property contentTypes in groups
func (m *ItemSitesItemContentTypesContentTypeItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, requestConfiguration *ItemSitesItemContentTypesContentTypeItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateContentTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable), nil
}
// ToDeleteRequestInformation delete navigation property contentTypes for groups
func (m *ItemSitesItemContentTypesContentTypeItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemContentTypesContentTypeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the collection of content types defined for this site.
func (m *ItemSitesItemContentTypesContentTypeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemContentTypesContentTypeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property contentTypes in groups
func (m *ItemSitesItemContentTypesContentTypeItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, requestConfiguration *ItemSitesItemContentTypesContentTypeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
