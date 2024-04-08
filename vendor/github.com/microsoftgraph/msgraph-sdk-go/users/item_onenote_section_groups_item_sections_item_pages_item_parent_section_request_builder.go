package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder provides operations to manage the parentSection property of the microsoft.graph.onenotePage entity.
type ItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderGetQueryParameters the section that contains the page. Read-only.
type ItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderGetQueryParameters
}
// NewItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderInternal instantiates a new ParentSectionRequestBuilder and sets the default values.
func NewItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder) {
    m := &ItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/onenote/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}/pages/{onenotePage%2Did}/parentSection{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder instantiates a new ParentSectionRequestBuilder and sets the default values.
func NewItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the section that contains the page. Read-only.
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenoteSectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenoteSectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenoteSectionable), nil
}
// ToGetRequestInformation the section that contains the page. Read-only.
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnenoteSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
