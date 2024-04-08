package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilder provides operations to manage the importedWindowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
type ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters collection of imported Windows autopilot devices.
type ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters
}
// ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderInternal instantiates a new ImportedWindowsAutopilotDeviceIdentityItemRequestBuilder and sets the default values.
func NewImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    m := &ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/importedWindowsAutopilotDeviceIdentities/{importedWindowsAutopilotDeviceIdentity%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilder instantiates a new ImportedWindowsAutopilotDeviceIdentityItemRequestBuilder and sets the default values.
func NewImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property importedWindowsAutopilotDeviceIdentities for deviceManagement
func (m *ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get collection of imported Windows autopilot devices.
func (m *ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateImportedWindowsAutopilotDeviceIdentityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityable), nil
}
// Patch update the navigation property importedWindowsAutopilotDeviceIdentities in deviceManagement
func (m *ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityable, requestConfiguration *ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateImportedWindowsAutopilotDeviceIdentityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityable), nil
}
// ToDeleteRequestInformation delete navigation property importedWindowsAutopilotDeviceIdentities for deviceManagement
func (m *ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation collection of imported Windows autopilot devices.
func (m *ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property importedWindowsAutopilotDeviceIdentities in deviceManagement
func (m *ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityable, requestConfiguration *ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
