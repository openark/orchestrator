package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMicrosoftGraphAddTokenSigningCertificateRequestBuilder provides operations to call the addTokenSigningCertificate method.
type ItemMicrosoftGraphAddTokenSigningCertificateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMicrosoftGraphAddTokenSigningCertificateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMicrosoftGraphAddTokenSigningCertificateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMicrosoftGraphAddTokenSigningCertificateRequestBuilderInternal instantiates a new MicrosoftGraphAddTokenSigningCertificateRequestBuilder and sets the default values.
func NewItemMicrosoftGraphAddTokenSigningCertificateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphAddTokenSigningCertificateRequestBuilder) {
    m := &ItemMicrosoftGraphAddTokenSigningCertificateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/microsoft.graph.addTokenSigningCertificate";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMicrosoftGraphAddTokenSigningCertificateRequestBuilder instantiates a new MicrosoftGraphAddTokenSigningCertificateRequestBuilder and sets the default values.
func NewItemMicrosoftGraphAddTokenSigningCertificateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphAddTokenSigningCertificateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMicrosoftGraphAddTokenSigningCertificateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a self-signed signing certificate and return a selfSignedCertificate object, which is the public part of the generated certificate.  The self-signed signing certificate is composed of the following objects, which are added to the servicePrincipal: + The keyCredentials object with the following objects:    + A private key object with **usage** set to `Sign`.    + A public key object with **usage** set to `Verify`.+ The passwordCredentials object.  All the objects have the same value of **customKeyIdentifier**. The **passwordCredential** is used to open the PFX file (private key). It and the associated private key object have the same value of **keyId**. When set during creation through the **displayName** property, the subject of the certificate cannot be updated. The **startDateTime** is set to the same time the certificate is created using the action. The **endDateTime** can be up to three years after the certificate is created.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/serviceprincipal-addtokensigningcertificate?view=graph-rest-1.0
func (m *ItemMicrosoftGraphAddTokenSigningCertificateRequestBuilder) Post(ctx context.Context, body ItemMicrosoftGraphAddTokenSigningCertificateAddTokenSigningCertificatePostRequestBodyable, requestConfiguration *ItemMicrosoftGraphAddTokenSigningCertificateRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SelfSignedCertificateable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSelfSignedCertificateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SelfSignedCertificateable), nil
}
// ToPostRequestInformation create a self-signed signing certificate and return a selfSignedCertificate object, which is the public part of the generated certificate.  The self-signed signing certificate is composed of the following objects, which are added to the servicePrincipal: + The keyCredentials object with the following objects:    + A private key object with **usage** set to `Sign`.    + A public key object with **usage** set to `Verify`.+ The passwordCredentials object.  All the objects have the same value of **customKeyIdentifier**. The **passwordCredential** is used to open the PFX file (private key). It and the associated private key object have the same value of **keyId**. When set during creation through the **displayName** property, the subject of the certificate cannot be updated. The **startDateTime** is set to the same time the certificate is created using the action. The **endDateTime** can be up to three years after the certificate is created.
func (m *ItemMicrosoftGraphAddTokenSigningCertificateRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemMicrosoftGraphAddTokenSigningCertificateAddTokenSigningCertificatePostRequestBodyable, requestConfiguration *ItemMicrosoftGraphAddTokenSigningCertificateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
