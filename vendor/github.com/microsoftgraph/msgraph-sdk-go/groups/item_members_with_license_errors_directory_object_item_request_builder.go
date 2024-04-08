package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder provides operations to manage the membersWithLicenseErrors property of the microsoft.graph.group entity.
type ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderGetQueryParameters a list of group members with license errors from this group-based license assignment. Read-only.
type ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderGetQueryParameters
}
// NewItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) {
    m := &ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/membersWithLicenseErrors/{directoryObject%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get a list of group members with license errors from this group-based license assignment. Read-only.
func (m *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable), nil
}
// MicrosoftGraphApplication casts the previous resource to application.
func (m *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) MicrosoftGraphApplication()(*ItemMembersWithLicenseErrorsItemMicrosoftGraphApplicationRequestBuilder) {
    return NewItemMembersWithLicenseErrorsItemMicrosoftGraphApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDevice casts the previous resource to device.
func (m *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) MicrosoftGraphDevice()(*ItemMembersWithLicenseErrorsItemMicrosoftGraphDeviceRequestBuilder) {
    return NewItemMembersWithLicenseErrorsItemMicrosoftGraphDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGroup casts the previous resource to group.
func (m *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) MicrosoftGraphGroup()(*ItemMembersWithLicenseErrorsItemMicrosoftGraphGroupRequestBuilder) {
    return NewItemMembersWithLicenseErrorsItemMicrosoftGraphGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphOrgContact casts the previous resource to orgContact.
func (m *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) MicrosoftGraphOrgContact()(*ItemMembersWithLicenseErrorsItemMicrosoftGraphOrgContactRequestBuilder) {
    return NewItemMembersWithLicenseErrorsItemMicrosoftGraphOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphServicePrincipal casts the previous resource to servicePrincipal.
func (m *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) MicrosoftGraphServicePrincipal()(*ItemMembersWithLicenseErrorsItemMicrosoftGraphServicePrincipalRequestBuilder) {
    return NewItemMembersWithLicenseErrorsItemMicrosoftGraphServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUser casts the previous resource to user.
func (m *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) MicrosoftGraphUser()(*ItemMembersWithLicenseErrorsItemMicrosoftGraphUserRequestBuilder) {
    return NewItemMembersWithLicenseErrorsItemMicrosoftGraphUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation a list of group members with license errors from this group-based license assignment. Read-only.
func (m *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
