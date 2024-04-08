package organization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// OrganizationItemRequestBuilder provides operations to manage the collection of organization entities.
type OrganizationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OrganizationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OrganizationItemRequestBuilderGetQueryParameters get the properties and relationships of the currently authenticated organization. Since the **organization** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in an **organization** instance.
type OrganizationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OrganizationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OrganizationItemRequestBuilderGetQueryParameters
}
// OrganizationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Branding provides operations to manage the branding property of the microsoft.graph.organization entity.
func (m *OrganizationItemRequestBuilder) Branding()(*ItemBrandingRequestBuilder) {
    return NewItemBrandingRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CertificateBasedAuthConfiguration provides operations to manage the certificateBasedAuthConfiguration property of the microsoft.graph.organization entity.
func (m *OrganizationItemRequestBuilder) CertificateBasedAuthConfiguration()(*ItemCertificateBasedAuthConfigurationRequestBuilder) {
    return NewItemCertificateBasedAuthConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CertificateBasedAuthConfigurationById provides operations to manage the certificateBasedAuthConfiguration property of the microsoft.graph.organization entity.
func (m *OrganizationItemRequestBuilder) CertificateBasedAuthConfigurationById(id string)(*ItemCertificateBasedAuthConfigurationCertificateBasedAuthConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["certificateBasedAuthConfiguration%2Did"] = id
    }
    return NewItemCertificateBasedAuthConfigurationCertificateBasedAuthConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewOrganizationItemRequestBuilderInternal instantiates a new OrganizationItemRequestBuilder and sets the default values.
func NewOrganizationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationItemRequestBuilder) {
    m := &OrganizationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/organization/{organization%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewOrganizationItemRequestBuilder instantiates a new OrganizationItemRequestBuilder and sets the default values.
func NewOrganizationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganizationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete entity from organization
func (m *OrganizationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *OrganizationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Extensions provides operations to manage the extensions property of the microsoft.graph.organization entity.
func (m *OrganizationItemRequestBuilder) Extensions()(*ItemExtensionsRequestBuilder) {
    return NewItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.organization entity.
func (m *OrganizationItemRequestBuilder) ExtensionsById(id string)(*ItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get get the properties and relationships of the currently authenticated organization. Since the **organization** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in an **organization** instance.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/organization-get?view=graph-rest-1.0
func (m *OrganizationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OrganizationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Organizationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOrganizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Organizationable), nil
}
// MicrosoftGraphCheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *OrganizationItemRequestBuilder) MicrosoftGraphCheckMemberGroups()(*ItemMicrosoftGraphCheckMemberGroupsRequestBuilder) {
    return NewItemMicrosoftGraphCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *OrganizationItemRequestBuilder) MicrosoftGraphCheckMemberObjects()(*ItemMicrosoftGraphCheckMemberObjectsRequestBuilder) {
    return NewItemMicrosoftGraphCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetMemberGroups provides operations to call the getMemberGroups method.
func (m *OrganizationItemRequestBuilder) MicrosoftGraphGetMemberGroups()(*ItemMicrosoftGraphGetMemberGroupsRequestBuilder) {
    return NewItemMicrosoftGraphGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetMemberObjects provides operations to call the getMemberObjects method.
func (m *OrganizationItemRequestBuilder) MicrosoftGraphGetMemberObjects()(*ItemMicrosoftGraphGetMemberObjectsRequestBuilder) {
    return NewItemMicrosoftGraphGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRestore provides operations to call the restore method.
func (m *OrganizationItemRequestBuilder) MicrosoftGraphRestore()(*ItemMicrosoftGraphRestoreRequestBuilder) {
    return NewItemMicrosoftGraphRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetMobileDeviceManagementAuthority provides operations to call the setMobileDeviceManagementAuthority method.
func (m *OrganizationItemRequestBuilder) MicrosoftGraphSetMobileDeviceManagementAuthority()(*ItemMicrosoftGraphSetMobileDeviceManagementAuthorityRequestBuilder) {
    return NewItemMicrosoftGraphSetMobileDeviceManagementAuthorityRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the properties of the currently authenticated organization. In this case, `organization` is defined as a collection of exactly one record, and so its **ID** must be specified in the request.  The **ID** is also known as the **tenantId** of the organization.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/organization-update?view=graph-rest-1.0
func (m *OrganizationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Organizationable, requestConfiguration *OrganizationItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Organizationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOrganizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Organizationable), nil
}
// ToDeleteRequestInformation delete entity from organization
func (m *OrganizationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *OrganizationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get the properties and relationships of the currently authenticated organization. Since the **organization** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in an **organization** instance.
func (m *OrganizationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OrganizationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of the currently authenticated organization. In this case, `organization` is defined as a collection of exactly one record, and so its **ID** must be specified in the request.  The **ID** is also known as the **tenantId** of the organization.
func (m *OrganizationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Organizationable, requestConfiguration *OrganizationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
