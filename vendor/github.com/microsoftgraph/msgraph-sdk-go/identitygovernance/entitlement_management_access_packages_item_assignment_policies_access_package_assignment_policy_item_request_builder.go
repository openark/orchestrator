package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder provides operations to manage the assignmentPolicies property of the microsoft.graph.accessPackage entity.
type EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetQueryParameters get assignmentPolicies from identityGovernance
type EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetQueryParameters
}
// EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackage provides operations to manage the accessPackage property of the microsoft.graph.accessPackageAssignmentPolicy entity.
func (m *EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) AccessPackage()(*EntitlementManagementAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilder) {
    return NewEntitlementManagementAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Catalog provides operations to manage the catalog property of the microsoft.graph.accessPackageAssignmentPolicy entity.
func (m *EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Catalog()(*EntitlementManagementAccessPackagesItemAssignmentPoliciesItemCatalogRequestBuilder) {
    return NewEntitlementManagementAccessPackagesItemAssignmentPoliciesItemCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewEntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal instantiates a new AccessPackageAssignmentPolicyItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) {
    m := &EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/assignmentPolicies/{accessPackageAssignmentPolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder instantiates a new AccessPackageAssignmentPolicyItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignmentPolicies for identityGovernance
func (m *EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable, error) {
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
func (m *EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable, requestConfiguration *EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable, error) {
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
func (m *EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) Questions()(*EntitlementManagementAccessPackagesItemAssignmentPoliciesItemQuestionsRequestBuilder) {
    return NewEntitlementManagementAccessPackagesItemAssignmentPoliciesItemQuestionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// QuestionsById provides operations to manage the questions property of the microsoft.graph.accessPackageAssignmentPolicy entity.
func (m *EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) QuestionsById(id string)(*EntitlementManagementAccessPackagesItemAssignmentPoliciesItemQuestionsAccessPackageQuestionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageQuestion%2Did"] = id
    }
    return NewEntitlementManagementAccessPackagesItemAssignmentPoliciesItemQuestionsAccessPackageQuestionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property assignmentPolicies for identityGovernance
func (m *EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable, requestConfiguration *EntitlementManagementAccessPackagesItemAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
