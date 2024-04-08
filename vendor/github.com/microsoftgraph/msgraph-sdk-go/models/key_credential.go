package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// KeyCredential 
type KeyCredential struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A 40-character binary type that can be used to identify the credential. Optional. When not provided in the payload, defaults to the thumbprint of the certificate.
    customKeyIdentifier []byte
    // Friendly name for the key. Optional.
    displayName *string
    // The date and time at which the credential expires. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The certificate's raw data in byte array converted to Base64 string. Returned only on $select for a single object, that is, GET applications/{applicationId}?$select=keyCredentials or GET servicePrincipals/{servicePrincipalId}?$select=keyCredentials; otherwise, it is always null.
    key []byte
    // The unique identifier (GUID) for the key.
    keyId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The OdataType property
    odataType *string
    // The date and time at which the credential becomes valid.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The type of key credential; for example, Symmetric, AsymmetricX509Cert.
    typeEscaped *string
    // A string that describes the purpose for which the key can be used; for example, Verify.
    usage *string
}
// NewKeyCredential instantiates a new keyCredential and sets the default values.
func NewKeyCredential()(*KeyCredential) {
    m := &KeyCredential{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateKeyCredentialFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateKeyCredentialFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKeyCredential(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *KeyCredential) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustomKeyIdentifier gets the customKeyIdentifier property value. A 40-character binary type that can be used to identify the credential. Optional. When not provided in the payload, defaults to the thumbprint of the certificate.
func (m *KeyCredential) GetCustomKeyIdentifier()([]byte) {
    return m.customKeyIdentifier
}
// GetDisplayName gets the displayName property value. Friendly name for the key. Optional.
func (m *KeyCredential) GetDisplayName()(*string) {
    return m.displayName
}
// GetEndDateTime gets the endDateTime property value. The date and time at which the credential expires. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *KeyCredential) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *KeyCredential) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["customKeyIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomKeyIdentifier(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    res["keyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyId(val)
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
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    res["usage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsage(val)
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. The certificate's raw data in byte array converted to Base64 string. Returned only on $select for a single object, that is, GET applications/{applicationId}?$select=keyCredentials or GET servicePrincipals/{servicePrincipalId}?$select=keyCredentials; otherwise, it is always null.
func (m *KeyCredential) GetKey()([]byte) {
    return m.key
}
// GetKeyId gets the keyId property value. The unique identifier (GUID) for the key.
func (m *KeyCredential) GetKeyId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.keyId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *KeyCredential) GetOdataType()(*string) {
    return m.odataType
}
// GetStartDateTime gets the startDateTime property value. The date and time at which the credential becomes valid.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *KeyCredential) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetType gets the type property value. The type of key credential; for example, Symmetric, AsymmetricX509Cert.
func (m *KeyCredential) GetType()(*string) {
    return m.typeEscaped
}
// GetUsage gets the usage property value. A string that describes the purpose for which the key can be used; for example, Verify.
func (m *KeyCredential) GetUsage()(*string) {
    return m.usage
}
// Serialize serializes information the current object
func (m *KeyCredential) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("customKeyIdentifier", m.GetCustomKeyIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("keyId", m.GetKeyId())
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
    {
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("usage", m.GetUsage())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *KeyCredential) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustomKeyIdentifier sets the customKeyIdentifier property value. A 40-character binary type that can be used to identify the credential. Optional. When not provided in the payload, defaults to the thumbprint of the certificate.
func (m *KeyCredential) SetCustomKeyIdentifier(value []byte)() {
    m.customKeyIdentifier = value
}
// SetDisplayName sets the displayName property value. Friendly name for the key. Optional.
func (m *KeyCredential) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEndDateTime sets the endDateTime property value. The date and time at which the credential expires. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *KeyCredential) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetKey sets the key property value. The certificate's raw data in byte array converted to Base64 string. Returned only on $select for a single object, that is, GET applications/{applicationId}?$select=keyCredentials or GET servicePrincipals/{servicePrincipalId}?$select=keyCredentials; otherwise, it is always null.
func (m *KeyCredential) SetKey(value []byte)() {
    m.key = value
}
// SetKeyId sets the keyId property value. The unique identifier (GUID) for the key.
func (m *KeyCredential) SetKeyId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.keyId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *KeyCredential) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStartDateTime sets the startDateTime property value. The date and time at which the credential becomes valid.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *KeyCredential) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetType sets the type property value. The type of key credential; for example, Symmetric, AsymmetricX509Cert.
func (m *KeyCredential) SetType(value *string)() {
    m.typeEscaped = value
}
// SetUsage sets the usage property value. A string that describes the purpose for which the key can be used; for example, Verify.
func (m *KeyCredential) SetUsage(value *string)() {
    m.usage = value
}
