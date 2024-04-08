package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrintersMicrosoftGraphCreateRequestBuilder provides operations to call the create method.
type PrintersMicrosoftGraphCreateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PrintersMicrosoftGraphCreateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintersMicrosoftGraphCreateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPrintersMicrosoftGraphCreateRequestBuilderInternal instantiates a new MicrosoftGraphCreateRequestBuilder and sets the default values.
func NewPrintersMicrosoftGraphCreateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersMicrosoftGraphCreateRequestBuilder) {
    m := &PrintersMicrosoftGraphCreateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/printers/microsoft.graph.create";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewPrintersMicrosoftGraphCreateRequestBuilder instantiates a new MicrosoftGraphCreateRequestBuilder and sets the default values.
func NewPrintersMicrosoftGraphCreateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersMicrosoftGraphCreateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintersMicrosoftGraphCreateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create (register) a printer with the Universal Print service. This is a long-running operation and as such, it returns a printerCreateOperation that can be used to track and verify the registration of the printer. For help creating the required Certificate Signing Request (CSR) for creating printer, see the CSR generation code sample.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/printer-create?view=graph-rest-1.0
func (m *PrintersMicrosoftGraphCreateRequestBuilder) Post(ctx context.Context, body PrintersMicrosoftGraphCreateCreatePostRequestBodyable, requestConfiguration *PrintersMicrosoftGraphCreateRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation create (register) a printer with the Universal Print service. This is a long-running operation and as such, it returns a printerCreateOperation that can be used to track and verify the registration of the printer. For help creating the required Certificate Signing Request (CSR) for creating printer, see the CSR generation code sample.
func (m *PrintersMicrosoftGraphCreateRequestBuilder) ToPostRequestInformation(ctx context.Context, body PrintersMicrosoftGraphCreateCreatePostRequestBodyable, requestConfiguration *PrintersMicrosoftGraphCreateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
