package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementManagementRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilder provides operations to call the cancel method.
type EntitlementManagementRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementManagementRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilderInternal instantiates a new MicrosoftGraphCancelRequestBuilder and sets the default values.
func NewEntitlementManagementRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilder) {
    m := &EntitlementManagementRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement/roleEligibilityScheduleRequests/{unifiedRoleEligibilityScheduleRequest%2Did}/microsoft.graph.cancel";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilder instantiates a new MicrosoftGraphCancelRequestBuilder and sets the default values.
func NewEntitlementManagementRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post immediately cancel a unifiedRoleEligibilityScheduleRequest object whose status is `Granted` and have the system automatically delete the cancelled request after 30 days. After calling this action, the **status** of the cancelled **unifiedRoleEligibilityScheduleRequest** changes to `Revoked`.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/unifiedroleeligibilityschedulerequest-cancel?view=graph-rest-1.0
func (m *EntitlementManagementRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilder) Post(ctx context.Context, requestConfiguration *EntitlementManagementRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation immediately cancel a unifiedRoleEligibilityScheduleRequest object whose status is `Granted` and have the system automatically delete the cancelled request after 30 days. After calling this action, the **status** of the cancelled **unifiedRoleEligibilityScheduleRequest** changes to `Revoked`.
func (m *EntitlementManagementRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
