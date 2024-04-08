package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder provides operations to call the getAvailableExtensionProperties method.
type EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilderInternal instantiates a new MicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder and sets the default values.
func NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder) {
    m := &EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/connectedOrganizations/{connectedOrganization%2Did}/internalSponsors/microsoft.graph.getAvailableExtensionProperties";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder instantiates a new MicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder and sets the default values.
func NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post return all directory extension definitions that have been registered in a directory, including through multi-tenant apps. The following entities support extension properties:+ user+ group+ administrativeUnit+ application+ device+ organization
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/directoryobject-getavailableextensionproperties?view=graph-rest-1.0
func (m *EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder) Post(ctx context.Context, body EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostRequestBodyable, requestConfiguration *EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration)(EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateEntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesResponseable), nil
}
// ToPostRequestInformation return all directory extension definitions that have been registered in a directory, including through multi-tenant apps. The following entities support extension properties:+ user+ group+ administrativeUnit+ application+ device+ organization
func (m *EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostRequestBodyable, requestConfiguration *EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
