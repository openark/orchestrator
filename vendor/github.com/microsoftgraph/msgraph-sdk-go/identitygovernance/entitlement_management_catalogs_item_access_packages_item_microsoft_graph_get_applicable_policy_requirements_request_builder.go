package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementManagementCatalogsItemAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilder provides operations to call the getApplicablePolicyRequirements method.
type EntitlementManagementCatalogsItemAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementCatalogsItemAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementCatalogsItemAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementManagementCatalogsItemAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilderInternal instantiates a new MicrosoftGraphGetApplicablePolicyRequirementsRequestBuilder and sets the default values.
func NewEntitlementManagementCatalogsItemAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementCatalogsItemAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilder) {
    m := &EntitlementManagementCatalogsItemAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/catalogs/{accessPackageCatalog%2Did}/accessPackages/{accessPackage%2Did}/microsoft.graph.getApplicablePolicyRequirements";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementCatalogsItemAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilder instantiates a new MicrosoftGraphGetApplicablePolicyRequirementsRequestBuilder and sets the default values.
func NewEntitlementManagementCatalogsItemAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementCatalogsItemAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementCatalogsItemAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post in Azure AD entitlement management, this action retrieves a list of accessPackageAssignmentRequestRequirements objects that the currently signed-in user can use to create an accessPackageAssignmentRequest.  Each requirement object corresponds to an access package assignment policy that the currently signed-in user is allowed to request an assignment for.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/accesspackage-getapplicablepolicyrequirements?view=graph-rest-1.0
func (m *EntitlementManagementCatalogsItemAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilder) Post(ctx context.Context, requestConfiguration *EntitlementManagementCatalogsItemAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration)(EntitlementManagementCatalogsItemAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsGetApplicablePolicyRequirementsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateEntitlementManagementCatalogsItemAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsGetApplicablePolicyRequirementsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementManagementCatalogsItemAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsGetApplicablePolicyRequirementsResponseable), nil
}
// ToPostRequestInformation in Azure AD entitlement management, this action retrieves a list of accessPackageAssignmentRequestRequirements objects that the currently signed-in user can use to create an accessPackageAssignmentRequest.  Each requirement object corresponds to an access package assignment policy that the currently signed-in user is allowed to request an assignment for.
func (m *EntitlementManagementCatalogsItemAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementCatalogsItemAccessPackagesItemMicrosoftGraphGetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
