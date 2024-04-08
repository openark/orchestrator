package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlRequestBuilder provides operations to call the getNotebookFromWebUrl method.
type ItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlRequestBuilderInternal instantiates a new MicrosoftGraphGetNotebookFromWebUrlRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlRequestBuilder) {
    m := &ItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/onenote/notebooks/microsoft.graph.getNotebookFromWebUrl";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlRequestBuilder instantiates a new MicrosoftGraphGetNotebookFromWebUrlRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Post retrieve the properties and relationships of a notebook object by using its URL path. The location can be user notebooks on Microsoft 365, group notebooks, or SharePoint site-hosted team notebooks on Microsoft 365.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/notebook-getnotebookfromweburl?view=graph-rest-1.0
func (m *ItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlRequestBuilder) Post(ctx context.Context, body ItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBodyable, requestConfiguration *ItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CopyNotebookModelable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCopyNotebookModelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CopyNotebookModelable), nil
}
// ToPostRequestInformation retrieve the properties and relationships of a notebook object by using its URL path. The location can be user notebooks on Microsoft 365, group notebooks, or SharePoint site-hosted team notebooks on Microsoft 365.
func (m *ItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBodyable, requestConfiguration *ItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
