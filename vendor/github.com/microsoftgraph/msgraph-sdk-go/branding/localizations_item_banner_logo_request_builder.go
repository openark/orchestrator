package branding

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// LocalizationsItemBannerLogoRequestBuilder provides operations to manage the media for the organizationalBranding entity.
type LocalizationsItemBannerLogoRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// LocalizationsItemBannerLogoRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LocalizationsItemBannerLogoRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// LocalizationsItemBannerLogoRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LocalizationsItemBannerLogoRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLocalizationsItemBannerLogoRequestBuilderInternal instantiates a new BannerLogoRequestBuilder and sets the default values.
func NewLocalizationsItemBannerLogoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LocalizationsItemBannerLogoRequestBuilder) {
    m := &LocalizationsItemBannerLogoRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/branding/localizations/{organizationalBrandingLocalization%2Did}/bannerLogo";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewLocalizationsItemBannerLogoRequestBuilder instantiates a new BannerLogoRequestBuilder and sets the default values.
func NewLocalizationsItemBannerLogoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LocalizationsItemBannerLogoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLocalizationsItemBannerLogoRequestBuilderInternal(urlParams, requestAdapter)
}
// Get a banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG no larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/organizationalbranding-list-localizations?view=graph-rest-1.0
func (m *LocalizationsItemBannerLogoRequestBuilder) Get(ctx context.Context, requestConfiguration *LocalizationsItemBannerLogoRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Put a banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG no larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
func (m *LocalizationsItemBannerLogoRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *LocalizationsItemBannerLogoRequestBuilderPutRequestConfiguration)(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation a banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG no larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
func (m *LocalizationsItemBannerLogoRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LocalizationsItemBannerLogoRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPutRequestInformation a banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG no larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
func (m *LocalizationsItemBannerLogoRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *LocalizationsItemBannerLogoRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.SetStreamContent(body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
