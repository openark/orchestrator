package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilder provides operations to manage the defaultManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
type DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilderGetQueryParameters default managed app policies.
type DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilderGetQueryParameters
}
// DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Apps provides operations to manage the apps property of the microsoft.graph.defaultManagedAppProtection entity.
func (m *DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilder) Apps()(*DefaultManagedAppProtectionsItemAppsRequestBuilder) {
    return NewDefaultManagedAppProtectionsItemAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AppsById provides operations to manage the apps property of the microsoft.graph.defaultManagedAppProtection entity.
func (m *DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilder) AppsById(id string)(*DefaultManagedAppProtectionsItemAppsManagedMobileAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedMobileApp%2Did"] = id
    }
    return NewDefaultManagedAppProtectionsItemAppsManagedMobileAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewDefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilderInternal instantiates a new DefaultManagedAppProtectionItemRequestBuilder and sets the default values.
func NewDefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilder) {
    m := &DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/defaultManagedAppProtections/{defaultManagedAppProtection%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilder instantiates a new DefaultManagedAppProtectionItemRequestBuilder and sets the default values.
func NewDefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property defaultManagedAppProtections for deviceAppManagement
func (m *DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeploymentSummary provides operations to manage the deploymentSummary property of the microsoft.graph.defaultManagedAppProtection entity.
func (m *DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilder) DeploymentSummary()(*DefaultManagedAppProtectionsItemDeploymentSummaryRequestBuilder) {
    return NewDefaultManagedAppProtectionsItemDeploymentSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get default managed app policies.
func (m *DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DefaultManagedAppProtectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDefaultManagedAppProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DefaultManagedAppProtectionable), nil
}
// Patch update the navigation property defaultManagedAppProtections in deviceAppManagement
func (m *DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DefaultManagedAppProtectionable, requestConfiguration *DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DefaultManagedAppProtectionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDefaultManagedAppProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DefaultManagedAppProtectionable), nil
}
// ToDeleteRequestInformation delete navigation property defaultManagedAppProtections for deviceAppManagement
func (m *DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation default managed app policies.
func (m *DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property defaultManagedAppProtections in deviceAppManagement
func (m *DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DefaultManagedAppProtectionable, requestConfiguration *DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
