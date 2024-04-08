package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SelfSignedCertificate 
type SelfSignedCertificate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The customKeyIdentifier property
    customKeyIdentifier []byte
    // The displayName property
    displayName *string
    // The endDateTime property
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The key property
    key []byte
    // The keyId property
    keyId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The OdataType property
    odataType *string
    // The startDateTime property
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The thumbprint property
    thumbprint *string
    // The type property
    typeEscaped *string
    // The usage property
    usage *string
}
// NewSelfSignedCertificate instantiates a new SelfSignedCertificate and sets the default values.
func NewSelfSignedCertificate()(*SelfSignedCertificate) {
    m := &SelfSignedCertificate{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSelfSignedCertificateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSelfSignedCertificateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSelfSignedCertificate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SelfSignedCertificate) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustomKeyIdentifier gets the customKeyIdentifier property value. The customKeyIdentifier property
func (m *SelfSignedCertificate) GetCustomKeyIdentifier()([]byte) {
    return m.customKeyIdentifier
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *SelfSignedCertificate) GetDisplayName()(*string) {
    return m.displayName
}
// GetEndDateTime gets the endDateTime property value. The endDateTime property
func (m *SelfSignedCertificate) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SelfSignedCertificate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["thumbprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbprint(val)
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
// GetKey gets the key property value. The key property
func (m *SelfSignedCertificate) GetKey()([]byte) {
    return m.key
}
// GetKeyId gets the keyId property value. The keyId property
func (m *SelfSignedCertificate) GetKeyId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.keyId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SelfSignedCertificate) GetOdataType()(*string) {
    return m.odataType
}
// GetStartDateTime gets the startDateTime property value. The startDateTime property
func (m *SelfSignedCertificate) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetThumbprint gets the thumbprint property value. The thumbprint property
func (m *SelfSignedCertificate) GetThumbprint()(*string) {
    return m.thumbprint
}
// GetType gets the type property value. The type property
func (m *SelfSignedCertificate) GetType()(*string) {
    return m.typeEscaped
}
// GetUsage gets the usage property value. The usage property
func (m *SelfSignedCertificate) GetUsage()(*string) {
    return m.usage
}
// Serialize serializes information the current object
func (m *SelfSignedCertificate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("thumbprint", m.GetThumbprint())
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
func (m *SelfSignedCertificate) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustomKeyIdentifier sets the customKeyIdentifier property value. The customKeyIdentifier property
func (m *SelfSignedCertificate) SetCustomKeyIdentifier(value []byte)() {
    m.customKeyIdentifier = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *SelfSignedCertificate) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEndDateTime sets the endDateTime property value. The endDateTime property
func (m *SelfSignedCertificate) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetKey sets the key property value. The key property
func (m *SelfSignedCertificate) SetKey(value []byte)() {
    m.key = value
}
// SetKeyId sets the keyId property value. The keyId property
func (m *SelfSignedCertificate) SetKeyId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.keyId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SelfSignedCertificate) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStartDateTime sets the startDateTime property value. The startDateTime property
func (m *SelfSignedCertificate) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetThumbprint sets the thumbprint property value. The thumbprint property
func (m *SelfSignedCertificate) SetThumbprint(value *string)() {
    m.thumbprint = value
}
// SetType sets the type property value. The type property
func (m *SelfSignedCertificate) SetType(value *string)() {
    m.typeEscaped = value
}
// SetUsage sets the usage property value. The usage property
func (m *SelfSignedCertificate) SetUsage(value *string)() {
    m.usage = value
}
