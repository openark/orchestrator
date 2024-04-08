package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilder provides operations to manage the partners property of the microsoft.graph.crossTenantAccessPolicy entity.
type CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilderGetQueryParameters defines partner-specific configurations for external Azure Active Directory organizations.
type CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilderGetQueryParameters
}
// CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilderInternal instantiates a new CrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilder and sets the default values.
func NewCrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilder) {
    m := &CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/crossTenantAccessPolicy/partners/{crossTenantAccessPolicyConfigurationPartner%2DtenantId}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilder instantiates a new CrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilder and sets the default values.
func NewCrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property partners for policies
func (m *CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get defines partner-specific configurations for external Azure Active Directory organizations.
func (m *CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyConfigurationPartnerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCrossTenantAccessPolicyConfigurationPartnerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyConfigurationPartnerable), nil
}
// Patch update the navigation property partners in policies
func (m *CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyConfigurationPartnerable, requestConfiguration *CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyConfigurationPartnerable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCrossTenantAccessPolicyConfigurationPartnerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyConfigurationPartnerable), nil
}
// ToDeleteRequestInformation delete navigation property partners for policies
func (m *CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation defines partner-specific configurations for external Azure Active Directory organizations.
func (m *CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property partners in policies
func (m *CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyConfigurationPartnerable, requestConfiguration *CrossTenantAccessPolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
