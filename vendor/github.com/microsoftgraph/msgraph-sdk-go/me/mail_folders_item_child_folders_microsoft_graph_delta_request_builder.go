package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MailFoldersItemChildFoldersMicrosoftGraphDeltaRequestBuilder provides operations to call the delta method.
type MailFoldersItemChildFoldersMicrosoftGraphDeltaRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MailFoldersItemChildFoldersMicrosoftGraphDeltaRequestBuilderGetQueryParameters get a set of mail folders that have been added, deleted, or removed from the user's mailbox. A **delta** function call for mail folders in a mailbox is similar to a GET request, except that by appropriately applying state tokens in one or more of these calls, you can query for incremental changes in the mail folders. This allows you to maintain and synchronize a local store of a user's mail folders without having to fetch all the mail folders of that mailbox from the server every time.
type MailFoldersItemChildFoldersMicrosoftGraphDeltaRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// MailFoldersItemChildFoldersMicrosoftGraphDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MailFoldersItemChildFoldersMicrosoftGraphDeltaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MailFoldersItemChildFoldersMicrosoftGraphDeltaRequestBuilderGetQueryParameters
}
// NewMailFoldersItemChildFoldersMicrosoftGraphDeltaRequestBuilderInternal instantiates a new MicrosoftGraphDeltaRequestBuilder and sets the default values.
func NewMailFoldersItemChildFoldersMicrosoftGraphDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MailFoldersItemChildFoldersMicrosoftGraphDeltaRequestBuilder) {
    m := &MailFoldersItemChildFoldersMicrosoftGraphDeltaRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/mailFolders/{mailFolder%2Did}/childFolders/microsoft.graph.delta(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMailFoldersItemChildFoldersMicrosoftGraphDeltaRequestBuilder instantiates a new MicrosoftGraphDeltaRequestBuilder and sets the default values.
func NewMailFoldersItemChildFoldersMicrosoftGraphDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MailFoldersItemChildFoldersMicrosoftGraphDeltaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMailFoldersItemChildFoldersMicrosoftGraphDeltaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a set of mail folders that have been added, deleted, or removed from the user's mailbox. A **delta** function call for mail folders in a mailbox is similar to a GET request, except that by appropriately applying state tokens in one or more of these calls, you can query for incremental changes in the mail folders. This allows you to maintain and synchronize a local store of a user's mail folders without having to fetch all the mail folders of that mailbox from the server every time.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/mailfolder-delta?view=graph-rest-1.0
func (m *MailFoldersItemChildFoldersMicrosoftGraphDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *MailFoldersItemChildFoldersMicrosoftGraphDeltaRequestBuilderGetRequestConfiguration)(MailFoldersItemChildFoldersMicrosoftGraphDeltaDeltaResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateMailFoldersItemChildFoldersMicrosoftGraphDeltaDeltaResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MailFoldersItemChildFoldersMicrosoftGraphDeltaDeltaResponseable), nil
}
// ToGetRequestInformation get a set of mail folders that have been added, deleted, or removed from the user's mailbox. A **delta** function call for mail folders in a mailbox is similar to a GET request, except that by appropriately applying state tokens in one or more of these calls, you can query for incremental changes in the mail folders. This allows you to maintain and synchronize a local store of a user's mail folders without having to fetch all the mail folders of that mailbox from the server every time.
func (m *MailFoldersItemChildFoldersMicrosoftGraphDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MailFoldersItemChildFoldersMicrosoftGraphDeltaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
