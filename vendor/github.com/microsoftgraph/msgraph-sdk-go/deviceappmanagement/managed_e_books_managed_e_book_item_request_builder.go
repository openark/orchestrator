package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ManagedEBooksManagedEBookItemRequestBuilder provides operations to manage the managedEBooks property of the microsoft.graph.deviceAppManagement entity.
type ManagedEBooksManagedEBookItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedEBooksManagedEBookItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedEBooksManagedEBookItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedEBooksManagedEBookItemRequestBuilderGetQueryParameters the Managed eBook.
type ManagedEBooksManagedEBookItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedEBooksManagedEBookItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedEBooksManagedEBookItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedEBooksManagedEBookItemRequestBuilderGetQueryParameters
}
// ManagedEBooksManagedEBookItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedEBooksManagedEBookItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.managedEBook entity.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) Assignments()(*ManagedEBooksItemAssignmentsRequestBuilder) {
    return NewManagedEBooksItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.managedEBook entity.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) AssignmentsById(id string)(*ManagedEBooksItemAssignmentsManagedEBookAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedEBookAssignment%2Did"] = id
    }
    return NewManagedEBooksItemAssignmentsManagedEBookAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewManagedEBooksManagedEBookItemRequestBuilderInternal instantiates a new ManagedEBookItemRequestBuilder and sets the default values.
func NewManagedEBooksManagedEBookItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedEBooksManagedEBookItemRequestBuilder) {
    m := &ManagedEBooksManagedEBookItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedEBooks/{managedEBook%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewManagedEBooksManagedEBookItemRequestBuilder instantiates a new ManagedEBookItemRequestBuilder and sets the default values.
func NewManagedEBooksManagedEBookItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedEBooksManagedEBookItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedEBooksManagedEBookItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property managedEBooks for deviceAppManagement
func (m *ManagedEBooksManagedEBookItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedEBooksManagedEBookItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceStates provides operations to manage the deviceStates property of the microsoft.graph.managedEBook entity.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) DeviceStates()(*ManagedEBooksItemDeviceStatesRequestBuilder) {
    return NewManagedEBooksItemDeviceStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceStatesById provides operations to manage the deviceStates property of the microsoft.graph.managedEBook entity.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) DeviceStatesById(id string)(*ManagedEBooksItemDeviceStatesDeviceInstallStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceInstallState%2Did"] = id
    }
    return NewManagedEBooksItemDeviceStatesDeviceInstallStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the Managed eBook.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedEBooksManagedEBookItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedEBookable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedEBookFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedEBookable), nil
}
// InstallSummary provides operations to manage the installSummary property of the microsoft.graph.managedEBook entity.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) InstallSummary()(*ManagedEBooksItemInstallSummaryRequestBuilder) {
    return NewManagedEBooksItemInstallSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAssign provides operations to call the assign method.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) MicrosoftGraphAssign()(*ManagedEBooksItemMicrosoftGraphAssignRequestBuilder) {
    return NewManagedEBooksItemMicrosoftGraphAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property managedEBooks in deviceAppManagement
func (m *ManagedEBooksManagedEBookItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedEBookable, requestConfiguration *ManagedEBooksManagedEBookItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedEBookable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedEBookFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedEBookable), nil
}
// ToDeleteRequestInformation delete navigation property managedEBooks for deviceAppManagement
func (m *ManagedEBooksManagedEBookItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedEBooksManagedEBookItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the Managed eBook.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedEBooksManagedEBookItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property managedEBooks in deviceAppManagement
func (m *ManagedEBooksManagedEBookItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedEBookable, requestConfiguration *ManagedEBooksManagedEBookItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserStateSummary provides operations to manage the userStateSummary property of the microsoft.graph.managedEBook entity.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) UserStateSummary()(*ManagedEBooksItemUserStateSummaryRequestBuilder) {
    return NewManagedEBooksItemUserStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserStateSummaryById provides operations to manage the userStateSummary property of the microsoft.graph.managedEBook entity.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) UserStateSummaryById(id string)(*ManagedEBooksItemUserStateSummaryUserInstallStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userInstallStateSummary%2Did"] = id
    }
    return NewManagedEBooksItemUserStateSummaryUserInstallStateSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
