package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder provides operations to call the additionalAccess method.
type EntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetQueryParameters invoke function additionalAccess
type EntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
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
// EntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetQueryParameters
}
// NewEntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderInternal instantiates a new MicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder and sets the default values.
func NewEntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, accessPackageId *string, incompatibleAccessPackageId *string)(*EntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) {
    m := &EntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/assignments/microsoft.graph.additionalAccess(accessPackageId='{accessPackageId}',incompatibleAccessPackageId='{incompatibleAccessPackageId}'){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if accessPackageId != nil {
        urlTplParams["accessPackageId"] = *accessPackageId
    }
    if incompatibleAccessPackageId != nil {
        urlTplParams["incompatibleAccessPackageId"] = *incompatibleAccessPackageId
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder instantiates a new MicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder and sets the default values.
func NewEntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function additionalAccess
func (m *EntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration)(EntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateEntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponseable), nil
}
// ToGetRequestInformation invoke function additionalAccess
func (m *EntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAssignmentsMicrosoftGraphAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
