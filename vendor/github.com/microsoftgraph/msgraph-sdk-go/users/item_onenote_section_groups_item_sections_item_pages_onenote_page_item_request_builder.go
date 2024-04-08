package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder provides operations to manage the pages property of the microsoft.graph.onenoteSection entity.
type ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetQueryParameters the collection of pages in the section.  Read-only. Nullable.
type ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetQueryParameters
}
// ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderInternal instantiates a new OnenotePageItemRequestBuilder and sets the default values.
func NewItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) {
    m := &ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/onenote/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}/pages/{onenotePage%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder instantiates a new OnenotePageItemRequestBuilder and sets the default values.
func NewItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the user entity.
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) Content()(*ItemOnenoteSectionGroupsItemSectionsItemPagesItemContentRequestBuilder) {
    return NewItemOnenoteSectionGroupsItemSectionsItemPagesItemContentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete navigation property pages for users
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, error) {
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
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) MicrosoftGraphCopyToSection()(*ItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionRequestBuilder) {
    return NewItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphOnenotePatchContent provides operations to call the onenotePatchContent method.
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) MicrosoftGraphOnenotePatchContent()(*ItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentRequestBuilder) {
    return NewItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPreview provides operations to call the preview method.
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) MicrosoftGraphPreview()(*ItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphPreviewRequestBuilder) {
    return NewItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ParentNotebook provides operations to manage the parentNotebook property of the microsoft.graph.onenotePage entity.
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) ParentNotebook()(*ItemOnenoteSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilder) {
    return NewItemOnenoteSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ParentSection provides operations to manage the parentSection property of the microsoft.graph.onenotePage entity.
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) ParentSection()(*ItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder) {
    return NewItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property pages in users
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, requestConfiguration *ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, error) {
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
// ToDeleteRequestInformation delete navigation property pages for users
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property pages in users
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, requestConfiguration *ItemOnenoteSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
