package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilder provides operations to manage the pages property of the microsoft.graph.onenoteSection entity.
type ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilderGetQueryParameters the collection of pages in the section.  Read-only. Nullable.
type ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilderGetQueryParameters
}
// ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilderInternal instantiates a new OnenotePageItemRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilder) {
    m := &ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/onenote/notebooks/{notebook%2Did}/sections/{onenoteSection%2Did}/pages/{onenotePage%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilder instantiates a new OnenotePageItemRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the group entity.
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilder) Content()(*ItemSitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectionsItemPagesItemContentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete navigation property pages for groups
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the collection of pages in the section.  Read-only. Nullable.
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenotePageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable), nil
}
// MicrosoftGraphCopyToSection provides operations to call the copyToSection method.
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilder) MicrosoftGraphCopyToSection()(*ItemSitesItemOnenoteNotebooksItemSectionsItemPagesItemMicrosoftGraphCopyToSectionRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectionsItemPagesItemMicrosoftGraphCopyToSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphOnenotePatchContent provides operations to call the onenotePatchContent method.
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilder) MicrosoftGraphOnenotePatchContent()(*ItemSitesItemOnenoteNotebooksItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPreview provides operations to call the preview method.
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilder) MicrosoftGraphPreview()(*ItemSitesItemOnenoteNotebooksItemSectionsItemPagesItemMicrosoftGraphPreviewRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectionsItemPagesItemMicrosoftGraphPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ParentNotebook provides operations to manage the parentNotebook property of the microsoft.graph.onenotePage entity.
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilder) ParentNotebook()(*ItemSitesItemOnenoteNotebooksItemSectionsItemPagesItemParentNotebookRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectionsItemPagesItemParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ParentSection provides operations to manage the parentSection property of the microsoft.graph.onenotePage entity.
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilder) ParentSection()(*ItemSitesItemOnenoteNotebooksItemSectionsItemPagesItemParentSectionRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectionsItemPagesItemParentSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property pages in groups
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenotePageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable), nil
}
// ToDeleteRequestInformation delete navigation property pages for groups
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the collection of pages in the section.  Read-only. Nullable.
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property pages in groups
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
