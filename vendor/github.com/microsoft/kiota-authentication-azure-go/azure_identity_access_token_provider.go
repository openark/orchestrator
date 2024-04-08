package microsoft_kiota_authentication_azure

import (
	"context"
	"encoding/base64"
	"errors"
	"strings"

	u "net/url"

	azcore "github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azpolicy "github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	absauth "github.com/microsoft/kiota-abstractions-go/authentication"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

// AzureIdentityAccessTokenProvider implementation of AccessTokenProvider that supports implementations of TokenCredential from Azure.Identity.
type AzureIdentityAccessTokenProvider struct {
	scopes                []string
	credential            azcore.TokenCredential
	allowedHostsValidator *absauth.AllowedHostsValidator
	// The observation options for the request adapter.
	observabilityOptions ObservabilityOptions
}

// ObservabilityOptions holds the tracing, metrics and logging configuration for the request adapter
type ObservabilityOptions struct {
}

func (o ObservabilityOptions) GetTracerInstrumentationName() string {
	return "github.com/microsoft/kiota-authentication-azure-go"
}

// NewAzureIdentityAccessTokenProvider creates a new instance of the AzureIdentityAccessTokenProvider using "<scheme>://<host>/.default" as the default scope.
func NewAzureIdentityAccessTokenProvider(credential azcore.TokenCredential) (*AzureIdentityAccessTokenProvider, error) {
	return NewAzureIdentityAccessTokenProviderWithScopes(credential, nil)
}

// NewAzureIdentityAccessTokenProviderWithScopes creates a new instance of the AzureIdentityAccessTokenProvider.
func NewAzureIdentityAccessTokenProviderWithScopes(credential azcore.TokenCredential, scopes []string) (*AzureIdentityAccessTokenProvider, error) {
	return NewAzureIdentityAccessTokenProviderWithScopesAndValidHosts(credential, scopes, nil)
}

// NewAzureIdentityAccessTokenProviderWithScopesAndValidHosts creates a new instance of the AzureIdentityAccessTokenProvider.
func NewAzureIdentityAccessTokenProviderWithScopesAndValidHosts(credential azcore.TokenCredential, scopes []string, validHosts []string) (*AzureIdentityAccessTokenProvider, error) {
	return NewAzureIdentityAccessTokenProviderWithScopesAndValidHostsAndObservabilityOptions(credential, scopes, validHosts, ObservabilityOptions{})
}

// NewAzureIdentityAccessTokenProviderWithScopesAndValidHosts creates a new instance of the AzureIdentityAccessTokenProvider.
func NewAzureIdentityAccessTokenProviderWithScopesAndValidHostsAndObservabilityOptions(credential azcore.TokenCredential, scopes []string, validHosts []string, observabilityOptions ObservabilityOptions) (*AzureIdentityAccessTokenProvider, error) {
	if credential == nil {
		return nil, errors.New("credential cannot be nil")
	}
	scopesLen := len(scopes)
	finalScopes := make([]string, scopesLen)
	if scopesLen > 0 {
		copy(finalScopes, scopes)
	}
	validator := absauth.NewAllowedHostsValidator(validHosts)
	result := &AzureIdentityAccessTokenProvider{
		credential:            credential,
		scopes:                finalScopes,
		allowedHostsValidator: &validator,
		observabilityOptions:  observabilityOptions,
	}

	return result, nil
}

const claimsKey = "claims"

// GetAuthorizationToken returns the access token for the provided url.
func (p *AzureIdentityAccessTokenProvider) GetAuthorizationToken(ctx context.Context, url *u.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
	ctx, span := otel.GetTracerProvider().Tracer(p.observabilityOptions.GetTracerInstrumentationName()).Start(ctx, "GetAuthorizationToken")
	defer span.End()
	if !(*(p.allowedHostsValidator)).IsUrlHostValid(url) {
		span.SetAttributes(attribute.Bool("com.microsoft.kiota.authentication.is_url_valid", false))
		return "", nil
	}
	if !strings.EqualFold(url.Scheme, "https") {
		span.SetAttributes(attribute.Bool("com.microsoft.kiota.authentication.is_url_valid", false))
		err := errors.New("url scheme must be https")
		span.RecordError(err)
		return "", err
	}
	span.SetAttributes(attribute.Bool("com.microsoft.kiota.authentication.is_url_valid", true))

	claims := ""

	if additionalAuthenticationContext != nil &&
		additionalAuthenticationContext[claimsKey] != nil {
		if rawClaims, ok := additionalAuthenticationContext[claimsKey].(string); ok {
			decodedClaims, err := base64.StdEncoding.DecodeString(rawClaims)
			if err != nil {
				span.RecordError(err)
				return "", err
			}
			claims = string(decodedClaims)
			err = errors.New("received a claim for CAE but azure identity doesn't support claims: " + claims + " https://github.com/Azure/azure-sdk-for-go/issues/14284")
			span.RecordError(err)
			return "", err
		}
	}
	span.SetAttributes(attribute.Bool("com.microsoft.kiota.authentication.additional_claims_provided", claims != ""))

	if len(p.scopes) == 0 {
		p.scopes = append(p.scopes, url.Scheme+"://"+url.Host+"/.default")
	}

	options := azpolicy.TokenRequestOptions{
		Scopes: p.scopes,
		//TODO pass the claims once the API is updated to support it https://github.com/Azure/azure-sdk-for-go/issues/14284
	}
	span.SetAttributes(attribute.String("com.microsoft.kiota.authentication.scopes", strings.Join(p.scopes, ",")))
	token, err := p.credential.GetToken(ctx, options)
	if err != nil {
		span.RecordError(err)
		return "", err
	}
	return token.Token, nil
}

// GetAllowedHostsValidator returns the hosts validator.
func (p *AzureIdentityAccessTokenProvider) GetAllowedHostsValidator() *absauth.AllowedHostsValidator {
	return p.allowedHostsValidator
}
