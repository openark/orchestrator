package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivacyProfile 
type PrivacyProfile struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A valid smtp email address for the privacy statement contact. Not required.
    contactEmail *string
    // The OdataType property
    odataType *string
    // A valid URL format that begins with http:// or https://. Maximum length is 255 characters. The URL that directs to the company's privacy statement. Not required.
    statementUrl *string
}
// NewPrivacyProfile instantiates a new privacyProfile and sets the default values.
func NewPrivacyProfile()(*PrivacyProfile) {
    m := &PrivacyProfile{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePrivacyProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivacyProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivacyProfile(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrivacyProfile) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContactEmail gets the contactEmail property value. A valid smtp email address for the privacy statement contact. Not required.
func (m *PrivacyProfile) GetContactEmail()(*string) {
    return m.contactEmail
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivacyProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contactEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactEmail(val)
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
    res["statementUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatementUrl(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PrivacyProfile) GetOdataType()(*string) {
    return m.odataType
}
// GetStatementUrl gets the statementUrl property value. A valid URL format that begins with http:// or https://. Maximum length is 255 characters. The URL that directs to the company's privacy statement. Not required.
func (m *PrivacyProfile) GetStatementUrl()(*string) {
    return m.statementUrl
}
// Serialize serializes information the current object
func (m *PrivacyProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("contactEmail", m.GetContactEmail())
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
        err := writer.WriteStringValue("statementUrl", m.GetStatementUrl())
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
func (m *PrivacyProfile) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContactEmail sets the contactEmail property value. A valid smtp email address for the privacy statement contact. Not required.
func (m *PrivacyProfile) SetContactEmail(value *string)() {
    m.contactEmail = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PrivacyProfile) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStatementUrl sets the statementUrl property value. A valid URL format that begins with http:// or https://. Maximum length is 255 characters. The URL that directs to the company's privacy statement. Not required.
func (m *PrivacyProfile) SetStatementUrl(value *string)() {
    m.statementUrl = value
}
