package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// OnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentRequestBuilder provides operations to call the onenotePatchContent method.
type OnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentRequestBuilderInternal instantiates a new MicrosoftGraphOnenotePatchContentRequestBuilder and sets the default values.
func NewOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentRequestBuilder) {
    m := &OnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/onenote/notebooks/{notebook%2Did}/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}/pages/{onenotePage%2Did}/microsoft.graph.onenotePatchContent";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentRequestBuilder instantiates a new MicrosoftGraphOnenotePatchContentRequestBuilder and sets the default values.
func NewOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action onenotePatchContent
func (m *OnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentRequestBuilder) Post(ctx context.Context, body OnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentPostRequestBodyable, requestConfiguration *OnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action onenotePatchContent
func (m *OnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentRequestBuilder) ToPostRequestInformation(ctx context.Context, body OnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentOnenotePatchContentPostRequestBodyable, requestConfiguration *OnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemMicrosoftGraphOnenotePatchContentRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
