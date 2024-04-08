package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsInformationProtectionDataRecoveryCertificate windows Information Protection DataRecoveryCertificate
type WindowsInformationProtectionDataRecoveryCertificate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Data recovery Certificate
    certificate []byte
    // Data recovery Certificate description
    description *string
    // Data recovery Certificate expiration datetime
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // Data recovery Certificate subject name
    subjectName *string
}
// NewWindowsInformationProtectionDataRecoveryCertificate instantiates a new windowsInformationProtectionDataRecoveryCertificate and sets the default values.
func NewWindowsInformationProtectionDataRecoveryCertificate()(*WindowsInformationProtectionDataRecoveryCertificate) {
    m := &WindowsInformationProtectionDataRecoveryCertificate{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsInformationProtectionDataRecoveryCertificateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsInformationProtectionDataRecoveryCertificateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsInformationProtectionDataRecoveryCertificate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsInformationProtectionDataRecoveryCertificate) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCertificate gets the certificate property value. Data recovery Certificate
func (m *WindowsInformationProtectionDataRecoveryCertificate) GetCertificate()([]byte) {
    return m.certificate
}
// GetDescription gets the description property value. Data recovery Certificate description
func (m *WindowsInformationProtectionDataRecoveryCertificate) GetDescription()(*string) {
    return m.description
}
// GetExpirationDateTime gets the expirationDateTime property value. Data recovery Certificate expiration datetime
func (m *WindowsInformationProtectionDataRecoveryCertificate) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsInformationProtectionDataRecoveryCertificate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["certificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificate(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
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
    res["subjectName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectName(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WindowsInformationProtectionDataRecoveryCertificate) GetOdataType()(*string) {
    return m.odataType
}
// GetSubjectName gets the subjectName property value. Data recovery Certificate subject name
func (m *WindowsInformationProtectionDataRecoveryCertificate) GetSubjectName()(*string) {
    return m.subjectName
}
// Serialize serializes information the current object
func (m *WindowsInformationProtectionDataRecoveryCertificate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("certificate", m.GetCertificate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
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
        err := writer.WriteStringValue("subjectName", m.GetSubjectName())
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
func (m *WindowsInformationProtectionDataRecoveryCertificate) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCertificate sets the certificate property value. Data recovery Certificate
func (m *WindowsInformationProtectionDataRecoveryCertificate) SetCertificate(value []byte)() {
    m.certificate = value
}
// SetDescription sets the description property value. Data recovery Certificate description
func (m *WindowsInformationProtectionDataRecoveryCertificate) SetDescription(value *string)() {
    m.description = value
}
// SetExpirationDateTime sets the expirationDateTime property value. Data recovery Certificate expiration datetime
func (m *WindowsInformationProtectionDataRecoveryCertificate) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsInformationProtectionDataRecoveryCertificate) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSubjectName sets the subjectName property value. Data recovery Certificate subject name
func (m *WindowsInformationProtectionDataRecoveryCertificate) SetSubjectName(value *string)() {
    m.subjectName = value
}
