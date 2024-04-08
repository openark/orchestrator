package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder provides operations to manage the defaultPages property of the microsoft.graph.userFlowLanguageConfiguration entity.
type B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderGetQueryParameters collection of pages with the default content to display in a user flow for a specified language. This collection does not allow any kind of modification.
type B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderGetQueryParameters
}
// B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderInternal instantiates a new UserFlowLanguagePageItemRequestBuilder and sets the default values.
func NewB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) {
    m := &B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/languages/{userFlowLanguageConfiguration%2Did}/defaultPages/{userFlowLanguagePage%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder instantiates a new UserFlowLanguagePageItemRequestBuilder and sets the default values.
func NewB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the identityContainer entity.
func (m *B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) Content()(*B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilder) {
    return NewB2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete navigation property defaultPages for identity
func (m *B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get collection of pages with the default content to display in a user flow for a specified language. This collection does not allow any kind of modification.
func (m *B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserFlowLanguagePageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserFlowLanguagePageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserFlowLanguagePageable), nil
}
// Patch update the navigation property defaultPages in identity
func (m *B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserFlowLanguagePageable, requestConfiguration *B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserFlowLanguagePageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserFlowLanguagePageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserFlowLanguagePageable), nil
}
// ToDeleteRequestInformation delete navigation property defaultPages for identity
func (m *B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation collection of pages with the default content to display in a user flow for a specified language. This collection does not allow any kind of modification.
func (m *B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property defaultPages in identity
func (m *B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserFlowLanguagePageable, requestConfiguration *B2xUserFlowsItemLanguagesItemDefaultPagesUserFlowLanguagePageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
