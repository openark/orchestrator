package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApiApplication 
type ApiApplication struct {
    // When true, allows an application to use claims mapping without specifying a custom signing key.
    acceptMappedClaims *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app. If you set the appID of the client app to this value, the user only consents once to the client app. Azure AD knows that consenting to the client means implicitly consenting to the web API and automatically provisions service principals for both APIs at the same time. Both the client and the web API app must be registered in the same tenant.
    knownClientApplications []i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The definition of the delegated permissions exposed by the web API represented by this application registration. These delegated permissions may be requested by a client application, and may be granted by users or administrators during consent. Delegated permissions are sometimes referred to as OAuth 2.0 scopes.
    oauth2PermissionScopes []PermissionScopeable
    // The OdataType property
    odataType *string
    // Lists the client applications that are pre-authorized with the specified delegated permissions to access this application's APIs. Users are not required to consent to any pre-authorized application (for the permissions specified). However, any additional permissions not listed in preAuthorizedApplications (requested through incremental consent for example) will require user consent.
    preAuthorizedApplications []PreAuthorizedApplicationable
    // Specifies the access token version expected by this resource. This changes the version and format of the JWT produced independent of the endpoint or client used to request the access token.  The endpoint used, v1.0 or v2.0, is chosen by the client and only impacts the version of id_tokens. Resources need to explicitly configure requestedAccessTokenVersion to indicate the supported access token format.  Possible values for requestedAccessTokenVersion are 1, 2, or null. If the value is null, this defaults to 1, which corresponds to the v1.0 endpoint.  If signInAudience on the application is configured as AzureADandPersonalMicrosoftAccount, the value for this property must be 2
    requestedAccessTokenVersion *int32
}
// NewApiApplication instantiates a new apiApplication and sets the default values.
func NewApiApplication()(*ApiApplication) {
    m := &ApiApplication{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateApiApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApiApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApiApplication(), nil
}
// GetAcceptMappedClaims gets the acceptMappedClaims property value. When true, allows an application to use claims mapping without specifying a custom signing key.
func (m *ApiApplication) GetAcceptMappedClaims()(*bool) {
    return m.acceptMappedClaims
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApiApplication) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApiApplication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["acceptMappedClaims"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceptMappedClaims(val)
        }
        return nil
    }
    res["knownClientApplications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID, len(val))
            for i, v := range val {
                res[i] = *(v.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID))
            }
            m.SetKnownClientApplications(res)
        }
        return nil
    }
    res["oauth2PermissionScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PermissionScopeable, len(val))
            for i, v := range val {
                res[i] = v.(PermissionScopeable)
            }
            m.SetOauth2PermissionScopes(res)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["preAuthorizedApplications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePreAuthorizedApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PreAuthorizedApplicationable, len(val))
            for i, v := range val {
                res[i] = v.(PreAuthorizedApplicationable)
            }
            m.SetPreAuthorizedApplications(res)
        }
        return nil
    }
    res["requestedAccessTokenVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedAccessTokenVersion(val)
        }
        return nil
    }
    return res
}
// GetKnownClientApplications gets the knownClientApplications property value. Used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app. If you set the appID of the client app to this value, the user only consents once to the client app. Azure AD knows that consenting to the client means implicitly consenting to the web API and automatically provisions service principals for both APIs at the same time. Both the client and the web API app must be registered in the same tenant.
func (m *ApiApplication) GetKnownClientApplications()([]i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.knownClientApplications
}
// GetOauth2PermissionScopes gets the oauth2PermissionScopes property value. The definition of the delegated permissions exposed by the web API represented by this application registration. These delegated permissions may be requested by a client application, and may be granted by users or administrators during consent. Delegated permissions are sometimes referred to as OAuth 2.0 scopes.
func (m *ApiApplication) GetOauth2PermissionScopes()([]PermissionScopeable) {
    return m.oauth2PermissionScopes
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ApiApplication) GetOdataType()(*string) {
    return m.odataType
}
// GetPreAuthorizedApplications gets the preAuthorizedApplications property value. Lists the client applications that are pre-authorized with the specified delegated permissions to access this application's APIs. Users are not required to consent to any pre-authorized application (for the permissions specified). However, any additional permissions not listed in preAuthorizedApplications (requested through incremental consent for example) will require user consent.
func (m *ApiApplication) GetPreAuthorizedApplications()([]PreAuthorizedApplicationable) {
    return m.preAuthorizedApplications
}
// GetRequestedAccessTokenVersion gets the requestedAccessTokenVersion property value. Specifies the access token version expected by this resource. This changes the version and format of the JWT produced independent of the endpoint or client used to request the access token.  The endpoint used, v1.0 or v2.0, is chosen by the client and only impacts the version of id_tokens. Resources need to explicitly configure requestedAccessTokenVersion to indicate the supported access token format.  Possible values for requestedAccessTokenVersion are 1, 2, or null. If the value is null, this defaults to 1, which corresponds to the v1.0 endpoint.  If signInAudience on the application is configured as AzureADandPersonalMicrosoftAccount, the value for this property must be 2
func (m *ApiApplication) GetRequestedAccessTokenVersion()(*int32) {
    return m.requestedAccessTokenVersion
}
// Serialize serializes information the current object
func (m *ApiApplication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("acceptMappedClaims", m.GetAcceptMappedClaims())
        if err != nil {
            return err
        }
    }
    if m.GetKnownClientApplications() != nil {
        err := writer.WriteCollectionOfUUIDValues("knownClientApplications", m.GetKnownClientApplications())
        if err != nil {
            return err
        }
    }
    if m.GetOauth2PermissionScopes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOauth2PermissionScopes()))
        for i, v := range m.GetOauth2PermissionScopes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("oauth2PermissionScopes", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetPreAuthorizedApplications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPreAuthorizedApplications()))
        for i, v := range m.GetPreAuthorizedApplications() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("preAuthorizedApplications", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("requestedAccessTokenVersion", m.GetRequestedAccessTokenVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcceptMappedClaims sets the acceptMappedClaims property value. When true, allows an application to use claims mapping without specifying a custom signing key.
func (m *ApiApplication) SetAcceptMappedClaims(value *bool)() {
    m.acceptMappedClaims = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApiApplication) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKnownClientApplications sets the knownClientApplications property value. Used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app. If you set the appID of the client app to this value, the user only consents once to the client app. Azure AD knows that consenting to the client means implicitly consenting to the web API and automatically provisions service principals for both APIs at the same time. Both the client and the web API app must be registered in the same tenant.
func (m *ApiApplication) SetKnownClientApplications(value []i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.knownClientApplications = value
}
// SetOauth2PermissionScopes sets the oauth2PermissionScopes property value. The definition of the delegated permissions exposed by the web API represented by this application registration. These delegated permissions may be requested by a client application, and may be granted by users or administrators during consent. Delegated permissions are sometimes referred to as OAuth 2.0 scopes.
func (m *ApiApplication) SetOauth2PermissionScopes(value []PermissionScopeable)() {
    m.oauth2PermissionScopes = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ApiApplication) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPreAuthorizedApplications sets the preAuthorizedApplications property value. Lists the client applications that are pre-authorized with the specified delegated permissions to access this application's APIs. Users are not required to consent to any pre-authorized application (for the permissions specified). However, any additional permissions not listed in preAuthorizedApplications (requested through incremental consent for example) will require user consent.
func (m *ApiApplication) SetPreAuthorizedApplications(value []PreAuthorizedApplicationable)() {
    m.preAuthorizedApplications = value
}
// SetRequestedAccessTokenVersion sets the requestedAccessTokenVersion property value. Specifies the access token version expected by this resource. This changes the version and format of the JWT produced independent of the endpoint or client used to request the access token.  The endpoint used, v1.0 or v2.0, is chosen by the client and only impacts the version of id_tokens. Resources need to explicitly configure requestedAccessTokenVersion to indicate the supported access token format.  Possible values for requestedAccessTokenVersion are 1, 2, or null. If the value is null, this defaults to 1, which corresponds to the v1.0 endpoint.  If signInAudience on the application is configured as AzureADandPersonalMicrosoftAccount, the value for this property must be 2
func (m *ApiApplication) SetRequestedAccessTokenVersion(value *int32)() {
    m.requestedAccessTokenVersion = value
}
