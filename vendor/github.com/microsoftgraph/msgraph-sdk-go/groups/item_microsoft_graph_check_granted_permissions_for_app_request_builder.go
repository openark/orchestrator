package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMicrosoftGraphCheckGrantedPermissionsForAppRequestBuilder provides operations to call the checkGrantedPermissionsForApp method.
type ItemMicrosoftGraphCheckGrantedPermissionsForAppRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMicrosoftGraphCheckGrantedPermissionsForAppRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMicrosoftGraphCheckGrantedPermissionsForAppRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMicrosoftGraphCheckGrantedPermissionsForAppRequestBuilderInternal instantiates a new MicrosoftGraphCheckGrantedPermissionsForAppRequestBuilder and sets the default values.
func NewItemMicrosoftGraphCheckGrantedPermissionsForAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphCheckGrantedPermissionsForAppRequestBuilder) {
    m := &ItemMicrosoftGraphCheckGrantedPermissionsForAppRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/microsoft.graph.checkGrantedPermissionsForApp";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMicrosoftGraphCheckGrantedPermissionsForAppRequestBuilder instantiates a new MicrosoftGraphCheckGrantedPermissionsForAppRequestBuilder and sets the default values.
func NewItemMicrosoftGraphCheckGrantedPermissionsForAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphCheckGrantedPermissionsForAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMicrosoftGraphCheckGrantedPermissionsForAppRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action checkGrantedPermissionsForApp
func (m *ItemMicrosoftGraphCheckGrantedPermissionsForAppRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemMicrosoftGraphCheckGrantedPermissionsForAppRequestBuilderPostRequestConfiguration)(ItemMicrosoftGraphCheckGrantedPermissionsForAppCheckGrantedPermissionsForAppResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemMicrosoftGraphCheckGrantedPermissionsForAppCheckGrantedPermissionsForAppResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemMicrosoftGraphCheckGrantedPermissionsForAppCheckGrantedPermissionsForAppResponseable), nil
}
// ToPostRequestInformation invoke action checkGrantedPermissionsForApp
func (m *ItemMicrosoftGraphCheckGrantedPermissionsForAppRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemMicrosoftGraphCheckGrantedPermissionsForAppRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
