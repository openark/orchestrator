package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrintRequestBuilder provides operations to manage the print singleton.
type PrintRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PrintRequestBuilderGetQueryParameters get print
type PrintRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrintRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrintRequestBuilderGetQueryParameters
}
// PrintRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Connectors provides operations to manage the connectors property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) Connectors()(*ConnectorsRequestBuilder) {
    return NewConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ConnectorsById provides operations to manage the connectors property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) ConnectorsById(id string)(*ConnectorsPrintConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printConnector%2Did"] = id
    }
    return NewConnectorsPrintConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewPrintRequestBuilderInternal instantiates a new PrintRequestBuilder and sets the default values.
func NewPrintRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintRequestBuilder) {
    m := &PrintRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewPrintRequestBuilder instantiates a new PrintRequestBuilder and sets the default values.
func NewPrintRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get print
func (m *PrintRequestBuilder) Get(ctx context.Context, requestConfiguration *PrintRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Printable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrintFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Printable), nil
}
// Operations provides operations to manage the operations property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) Operations()(*OperationsRequestBuilder) {
    return NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) OperationsById(id string)(*OperationsPrintOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printOperation%2Did"] = id
    }
    return NewOperationsPrintOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update print
func (m *PrintRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Printable, requestConfiguration *PrintRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Printable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrintFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Printable), nil
}
// Printers provides operations to manage the printers property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) Printers()(*PrintersRequestBuilder) {
    return NewPrintersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PrintersById provides operations to manage the printers property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) PrintersById(id string)(*PrintersPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printer%2Did"] = id
    }
    return NewPrintersPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Services provides operations to manage the services property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) Services()(*ServicesRequestBuilder) {
    return NewServicesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ServicesById provides operations to manage the services property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) ServicesById(id string)(*ServicesPrintServiceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printService%2Did"] = id
    }
    return NewServicesPrintServiceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Shares provides operations to manage the shares property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) Shares()(*SharesRequestBuilder) {
    return NewSharesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SharesById provides operations to manage the shares property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) SharesById(id string)(*SharesPrinterShareItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printerShare%2Did"] = id
    }
    return NewSharesPrinterShareItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// TaskDefinitions provides operations to manage the taskDefinitions property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) TaskDefinitions()(*TaskDefinitionsRequestBuilder) {
    return NewTaskDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TaskDefinitionsById provides operations to manage the taskDefinitions property of the microsoft.graph.print entity.
func (m *PrintRequestBuilder) TaskDefinitionsById(id string)(*TaskDefinitionsPrintTaskDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printTaskDefinition%2Did"] = id
    }
    return NewTaskDefinitionsPrintTaskDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation get print
func (m *PrintRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrintRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update print
func (m *PrintRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Printable, requestConfiguration *PrintRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
