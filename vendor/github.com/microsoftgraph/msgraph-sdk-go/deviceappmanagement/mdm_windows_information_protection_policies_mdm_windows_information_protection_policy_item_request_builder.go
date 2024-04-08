package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder provides operations to manage the mdmWindowsInformationProtectionPolicies property of the microsoft.graph.deviceAppManagement entity.
type MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderGetQueryParameters windows information protection for apps running on devices which are MDM enrolled.
type MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderGetQueryParameters
}
// MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderInternal instantiates a new MdmWindowsInformationProtectionPolicyItemRequestBuilder and sets the default values.
func NewMdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) {
    m := &MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mdmWindowsInformationProtectionPolicies/{mdmWindowsInformationProtectionPolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder instantiates a new MdmWindowsInformationProtectionPolicyItemRequestBuilder and sets the default values.
func NewMdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property mdmWindowsInformationProtectionPolicies for deviceAppManagement
func (m *MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get windows information protection for apps running on devices which are MDM enrolled.
func (m *MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MdmWindowsInformationProtectionPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMdmWindowsInformationProtectionPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MdmWindowsInformationProtectionPolicyable), nil
}
// Patch update the navigation property mdmWindowsInformationProtectionPolicies in deviceAppManagement
func (m *MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MdmWindowsInformationProtectionPolicyable, requestConfiguration *MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MdmWindowsInformationProtectionPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMdmWindowsInformationProtectionPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MdmWindowsInformationProtectionPolicyable), nil
}
// ToDeleteRequestInformation delete navigation property mdmWindowsInformationProtectionPolicies for deviceAppManagement
func (m *MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation windows information protection for apps running on devices which are MDM enrolled.
func (m *MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property mdmWindowsInformationProtectionPolicies in deviceAppManagement
func (m *MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MdmWindowsInformationProtectionPolicyable, requestConfiguration *MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
