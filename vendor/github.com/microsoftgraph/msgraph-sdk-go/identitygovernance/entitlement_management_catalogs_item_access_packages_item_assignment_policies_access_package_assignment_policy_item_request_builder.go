package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder provides operations to manage the assignmentPolicies property of the microsoft.graph.accessPackage entity.
type EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetQueryParameters get assignmentPolicies from identityGovernance
type EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetQueryParameters
}
// EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackage provides operations to manage the accessPackage property of the microsoft.graph.accessPackageAssignmentPolicy entity.
func (m *EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) AccessPackage()(*EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilder) {
    return NewEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Catalog provides operations to manage the catalog property of the microsoft.graph.accessPackageAssignmentPolicy entity.
func (m *EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Catalog()(*EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemCatalogRequestBuilder) {
    return NewEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal instantiates a new AccessPackageAssignmentPolicyItemRequestBuilder and sets the default values.
func NewEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) {
    m := &EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/catalogs/{accessPackageCatalog%2Did}/accessPackages/{accessPackage%2Did}/assignmentPolicies/{accessPackageAssignmentPolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder instantiates a new AccessPackageAssignmentPolicyItemRequestBuilder and sets the default values.
func NewEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignmentPolicies for identityGovernance
func (m *EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get assignmentPolicies from identityGovernance
func (m *EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageAssignmentPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable), nil
}
// Patch update the navigation property assignmentPolicies in identityGovernance
func (m *EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable, requestConfiguration *EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageAssignmentPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable), nil
}
// Questions provides operations to manage the questions property of the microsoft.graph.accessPackageAssignmentPolicy entity.
func (m *EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Questions()(*EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemQuestionsRequestBuilder) {
    return NewEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemQuestionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// QuestionsById provides operations to manage the questions property of the microsoft.graph.accessPackageAssignmentPolicy entity.
func (m *EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) QuestionsById(id string)(*EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemQuestionsAccessPackageQuestionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageQuestion%2Did"] = id
    }
    return NewEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemQuestionsAccessPackageQuestionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property assignmentPolicies for identityGovernance
func (m *EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get assignmentPolicies from identityGovernance
func (m *EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignmentPolicies in identityGovernance
func (m *EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable, requestConfiguration *EntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
