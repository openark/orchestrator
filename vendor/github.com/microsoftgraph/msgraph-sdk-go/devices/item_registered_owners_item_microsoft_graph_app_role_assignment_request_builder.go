package devices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemRegisteredOwnersItemMicrosoftGraphAppRoleAssignmentRequestBuilder casts the previous resource to appRoleAssignment.
type ItemRegisteredOwnersItemMicrosoftGraphAppRoleAssignmentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemRegisteredOwnersItemMicrosoftGraphAppRoleAssignmentRequestBuilderGetQueryParameters get the item of type microsoft.graph.directoryObject as microsoft.graph.appRoleAssignment
type ItemRegisteredOwnersItemMicrosoftGraphAppRoleAssignmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemRegisteredOwnersItemMicrosoftGraphAppRoleAssignmentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRegisteredOwnersItemMicrosoftGraphAppRoleAssignmentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemRegisteredOwnersItemMicrosoftGraphAppRoleAssignmentRequestBuilderGetQueryParameters
}
// NewItemRegisteredOwnersItemMicrosoftGraphAppRoleAssignmentRequestBuilderInternal instantiates a new MicrosoftGraphAppRoleAssignmentRequestBuilder and sets the default values.
func NewItemRegisteredOwnersItemMicrosoftGraphAppRoleAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredOwnersItemMicrosoftGraphAppRoleAssignmentRequestBuilder) {
    m := &ItemRegisteredOwnersItemMicrosoftGraphAppRoleAssignmentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/devices/{device%2Did}/registeredOwners/{directoryObject%2Did}/microsoft.graph.appRoleAssignment{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemRegisteredOwnersItemMicrosoftGraphAppRoleAssignmentRequestBuilder instantiates a new MicrosoftGraphAppRoleAssignmentRequestBuilder and sets the default values.
func NewItemRegisteredOwnersItemMicrosoftGraphAppRoleAssignmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredOwnersItemMicrosoftGraphAppRoleAssignmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRegisteredOwnersItemMicrosoftGraphAppRoleAssignmentRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.appRoleAssignment
func (m *ItemRegisteredOwnersItemMicrosoftGraphAppRoleAssignmentRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemRegisteredOwnersItemMicrosoftGraphAppRoleAssignmentRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppRoleAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAppRoleAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppRoleAssignmentable), nil
}
// ToGetRequestInformation get the item of type microsoft.graph.directoryObject as microsoft.graph.appRoleAssignment
func (m *ItemRegisteredOwnersItemMicrosoftGraphAppRoleAssignmentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemRegisteredOwnersItemMicrosoftGraphAppRoleAssignmentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
