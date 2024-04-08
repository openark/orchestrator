package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemListsItemContentTypesContentTypeItemRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.list entity.
type ItemListsItemContentTypesContentTypeItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemListsItemContentTypesContentTypeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemContentTypesContentTypeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemListsItemContentTypesContentTypeItemRequestBuilderGetQueryParameters the collection of content types present in this list.
type ItemListsItemContentTypesContentTypeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemListsItemContentTypesContentTypeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemContentTypesContentTypeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemListsItemContentTypesContentTypeItemRequestBuilderGetQueryParameters
}
// ItemListsItemContentTypesContentTypeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemContentTypesContentTypeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Base provides operations to manage the base property of the microsoft.graph.contentType entity.
func (m *ItemListsItemContentTypesContentTypeItemRequestBuilder) Base()(*ItemListsItemContentTypesItemBaseRequestBuilder) {
    return NewItemListsItemContentTypesItemBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// BaseTypes provides operations to manage the baseTypes property of the microsoft.graph.contentType entity.
func (m *ItemListsItemContentTypesContentTypeItemRequestBuilder) BaseTypes()(*ItemListsItemContentTypesItemBaseTypesRequestBuilder) {
    return NewItemListsItemContentTypesItemBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// BaseTypesById provides operations to manage the baseTypes property of the microsoft.graph.contentType entity.
func (m *ItemListsItemContentTypesContentTypeItemRequestBuilder) BaseTypesById(id string)(*ItemListsItemContentTypesItemBaseTypesContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did1"] = id
    }
    return NewItemListsItemContentTypesItemBaseTypesContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ColumnLinks provides operations to manage the columnLinks property of the microsoft.graph.contentType entity.
func (m *ItemListsItemContentTypesContentTypeItemRequestBuilder) ColumnLinks()(*ItemListsItemContentTypesItemColumnLinksRequestBuilder) {
    return NewItemListsItemContentTypesItemColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ColumnLinksById provides operations to manage the columnLinks property of the microsoft.graph.contentType entity.
func (m *ItemListsItemContentTypesContentTypeItemRequestBuilder) ColumnLinksById(id string)(*ItemListsItemContentTypesItemColumnLinksColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink%2Did"] = id
    }
    return NewItemListsItemContentTypesItemColumnLinksColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ColumnPositions provides operations to manage the columnPositions property of the microsoft.graph.contentType entity.
func (m *ItemListsItemContentTypesContentTypeItemRequestBuilder) ColumnPositions()(*ItemListsItemContentTypesItemColumnPositionsRequestBuilder) {
    return NewItemListsItemContentTypesItemColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ColumnPositionsById provides operations to manage the columnPositions property of the microsoft.graph.contentType entity.
func (m *ItemListsItemContentTypesContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*ItemListsItemContentTypesItemColumnPositionsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return NewItemListsItemContentTypesItemColumnPositionsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.contentType entity.
func (m *ItemListsItemContentTypesContentTypeItemRequestBuilder) Columns()(*ItemListsItemContentTypesItemColumnsRequestBuilder) {
    return NewItemListsItemContentTypesItemColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ColumnsById provides operations to manage the columns property of the microsoft.graph.contentType entity.
func (m *ItemListsItemContentTypesContentTypeItemRequestBuilder) ColumnsById(id string)(*ItemListsItemContentTypesItemColumnsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return NewItemListsItemContentTypesItemColumnsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemListsItemContentTypesContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewItemListsItemContentTypesContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemContentTypesContentTypeItemRequestBuilder) {
    m := &ItemListsItemContentTypesContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/contentTypes/{contentType%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemListsItemContentTypesContentTypeItemRequestBuilder instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewItemListsItemContentTypesContentTypeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemContentTypesContentTypeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListsItemContentTypesContentTypeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property contentTypes for sites
func (m *ItemListsItemContentTypesContentTypeItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemListsItemContentTypesContentTypeItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the collection of content types present in this list.
func (m *ItemListsItemContentTypesContentTypeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemListsItemContentTypesContentTypeItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, error) {
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
func (m *ItemListsItemContentTypesContentTypeItemRequestBuilder) MicrosoftGraphAssociateWithHubSites()(*ItemListsItemContentTypesItemMicrosoftGraphAssociateWithHubSitesRequestBuilder) {
    return NewItemListsItemContentTypesItemMicrosoftGraphAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCopyToDefaultContentLocation provides operations to call the copyToDefaultContentLocation method.
func (m *ItemListsItemContentTypesContentTypeItemRequestBuilder) MicrosoftGraphCopyToDefaultContentLocation()(*ItemListsItemContentTypesItemMicrosoftGraphCopyToDefaultContentLocationRequestBuilder) {
    return NewItemListsItemContentTypesItemMicrosoftGraphCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphIsPublished provides operations to call the isPublished method.
func (m *ItemListsItemContentTypesContentTypeItemRequestBuilder) MicrosoftGraphIsPublished()(*ItemListsItemContentTypesItemMicrosoftGraphIsPublishedRequestBuilder) {
    return NewItemListsItemContentTypesItemMicrosoftGraphIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPublish provides operations to call the publish method.
func (m *ItemListsItemContentTypesContentTypeItemRequestBuilder) MicrosoftGraphPublish()(*ItemListsItemContentTypesItemMicrosoftGraphPublishRequestBuilder) {
    return NewItemListsItemContentTypesItemMicrosoftGraphPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUnpublish provides operations to call the unpublish method.
func (m *ItemListsItemContentTypesContentTypeItemRequestBuilder) MicrosoftGraphUnpublish()(*ItemListsItemContentTypesItemMicrosoftGraphUnpublishRequestBuilder) {
    return NewItemListsItemContentTypesItemMicrosoftGraphUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property contentTypes in sites
func (m *ItemListsItemContentTypesContentTypeItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, requestConfiguration *ItemListsItemContentTypesContentTypeItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, error) {
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
// ToDeleteRequestInformation delete navigation property contentTypes for sites
func (m *ItemListsItemContentTypesContentTypeItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemListsItemContentTypesContentTypeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the collection of content types present in this list.
func (m *ItemListsItemContentTypesContentTypeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemListsItemContentTypesContentTypeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property contentTypes in sites
func (m *ItemListsItemContentTypesContentTypeItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, requestConfiguration *ItemListsItemContentTypesContentTypeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
