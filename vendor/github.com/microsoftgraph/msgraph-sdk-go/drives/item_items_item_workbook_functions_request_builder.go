package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookFunctionsRequestBuilder provides operations to manage the functions property of the microsoft.graph.workbook entity.
type ItemItemsItemWorkbookFunctionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemWorkbookFunctionsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookFunctionsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemsItemWorkbookFunctionsRequestBuilderGetQueryParameters get functions from drives
type ItemItemsItemWorkbookFunctionsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemWorkbookFunctionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookFunctionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookFunctionsRequestBuilderGetQueryParameters
}
// ItemItemsItemWorkbookFunctionsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookFunctionsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookFunctionsRequestBuilderInternal instantiates a new FunctionsRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsRequestBuilder) {
    m := &ItemItemsItemWorkbookFunctionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/functions{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsItemWorkbookFunctionsRequestBuilder instantiates a new FunctionsRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookFunctionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property functions for drives
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookFunctionsRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get functions from drives
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookFunctionsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookFunctionsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionsable), nil
}
// MicrosoftGraphAbs provides operations to call the abs method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAbs()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAbsRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAbsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAccrInt provides operations to call the accrInt method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAccrInt()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAccrIntM provides operations to call the accrIntM method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAccrIntM()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntMRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAccrIntMRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAcos provides operations to call the acos method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAcos()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAcosRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAcosRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAcosh provides operations to call the acosh method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAcosh()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAcoshRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAcoshRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAcot provides operations to call the acot method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAcot()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAcotRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAcotRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAcoth provides operations to call the acoth method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAcoth()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAcothRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAcothRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAmorDegrc provides operations to call the amorDegrc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAmorDegrc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAmorLinc provides operations to call the amorLinc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAmorLinc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorLincRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAmorLincRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAnd provides operations to call the and method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAnd()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAndRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAndRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphArabic provides operations to call the arabic method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphArabic()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphArabicRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphArabicRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAreas provides operations to call the areas method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAreas()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAreasRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAreasRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAsc provides operations to call the asc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAsc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAscRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAscRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAsin provides operations to call the asin method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAsin()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAsinRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAsinRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAsinh provides operations to call the asinh method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAsinh()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAsinhRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAsinhRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAtan provides operations to call the atan method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAtan()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAtanRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAtanRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAtan2 provides operations to call the atan2 method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAtan2()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAtan2RequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAtan2RequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAtanh provides operations to call the atanh method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAtanh()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAtanhRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAtanhRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAveDev provides operations to call the aveDev method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAveDev()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAveDevRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAveDevRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAverage provides operations to call the average method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAverage()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAverageRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAverageA provides operations to call the averageA method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAverageA()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageARequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAverageARequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAverageIf provides operations to call the averageIf method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAverageIf()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAverageIfs provides operations to call the averageIfs method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphAverageIfs()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfsRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBahtText provides operations to call the bahtText method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBahtText()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBahtTextRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBahtTextRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBase provides operations to call the base method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBase()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBaseRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBesselI provides operations to call the besselI method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBesselI()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBesselIRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBesselIRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBesselJ provides operations to call the besselJ method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBesselJ()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBesselJRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBesselJRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBesselK provides operations to call the besselK method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBesselK()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBesselKRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBesselKRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBesselY provides operations to call the besselY method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBesselY()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBesselYRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBesselYRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBeta_Dist provides operations to call the beta_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBeta_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBeta_Inv provides operations to call the beta_Inv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBeta_Inv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_InvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBin2Dec provides operations to call the bin2Dec method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBin2Dec()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBin2DecRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBin2DecRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBin2Hex provides operations to call the bin2Hex method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBin2Hex()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBin2HexRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBin2HexRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBin2Oct provides operations to call the bin2Oct method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBin2Oct()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBin2OctRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBin2OctRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBinom_Dist provides operations to call the binom_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBinom_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBinom_Dist_Range provides operations to call the binom_Dist_Range method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBinom_Dist_Range()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBinom_Inv provides operations to call the binom_Inv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBinom_Inv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBitand provides operations to call the bitand method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBitand()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBitandRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBitandRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBitlshift provides operations to call the bitlshift method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBitlshift()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBitlshiftRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBitlshiftRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBitor provides operations to call the bitor method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBitor()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBitorRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBitorRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBitrshift provides operations to call the bitrshift method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBitrshift()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBitrshiftRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBitrshiftRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBitxor provides operations to call the bitxor method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphBitxor()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBitxorRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBitxorRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCeiling_Math provides operations to call the ceiling_Math method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCeiling_Math()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCeiling_Precise provides operations to call the ceiling_Precise method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCeiling_Precise()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_PreciseRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_PreciseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphChar provides operations to call the char method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphChar()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCharRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCharRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphChiSq_Dist provides operations to call the chiSq_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphChiSq_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphChiSq_Dist_RT provides operations to call the chiSq_Dist_RT method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphChiSq_Dist_RT()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Dist_RTRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphChiSq_Inv provides operations to call the chiSq_Inv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphChiSq_Inv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_InvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphChiSq_Inv_RT provides operations to call the chiSq_Inv_RT method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphChiSq_Inv_RT()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphChoose provides operations to call the choose method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphChoose()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphChooseRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphChooseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphClean provides operations to call the clean method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphClean()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCleanRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCleanRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCode provides operations to call the code method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCode()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCodeRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCodeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphColumns provides operations to call the columns method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphColumns()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphColumnsRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCombin provides operations to call the combin method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCombin()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCombinRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCombinRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCombina provides operations to call the combina method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCombina()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCombinaRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCombinaRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphComplex provides operations to call the complex method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphComplex()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphComplexRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphComplexRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphConcatenate provides operations to call the concatenate method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphConcatenate()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphConcatenateRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphConcatenateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphConfidence_Norm provides operations to call the confidence_Norm method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphConfidence_Norm()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_NormRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_NormRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphConfidence_T provides operations to call the confidence_T method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphConfidence_T()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphConvert provides operations to call the convert method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphConvert()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphConvertRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphConvertRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCos provides operations to call the cos method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCos()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCosRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCosRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCosh provides operations to call the cosh method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCosh()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCoshRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCoshRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCot provides operations to call the cot method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCot()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCotRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCotRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCoth provides operations to call the coth method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCoth()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCothRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCothRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCount provides operations to call the count method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCount()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCountRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCountA provides operations to call the countA method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCountA()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCountARequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCountARequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCountBlank provides operations to call the countBlank method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCountBlank()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCountBlankRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCountBlankRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCountIf provides operations to call the countIf method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCountIf()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCountIfRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCountIfRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCountIfs provides operations to call the countIfs method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCountIfs()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCountIfsRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCountIfsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCoupDayBs provides operations to call the coupDayBs method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCoupDayBs()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCoupDayBsRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCoupDayBsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCoupDays provides operations to call the coupDays method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCoupDays()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCoupDaysRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCoupDaysRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCoupDaysNc provides operations to call the coupDaysNc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCoupDaysNc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCoupDaysNcRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCoupDaysNcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCoupNcd provides operations to call the coupNcd method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCoupNcd()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCoupNcdRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCoupNcdRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCoupNum provides operations to call the coupNum method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCoupNum()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCoupNumRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCoupNumRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCoupPcd provides operations to call the coupPcd method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCoupPcd()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCoupPcdRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCoupPcdRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCsc provides operations to call the csc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCsc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCscRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCscRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCsch provides operations to call the csch method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCsch()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCschRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCschRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCumIPmt provides operations to call the cumIPmt method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCumIPmt()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCumIPmtRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCumIPmtRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCumPrinc provides operations to call the cumPrinc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphCumPrinc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDate provides operations to call the date method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDate()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDateRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDatevalue provides operations to call the datevalue method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDatevalue()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDatevalueRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDatevalueRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDaverage provides operations to call the daverage method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDaverage()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDaverageRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDaverageRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDay provides operations to call the day method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDay()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDayRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDayRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDays provides operations to call the days method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDays()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDaysRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDaysRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDays360 provides operations to call the days360 method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDays360()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDays360RequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDays360RequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDb provides operations to call the db method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDb()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDbRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDbRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDbcs provides operations to call the dbcs method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDbcs()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDbcsRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDbcsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDcount provides operations to call the dcount method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDcount()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDcountRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDcountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDcountA provides operations to call the dcountA method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDcountA()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDcountARequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDcountARequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDdb provides operations to call the ddb method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDdb()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDdbRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDdbRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDec2Bin provides operations to call the dec2Bin method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDec2Bin()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDec2BinRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDec2BinRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDec2Hex provides operations to call the dec2Hex method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDec2Hex()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDec2HexRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDec2HexRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDec2Oct provides operations to call the dec2Oct method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDec2Oct()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDec2OctRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDec2OctRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDecimal provides operations to call the decimal method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDecimal()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDecimalRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDecimalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDegrees provides operations to call the degrees method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDegrees()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDegreesRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDegreesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDelta provides operations to call the delta method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDelta()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDeltaRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDevSq provides operations to call the devSq method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDevSq()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDevSqRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDevSqRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDget provides operations to call the dget method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDget()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDgetRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDgetRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDisc provides operations to call the disc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDisc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDiscRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDiscRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDmax provides operations to call the dmax method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDmax()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDmaxRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDmaxRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDmin provides operations to call the dmin method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDmin()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDminRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDminRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDollar provides operations to call the dollar method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDollar()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDollarRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDollarRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDollarDe provides operations to call the dollarDe method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDollarDe()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDollarDeRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDollarDeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDollarFr provides operations to call the dollarFr method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDollarFr()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDproduct provides operations to call the dproduct method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDproduct()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDproductRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDproductRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDstDev provides operations to call the dstDev method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDstDev()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDstDevRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDstDevRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDstDevP provides operations to call the dstDevP method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDstDevP()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDstDevPRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDstDevPRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDsum provides operations to call the dsum method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDsum()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDsumRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDsumRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDuration provides operations to call the duration method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDuration()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDurationRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDurationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDvar provides operations to call the dvar method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDvar()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDvarRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDvarRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDvarP provides operations to call the dvarP method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphDvarP()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEcma_Ceiling provides operations to call the ecma_Ceiling method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphEcma_Ceiling()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphEcma_CeilingRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphEcma_CeilingRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEdate provides operations to call the edate method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphEdate()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphEdateRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphEdateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEffect provides operations to call the effect method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphEffect()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphEffectRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphEffectRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEoMonth provides operations to call the eoMonth method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphEoMonth()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphEoMonthRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphEoMonthRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphErf provides operations to call the erf method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphErf()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphErfRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphErfRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphErf_Precise provides operations to call the erf_Precise method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphErf_Precise()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphErf_PreciseRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphErf_PreciseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphErfC provides operations to call the erfC method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphErfC()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphErfCRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphErfCRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphErfC_Precise provides operations to call the erfC_Precise method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphErfC_Precise()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphErfC_PreciseRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphErfC_PreciseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphError_Type provides operations to call the error_Type method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphError_Type()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphError_TypeRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphError_TypeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEven provides operations to call the even method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphEven()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphEvenRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphEvenRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphExact provides operations to call the exact method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphExact()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphExactRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphExactRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphExp provides operations to call the exp method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphExp()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphExpRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphExpRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphExpon_Dist provides operations to call the expon_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphExpon_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphExpon_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphExpon_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphF_Dist provides operations to call the f_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphF_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphF_Dist_RT provides operations to call the f_Dist_RT method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphF_Dist_RT()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphF_Inv provides operations to call the f_Inv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphF_Inv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphF_InvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphF_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphF_Inv_RT provides operations to call the f_Inv_RT method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphF_Inv_RT()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphF_Inv_RTRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphF_Inv_RTRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphFact provides operations to call the fact method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFact()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFactRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFactRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphFactDouble provides operations to call the factDouble method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFactDouble()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFactDoubleRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFactDoubleRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphFalse provides operations to call the false method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFalse()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFalseRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFalseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphFind provides operations to call the find method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFind()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFindRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFindRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphFindB provides operations to call the findB method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFindB()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFindBRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFindBRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphFisher provides operations to call the fisher method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFisher()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFisherRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFisherRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphFisherInv provides operations to call the fisherInv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFisherInv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFisherInvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFisherInvRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphFixed provides operations to call the fixed method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFixed()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFixedRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFixedRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphFloor_Math provides operations to call the floor_Math method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFloor_Math()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_MathRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_MathRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphFloor_Precise provides operations to call the floor_Precise method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFloor_Precise()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFloor_PreciseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphFv provides operations to call the fv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFvRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphFvschedule provides operations to call the fvschedule method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphFvschedule()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphFvscheduleRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphFvscheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGamma provides operations to call the gamma method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphGamma()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphGammaRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphGammaRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGamma_Dist provides operations to call the gamma_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphGamma_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphGamma_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphGamma_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGamma_Inv provides operations to call the gamma_Inv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphGamma_Inv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphGamma_InvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphGamma_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGammaLn provides operations to call the gammaLn method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphGammaLn()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphGammaLnRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphGammaLnRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGammaLn_Precise provides operations to call the gammaLn_Precise method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphGammaLn_Precise()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphGammaLn_PreciseRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphGammaLn_PreciseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGauss provides operations to call the gauss method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphGauss()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphGaussRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphGaussRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGcd provides operations to call the gcd method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphGcd()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphGcdRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphGcdRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGeoMean provides operations to call the geoMean method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphGeoMean()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphGeoMeanRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphGeoMeanRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGeStep provides operations to call the geStep method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphGeStep()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphGeStepRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphGeStepRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphHarMean provides operations to call the harMean method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphHarMean()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphHarMeanRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphHarMeanRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphHex2Bin provides operations to call the hex2Bin method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphHex2Bin()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphHex2BinRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphHex2BinRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphHex2Dec provides operations to call the hex2Dec method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphHex2Dec()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphHex2DecRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphHex2DecRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphHex2Oct provides operations to call the hex2Oct method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphHex2Oct()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphHex2OctRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphHex2OctRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphHlookup provides operations to call the hlookup method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphHlookup()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphHour provides operations to call the hour method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphHour()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphHourRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphHourRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphHyperlink provides operations to call the hyperlink method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphHyperlink()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphHypGeom_Dist provides operations to call the hypGeom_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphHypGeom_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphIf provides operations to call the if method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIf()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIfRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIfRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImAbs provides operations to call the imAbs method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImAbs()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImAbsRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImAbsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImaginary provides operations to call the imaginary method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImaginary()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImaginaryRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImaginaryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImArgument provides operations to call the imArgument method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImArgument()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImArgumentRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImArgumentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImConjugate provides operations to call the imConjugate method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImConjugate()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImConjugateRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImConjugateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImCos provides operations to call the imCos method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImCos()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImCosRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImCosRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImCosh provides operations to call the imCosh method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImCosh()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImCoshRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImCoshRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImCot provides operations to call the imCot method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImCot()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImCotRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImCotRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImCsc provides operations to call the imCsc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImCsc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImCscRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImCscRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImCsch provides operations to call the imCsch method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImCsch()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImCschRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImCschRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImDiv provides operations to call the imDiv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImDiv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImDivRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImDivRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImExp provides operations to call the imExp method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImExp()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImExpRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImExpRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImLn provides operations to call the imLn method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImLn()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImLnRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImLnRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImLog10 provides operations to call the imLog10 method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImLog10()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImLog10RequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImLog10RequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImLog2 provides operations to call the imLog2 method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImLog2()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImLog2RequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImLog2RequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImPower provides operations to call the imPower method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImPower()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImPowerRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImPowerRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImProduct provides operations to call the imProduct method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImProduct()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImProductRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImProductRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImReal provides operations to call the imReal method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImReal()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImRealRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImRealRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImSec provides operations to call the imSec method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImSec()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImSecRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImSecRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImSech provides operations to call the imSech method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImSech()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImSechRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImSechRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImSin provides operations to call the imSin method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImSin()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImSinRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImSinRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImSinh provides operations to call the imSinh method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImSinh()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImSinhRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImSinhRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImSqrt provides operations to call the imSqrt method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImSqrt()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImSqrtRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImSqrtRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImSub provides operations to call the imSub method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImSub()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImSubRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImSubRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImSum provides operations to call the imSum method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImSum()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImSumRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImSumRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphImTan provides operations to call the imTan method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphImTan()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphImTanRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphImTanRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphInt provides operations to call the int method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphInt()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIntRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIntRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphIntRate provides operations to call the intRate method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIntRate()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIntRateRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIntRateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphIpmt provides operations to call the ipmt method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIpmt()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIpmtRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIpmtRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphIrr provides operations to call the irr method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIrr()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIrrRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIrrRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphIsErr provides operations to call the isErr method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsErr()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsErrRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsErrRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphIsError provides operations to call the isError method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsError()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsErrorRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsErrorRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphIsEven provides operations to call the isEven method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsEven()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsEvenRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsEvenRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphIsFormula provides operations to call the isFormula method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsFormula()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsFormulaRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsFormulaRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphIsLogical provides operations to call the isLogical method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsLogical()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsLogicalRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsLogicalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphIsNA provides operations to call the isNA method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsNA()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsNARequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsNARequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphIsNonText provides operations to call the isNonText method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsNonText()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsNonTextRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsNonTextRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphIsNumber provides operations to call the isNumber method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsNumber()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsNumberRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsNumberRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphIso_Ceiling provides operations to call the iso_Ceiling method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIso_Ceiling()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIso_CeilingRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIso_CeilingRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphIsOdd provides operations to call the isOdd method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsOdd()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsOddRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsOddRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphIsoWeekNum provides operations to call the isoWeekNum method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsoWeekNum()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsoWeekNumRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsoWeekNumRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphIspmt provides operations to call the ispmt method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIspmt()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIspmtRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIspmtRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphIsref provides operations to call the isref method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsref()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsrefRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsrefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphIsText provides operations to call the isText method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphIsText()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIsTextRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIsTextRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphKurt provides operations to call the kurt method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphKurt()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphKurtRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphKurtRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLarge provides operations to call the large method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLarge()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLargeRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLargeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLcm provides operations to call the lcm method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLcm()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLcmRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLcmRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLeft provides operations to call the left method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLeft()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLeftRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLeftRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLeftb provides operations to call the leftb method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLeftb()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLeftbRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLeftbRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLen provides operations to call the len method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLen()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLenRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLenRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLenb provides operations to call the lenb method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLenb()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLenbRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLenbRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLn provides operations to call the ln method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLn()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLnRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLnRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLog provides operations to call the log method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLog()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLogRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLogRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLog10 provides operations to call the log10 method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLog10()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLog10RequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLog10RequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLogNorm_Dist provides operations to call the logNorm_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLogNorm_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLogNorm_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLogNorm_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLogNorm_Inv provides operations to call the logNorm_Inv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLogNorm_Inv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLogNorm_InvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLogNorm_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLookup provides operations to call the lookup method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLookup()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLookupRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLookupRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLower provides operations to call the lower method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphLower()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphLowerRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphLowerRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMatch provides operations to call the match method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMatch()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMatchRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMatchRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMax provides operations to call the max method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMax()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMaxRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMaxRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMaxA provides operations to call the maxA method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMaxA()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMaxARequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMaxARequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMduration provides operations to call the mduration method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMduration()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMdurationRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMdurationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMedian provides operations to call the median method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMedian()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMedianRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMedianRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMid provides operations to call the mid method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMid()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMidRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMidRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMidb provides operations to call the midb method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMidb()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMidbRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMidbRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMin provides operations to call the min method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMin()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMinRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMinRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMinA provides operations to call the minA method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMinA()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMinARequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMinARequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMinute provides operations to call the minute method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMinute()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMinuteRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMinuteRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMirr provides operations to call the mirr method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMirr()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMirrRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMirrRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMod provides operations to call the mod method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMod()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphModRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphModRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMonth provides operations to call the month method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMonth()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMonthRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMonthRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMround provides operations to call the mround method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMround()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMroundRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMroundRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMultiNomial provides operations to call the multiNomial method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphMultiNomial()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphMultiNomialRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphMultiNomialRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphN provides operations to call the n method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphN()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphNa provides operations to call the na method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNa()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNaRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNaRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphNegBinom_Dist provides operations to call the negBinom_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNegBinom_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNegBinom_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNegBinom_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphNetworkDays provides operations to call the networkDays method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNetworkDays()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphNetworkDays_Intl provides operations to call the networkDays_Intl method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNetworkDays_Intl()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDays_IntlRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDays_IntlRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphNominal provides operations to call the nominal method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNominal()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNominalRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNominalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphNorm_Dist provides operations to call the norm_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNorm_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphNorm_Inv provides operations to call the norm_Inv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNorm_Inv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphNorm_S_Dist provides operations to call the norm_S_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNorm_S_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_S_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_S_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphNorm_S_Inv provides operations to call the norm_S_Inv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNorm_S_Inv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_S_InvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_S_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphNot provides operations to call the not method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNot()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNotRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNotRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphNow provides operations to call the now method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNow()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNowRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNowRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphNper provides operations to call the nper method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNper()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNperRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNperRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphNpv provides operations to call the npv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNpv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNpvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNpvRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphNumberValue provides operations to call the numberValue method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphNumberValue()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphOct2Bin provides operations to call the oct2Bin method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphOct2Bin()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphOct2BinRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphOct2BinRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphOct2Dec provides operations to call the oct2Dec method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphOct2Dec()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphOct2DecRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphOct2DecRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphOct2Hex provides operations to call the oct2Hex method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphOct2Hex()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphOct2HexRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphOct2HexRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphOdd provides operations to call the odd method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphOdd()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphOddRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphOddRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphOddFPrice provides operations to call the oddFPrice method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphOddFPrice()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphOddFPriceRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphOddFPriceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphOddFYield provides operations to call the oddFYield method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphOddFYield()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphOddFYieldRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphOddFYieldRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphOddLPrice provides operations to call the oddLPrice method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphOddLPrice()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphOddLPriceRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphOddLPriceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphOddLYield provides operations to call the oddLYield method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphOddLYield()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphOddLYieldRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphOddLYieldRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphOr provides operations to call the or method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphOr()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphOrRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphOrRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPduration provides operations to call the pduration method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPduration()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPdurationRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPdurationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPercentile_Exc provides operations to call the percentile_Exc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPercentile_Exc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPercentile_ExcRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPercentile_ExcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPercentile_Inc provides operations to call the percentile_Inc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPercentile_Inc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPercentile_IncRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPercentile_IncRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPercentRank_Exc provides operations to call the percentRank_Exc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPercentRank_Exc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPercentRank_ExcRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPercentRank_ExcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPercentRank_Inc provides operations to call the percentRank_Inc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPercentRank_Inc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPercentRank_IncRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPercentRank_IncRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPermut provides operations to call the permut method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPermut()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPermutRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPermutRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPermutationa provides operations to call the permutationa method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPermutationa()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPermutationaRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPermutationaRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPhi provides operations to call the phi method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPhi()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPhiRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPhiRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPi provides operations to call the pi method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPi()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPiRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPiRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPmt provides operations to call the pmt method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPmt()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPmtRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPmtRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPoisson_Dist provides operations to call the poisson_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPoisson_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPoisson_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPoisson_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPower provides operations to call the power method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPower()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPowerRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPowerRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPpmt provides operations to call the ppmt method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPpmt()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPpmtRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPpmtRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPrice provides operations to call the price method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPrice()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPriceRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPriceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPriceDisc provides operations to call the priceDisc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPriceDisc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPriceDiscRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPriceDiscRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPriceMat provides operations to call the priceMat method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPriceMat()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPriceMatRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPriceMatRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphProduct provides operations to call the product method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphProduct()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphProductRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphProductRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphProper provides operations to call the proper method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphProper()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphProperRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphProperRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPv provides operations to call the pv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphPv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphPvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphPvRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphQuartile_Exc provides operations to call the quartile_Exc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphQuartile_Exc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_ExcRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_ExcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphQuartile_Inc provides operations to call the quartile_Inc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphQuartile_Inc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphQuotient provides operations to call the quotient method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphQuotient()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRadians provides operations to call the radians method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRadians()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRadiansRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRadiansRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRand provides operations to call the rand method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRand()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRandRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRandRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRandBetween provides operations to call the randBetween method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRandBetween()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRandBetweenRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRandBetweenRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRank_Avg provides operations to call the rank_Avg method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRank_Avg()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRank_AvgRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRank_AvgRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRank_Eq provides operations to call the rank_Eq method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRank_Eq()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRank_EqRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRank_EqRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRate provides operations to call the rate method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRate()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRateRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReceived provides operations to call the received method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphReceived()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphReceivedRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphReceivedRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReplace provides operations to call the replace method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphReplace()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReplaceB provides operations to call the replaceB method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphReplaceB()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRept provides operations to call the rept method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRept()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphReptRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphReptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRight provides operations to call the right method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRight()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRightRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRightRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRightb provides operations to call the rightb method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRightb()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRightbRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRightbRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRoman provides operations to call the roman method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRoman()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRomanRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRomanRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRound provides operations to call the round method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRound()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRoundRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRoundRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRoundDown provides operations to call the roundDown method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRoundDown()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRoundDownRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRoundDownRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRoundUp provides operations to call the roundUp method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRoundUp()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRoundUpRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRoundUpRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRows provides operations to call the rows method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRows()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRowsRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRowsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRri provides operations to call the rri method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphRri()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphRriRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphRriRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSec provides operations to call the sec method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSec()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSecRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSecRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSech provides operations to call the sech method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSech()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSechRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSechRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSecond provides operations to call the second method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSecond()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSecondRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSecondRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSeriesSum provides operations to call the seriesSum method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSeriesSum()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSheet provides operations to call the sheet method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSheet()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSheetRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSheetRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSheets provides operations to call the sheets method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSheets()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSheetsRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSheetsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSign provides operations to call the sign method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSign()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSignRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSignRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSin provides operations to call the sin method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSin()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSinRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSinRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSinh provides operations to call the sinh method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSinh()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSinhRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSinhRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSkew provides operations to call the skew method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSkew()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSkewRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSkewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSkew_p provides operations to call the skew_p method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSkew_p()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSkew_pRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSkew_pRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSln provides operations to call the sln method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSln()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSlnRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSlnRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSmall provides operations to call the small method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSmall()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSmallRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSmallRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSqrt provides operations to call the sqrt method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSqrt()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSqrtRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSqrtRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSqrtPi provides operations to call the sqrtPi method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSqrtPi()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSqrtPiRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSqrtPiRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphStandardize provides operations to call the standardize method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphStandardize()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphStandardizeRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphStandardizeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphStDev_P provides operations to call the stDev_P method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphStDev_P()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphStDev_PRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphStDev_PRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphStDev_S provides operations to call the stDev_S method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphStDev_S()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphStDev_SRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphStDev_SRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphStDevA provides operations to call the stDevA method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphStDevA()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphStDevARequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphStDevARequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphStDevPA provides operations to call the stDevPA method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphStDevPA()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphStDevPARequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphStDevPARequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSubstitute provides operations to call the substitute method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSubstitute()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSubtotal provides operations to call the subtotal method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSubtotal()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSubtotalRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSubtotalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSum provides operations to call the sum method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSum()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSumRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSumRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSumIf provides operations to call the sumIf method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSumIf()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSumIfRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSumIfRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSumIfs provides operations to call the sumIfs method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSumIfs()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSumIfsRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSumIfsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSumSq provides operations to call the sumSq method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSumSq()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSumSqRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSumSqRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSyd provides operations to call the syd method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphSyd()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSydRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSydRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphT provides operations to call the t method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphT()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphT_Dist provides operations to call the t_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphT_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphT_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphT_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphT_Dist_2T provides operations to call the t_Dist_2T method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphT_Dist_2T()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphT_Dist_2TRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphT_Dist_2TRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphT_Dist_RT provides operations to call the t_Dist_RT method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphT_Dist_RT()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphT_Dist_RTRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphT_Dist_RTRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphT_Inv provides operations to call the t_Inv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphT_Inv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphT_InvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphT_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphT_Inv_2T provides operations to call the t_Inv_2T method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphT_Inv_2T()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphT_Inv_2TRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphT_Inv_2TRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTan provides operations to call the tan method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTan()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTanRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTanRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTanh provides operations to call the tanh method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTanh()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTanhRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTanhRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTbillEq provides operations to call the tbillEq method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTbillEq()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTbillEqRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTbillEqRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTbillPrice provides operations to call the tbillPrice method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTbillPrice()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTbillPriceRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTbillPriceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTbillYield provides operations to call the tbillYield method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTbillYield()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTbillYieldRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTbillYieldRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphText provides operations to call the text method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphText()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTextRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTextRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTime provides operations to call the time method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTime()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTimeRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTimevalue provides operations to call the timevalue method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTimevalue()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTimevalueRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTimevalueRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphToday provides operations to call the today method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphToday()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTodayRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTodayRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTrim provides operations to call the trim method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTrim()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTrimRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTrimRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTrimMean provides operations to call the trimMean method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTrimMean()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTrimMeanRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTrimMeanRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTrue provides operations to call the true method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTrue()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTrueRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTrueRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTrunc provides operations to call the trunc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphTrunc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTruncRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTruncRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphType provides operations to call the type method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphType()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphTypeRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphTypeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUnichar provides operations to call the unichar method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphUnichar()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphUnicharRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphUnicharRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUnicode provides operations to call the unicode method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphUnicode()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphUnicodeRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphUnicodeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUpper provides operations to call the upper method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphUpper()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphUpperRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphUpperRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUsdollar provides operations to call the usdollar method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphUsdollar()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphUsdollarRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphUsdollarRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphValue provides operations to call the value method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphValue()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphValueRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphValueRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphVar_P provides operations to call the var_P method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphVar_P()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphVar_PRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphVar_PRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphVar_S provides operations to call the var_S method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphVar_S()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphVar_SRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphVar_SRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphVarA provides operations to call the varA method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphVarA()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphVarARequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphVarARequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphVarPA provides operations to call the varPA method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphVarPA()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphVarPARequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphVarPARequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphVdb provides operations to call the vdb method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphVdb()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphVdbRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphVlookup provides operations to call the vlookup method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphVlookup()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWeekday provides operations to call the weekday method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphWeekday()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphWeekdayRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphWeekdayRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWeekNum provides operations to call the weekNum method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphWeekNum()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphWeekNumRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphWeekNumRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWeibull_Dist provides operations to call the weibull_Dist method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphWeibull_Dist()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphWeibull_DistRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphWeibull_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWorkDay provides operations to call the workDay method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphWorkDay()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDayRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDayRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWorkDay_Intl provides operations to call the workDay_Intl method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphWorkDay_Intl()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphXirr provides operations to call the xirr method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphXirr()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphXirrRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphXirrRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphXnpv provides operations to call the xnpv method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphXnpv()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphXnpvRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphXnpvRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphXor provides operations to call the xor method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphXor()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphXorRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphXorRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphYear provides operations to call the year method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphYear()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphYearRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphYearRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphYearFrac provides operations to call the yearFrac method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphYearFrac()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphYearFracRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphYearFracRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphYield provides operations to call the yield method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphYield()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphYieldRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphYieldRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphYieldDisc provides operations to call the yieldDisc method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphYieldDisc()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphYieldDiscRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphYieldDiscRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphYieldMat provides operations to call the yieldMat method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphYieldMat()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphYieldMatRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphYieldMatRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphZ_Test provides operations to call the z_Test method.
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) MicrosoftGraphZ_Test()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property functions in drives
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionsable, requestConfiguration *ItemItemsItemWorkbookFunctionsRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookFunctionsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionsable), nil
}
// ToDeleteRequestInformation delete navigation property functions for drives
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookFunctionsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get functions from drives
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookFunctionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property functions in drives
func (m *ItemItemsItemWorkbookFunctionsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionsable, requestConfiguration *ItemItemsItemWorkbookFunctionsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
