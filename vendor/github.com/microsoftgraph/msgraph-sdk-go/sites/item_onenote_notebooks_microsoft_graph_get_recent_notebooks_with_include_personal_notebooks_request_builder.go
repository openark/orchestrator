package sites

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder provides operations to call the getRecentNotebooks method.
type ItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderGetQueryParameters invoke function getRecentNotebooks
type ItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderGetQueryParameters
}
// NewItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderInternal instantiates a new MicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, includePersonalNotebooks *bool)(*ItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder) {
    m := &ItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/onenote/notebooks/microsoft.graph.getRecentNotebooks(includePersonalNotebooks={includePersonalNotebooks}){?%24top,%24skip,%24search,%24filter,%24count}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if includePersonalNotebooks != nil {
        urlTplParams["includePersonalNotebooks"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatBool(*includePersonalNotebooks)
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder instantiates a new MicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getRecentNotebooks
func (m *ItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderGetRequestConfiguration)(ItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponseable), nil
}
// ToGetRequestInformation invoke function getRecentNotebooks
func (m *ItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnenoteNotebooksMicrosoftGraphGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
