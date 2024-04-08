package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackagesAccessPackageItemRequestBuilder provides operations to manage the accessPackages property of the microsoft.graph.entitlementManagement entity.
type EntitlementManagementAccessPackagesAccessPackageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementAccessPackagesAccessPackageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackagesAccessPackageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementManagementAccessPackagesAccessPackageItemRequestBuilderGetQueryParameters access packages define the collection of resource roles and the policies for which subjects can request or be assigned access to those resources.
type EntitlementManagementAccessPackagesAccessPackageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementAccessPackagesAccessPackageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackagesAccessPackageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilderGetQueryParameters
}
// EntitlementManagementAccessPackagesAccessPackageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackagesAccessPackageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackagesIncompatibleWith provides operations to manage the accessPackagesIncompatibleWith property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilder) AccessPackagesIncompatibleWith()(*EntitlementManagementAccessPackagesItemAccessPackagesIncompatibleWithRequestBuilder) {
    return NewEntitlementManagementAccessPackagesItemAccessPackagesIncompatibleWithRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AccessPackagesIncompatibleWithById provides operations to manage the accessPackagesIncompatibleWith property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilder) AccessPackagesIncompatibleWithById(id string)(*EntitlementManagementAccessPackagesItemAccessPackagesIncompatibleWithAccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage%2Did1"] = id
    }
    return NewEntitlementManagementAccessPackagesItemAccessPackagesIncompatibleWithAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// AssignmentPolicies provides operations to manage the assignmentPolicies property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilder) AssignmentPolicies()(*EntitlementManagementAccessPackagesItemAssignmentPoliciesRequestBuilder) {
    return NewEntitlementManagementAccessPackagesItemAssignmentPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignmentPoliciesById provides operations to manage the assignmentPolicies property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilder) AssignmentPoliciesById(id string)(*EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentPolicy%2Did"] = id
    }
    return NewEntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Catalog provides operations to manage the catalog property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilder) Catalog()(*EntitlementManagementAccessPackagesItemCatalogRequestBuilder) {
    return NewEntitlementManagementAccessPackagesItemCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewEntitlementManagementAccessPackagesAccessPackageItemRequestBuilderInternal instantiates a new AccessPackageItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesAccessPackageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesAccessPackageItemRequestBuilder) {
    m := &EntitlementManagementAccessPackagesAccessPackageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementAccessPackagesAccessPackageItemRequestBuilder instantiates a new AccessPackageItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesAccessPackageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesAccessPackageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackagesAccessPackageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property accessPackages for identityGovernance
func (m *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get access packages define the collection of resource roles and the policies for which subjects can request or be assigned access to those resources.
func (m *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable, error) {
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
func (m *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilder) IncompatibleAccessPackages()(*EntitlementManagementAccessPackagesItemIncompatibleAccessPackagesRequestBuilder) {
    return NewEntitlementManagementAccessPackagesItemIncompatibleAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IncompatibleAccessPackagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.accessPackages.item.incompatibleAccessPackages.item collection
func (m *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilder) IncompatibleAccessPackagesById(id string)(*EntitlementManagementAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage%2Did1"] = id
    }
    return NewEntitlementManagementAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// IncompatibleGroups provides operations to manage the incompatibleGroups property of the microsoft.graph.accessPackage entity.
func (m *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilder) IncompatibleGroups()(*EntitlementManagementAccessPackagesItemIncompatibleGroupsRequestBuilder) {
    return NewEntitlementManagementAccessPackagesItemIncompatibleGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IncompatibleGroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.accessPackages.item.incompatibleGroups.item collection
func (m *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilder) IncompatibleGroupsById(id string)(*EntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group%2Did"] = id
    }
    return NewEntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphGetApplicablePolicyRequirements provides operations to call the getApplicablePolicyRequirements method.
func (m *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilder) MicrosoftGraphGetApplicablePolicyRequirements()(*EntitlementManagementAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilder) {
    return NewEntitlementManagementAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property accessPackages in identityGovernance
func (m *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable, requestConfiguration *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable, error) {
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
func (m *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation access packages define the collection of resource roles and the policies for which subjects can request or be assigned access to those resources.
func (m *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable, requestConfiguration *EntitlementManagementAccessPackagesAccessPackageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
