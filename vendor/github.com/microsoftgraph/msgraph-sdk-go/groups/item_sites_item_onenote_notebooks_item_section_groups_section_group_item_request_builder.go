package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder provides operations to manage the sectionGroups property of the microsoft.graph.notebook entity.
type ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderGetQueryParameters the section groups in the notebook. Read-only. Nullable.
type ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderGetQueryParameters
}
// ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderInternal instantiates a new SectionGroupItemRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) {
    m := &ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/onenote/notebooks/{notebook%2Did}/sectionGroups/{sectionGroup%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder instantiates a new SectionGroupItemRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property sectionGroups for groups
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the section groups in the notebook. Read-only. Nullable.
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSectionGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable), nil
}
// ParentNotebook provides operations to manage the parentNotebook property of the microsoft.graph.sectionGroup entity.
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) ParentNotebook()(*ItemSitesItemOnenoteNotebooksItemSectionGroupsItemParentNotebookRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ParentSectionGroup provides operations to manage the parentSectionGroup property of the microsoft.graph.sectionGroup entity.
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) ParentSectionGroup()(*ItemSitesItemOnenoteNotebooksItemSectionGroupsItemParentSectionGroupRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property sectionGroups in groups
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSectionGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable), nil
}
// SectionGroups provides operations to manage the sectionGroups property of the microsoft.graph.sectionGroup entity.
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) SectionGroups()(*ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SectionGroupsById provides operations to manage the sectionGroups property of the microsoft.graph.sectionGroup entity.
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) SectionGroupsById(id string)(*ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup%2Did1"] = id
    }
    return NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Sections provides operations to manage the sections property of the microsoft.graph.sectionGroup entity.
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) Sections()(*ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SectionsById provides operations to manage the sections property of the microsoft.graph.sectionGroup entity.
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) SectionsById(id string)(*ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection%2Did"] = id
    }
    return NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property sectionGroups for groups
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the section groups in the notebook. Read-only. Nullable.
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property sectionGroups in groups
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
