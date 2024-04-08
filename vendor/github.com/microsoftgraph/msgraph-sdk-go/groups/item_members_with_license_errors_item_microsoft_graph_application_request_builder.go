package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMembersWithLicenseErrorsItemMicrosoftGraphApplicationRequestBuilder casts the previous resource to application.
type ItemMembersWithLicenseErrorsItemMicrosoftGraphApplicationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMembersWithLicenseErrorsItemMicrosoftGraphApplicationRequestBuilderGetQueryParameters get the item of type microsoft.graph.directoryObject as microsoft.graph.application
type ItemMembersWithLicenseErrorsItemMicrosoftGraphApplicationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMembersWithLicenseErrorsItemMicrosoftGraphApplicationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMembersWithLicenseErrorsItemMicrosoftGraphApplicationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMembersWithLicenseErrorsItemMicrosoftGraphApplicationRequestBuilderGetQueryParameters
}
// NewItemMembersWithLicenseErrorsItemMicrosoftGraphApplicationRequestBuilderInternal instantiates a new MicrosoftGraphApplicationRequestBuilder and sets the default values.
func NewItemMembersWithLicenseErrorsItemMicrosoftGraphApplicationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersWithLicenseErrorsItemMicrosoftGraphApplicationRequestBuilder) {
    m := &ItemMembersWithLicenseErrorsItemMicrosoftGraphApplicationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/membersWithLicenseErrors/{directoryObject%2Did}/microsoft.graph.application{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMembersWithLicenseErrorsItemMicrosoftGraphApplicationRequestBuilder instantiates a new MicrosoftGraphApplicationRequestBuilder and sets the default values.
func NewItemMembersWithLicenseErrorsItemMicrosoftGraphApplicationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersWithLicenseErrorsItemMicrosoftGraphApplicationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMembersWithLicenseErrorsItemMicrosoftGraphApplicationRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.application
func (m *ItemMembersWithLicenseErrorsItemMicrosoftGraphApplicationRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMembersWithLicenseErrorsItemMicrosoftGraphApplicationRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Applicationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Applicationable), nil
}
// ToGetRequestInformation get the item of type microsoft.graph.directoryObject as microsoft.graph.application
func (m *ItemMembersWithLicenseErrorsItemMicrosoftGraphApplicationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMembersWithLicenseErrorsItemMicrosoftGraphApplicationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
