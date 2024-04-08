package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder provides operations to manage the accessPackages property of the microsoft.graph.accessPackageCatalog entity.
type EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilderGetQueryParameters the access packages in this catalog. Read-only. Nullable.
type EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilderGetQueryParameters
}
// EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackagesIncompatibleWith provides operations to manage the accessPackagesIncompatibleWith property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) AccessPackagesIncompatibleWith()(*EntitlementManagementCatalogsItemAccessPackagesItemAccessPackagesIncompatibleWithRequestBuilder) {
    return NewEntitlementManagementCatalogsItemAccessPackagesItemAccessPackagesIncompatibleWithRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AccessPackagesIncompatibleWithById provides operations to manage the accessPackagesIncompatibleWith property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) AccessPackagesIncompatibleWithById(id string)(*EntitlementManagementCatalogsItemAccessPackagesItemAccessPackagesIncompatibleWithAccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage%2Did1"] = id
    }
    return NewEntitlementManagementCatalogsItemAccessPackagesItemAccessPackagesIncompatibleWithAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// AssignmentPolicies provides operations to manage the assignmentPolicies property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) AssignmentPolicies()(*EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesRequestBuilder) {
    return NewEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignmentPoliciesById provides operations to manage the assignmentPolicies property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) AssignmentPoliciesById(id string)(*EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentPolicy%2Did"] = id
    }
    return NewEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Catalog provides operations to manage the catalog property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) Catalog()(*EntitlementManagementCatalogsItemAccessPackagesItemCatalogRequestBuilder) {
    return NewEntitlementManagementCatalogsItemAccessPackagesItemCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewEntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilderInternal instantiates a new AccessPackageItemRequestBuilder and sets the default values.
func NewEntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) {
    m := &EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/catalogs/{accessPackageCatalog%2Did}/accessPackages/{accessPackage%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder instantiates a new AccessPackageItemRequestBuilder and sets the default values.
func NewEntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property accessPackages for identityGovernance
func (m *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the access packages in this catalog. Read-only. Nullable.
func (m *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable), nil
}
// IncompatibleAccessPackages provides operations to manage the incompatibleAccessPackages property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) IncompatibleAccessPackages()(*EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder) {
    return NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IncompatibleAccessPackagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.catalogs.item.accessPackages.item.incompatibleAccessPackages.item collection
func (m *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) IncompatibleAccessPackagesById(id string)(*EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage%2Did1"] = id
    }
    return NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// IncompatibleGroups provides operations to manage the incompatibleGroups property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) IncompatibleGroups()(*EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilder) {
    return NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IncompatibleGroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.catalogs.item.accessPackages.item.incompatibleGroups.item collection
func (m *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) IncompatibleGroupsById(id string)(*EntitlementManagementCatalogsItemAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group%2Did"] = id
    }
    return NewEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphGetApplicablePolicyRequirements provides operations to call the getApplicablePolicyRequirements method.
func (m *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) MicrosoftGraphGetApplicablePolicyRequirements()(*EntitlementManagementCatalogsItemAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilder) {
    return NewEntitlementManagementCatalogsItemAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property accessPackages in identityGovernance
func (m *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable, requestConfiguration *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable), nil
}
// ToDeleteRequestInformation delete navigation property accessPackages for identityGovernance
func (m *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the access packages in this catalog. Read-only. Nullable.
func (m *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property accessPackages in identityGovernance
func (m *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable, requestConfiguration *EntitlementManagementCatalogsItemAccessPackagesAccessPackageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
