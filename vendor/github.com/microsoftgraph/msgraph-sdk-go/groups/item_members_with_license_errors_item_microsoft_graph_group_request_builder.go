package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMembersWithLicenseErrorsItemMicrosoftGraphGroupRequestBuilder casts the previous resource to group.
type ItemMembersWithLicenseErrorsItemMicrosoftGraphGroupRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMembersWithLicenseErrorsItemMicrosoftGraphGroupRequestBuilderGetQueryParameters get the item of type microsoft.graph.directoryObject as microsoft.graph.group
type ItemMembersWithLicenseErrorsItemMicrosoftGraphGroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMembersWithLicenseErrorsItemMicrosoftGraphGroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMembersWithLicenseErrorsItemMicrosoftGraphGroupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMembersWithLicenseErrorsItemMicrosoftGraphGroupRequestBuilderGetQueryParameters
}
// NewItemMembersWithLicenseErrorsItemMicrosoftGraphGroupRequestBuilderInternal instantiates a new MicrosoftGraphGroupRequestBuilder and sets the default values.
func NewItemMembersWithLicenseErrorsItemMicrosoftGraphGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersWithLicenseErrorsItemMicrosoftGraphGroupRequestBuilder) {
    m := &ItemMembersWithLicenseErrorsItemMicrosoftGraphGroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/membersWithLicenseErrors/{directoryObject%2Did}/microsoft.graph.group{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMembersWithLicenseErrorsItemMicrosoftGraphGroupRequestBuilder instantiates a new MicrosoftGraphGroupRequestBuilder and sets the default values.
func NewItemMembersWithLicenseErrorsItemMicrosoftGraphGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersWithLicenseErrorsItemMicrosoftGraphGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMembersWithLicenseErrorsItemMicrosoftGraphGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.group
func (m *ItemMembersWithLicenseErrorsItemMicrosoftGraphGroupRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMembersWithLicenseErrorsItemMicrosoftGraphGroupRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable), nil
}
// ToGetRequestInformation get the item of type microsoft.graph.directoryObject as microsoft.graph.group
func (m *ItemMembersWithLicenseErrorsItemMicrosoftGraphGroupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMembersWithLicenseErrorsItemMicrosoftGraphGroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
