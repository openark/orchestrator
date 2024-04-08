package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AdministrativeUnitsItemMembersRequestBuilder provides operations to manage the members property of the microsoft.graph.administrativeUnit entity.
type AdministrativeUnitsItemMembersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AdministrativeUnitsItemMembersRequestBuilderGetQueryParameters users and groups that are members of this administrative unit. Supports $expand.
type AdministrativeUnitsItemMembersRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
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
// AdministrativeUnitsItemMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdministrativeUnitsItemMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdministrativeUnitsItemMembersRequestBuilderGetQueryParameters
}
// NewAdministrativeUnitsItemMembersRequestBuilderInternal instantiates a new MembersRequestBuilder and sets the default values.
func NewAdministrativeUnitsItemMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdministrativeUnitsItemMembersRequestBuilder) {
    m := &AdministrativeUnitsItemMembersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/administrativeUnits/{administrativeUnit%2Did}/members{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAdministrativeUnitsItemMembersRequestBuilder instantiates a new MembersRequestBuilder and sets the default values.
func NewAdministrativeUnitsItemMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdministrativeUnitsItemMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdministrativeUnitsItemMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *AdministrativeUnitsItemMembersRequestBuilder) Count()(*AdministrativeUnitsItemMembersCountRequestBuilder) {
    return NewAdministrativeUnitsItemMembersCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get users and groups that are members of this administrative unit. Supports $expand.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/administrativeunit-list-members?view=graph-rest-1.0
func (m *AdministrativeUnitsItemMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *AdministrativeUnitsItemMembersRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable), nil
}
// MicrosoftGraphApplication casts the previous resource to application.
func (m *AdministrativeUnitsItemMembersRequestBuilder) MicrosoftGraphApplication()(*AdministrativeUnitsItemMembersMicrosoftGraphApplicationRequestBuilder) {
    return NewAdministrativeUnitsItemMembersMicrosoftGraphApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDevice casts the previous resource to device.
func (m *AdministrativeUnitsItemMembersRequestBuilder) MicrosoftGraphDevice()(*AdministrativeUnitsItemMembersMicrosoftGraphDeviceRequestBuilder) {
    return NewAdministrativeUnitsItemMembersMicrosoftGraphDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGroup casts the previous resource to group.
func (m *AdministrativeUnitsItemMembersRequestBuilder) MicrosoftGraphGroup()(*AdministrativeUnitsItemMembersMicrosoftGraphGroupRequestBuilder) {
    return NewAdministrativeUnitsItemMembersMicrosoftGraphGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphOrgContact casts the previous resource to orgContact.
func (m *AdministrativeUnitsItemMembersRequestBuilder) MicrosoftGraphOrgContact()(*AdministrativeUnitsItemMembersMicrosoftGraphOrgContactRequestBuilder) {
    return NewAdministrativeUnitsItemMembersMicrosoftGraphOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphServicePrincipal casts the previous resource to servicePrincipal.
func (m *AdministrativeUnitsItemMembersRequestBuilder) MicrosoftGraphServicePrincipal()(*AdministrativeUnitsItemMembersMicrosoftGraphServicePrincipalRequestBuilder) {
    return NewAdministrativeUnitsItemMembersMicrosoftGraphServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUser casts the previous resource to user.
func (m *AdministrativeUnitsItemMembersRequestBuilder) MicrosoftGraphUser()(*AdministrativeUnitsItemMembersMicrosoftGraphUserRequestBuilder) {
    return NewAdministrativeUnitsItemMembersMicrosoftGraphUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Ref provides operations to manage the collection of directory entities.
func (m *AdministrativeUnitsItemMembersRequestBuilder) Ref()(*AdministrativeUnitsItemMembersRefRequestBuilder) {
    return NewAdministrativeUnitsItemMembersRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation users and groups that are members of this administrative unit. Supports $expand.
func (m *AdministrativeUnitsItemMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AdministrativeUnitsItemMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
