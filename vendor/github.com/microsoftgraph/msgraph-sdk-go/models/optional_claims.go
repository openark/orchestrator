package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OptionalClaims 
type OptionalClaims struct {
    // The optional claims returned in the JWT access token.
    accessToken []OptionalClaimable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The optional claims returned in the JWT ID token.
    idToken []OptionalClaimable
    // The OdataType property
    odataType *string
    // The optional claims returned in the SAML token.
    saml2Token []OptionalClaimable
}
// NewOptionalClaims instantiates a new optionalClaims and sets the default values.
func NewOptionalClaims()(*OptionalClaims) {
    m := &OptionalClaims{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOptionalClaimsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOptionalClaimsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOptionalClaims(), nil
}
// GetAccessToken gets the accessToken property value. The optional claims returned in the JWT access token.
func (m *OptionalClaims) GetAccessToken()([]OptionalClaimable) {
    return m.accessToken
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OptionalClaims) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OptionalClaims) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOptionalClaimFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OptionalClaimable, len(val))
            for i, v := range val {
                res[i] = v.(OptionalClaimable)
            }
            m.SetAccessToken(res)
        }
        return nil
    }
    res["idToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOptionalClaimFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OptionalClaimable, len(val))
            for i, v := range val {
                res[i] = v.(OptionalClaimable)
            }
            m.SetIdToken(res)
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
    res["saml2Token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOptionalClaimFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OptionalClaimable, len(val))
            for i, v := range val {
                res[i] = v.(OptionalClaimable)
            }
            m.SetSaml2Token(res)
        }
        return nil
    }
    return res
}
// GetIdToken gets the idToken property value. The optional claims returned in the JWT ID token.
func (m *OptionalClaims) GetIdToken()([]OptionalClaimable) {
    return m.idToken
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OptionalClaims) GetOdataType()(*string) {
    return m.odataType
}
// GetSaml2Token gets the saml2Token property value. The optional claims returned in the SAML token.
func (m *OptionalClaims) GetSaml2Token()([]OptionalClaimable) {
    return m.saml2Token
}
// Serialize serializes information the current object
func (m *OptionalClaims) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccessToken() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessToken()))
        for i, v := range m.GetAccessToken() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("accessToken", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIdToken() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIdToken()))
        for i, v := range m.GetIdToken() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("idToken", cast)
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
    if m.GetSaml2Token() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSaml2Token()))
        for i, v := range m.GetSaml2Token() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("saml2Token", cast)
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
// SetAccessToken sets the accessToken property value. The optional claims returned in the JWT access token.
func (m *OptionalClaims) SetAccessToken(value []OptionalClaimable)() {
    m.accessToken = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OptionalClaims) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIdToken sets the idToken property value. The optional claims returned in the JWT ID token.
func (m *OptionalClaims) SetIdToken(value []OptionalClaimable)() {
    m.idToken = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OptionalClaims) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSaml2Token sets the saml2Token property value. The optional claims returned in the SAML token.
func (m *OptionalClaims) SetSaml2Token(value []OptionalClaimable)() {
    m.saml2Token = value
}
